<script module lang="ts">
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Page from './page.svelte';

	const { Story } = defineMeta({
		title: 'Page',
		component: Page,
		tags: ['autodocs'],
		args: {},
		argTypes: {}
	});
</script>

<Story name="Wide">
	<div class="wide-container">
			<Page />
	</div>
</Story>

<Story name="Tall">
	<div class="tall-container">
			<Page />
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
