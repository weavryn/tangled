<script lang="ts">
    import TitleBar from './assets/components/TitleBar.svelte';
    import SubBar from './assets/components/SubBar.svelte';
    import { updater, watcherOn } from './assets/js/stores';
    import ModList from './assets/components/ModList.svelte';

    import MdiCheckboxBlankOutline from '~icons/mdi/checkbox-blank-outline'
    import MdiCheckboxOutline from '~icons/mdi/checkbox-outline'
    import MdiDelete from '~icons/mdi/delete'
    import MaterialSymbolsAddRounded from '~icons/material-symbols/add-rounded'
    import GgSpinnerTwoAlt from '~icons/gg/spinner-two-alt'
    import BiExclamationLg from '~icons/bi/exclamation-lg'

    
    import { SearchMods, StartWatcher, ToggleMod } from '../wailsjs/go/main/App';

    import { EventsOn } from '../wailsjs/runtime/runtime';

    // run watcher (as a promise)
    watcherOn.subscribe(v => {
      if(v == false) {
        watcherOn.set(true);
        Promise.all([StartWatcher()]).then((data) => {
          console.log(data);
        }).catch(e => {
          errorFound = true
          console.log(e);
        });
      }
    });

    let mods = [];
    let searched = false
    let errorFound = false

    async function search() {
      searched = false
      errorFound = false
      mods = []
      let result = await SearchMods()
      if (!result[0]) {
        errorFound = true
      } else {
        mods = result
      }
      searched = true
    }

    async function toggleMod(fileName) {
      await ToggleMod(fileName);
      search();
    }

    EventsOn('modFSChange', (data) => {
      search();
    });
    search();
</script>

<div id="container" class="cat-mocha h-screen w-screen bg-cat-base">
  <TitleBar />
  <main class="pt-11 pb-7 h-screen">
    <section class="p-2 h-full">
      <div class="h-full w-full bg-cat-surface0 rounded-tl-md rounded-bl-md flex flex-col overflow-y-scroll scrollbar">
        <!-- {#each Array(100) as mod}
        <div class="w-full h-12 min-h-[3rem] even:bg-cat-surface1 odd:bg-cat-surface2 first:rounded-tl-md last:rounded-bl-md flex flex-row justify-between items-center gap-2 px-2">
            <div>
                <MdiCheckboxOutline class="text-cat-pink text-2xl" />
            </div>
            <div class="w-full text-start">
                RavenWeave (RavenWeave-1.1.jar)
            </div>
            <div class="flex flex-row justify-end gap-2 text-2xl">
                <MaterialSymbolsAddRounded class="hover:text-cat-mantle hover:bg-cat-pink rounded-full transition-colors m-1 cursor-pointer" />
                <MdiDelete class="hover:text-cat-mantle hover:bg-cat-pink rounded-full transition-colors m-1 cursor-pointer" />
            </div>
        </div>
        {/each} -->
        {#if errorFound}
        <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
          <BiExclamationLg class="text-6xl text-cat-red" />
        </div>
        {/if}
        {#if !searched}
          <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
            <GgSpinnerTwoAlt class="animate-spin text-6xl text-cat-pink" />
          </div>
        {:else}
          {#if !errorFound}
          {#each mods as mod}
            <div class="w-full h-12 min-h-[3rem] even:bg-cat-surface1 odd:bg-cat-surface2 first:rounded-tl-md last:rounded-bl-md flex flex-row justify-between items-center gap-2 px-2">
                <div>
                  <button class="cursor-pointer flex items-center justify-center text-cat-pink text-2xl hover:text-cat-mantle hover:bg-cat-pink rounded-full p-1 transition-colors" on:click={toggleMod(mod.fileName)}>
                  {#if mod.fileName.endsWith('.jar')}
                    <MdiCheckboxOutline />
                  {:else}
                    <MdiCheckboxBlankOutline />
                  {/if}
                  </button>
                </div>
                <div class="w-full text-start">
                    {mod.name} {mod.version} ({mod.fileName})
                </div>
                <div class="flex flex-row justify-end gap-2 text-2xl">
                    <MaterialSymbolsAddRounded class="hover:text-cat-mantle hover:bg-cat-pink rounded-full transition-colors p-1 cursor-pointer" />
                    <MdiDelete class="hover:text-cat-mantle hover:bg-cat-pink rounded-full transition-colors p-1 cursor-pointer" />
                </div>
            </div>
          {/each}
          {/if}
        {/if}
        <div class="absolute top-0 left-0 right-0 m-auto w-[21rem] h-11 flex flex-row justify-center">
          <a href="#" class="bg-cat-red text-cat-mantle h-11 w-28 flex items-center justify-center gap-2 rounded-l-full">
            Launch
          </a>
          <a href="#" class="bg-cat-mauve text-cat-mantle h-11 w-28 flex items-center justify-center gap-2">
            Profiles
          </a>
          <a href="#" class="bg-cat-teal text-cat-mantle h-11 w-28 flex items-center justify-center gap-2 rounded-r-full">
            Mods
          </a>
        </div>
    </section>
  </main>
  <SubBar />
</div>

<style>
  .scrollbar::-webkit-scrollbar {
    width: 4px;
  }
  .scrollbar::-webkit-scrollbar-track {
    background: #6c7086;
    border-top-right-radius: 0.375rem;
    border-bottom-right-radius: 0.375rem;
  } 
  .scrollbar::-webkit-scrollbar-thumb {
    background-color: #f5c2e7;
    border-radius: 0.375rem;
  }
</style>
