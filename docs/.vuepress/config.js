module.exports = {
    title: 'go cmb 招商银行接口服务 文档',
    // theme: "@vuepress/theme-blog",
    base: "/go-cmb-service/",
    plugins: [
        [
            '@vuepress/google-analytics',
            {
                'ga': 'G-RXDN42SD0T' // UA-00000000-0
            }
        ],
        '@vuepress/last-updated',
        '@vuepress/back-to-top',
        '@vuepress/nprogress',
        '@vuepress/medium-zoom',
        ['@vuepress/search', {
            searchMaxSuggestions: 10
        }]
        // '@vuepress/pwa'
    ],
    themeConfig: {

        repo:'https://github.com/ahKevinXy/go-cmb-service',
        search: true,
        smoothScroll: true,

        sidebar:'auto',
        lastUpdated:"最新更新时间",
        dateFormat: 'YYYY-MM-DD',

        /**
         * Ref: https://vuepress-theme-blog.ulivz.com/#footer
         */
        footer: {
            contact: [
                {
                    type: 'github',
                    link: 'https://github.com/jenkey2011',
                }
            ],
            copyright: [{
                text: 'opencodes',
                link: '/',
            }
            ],
        },
        directories:[
            {
                id: 'post',
                dirname: '_posts',
                path: '/',
                itemPermalink: '/:year/:month/:day/:slug',
            },
            {
                id: 'writing',
                dirname: '_writings',
                path: '/',
                itemPermalink: '/:year/:month/:day/:slug',
            },
        ],
        /**
         * Ref: https://vuepress-theme-blog.ulivz.com/#sitemap
         */
        sitemap: {
            hostname: 'https://ahkevinxy.github.io'
        },
        // smoothScroll: true
    },
}