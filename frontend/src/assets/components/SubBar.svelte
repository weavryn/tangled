<script lang="ts">
    import IcRoundCheck from '~icons/ic/round-check'
    import GgSpinnerTwoAlt from '~icons/gg/spinner-two-alt'
    import BiExclamationLg from '~icons/bi/exclamation-lg'
    import IcRoundClose from '~icons/ic/round-close'
    import BxsInjection from '~icons/bxs/injection'

    import { updater } from '../js/stores'

    let updaterData;

    updater.subscribe(v => {
        updaterData = v
    })
</script>

<footer class="fixed bottom-0 left-0 h-7 w-screen bg-cat-crust flex flex-row items-center justify-between">
    <section class="flex flex-row h-full">
        <a href="#" class="bg-cat-surface0 flex px-2 h-full items-center text-cat-text hover:text-cat-mantle hover:bg-cat-pink transition-colors">
            <span>Settings</span>
        </a>
        <a href="#" class="bg-cat-surface0 flex px-2 h-full items-center text-cat-text hover:text-cat-mantle hover:bg-cat-pink gap-2 transition-colors border-l border-l-cat-pink">
            <BxsInjection />
            <span>Waiting for Minecraft...</span>
        </a>
    </section>
    {#if updaterData == 0}
    <a href="#" class="bg-cat-surface0 flex px-2 h-full items-center text-cat-text cursor-default transition-colors gap-2">
        <GgSpinnerTwoAlt class="animate-spin" />
        <span>Checking for updates...</span>
    </a>
    {:else if updaterData == 1}
    <a href="#" class="bg-cat-surface0 flex px-2 h-full items-center text-cat-text hover:text-cat-mantle hover:bg-cat-pink transition-colors gap-2">
        <IcRoundCheck />
        <span>You're up to date!</span>
    </a>
    {:else if updaterData == 2}
    <a href="#" class="bg-cat-pink flex px-2 h-full items-center text-cat-mantle transition-colors gap-2">
        <BiExclamationLg class="animate-pulse" />
        <span>Weave is out of date! (v0.2.4)</span>
    </a>
    {:else if updaterData == 3}
    <a href="#" class="bg-cat-red flex px-2 h-full items-center text-cat-mantle transition-colors gap-2">
        <IcRoundClose class="animate-pulse" />
        <span>Failed to check for updates.</span>
    </a>
    {:else if updaterData == 4}
    <a href="#" class="bg-cat-red flex px-2 h-full items-center text-cat-mantle transition-colors gap-2">
        <IcRoundClose class="animate-pulse" />
        <span>Failed to update. Click here for log.</span>
    </a>
    {/if}
</footer>