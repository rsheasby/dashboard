<script lang="ts">
    import { onMount, onDestroy } from 'svelte';

    let {underText} : {underText?: string} = $props();
    let currentTime = $state(new Date());
    let intervalId: NodeJS.Timeout;

    onMount(() => {
        intervalId = setInterval(() => {
            currentTime = new Date();
        }, 200);

        return () => {
            clearInterval(intervalId);
        };
    });

    const twelveHour = (t: Date) => {
        let h = t.getHours();
        return (h % 12 || 12).toString();
    };

    let hours: string = $derived(twelveHour(currentTime));
    let minutes = $derived(currentTime.getMinutes().toString().padStart(2, '0'));
    let seconds = $derived(currentTime.getSeconds().toString().padStart(2, '0'));
    let ampm = $derived(currentTime.getHours() >= 12 ? 'PM' : 'AM');

</script>

<style>
</style>

<div class="flex flex-row justify-center gap-2 p-3">
    <div class="flex-1 text-xs"></div>
    <div class="text-6xl flex flex-row gap-0.5">
        <div class="font-bold mr-0.5">{hours}</div>
        <div class="text-rabbit relative top-[3px] text-5xl font-bold">:</div>
        <div class="text-rabbit font-thin">{minutes}</div>
    </div>
    <div class="flex-1 justify-self-end self-end text-md">{ampm}</div>
</div>
<div class="text-center text-xs font-thin">
    {underText}
</div>