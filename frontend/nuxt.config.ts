// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  app: {
    head: {
      title: 'Build With Angga - Crowdfunding App',
      meta: [
        { charset: 'UTF-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'This is a description of my Nuxt app' },
        { name: 'author', content: 'fauzan264' },
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        { 
          rel: 'stylesheet', 
          href: 'https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap' 
        },
      ],
      script: [],
    },
  },

  css: [
    '@/assets/css/tailwind.css', // Add your Tailwind CSS here
  ],

  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },

  pages: true,

  runtimeConfig: {
    public: {
      apiBaseURL: 'http://localhost:8080/api/v1',
    },
  },

  plugins: [
    '~/plugins/auth.js'
  ],

  compatibilityDate: '2025-01-20'
})