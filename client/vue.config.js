const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api/v1': {
        target: "http://localhost:8000",
        changeOrigin: true
      }
    }
  }
})
