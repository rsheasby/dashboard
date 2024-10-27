<script module>
  import { defineMeta } from '@storybook/addon-svelte-csf';
  import Widget from '@components/Widget.svelte';

  const { Story } = defineMeta({
    title: 'Components/Widget',
    component: Widget,
    tags: ['autodocs'],
    argTypes: {
    },
    args: {
    }
  });
</script>

<Story name="Primary" args={{}} />
