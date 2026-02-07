<template>
  <div class="h-screen bg-gray-50 font-sans antialiased flex flex-col">
    <!-- Navbar superior -->
    <TopNavbar
      :user-type="userType"
      @edit-profile="$emit('edit-profile')"
      @logout="$emit('logout')"
      @toggle-sidebar="toggleSidebar"
    />

    <!-- Conteudo principal com sidebar -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Barra lateral unificada -->
      <Sidebar
        :menu-items="menuItems"
        :active-view="activeView"
        :variant="variant"
        :is-open="isSidebarOpen"
        @navigate="(view) => $emit('navigate', view)"
        @close="closeSidebar"
      />

      <!-- Area de conteudo principal -->
      <main id="main-content" class="flex-1 overflow-y-auto p-4 sm:p-6 md:p-8">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup>
/**
 * DashboardLayout - Componente wrapper para dashboards
 * Encapsula TopNavbar + Sidebar com slots para conteúdo
 * Usado por PacienteDashboard e ProfissionalDashboard
 */
import { ref } from 'vue';
import TopNavbar from './TopNavbar.vue';
import Sidebar from './Sidebar.vue';

defineProps({
  /**
   * Tipo de usuário para navbar
   */
  userType: {
    type: String,
    required: true,
  },
  /**
   * Variante visual: 'paciente' ou 'profissional'
   */
  variant: {
    type: String,
    required: true,
    validator: (v) => ['paciente', 'profissional'].includes(v),
  },
  /**
   * Items do menu para sidebar
   */
  menuItems: {
    type: Array,
    required: true,
  },
  /**
   * View atualmente ativa
   */
  activeView: {
    type: String,
    default: '',
  },
});

defineEmits(['edit-profile', 'logout', 'navigate']);

const isSidebarOpen = ref(false);

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value;
};

const closeSidebar = () => {
  isSidebarOpen.value = false;
};
</script>
