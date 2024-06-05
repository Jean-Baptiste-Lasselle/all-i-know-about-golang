/** @type {import('tailwindcss').Config} */
// const plugin = import('tailwindcss/plugin');
const plugin = require('tailwindcss/plugin');

export default {
  content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
  theme: {
    extend: {
      listStyleType: {
        revert: 'revert',
      },
    },
  },
	plugins: [
		require('daisyui'),
    plugin(({ addVariant }) => {
      // https://tailwindcss.com/docs/plugins
      addVariant('ul', '& ul');
    }),
	],
};

