<script module>
  import { defineMeta } from '@storybook/addon-svelte-csf';
  import Clock from './clock.svelte';

  const { Story } = defineMeta({
    title: 'Components/Clock',
    component: Clock,
    tags: ['autodocs'],
    argTypes: {
    },
    args: {
    }
  });
</script>

<Story name="Primary" args={{}} />
