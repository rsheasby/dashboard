<script module>
  import { defineMeta } from '@storybook/addon-svelte-csf';
  import Widget from '@components/Widget.svelte';
  import { fn } from '@storybook/test';

  // More on how to set up stories at: https://storybook.js.org/docs/writing-stories
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
