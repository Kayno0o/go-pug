import { defineConfig, presetIcons } from 'unocss'
import presetWind from '@unocss/preset-wind'

export default defineConfig({
  cli: {
    entry: {
      patterns: [
        "./client/**/*.{css,js,html}",
        "./app/**/*.pug.go",
      ],
      outFile: "./public/style/main.css"
    },
  },
  presets: [
    presetWind(),
    presetIcons(),
  ],
})