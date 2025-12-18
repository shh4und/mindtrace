<template>
  <DashboardLayout
    :user-type="TipoUsuario.Paciente"
    variant="paciente"
    :menu-items="menuItems"
    :active-view="activeView"
    @edit-profile="navigateTo('paciente-editar-perfil')"
    @logout="handleLogout"
    @navigate="handleNavigation"
  >
    <router-view v-slot="{ Component, route: childRoute }">
      <Suspense>
        <template #default>
          <transition name="fade">
            <component v-if="Component" :is="Component" :key="childRoute.fullPath" />
          </transition>
        </template>
        <template #fallback>
          <div class="flex items-center justify-center h-64">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-emerald-600"></div>
          </div>
        </template>
      </Suspense>
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
  faHome,
  faFaceSmileBeam,
  faChartLine,
  faLink,
  faUserPen,
  faClipboardList,
  faUserDoctor
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const route = useRoute();
const userStore = useUserStore();

// Itens do menu para paciente
const menuItems = [
  { name: 'resumo', view: 'resumo', label: 'Resumo', icon: faHome },
  { name: 'humor', view: 'humor', label: 'Registro de Humor', icon: faFaceSmileBeam },
  { name: 'relatorios', view: 'relatorios', label: 'Relatórios', icon: faChartLine },
  { name: 'questionarios', view: 'questionarios', label: 'Questionários', icon: faClipboardList },
  { name: 'profissionais', view: 'profissionais', label: 'Meus Profissionais', icon: faUserDoctor },
  { name: 'vincular', view: 'vincular', label: 'Vincular Profissional', icon: faLink },
  { name: 'editar', view: 'editar-perfil', label: 'Editar Perfil', icon: faUserPen }
];

// Mapeamento de view para nome de rota
const viewToRoute = {
  'resumo': 'paciente-resumo',
  'humor': 'paciente-humor',
  'relatorios': 'paciente-relatorios',
  'questionarios': 'paciente-questionarios',
  'profissionais': 'paciente-profissionais',
  'vincular': 'paciente-vincular',
  'editar-perfil': 'paciente-editar-perfil'
};

// Mapeamento reverso para determinar view ativa baseado na rota atual
const routeToView = {
  'paciente-resumo': 'resumo',
  'paciente-humor': 'humor',
  'paciente-relatorios': 'relatorios',
  'paciente-questionarios': 'questionarios',
  'paciente-responder-questionario': 'questionarios',
  'paciente-profissionais': 'profissionais',
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