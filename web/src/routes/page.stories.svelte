<script module lang="ts">
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Overview from './overview.svelte';

	const { Story } = defineMeta({
		title: 'Page',
		component: Overview,
		tags: ['autodocs'],
		args: {},
		argTypes: {}
	});
</script>

<Story name="Wide">
	<div class="wide-container">
			<Overview />
	</div>
</Story>

<Story name="Tall">
	<div class="tall-container">
			<Overview />
	</div>
</Story>

<style>
	:root {
		--clock-height: 300px;
		--clock-width: 400px;
	}
	.wide-container {
		width: var(--clock-width);
		height: var(--clock-height);
	}

	.tall-container {
		width: var(--clock-height);
		height: var(--clock-width);
	}
</style>
