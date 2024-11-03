<script module lang="ts">
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Clock from './clock.svelte';
	import Background from './background.svelte';
	import type { ArgTypes } from '@storybook/svelte';

	const argTypes: ArgTypes = {
		underText: {
			control: 'text',
		}
	};

	const { Story } = defineMeta({
		title: 'Components/Clock',
		component: Clock,
		tags: ['autodocs'],
		argTypes,
		args: {
			underText: 'you have a meeting in 10 minutes.'
		}
	});
</script>

<Story name="Wide">
	{#snippet children(args: any)}
		<div class="wide-container">
			<Background>
				<Clock underText={args.underText} />
			</Background>
		</div>
	{/snippet}
</Story>

<Story name="Tall">
	{#snippet children(args: any)}
		<div class="tall-container">
			<Background>
				<Clock underText={args.underText} />
			</Background>
		</div>
	{/snippet}
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
