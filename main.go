package main

import (
	"embed"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Somehow, you do not have a home directory. Error: ", err)
	}
	if _, err := os.Stat(homeDir + "/.tangled"); os.IsNotExist(err) {
		err := os.Mkdir(homeDir+"/.tangled", 0755)
		if err != nil {
			log.Fatalln("Could not create the .tangled folder in your home directory. Error: ", err)
		}
	}

	if _, err := os.Stat(homeDir + "/.tangled/agent"); os.IsNotExist(err) {
		err := os.Mkdir(homeDir+"/.tangled/agent", 0755)
		if err != nil {
			log.Fatalln("Could not create the .tangled/agent folder in your home directory. Error: ", err)
		}
	}

	if _, err := os.Stat(homeDir + "/.tangled/mods"); os.IsNotExist(err) {
		err := os.Mkdir(homeDir+"/.tangled/mods", 0755)
		if err != nil {
			log.Fatalln("Could not create the .tangled/mods folder in your home directory. Error: ", err)
		}
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Tangled",
		Width:  800,
		Height: 350,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 30, G: 30, B: 46, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Frameless:     true,
		DisableResize: true,
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
