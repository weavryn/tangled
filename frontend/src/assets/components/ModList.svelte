<script lang="ts">
    import MdiCheckboxBlankOutline from '~icons/mdi/checkbox-blank-outline'
    import MdiCheckboxOutline from '~icons/mdi/checkbox-outline'
    import MdiDelete from '~icons/mdi/delete'
    import MaterialSymbolsAddRounded from '~icons/material-symbols/add-rounded'
</script>
<div class="h-full w-full bg-cat-surface0 rounded-tl-md rounded-bl-md flex flex-col overflow-y-scroll scrollbar">
    {#each Array(100) as mod}
    <div class="w-full h-12 min-h-[3rem] even:bg-cat-surface1 odd:bg-cat-surface2 first:rounded-tl-md last:rounded-bl-md flex flex-row justify-between items-center gap-2 px-2">
        <div>
            <MdiCheckboxOutline class="text-cat-pink text-2xl" />
        </div>
        <div class="w-full text-start">
            RavenWeave (RavenWeave-1.1.jar)
        </div>
        <div class="flex flex-row justify-end gap-2 text-2xl">
            <MaterialSymbolsAddRounded />
            <MdiDelete />
        </div>
    </div>
    {/each}
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