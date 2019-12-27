import { configure } from '@storybook/vue'

// automatically import all files ending in *.stories.js, *.stories.ts
configure(require.context('../stories', true, /\.stories\.(js|ts)$/), module)
