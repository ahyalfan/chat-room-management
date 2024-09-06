import './assets/main.css';
import 'primeicons/primeicons.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';
import primevue from 'primevue/config';
import Aura from '@primevue/themes/aura';

import App from './App.vue';
import router from './router';
import FocusTrap from 'primevue/focustrap';

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(primevue, {
  theme: {
    preset: Aura,
  },
  unstyled: false, // jika ingin style sendiri bisa kasih nilai true
});
app.directive('focustrap', FocusTrap);

app.mount('#app');
