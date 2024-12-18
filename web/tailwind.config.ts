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
      fontSize: {
        '6xs': '0.125em',
        '5xs': '0.1875em',
        '4xs': '0.25em',
        '3xs': '0.375em',
        '2xs': '0.5em',
        xs: '0.75em' /* 12px */,
        sm: '0.875em' /* 14px */,
        md: '1em' /* 16px */,
        lg: '1.125em' /* 18px */,
        xl: '1.25em' /* 20px */,
        '2xl': '1.5em' /* 24px */,
        '3xl': '1.875em' /* 30px */,
        '4xl': '2.25em' /* 36px */,
        '5xl': '2.625em' /* 42px */,
        '6xl': '3em' /* 48px */,
      },
      spacing: {
        0: '0em',
        1: '0.25em',
        1.5: '0.375em',
        2: '0.5em',
        3: '0.75em',
        4: '1em',
        5: '1.25em',
        6: '1.5em',
        7: '1.75em',
        8: '2em',
        9: '2.25em',
        10: '2.5em',
        11: '2.75em',
        12: '3em',
        14: '3.5em',
        16: '4em',
        20: '5em',
        24: '6em',
        28: '7em',
        32: '8em',
        36: '9em',
        40: '10em',
        44: '11em',
        48: '12em',
        52: '13em',
        56: '14em',
        60: '15em',
        64: '16em',
        72: '18em',
        80: '20em',
        96: '24em',
      },
    }
  },

  plugins: [typography, forms, containerQueries, aspectRatio]
} as Config;
