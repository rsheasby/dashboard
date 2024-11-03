<script lang="ts">
	import WeatherIcon from './weatherIcon.svelte';
    
    const { right }: {right?: boolean} = $props();
</script>

<div class="w-min h-min" class:justify-self-end={right}>
    weather goes here
    <!-- <WeatherIcon selectedIcon="drizzle"/> -->
</div>