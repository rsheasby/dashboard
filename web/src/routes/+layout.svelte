<script lang="ts">
	import Background from '@components/background.svelte';
	import '../app.css';

	let { children } = $props();
</script>

<Background>
	{@render children()}
</Background>
