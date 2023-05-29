package main

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type Mods struct {
	Mods []Mod `json:"mods"`
}

type Mod struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	FileName string `json:"fileName"`
}

type Updates struct {
	Tangled TangledUpdate `json:"tangled"`
	Weave   WeaveUpdate   `json:"weave"`
}

type TangledUpdate struct {
	CurrentVersion string `json:"currentVersion"`
	LatestVersion  string `json:"latestVersion"`
}

type WeaveUpdate struct {
	CurrentVersion string `json:"currentVersion"`
	LatestVersion  string `json:"latestVersion"`
}

type WeaveManifest struct {
	MixinConfigs []string `json:"mixinConfigs"`
	EntryPoints  []string `json:"entrypoints"`
	Hooks        []string `json:"hooks"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartWatcher() {
	log.Println("starting watcher")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Somehow, you do not have a home directory. Error: ", err)
	}
	modDir := homeDir + "/.weave/mods"
	log.Println(modDir)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	log.Println("the watcher has started.")

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					fmt.Println("Mod Watcher Closed on its own")
					return
				}
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
					runtime.EventsEmit(a.ctx, "modFSChange", event.Name)
				}
				if event.Has(fsnotify.Create) {
					log.Println("modified file:", event.Name)
					runtime.EventsEmit(a.ctx, "modFSChange", event.Name)
				}
				if event.Has(fsnotify.Rename) {
					log.Println("modified file:", event.Name)
					runtime.EventsEmit(a.ctx, "modFSChange", event.Name)
				}
				if event.Has(fsnotify.Remove) {
					log.Println("modified file:", event.Name)
					runtime.EventsEmit(a.ctx, "modFSChange", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					fmt.Println("Mod Watcher Closed", err)
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(modDir)
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}

func (a *App) CheckForUpdates() {

}

func (a *App) ToggleMod(fileName string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Somehow, you do not have a home directory. Error: ", err)
	}
	modDir := homeDir + "/.weave/mods"
	var newFilename string
	if strings.HasSuffix(fileName, ".jar") {
		newFilename = strings.TrimSuffix(fileName, ".jar")
		newFilename = newFilename + ".jar.disabled"
	} else if strings.HasSuffix(fileName, ".jar.disabled") {
		newFilename = strings.TrimSuffix(fileName, ".jar.disabled")
		newFilename = newFilename + ".jar"
	}

	err = os.Rename(modDir+"/"+fileName, modDir+"/"+newFilename)
	if err != nil {
		log.Println("Error renaming file: ", err)
	}
}

func (a *App) SearchMods() []Mod {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Somehow, you do not have a home directory. Error: ", err)
	}
	modDir := homeDir + "/.weave/mods"

	files, err := ioutil.ReadDir(modDir)
	if err != nil {
		log.Fatal(err)
	}

	var mods []Mod

	for _, file := range files {
		if !file.IsDir() && (strings.HasSuffix(file.Name(), ".jar") || strings.HasSuffix(file.Name(), ".jar.disabled")) {
			jarFilePath := filepath.Join(modDir, file.Name())

			jarFile, err := zip.OpenReader(jarFilePath)
			if err != nil {
				log.Println("Error opening jar file: ", err)
				continue
			}
			defer jarFile.Close()

			jarFileName := file.Name()

			var jarVersion string
			for _, f := range jarFile.File {
				if strings.EqualFold(f.Name, "META-INF/MANIFEST.MF") {
					manifestFile, err := f.Open()
					if err != nil {
						log.Println("Error opening manifest file: ", err)
						continue
					}
					defer manifestFile.Close()

					manifestBytes, err := ioutil.ReadAll(manifestFile)
					if err != nil {
						log.Println("Error reading manifest file: ", err)
						continue
					}

					manifestString := string(manifestBytes)

					versionIndex := strings.Index(manifestString, "Implementation-Version: ")
					if versionIndex >= 0 {
						versionLine := manifestString[versionIndex:]
						endIndex := strings.Index(versionLine, "\n")
						version := strings.TrimSpace(versionLine[len("Implementation-Version: ") : endIndex-1])
						jarVersion = version
					}
					break
				}
			}

			if jarVersion == "" {
				// Extract version from JAR file name
				versionRegex := regexp.MustCompile(`-(\d+(\.\d+)+)`)
				versionMatches := versionRegex.FindStringSubmatch(jarFileName)
				if len(versionMatches) > 1 {
					jarVersion = versionMatches[1]
				} else {
					// Default to 1.0.0 if no version found in file name or manifest
					jarVersion = "1.0.0"
				}
			}

			var modName string
			for _, f := range jarFile.File {
				if strings.EqualFold(f.Name, "weave.mod.json") {
					jsonFile, err := f.Open()
					if err != nil {
						log.Println("Error opening json file: ", err)
						continue
					}
					defer jsonFile.Close()

					jsonBytes, err := ioutil.ReadAll(jsonFile)
					if err != nil {
						log.Println("Error reading json file: ", err)
						continue
					}

					var weaveManifest WeaveManifest
					err = json.Unmarshal(jsonBytes, &weaveManifest)
					if err != nil {
						log.Println("Error unmarshalling json file: ", err)
						continue
					}

					parts := strings.Split(jarFileName, weaveManifest.EntryPoints[0])
					modName = parts[len(parts)-1]
					break
				}
			}

			if modName != "" {
				r := regexp.MustCompile(`-\d+(\.\d+)+`)
				modName = r.ReplaceAllString(jarFileName, "")

				modName = strings.TrimSuffix(modName, filepath.Ext(modName))

				modName = strings.TrimSuffix(modName, ".jar")
			}

			mods = append(mods, Mod{
				Name:     modName,
				Version:  jarVersion,
				FileName: jarFileName,
			})

		}
	}

	return mods

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
