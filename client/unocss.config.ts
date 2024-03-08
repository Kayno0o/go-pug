import { defineConfig, presetIcons, presetMini } from 'unocss'
import presetWind from '@unocss/preset-wind'

export default defineConfig({
  content: {
    filesystem: [
      "./assets/**/*.{js,ts}",
      "./**/*.pug",
    ]
  },
  presets: [
    presetMini(),
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
