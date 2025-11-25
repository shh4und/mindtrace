<template>
  <div class="min-h-screen bg-gray-50 font-sans antialiased flex flex-col">
    <!-- Navbar superior -->
    <TopNavbar 
      :user-type="TipoUsuario.Paciente" 
      @edit-profile="navigateTo('paciente-editar-perfil')"
      @logout="handleLogout"
    />

    <!-- Conteudo principal com sidebar -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Barra lateral unificada -->
      <Sidebar 
        :menu-items="menuItems"
        :active-view="activeView"
        variant="paciente"
        @navigate="handleNavigation" 
      />

      <!-- Area de conteudo principal com router-view -->
      <main id="main-content" class="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </main>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useUserStore } from '@/store/user';
import { TipoUsuario } from '@/types/usuario.js';
import TopNavbar from '@/components/layout/TopNavbar.vue';
import Sidebar from '@/components/layout/Sidebar.vue';
import { 
  faHome,
  faFaceSmileBeam,
  faChartLine,
  faLink,
  faUserPen
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const route = useRoute();
const userStore = useUserStore();

// Itens do menu para paciente
const menuItems = [
  { name: 'resumo', view: 'resumo', label: 'Resumo', icon: faHome },
  { name: 'humor', view: 'humor', label: 'Registro de Humor', icon: faFaceSmileBeam },
  { name: 'relatorios', view: 'relatorios', label: 'Relatórios', icon: faChartLine },
  { name: 'vincular', view: 'vincular', label: 'Vincular Profissional', icon: faLink },
  { name: 'editar', view: 'editar-perfil', label: 'Editar Perfil', icon: faUserPen }
];

// Mapeamento de view para nome de rota
const viewToRoute = {
  'resumo': 'paciente-resumo',
  'humor': 'paciente-humor',
  'relatorios': 'paciente-relatorios',
  'vincular': 'paciente-vincular',
  'editar-perfil': 'paciente-editar-perfil'
};

// Mapeamento reverso para determinar view ativa baseado na rota atual
const routeToView = {
  'paciente-resumo': 'resumo',
  'paciente-humor': 'humor',
  'paciente-relatorios': 'relatorios',
  'paciente-vincular': 'vincular',
  'paciente-editar-perfil': 'editar-perfil'
};

// View ativa baseada na rota atual
const activeView = computed(() => routeToView[route.name] || 'resumo');

const handleNavigation = (view) => {
  const routeName = viewToRoute[view];
  if (routeName) {
    router.push({ name: routeName });
  }
};

const navigateTo = (routeName) => {
  router.push({ name: routeName });
};

const handleLogout = () => {
  userStore.logout();
};

onMounted(async () => {
  // Busca dados do usuario se ainda nao estiverem carregados
  if (!userStore.user) {
    await userStore.fetchUser(TipoUsuario.Paciente);
  }
});
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.15s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>