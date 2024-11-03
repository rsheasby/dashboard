<script>
	import Clock from '@components/clock.svelte';
	import Background from '@components/background.svelte';
	import Weather from '@components/weather.svelte';
</script>

<Background>
	<div class="grid">
		<Clock />
        <Weather />
	</div>
</Background>
