import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/main.css'
import VueApexCharts from 'vue3-apexcharts'
import { createPinia } from 'pinia'

import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
const pinia = createPinia()

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core';

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

/* import specific icons */
import { faHouse, faUser, faSearch } from '@fortawesome/free-solid-svg-icons';

/* add icons to the library */
library.add(faHouse, faUser, faSearch);

const app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon);

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
