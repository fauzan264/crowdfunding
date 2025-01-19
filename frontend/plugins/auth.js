import { Auth } from "@auth/core";

export default defineNuxtPlugin(nuxtApp => {
    const config = useRuntimeConfig()
    const storage = process.client ? localStorage : null;


    const auth = Auth({
        baseUrl: config.public.apiBaseURL,
        tokenType: 'Bearer',
        storage,

        login: {
            url: '/sessions',
            method: 'POST',
            body: (credentials) => {
                return {
                    username: credentials.username,
                    password: credentials.password,
                }
            },
        },

        // logout: {
        //     url: '/logout',
        //     method: 'POST',
        // }

        user: {
            url: '/users/fetch',
            method: 'GET',
            property: 'data'
        },

        token: {
            property: 'data.token',
        }
    })

    nuxtApp.provide('auth', auth)
})