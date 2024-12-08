<script lang="ts">
	import { nthNumber } from '$lib/nthNumber';
	import { onMount } from 'svelte';

	type Meeting = {
		name: string;
		startTime: Date;
		secondsToStart: number;
		endTime: Date;
		secondsToEnd: number;
	};

	let {
		forceTime,
		large,
		onclick,
		showDetail
	}: { forceTime?: Date; large?: boolean; onclick?: () => void; showDetail?: boolean } = $props();
	let currentTime = $state(new Date());
	let pulsePlaying: boolean = $state(false);

	let meetings = $state<Meeting[]>([]);
	const currentMeeting: Meeting | undefined = $derived(
		meetings.find((meeting) => meeting.startTime <= currentTime && meeting.endTime >= currentTime)
	);
	const currentMeetingSecondsToEnd = $derived.by(() => {
		if (currentMeeting?.endTime?.getTime() === undefined) {
			return undefined;
		}
		return (currentMeeting.endTime.getTime() - currentTime.getTime()) / 1000;
	});
	const currentMeetingDuration = $derived.by(() => {
		if (currentMeeting?.startTime?.getTime() === undefined || currentMeeting?.endTime?.getTime() === undefined) {
			return undefined;
		}
		return (currentMeeting.endTime.getTime() - currentMeeting.startTime.getTime()) / 1000;
	});
	const currentMeetingElapsedTime = $derived.by(() => {
		if (currentMeeting?.startTime?.getTime() === undefined) {
			return undefined;
		}
		return (currentTime.getTime() - currentMeeting.startTime.getTime()) / 1000;
	});
	let nextMeeting: Meeting | undefined = $derived(
		meetings.find((meeting) => meeting.startTime > currentTime)
	);
	const nextMeetingSecondsToStart = $derived.by(() => {
		if (nextMeeting?.startTime?.getTime() === undefined) {
			return undefined;
		}
		return (nextMeeting.startTime.getTime() - currentTime.getTime()) / 1000;
	});

	async function fetchMeetings() {
		try {
			const response = await fetch('https://dashboard.lobster-drum.ts.net/meetings');
			const data = await response.json();
			const parsedData = data.map((meeting: any) => ({
				...meeting,
				startTime: new Date(meeting.startTime),
				endTime: new Date(meeting.endTime)
			}));
			if (JSON.stringify(meetings) !== JSON.stringify(parsedData)) {
				meetings = parsedData;
			}
		} catch (error) {
			console.error('Error fetching meetings:', error);
		}
		console.log({ currentMeeting, nextMeeting });
	}

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
		}, 1000 / 60);

		fetchMeetings();
		const fetchIntervalId = setInterval(fetchMeetings, 5000);

		return () => {
			clearInterval(intervalId);
			clearInterval(fetchIntervalId);
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

	let remainingMinutes = $derived(
		currentMeeting
			? Math.ceil((currentMeeting.endTime.getTime() - currentTime.getTime()) / (60 * 1000))
			: 0
	);
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="flex w-fit max-w-full flex-col gap-1 p-2" {onclick}>
	<div class="grid-date text-xs font-light text-gray-500">
		it is the <strong class="rabbit-text">{day}</strong> of {month}
		{year},
	</div>
	<div class="grid-time flex flex-row flex-wrap justify-center gap-1">
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
		{#if currentMeeting}
			<div class="grid-meet-time text-right text-xs text-gray-300">
				and your meeting ends in <strong class="highlighted underarrow right"
					>{remainingMinutes} minutes</strong
				>.
			</div>
			<div
				class="grid-meet-name flex flex-col items-end justify-center pr-1 text-right text-xs transition-opacity"
			>
				<!-- <div class="loader"></div> -->
				<em>{currentMeeting.name}</em>
				<progress class="w-full h-0.5 progress-rabbit" value={currentMeetingElapsedTime} max={currentMeetingDuration}></progress>
			</div>
		{:else if nextMeeting}
			{#if nextMeetingSecondsToStart && nextMeetingSecondsToStart < 60 * 60 * 1}
				<div class="grid-meet-time text-right text-xs text-gray-300">
					and your next meeting is in
					<strong class="highlighted underarrow right" class:underarrow-hide={!showDetail} class:animate-pulse={nextMeetingSecondsToStart < 60 * 10}>
						{Math.ceil(nextMeetingSecondsToStart / 60)} minutes
					</strong>.
				</div>
				<div
					class="grid-meet-name flex flex-row items-center justify-end pr-1 text-right text-xs transition-opacity"
					class:opacity-0={!showDetail}
				>
					<em>{nextMeeting?.name}</em>
				</div>
			{:else}
				<div class="grid-meet-time text-right text-xs text-gray-300">
					and your next meeting is at
					<strong>
						{nextMeeting?.startTime?.toLocaleTimeString([], {
							hour: 'numeric',
							minute: '2-digit',
							hour12: true
						})}
					</strong>.
				</div>
				<div
					class="grid-meet-name flex flex-row items-center justify-end pr-1 text-right text-xs transition-opacity"
					class:opacity-0={!showDetail}
				>
				</div>
			{/if}
		{:else}
			<div class="grid-meet-time text-right text-xs text-gray-300">and your calendar is free.</div>
		{/if}
	</div>
</div>

<style>
	.progress-rabbit::-webkit-progress-value {
		background: var(--primary);
	}
	@keyframes saturatePulse {
		0%,
		45%,
		100% {
			filter: saturate(100%);
		}
		50%,
		60% {
			filter: saturate(150%);
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
		z-index: 0;
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
		border: 0.15em solid var(--primary);
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
		width: 0.15em;
		height: 0.15em;
		transform: translate(50%, 50%);
		border-radius: 50%;
	}
	.loader::before {
		left: auto;
		top: auto;
		right: 0;
		bottom: 0;
		transform: translate(-50%, -50%);
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
