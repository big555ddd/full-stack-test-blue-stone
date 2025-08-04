// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  runtimeConfig: {
    public: {
      WEB_API:
        process.env.NUXT_PUBLIC_WEB_API || "http://localhost:8080",
    },
  },
  modules: ["@pinia/nuxt", "@nuxtjs/tailwindcss"],
});
