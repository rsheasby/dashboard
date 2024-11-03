<script lang="ts">
	import WeatherIcon from './weatherIcon.svelte';
    
    const { right }: {right?: boolean} = $props();
</script>

<div class="w-min h-min" class:justify-self-end={right}>
    blahblah
    <!-- <WeatherIcon selectedIcon="drizzle"/> -->
</div>