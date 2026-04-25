import { defineConfig } from 'vitepress'

export default defineConfig({
  base: '/your-cli/',
  lang: 'zh-CN',
  title: 'your-cli',
  description: 'Your CLI description',

  themeConfig: {
    nav: [
      { text: '指南', link: '/guide/getting-started' },
      { text: '开发', link: '/guide/development' },
      { text: 'GitHub', link: 'https://github.com/example/your-cli' }
    ],

    sidebar: [
      {
        text: '指南',
        items: [
          { text: '快速开始', link: '/guide/getting-started' },
          { text: '开发指南', link: '/guide/development' },
          { text: 'Shell 补全', link: '/guide/completion' },
          { text: '发布流程', link: '/guide/release' }
        ]
      }
    ],

    search: {
      provider: 'local'
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/example/your-cli' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © example'
    },

    docFooter: {
      prev: '上一页',
      next: '下一页'
    },

    outline: {
      label: '本页目录'
    },

    lastUpdated: {
      text: '最后更新'
    }
  }
})
