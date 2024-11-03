<script lang="ts">
	import { nthNumber } from '$lib/nthNumber';
    import { onMount } from 'svelte';

    let {forceTime } : {forceTime?: Date, underText?: string} = $props();
    let currentTime = $state(new Date());
    let pulsePlaying: boolean = $state(false);

    onMount(() => {
        if (forceTime !== undefined) {
            currentTime = forceTime;
            return;
        }
        const now = new Date();
        const msToNextHalfSecond = 500 - (now.getMilliseconds() % 1000);
        
        setTimeout(() => {
            pulsePlaying = true;
        }, msToNextHalfSecond);

        const intervalId = setInterval(() => {
            const newTime = new Date();
            if (newTime.getSeconds() !== currentTime.getSeconds()) {
                currentTime = newTime;
            }
        }, 1000/60);

        return () => {
            clearInterval(intervalId);
        };
    });

    const twelveHour = (t: Date) => {
        let h = t.getHours();
        return (h % 12 || 12).toString();
    };
    
    let day: string = $derived(nthNumber(currentTime.getDate()));
    let month: string = $derived(currentTime.toLocaleString('default', { month: 'long' }));
    let year: string = $derived(currentTime.getFullYear().toString());

    let hours: string = $derived(twelveHour(currentTime));
    let minutes = $derived(currentTime.getMinutes().toString().padStart(2, '0'));
    let seconds = $derived(currentTime.getSeconds().toString().padStart(2, '0'));
    let ampm = $derived(currentTime.getHours() >= 12 ? 'PM' : 'AM');
</script>

<style>
@keyframes brightnessAlternate{
    0% {
        filter:brightness(0.8);
    }
    35% {
        filter:brightness(0.8);
    }
    65% {
        filter:brightness(1);
    }
    100% {
        filter:brightness(1);
    }
}

@keyframes saturatePulse{
    0%, 45%, 100% {
        filter:saturate(100%);
    }
    50%, 60% {
        filter:saturate(175%);
    }
}

.fancy-pulse.playing {
    animation-play-state: running;
}

.fancy-pulse {
    animation-play-state: paused;
    animation-name: saturatePulse;
    animation-duration: 1s;
    animation-iteration-count: infinite;
}
</style>

<div class="p-3 w-fit">
    <div class="flex flex-row font-extralight text-xs text-gray-400">
        <div class="contents">it is the</div>
        <div class="contents text-primary font-medium">{day}</div>
        <div class="contents">of</div>
        <div class="contents font-light">{month}</div>
        <div class="contents font-light">{year}</div>
        <div class="contents">.</div>
    </div>
    <div class="flex flex-row justify-center gap-2">
        <!-- <div class="flex-1"></div> -->
        <div class="text-6xl flex flex-row gap-0.5">
            <div class="font-bold mr-0.5">{hours}</div>
            <div class="relative top-[3px] text-5xl font-bold text-rabbit fancy-pulse" class:playing={pulsePlaying}>:</div>
            <div class="text-primary font-thin">{minutes}</div>
            <!-- <div class="relative top-[3px] text-5xl font-bold text-rabbit fancy-pulse" class:playing={pulsePlaying}>:</div>
            <div class="text-primary font-thin">{seconds}</div> -->
        </div>
        <div class="flex-1 text-base" class:self-end={ampm === 'PM'}>{ampm}</div>
    </div>
</div>