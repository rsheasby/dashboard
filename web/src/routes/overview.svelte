<script lang="ts">
	import Clock from '@components/clock.svelte';
	import Weather from '@components/weather.svelte';

	type View = 'overview' | 'clock' | 'weather';
	let currentView = $state<View>('overview');
	let clockShown: boolean = $derived(currentView === 'overview' || currentView === 'clock');
	let weatherShown: boolean = $derived(currentView === 'overview' || currentView === 'weather');

	const toggleClock = () => {
		if (currentView === 'clock') {
			currentView = 'overview';
		} else {
			currentView = 'clock';
		}
	};
</script>

<div
	class="grid-overview grid h-full w-full place-content-between transition-zoom bg-black"
	class:grid-clock={currentView === 'clock'}
>
	<div
		id="clock"
		class="h-full flex justify-center items-center"
		hidden={!clockShown}
	>
        <div class="transition-zoom {currentView === 'clock' ? 'max-w-full' : 'w-full h-full'}"
		class:text-3xl={currentView === 'clock'}
        >
            <Clock onclick={toggleClock} showDetail={currentView === 'clock'} />
        </div>
	</div>
	<div id="weather" class="min-w-fit flex justify-end">
		<Weather right />
	</div>
</div>

<style>
	.grid-overview {
		grid-template-columns: 60% 40%;
		grid-template-rows: 50% 50%;
	}
	.grid-clock {
		grid-template-columns: 100% 0;
		grid-template-rows: 100% 0;
	}
    
    .transition-zoom {
		transition-duration: 0.5s;
    }

	#clock {
		/* grid-area: clock; */
	}
	#weather {
		/* grid-area: weather; */
	}
</style>
