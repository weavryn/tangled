<script lang="ts">
    import logo from '../images/logo.png'
    import IcRoundMinimize from '~icons/ic/round-minimize'
    import IcRoundClose from '~icons/ic/round-close'
    import { WindowMinimise, WindowHide, Quit,  } from '../../../wailsjs/runtime';
    import { keyHandle } from '../js/keyHandle'
</script>

<style>
    nav {
        --wails-draggable: drag;
    }
</style>

<nav class="h-11 bg-cat-crust fixed top-0 left-0 w-screen flex flex-row justify-between px-3 items-center text-cat-text">
    <section class="flex flex-row items-center gap-2 select-none">
        <img src={logo} alt="tangled logo" class="w-8 h-8" />
        <span class="text-lg">Tangled</span>
    </section>
    <section class="flex flex-row items-center gap-2 select-none cursor-pointer text-lg">
        <btn on:click={WindowMinimise} on:keydown={(e) => { keyHandle(e, WindowMinimise) }} class="text-xl hover:text-cat-pink rounded-full transition-colors">
            <IcRoundMinimize />
        </btn>
        <btn on:click={Quit} on:keydown={(e) => { keyHandle(e, Quit) }} class="text-xl hover:text-cat-pink rounded-full transition-colors">
            <IcRoundClose />
        </btn>
    </section>
</nav>