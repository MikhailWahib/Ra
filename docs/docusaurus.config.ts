import { themes as prismThemes } from 'prism-react-renderer';
import type { Config } from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const config: Config = {
  title: 'Ra Programming Language 𓋹',
  tagline: 'A simple interpreted language built in Go',
  favicon: 'img/favicon.ico',

  url: 'https://mikhailwahib.github.io', // Replace with your GitHub Pages URL
  baseUrl: '/Ra/', // Replace with your repository name

  organizationName: 'MikhailWahib', // GitHub org/user name
  projectName: 'Ra', // Repository name

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          editUrl: 'https://github.com/MikhailWahib/Ra/edit/main/',
        },
        blog: {
          showReadingTime: true,
          editUrl: 'https://github.com/MikhailWahib/Ra/edit/main/',
        },
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    colorMode: {
      defaultMode: 'light',
      respectPrefersColorScheme: true,
    },
    navbar: {
      title: 'Ra',
      logo: {
        alt: 'Ra Logo',
        src: 'img/ankh-logo.svg',
      },
      style: 'dark', // Use 'dark' or 'light' to match the theme
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'docsSidebar',
          position: 'left',
          label: 'Docs',
        },
        {
          href: 'https://github.com/MikhailWahib/Ra',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark', // Use 'dark' or 'light' to match the theme
      // links: [
      // ],
      copyright: `Copyright © ${new Date().getFullYear()} Ra. Built with Docusaurus.`,
  },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ['go'], // Add Go syntax highlighting for code blocks
    },
  } satisfies Preset.ThemeConfig,
};

export default config;