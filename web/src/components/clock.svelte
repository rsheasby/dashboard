<script lang="ts">
	import { nthNumber } from '$lib/nthNumber';
    import { onMount } from 'svelte';

    let {forceTime, large, onclick} : {forceTime?: Date, large?: boolean, onclick?: () => void} = $props();
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
        }, 200);

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
        filter:saturate(200%);
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

strong {
    @apply text-primary font-medium;
}

em {
    @apply text-primary font-light;
}
</style>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="p-3 w-fit" {onclick}>
    <div class="font-light text-xs text-gray-400">
    it is the <strong>{day}</strong> of {month} {year}.
    </div>
    <div class="flex flex-row justify-center gap-2 leading-none">
        <!-- <div class="flex-1"></div> -->
        <div class="flex flex-row gap-0.5">
            <div class="text-6xl font-bold mr-0.5">{hours}</div>
            <div class="text-5xl font-bold text-rabbit fancy-pulse" class:playing={pulsePlaying}>:</div>
            <div class="text-6xl text-primary font-thin">{minutes}</div>
            <!-- <div class="relative top-[3px] text-5xl font-bold text-rabbit fancy-pulse" class:playing={pulsePlaying}>:</div>
            <div class="text-primary font-thin">{seconds}</div> -->
        </div>
        <div class="flex-1 text-md" class:self-end={ampm === 'PM'}>{ampm}</div>
    </div>
    <div class="font-extralight text-xs text-gray-500">
        <!-- you have a meeting in <strong class="animate-pulse">10</strong> minutes. -->
        your next meeting is at <em>12:30 AM</em>.
        <!-- your next meeting is in <em>two hours</em>. -->
    </div>
</div>