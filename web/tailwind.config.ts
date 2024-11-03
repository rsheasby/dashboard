import aspectRatio from '@tailwindcss/aspect-ratio';
import containerQueries from '@tailwindcss/container-queries';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';
import type { Config } from 'tailwindcss';

export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],

  theme: {
    extend: {
      colors: {
        'primary': 'var(--primary)',
        'rabbit': 'var(--rabbit)',
        'rabbit-dark': 'var(--rabbit-dark)',
      },
    }
  },

  plugins: [typography, forms, containerQueries, aspectRatio]
} as Config;
