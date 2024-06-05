import { defineConfig } from 'astro/config';
import preact from '@astrojs/preact';
import tailwind from '@astrojs/tailwind';
import robotsTxt from 'astro-robots-txt';
import sitemap from '@astrojs/sitemap';

// https://astro.build/config
export default defineConfig({
	// Enable Preact to support Preact JSX components.
	integrations: [preact({
		compat: true
	  }),
	  tailwind({
		// Example: Allow writing nested CSS declarations
		// alongside Tailwind's syntax
		nesting: true, // that's useful for postcss
		// applyBaseStyles: false,
	  }),
	  sitemap(),
	  robotsTxt()
	],

	/**
	 * MARKDOWN SHIKI CONFIG
	 */
	markdown: {
		shikiConfig: {
		  // Choose from Shiki's built-in themes (or add your own)
		  // https://shiki.style/themes
		  theme: 'github-dark',
		  // Alternatively, provide multiple themes
		  // See note below for using dual light/dark themes
		  themes: {
			// light: 'github-light',
			// light: 'solarized-dark',
			light: 'solarized-dark',
			dark: 'tokyo-night'
			// dark: 'github-dark',
		  },
		  // Add custom languages
		  // Note: Shiki has countless langs built-in, including .astro!
		  // https://shiki.style/languages
		  langs: [
			'css',
			'typescript',
			'javascript',
			'tsx',
			'terraform',
			'tex',
			'latex',
			'go',
			'shellscript',
			'yaml',
			'toml',
			'xml',
			'xsl',
			'wasm',
			'sql',
			'regexp',
			'markdown',
			'mermaid',
			'nginx',
			'ini',
			'json',
			'html'
			
		  ],
		  // Enable word wrap to prevent horizontal scrolling
		  wrap: true,
		  // Add custom transformers: https://shiki.style/guide/transformers
		  // Find common transformers: https://shiki.style/packages/transformers
		  transformers: [],
		},
	  },
});
