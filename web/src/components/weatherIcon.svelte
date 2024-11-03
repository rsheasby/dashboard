<script lang="ts">
    const { selectedIcon } = $props();

	const weatherIcons: Record<string, string> = import.meta.glob('/node_modules/@bybas/weather-icons/production/fill/all/*.svg', {
		query: 'url',
		eager: true
	});
	import { onMount } from 'svelte';
    onMount(() => {
        console.log(weatherIcons);
    });

	let selectedIconPath: string = $state(`/node_modules/@bybas/weather-icons/production/fill/all/${selectedIcon}.svg`);
</script>
<object data={`${selectedIconPath}`} class="h-20 w-20"></object>
