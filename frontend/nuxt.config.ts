export default {
  /*
   ** Nuxt rendering mode
   ** See https://nuxtjs.org/api/configuration-mode
   */
  mode: 'universal',
  /*
   ** Headers of the page
   ** See https://nuxtjs.org/api/configuration-head
   */
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   ** https://nuxtjs.org/guide/plugins
   */
  plugins: [],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: ['@nuxt/typescript-build', '@nuxtjs/vuetify'],
  /*
   ** Nuxt.js modules ('@nuxtjs/pwa')
   */
  modules: ['@nuxtjs/apollo', '@nuxtjs/proxy'],
  /*
   ** vuetify module configuration
   ** https://github.com/nuxt-community/vuetify-module
   */
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      light: true,
    },
  },
  apollo: {
    clientConfigs: {
      default: {
        httpEndpoint: 'http://localhost:8080/query',
        httpLinkOptions: {
          fetchOptions: {
            mode: 'cors',
          },
          credentials: 'omit',
        },
      },
    },
    errorHandler: '~/plugins/apollo-error-handler.js',
  },
  /*
   ** Build configuration
   ** See https://nuxtjs.org/api/configuration-build/
   */
  build: {},
  typescript: {
    typeCheck: {
      eslint: true,
    },
  },
}
