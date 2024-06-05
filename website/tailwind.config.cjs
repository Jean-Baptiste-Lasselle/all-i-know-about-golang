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
  daisyui: {
    themes: [
      "light",
      "dark",
      "cupcake",
      "bumblebee",
      "emerald",
      "corporate",
      "synthwave",
      "retro",
      "cyberpunk",
      "valentine",
      "halloween",
      "garden",
      "forest",
      "aqua",
      "lofi",
      "pastel",
      "fantasy",
      "wireframe",
      "black",
      "luxury",
      "dracula",
      "cmyk",
      "autumn",
      "business",
      "acid",
      "lemonade",
      "night",
      "coffee",
      "winter",
      "dim",
      "nord",
      "sunset",
    ],
  },
};

