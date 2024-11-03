<script module lang="ts">
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Clock from './clock.svelte';
	import type { ArgTypes } from '@storybook/svelte';

	const argTypes: ArgTypes = {
		underText: {
			control: 'text'
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

<Story name="Time">
	{#snippet children(args: any)}
		<div class="w-fit">
			<Clock underText={args.underText} />
		</div>
	{/snippet}
</Story>
