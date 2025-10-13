import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/main.css'
import VueApexCharts from 'vue3-apexcharts'
import { createPinia } from 'pinia'

import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
const pinia = createPinia()

/* importa o nucleo do fontawesome */
import { library } from '@fortawesome/fontawesome-svg-core';

/* importa o componente de icone do fontawesome */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

/* importa os icones utilizados na aplicacao */
import { faHouse, faUser, faSearch } from '@fortawesome/free-solid-svg-icons';

/* registra os icones na biblioteca compartilhada */
library.add(faHouse, faUser, faSearch);

// cria a aplicacao principal com dependencias globais
const app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon);

app.use(pinia)
app.use(router)
// habilita componentes apexcharts como plugin global
app.use(VueApexCharts)

// configuracoes padrao para notificacoes toast
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
// inicia o aplicativo montando na raiz html
app.mount('#app')
