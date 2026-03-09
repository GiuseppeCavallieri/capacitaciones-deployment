export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  ssr: false,
  css: ["bulma/css/bulma.min.css"],
  runtimeConfig: {
    public: {
      backUrl: process.env.BACK_URL || "http://localhost:8080",
    },
  },
});
