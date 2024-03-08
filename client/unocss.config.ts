import { defineConfig, presetIcons } from 'unocss'
import presetWind from '@unocss/preset-wind'

export default defineConfig({
  cli: {
    entry: {
      patterns: [
        "../client/assets/**/*.{css,js,ts}",
        "../client/**/*.pug",
        // "../server/src/jade/**/*.pug.go",
      ],
      outFile: "../server/public/css/main.css"
    },
  },
  presets: [
    presetWind(),
    presetIcons(),
  ],
  extractors: [
    {
      name: 'pug',
      order: -1,
      extract(args) {
        return args.code.split(/[\\:]?[\s'"`;{}.]+/g)
      },
    }
  ]
})
