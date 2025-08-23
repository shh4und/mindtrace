import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/main.css'
import VueApexCharts from 'vue3-apexcharts'
import { createPinia } from 'pinia'

import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
const pinia = createPinia()

const app = createApp(App)

app.use(pinia)
app.use(router)
app.use(VueApexCharts)

const options = {
    position: 'top-right',
    timeout: 5000,
    closeOnClick: true,
    pauseOnFocusLoss: true,
    pauseOnHover: true,
    draggable: true,
    draggablePercent: 0.6,
    showCloseButtonOnHover: false,
    hideProgressBar: false,
    closeButton: 'button',
    icon: true,
    rtl: false
};

app.use(Toast, options);

app.mount('#app')
