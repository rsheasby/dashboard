<script lang="ts">
	import { DotLottieSvelte } from '@lottiefiles/dotlottie-svelte';
	const { selectedIcon } = $props();

	const weatherIcons: Record<string, string> = import.meta.glob(
		'/node_modules/@lxg/weather-icons/production/fill/lottie/*.json',
		{
			query: 'url',
			eager: true
		}
	);
	import { onMount } from 'svelte';
	onMount(() => {
		console.log(weatherIcons);
	});

	let selectedIconPath: string = $state(
		`/node_modules/@lxg/weather-icons/production/fill/lottie/${selectedIcon}.json`
	);
</script>

<div id="weather-container" class="relative h-fit w-fit">
	<!-- svelte-ignore a11y_missing_attribute -->
	<DotLottieSvelte src={selectedIconPath} loop autoplay />
</div>

<style>
	/* #weather-icon {
		filter: grayscale(100%);
	}

	#weather-container:before {
		content: '';
		width: 100%;
		height: 100%;
		background-color: var(--primary);
		z-index: 1;
		position: absolute;
		mix-blend-mode: multiply;
	} */
</style>
