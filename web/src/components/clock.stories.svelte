<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Clock from './clock.svelte';
	import Background from './background.svelte';

	const { Story } = defineMeta({
		title: 'Components/Clock',
		component: Clock,
		tags: ['autodocs'],
		argTypes: {},
		args: {}
	});
</script>

<Story name="Wide">
	<div class="wide-container">
		<Background>
			<Clock />
		</Background>
	</div>
</Story>

<Story name="Tall">
	<div class="tall-container">
		<Background>
			<Clock />
		</Background>
	</div>
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
