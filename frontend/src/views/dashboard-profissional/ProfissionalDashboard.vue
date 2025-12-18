<template>
  <DashboardLayout
    :user-type="TipoUsuario.Profissional"
    variant="profissional"
    :menu-items="menuItems"
    :active-view="activeView"
    @edit-profile="navigateTo('profissional-editar-perfil')"
    @logout="handleLogout"
    @navigate="handleNavigation"
  >
    <router-view v-slot="{ Component, route: childRoute }">
      <transition name="fade" mode="out-in">
        <component :is="Component" :key="childRoute.fullPath" />
      </transition>
    </router-view>
  </DashboardLayout>
</template>

<script setup>
import { computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useUserStore } from '@/store/user';
import { TipoUsuario } from '@/types/usuario.js';
import DashboardLayout from '@/components/layout/DashboardLayout.vue';
import { 
  faUsers,
  faEnvelope,
  faUserPen,
  faClipboardList
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const route = useRoute();
const userStore = useUserStore();

// Itens do menu para profissional
const menuItems = [
  { name: 'pacientes', view: 'pacientes', label: 'Meus Pacientes', icon: faUsers },
  { name: 'questionarios', view: 'questionarios-atribuidos', label: 'Questionários', icon: faClipboardList },
  { name: 'convite', view: 'convite', label: 'Gerar Convite', icon: faEnvelope },
  { name: 'editar', view: 'editar-perfil', label: 'Editar Perfil', icon: faUserPen }
];

// Mapeamento de view para nome de rota
const viewToRoute = {
  'pacientes': 'profissional-pacientes',
  'questionarios-atribuidos': 'profissional-questionarios-atribuidos',
  'convite': 'profissional-convite',
  'editar-perfil': 'profissional-editar-perfil'
};

// Mapeamento reverso para determinar view ativa baseado na rota atual
const routeToView = {
  'profissional-pacientes': 'pacientes',
  'profissional-paciente-relatorio': 'pacientes', // Relatório faz parte da seção pacientes
  'profissional-atribuir-questionario': 'pacientes', // Atribuir faz parte da seção pacientes
  'profissional-questionarios-atribuidos': 'questionarios-atribuidos',
  'profissional-convite': 'convite',
  'profissional-editar-perfil': 'editar-perfil'
};

// View ativa baseada na rota atual
const activeView = computed(() => routeToView[route.name] || 'pacientes');

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
    await userStore.fetchUser(TipoUsuario.Profissional);
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