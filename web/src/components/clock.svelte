<script lang="ts">
	import { nthNumber } from '$lib/nthNumber';
	import { onMount } from 'svelte';

	let {
		forceTime,
		large,
		onclick,
		showDetail
	}: { forceTime?: Date; large?: boolean; onclick?: () => void; showDetail?: boolean } = $props();
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
	let ampm = 'AM';
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="flex flex-col gap-1 p-2" {onclick}>
	<div class="grid-date text-xs font-light text-gray-500">
		it is the <strong class="rabbit-text">{day}</strong> of {month}
		{year}.
	</div>
	<div class="grid-time flex flex-row justify-center gap-2">
		<!-- <div class="flex-1"></div> -->
		<div class="flex flex-row gap-0.5 leading-[1]">
			<div class="mr-0.5 text-6xl font-bold">{hours}</div>
			<div class="text-rabbit fancy-pulse text-5xl font-bold" class:playing={pulsePlaying}>:</div>
			<div class="text-primary text-6xl font-thin">{minutes}</div>
			<!-- <div class="relative top-[3px] text-5xl font-bold text-rabbit fancy-pulse" class:playing={pulsePlaying}>:</div>
            <div class="text-primary font-thin">{seconds}</div> -->
		</div>
		<div class="text-md relative top-0.5 flex-1" class:self-end={ampm === 'PM'}>{ampm}</div>
	</div>
	<div class="text-xs font-extralight">
		<div class="grid-meet-time text-gray-300">
			<!-- you have a meeting in <strong class="highlighted underarrow animate-pulse transition-opacity" class:underarrow-hide={!showDetail}>10 mins</strong>. -->
			<!-- your next meeting is at <em>12:30 AM</em>. -->
			<!-- your next meeting is in <em>two hours</em>. -->
			<!-- you should be in <strong class="highlighted underarrow right transition-opacity" class:underarrow-hide={!showDetail}>a meeting</strong>. -->
		</div>
		<div
			class="grid-meet-name flex flex-row items-center justify-end pr-1 text-right transition-opacity"
			class:opacity-0={!showDetail}
		>
        <div class="loader"></div>
			<em>Experience Standup</em>
		</div>
	</div>
</div>

<style>
	@keyframes saturatePulse {
		0%,
		45%,
		100% {
			filter: saturate(100%);
		}
		50%,
		60% {
			filter: saturate(200%);
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

	strong.rabbit-text {
		@apply text-primary font-medium;
	}

	strong.highlighted {
		@apply bg-rabbit px-0.5 text-black;
	}

	strong.underarrow {
		@apply relative;
	}

	strong.underarrow:after {
		content: '';
		@apply absolute -bottom-1 left-0 -z-10 h-0 w-0;

		border-left: 0.5em solid transparent;
		border-right: 0.5em solid transparent;
		border-top: 0.5em solid var(--primary);
	}

	strong.underarrow.right:after {
		left: unset;
		@apply right-0;
	}

	strong.underarrow-hide:after {
		@apply opacity-0;
	}

	em {
		@apply text-primary font-light;
	}

	.loader {
        @apply bg-gray-950;
		width: 0.8em;
		height: 0.8em;
		border-radius: 50%;
        border: 3px solid var(--primary);
		display: inline-block;
		position: relative;
		box-sizing: border-box;
		animation: rotation 10s linear infinite;
        margin-right: 0.25em;
	}
	.loader::after,
	.loader::before {
		content: '';
		box-sizing: border-box;
		position: absolute;
		left: 0;
		top: 0;
		background: var(--primary);
		width: 3px;
		height: 3px;
		transform: translate(100%, 100%);
		border-radius: 50%;
	}
	.loader::before {
		left: auto;
		top: auto;
		right: 0;
		bottom: 0;
		transform: translate(-100%, -100%);
	}
    
    /* .loader {
  width: 1em;
  height: 1em;
  display: inline-block;
  position: relative;
}
.loader::after,
.loader::before {
  content: '';  
  box-sizing: border-box;
  width: 1em;
  height: 1em;
  border-radius: 50%;
  background: white;
  position: absolute;
  left: 0;
  top: 0;
  animation: animloader 5s linear infinite;
}
.loader::after {
  animation-delay: 1s;
} */

@keyframes animloader {
  0% {
    transform: scale(0);
    opacity: 1;
  }
  100% {
    transform: scale(1);
    opacity: 0;
  }
}
    

	@keyframes rotation {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}
</style>
