import { createApp } from 'vue'
import App from './App.vue'
import { CreateWailsApp } from '../wailsjs/go/main/App'

const app = createApp(App)
const wails = CreateWailsApp()

app.config.globalProperties.$wails = wails

app.mount('#app')
