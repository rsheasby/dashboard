<script module lang="ts">
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Clock from './clock.svelte';
	import Background from './background.svelte';

	const args = {
		underText: 'you have a meeting in 10 minutes.'
	};
	type Args = typeof args;

	const { Story } = defineMeta({
		title: 'Components/Clock',
		component: Clock,
		tags: ['autodocs'],
		argTypes: {},
		args
	});
</script>

<Story name="Wide">
	{#snippet children(args: Args)}
		<div class="wide-container">
			<Background>
				<Clock underText={args.underText} />
			</Background>
		</div>
	{/snippet}
</Story>

<Story name="Tall">
	{#snippet children(args: Args)}
		<div class="tall-container">
			<Background>
				<Clock underText={args.underText} />
			</Background>
		</div>
	{/snippet}
</Story>

<style>
	.wide-container {
		width: 320px;
		height: 240px;
	}

	.tall-container {
		width: 240px;
		height: 320px;
	}
</style>
