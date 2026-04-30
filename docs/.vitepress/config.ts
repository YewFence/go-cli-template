import { defineConfig } from 'vitepress'

export default defineConfig({
  base: '/{{REPO_NAME}}/',
  lang: 'zh-CN',
  title: '{{PROJECT_NAME}}',
  description: '{{PROJECT_DESCRIPTION}}',

  themeConfig: {
    nav: [
      { text: 'Shell 补全', link: '/guide/completion' },
      { text: 'GitHub', link: 'https://github.com/{{GITHUB_OWNER}}/{{REPO_NAME}}' }
    ],

    sidebar: [
      {
        text: '指南',
        items: [
          { text: 'Shell 补全', link: '/guide/completion' }
        ]
      }
    ],

    search: {
      provider: 'local'
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/{{GITHUB_OWNER}}/{{REPO_NAME}}' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © {{GITHUB_OWNER}}'
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
