<template>
  <!-- Sidebar Container -->
  <aside 
    :class="[
      'fixed lg:static inset-y-0 left-0 z-30 transform transition-transform duration-300 ease-in-out lg:translate-x-0',
      isOpen ? 'translate-x-0' : '-translate-x-full',
      isCollapsed ? 'lg:w-20' : 'lg:w-64',
      'w-64 bg-white border-r border-gray-200 flex flex-col shadow-lg lg:shadow-none'
    ]"
  >
    <!-- Toggle Button (Desktop) -->
    <div class="hidden lg:flex justify-end p-2 border-b border-gray-200">
      <button 
        @click="toggleCollapse"
        class="p-2 rounded-lg hover:bg-gray-100 transition-colors text-gray-500 hover:text-gray-700"
      >
        <font-awesome-icon :icon="isCollapsed ? faChevronRight : faChevronLeft" class="w-4 h-4" />
      </button>
    </div>

    <!-- Navigation -->
    <nav class="flex-1 p-3 overflow-y-auto">
      <ul class="space-y-1">
        <li v-for="item in menuItems" :key="item.name">
          <button
            @click="handleNavigate(item.view)"
            :class="[
              'w-full flex items-center px-3 py-3 rounded-lg font-medium transition-all duration-200',
              activeView === item.view 
                ? 'bg-rose-50 text-rose-700 shadow-sm' 
                : 'text-gray-600 hover:bg-gray-50 hover:text-gray-900'
            ]"
            :title="isCollapsed ? item.label : ''"
          >
            <font-awesome-icon 
              :icon="item.icon" 
              :class="[
                'w-5 h-5 flex-shrink-0',
                isCollapsed ? 'mx-auto' : 'mr-3'
              ]"
            />
            <transition name="fade">
              <span v-if="!isCollapsed" class="truncate">{{ item.label }}</span>
            </transition>
            <font-awesome-icon 
              v-if="!isCollapsed && activeView === item.view"
              :icon="faCheck"
              class="w-4 h-4 ml-auto text-rose-600"
            />
          </button>
        </li>
      </ul>
    </nav>

    <!-- Footer Info (Optional) -->
    <div 
      v-if="!isCollapsed"
      class="p-4 border-t border-gray-200 bg-gray-50"
    >
      <div class="text-xs text-gray-500 text-center">
        <p class="font-medium">🧠 MindTrace</p>
        <p>Profissional</p>
      </div>
    </div>
  </aside>

  <!-- Mobile Overlay -->
  <div 
    v-if="isOpen"
    @click="closeSidebar"
    class="fixed inset-0 bg-black bg-opacity-50 z-20 lg:hidden"
  ></div>

  <!-- Mobile Toggle Button -->
  <button
    @click="toggleSidebar"
    class="fixed bottom-4 left-4 z-40 lg:hidden bg-rose-500 text-white p-4 rounded-full shadow-lg hover:bg-rose-600 transition-colors"
  >
    <font-awesome-icon :icon="isOpen ? faTimes : faBars" class="w-5 h-5" />
  </button>
</template>

<script setup>
import { ref, computed } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faUsers,
  faEnvelope,
  faUserPen,
  faChevronLeft,
  faChevronRight,
  faCheck,
  faBars,
  faTimes
} from '@fortawesome/free-solid-svg-icons';

// Props
const props = defineProps({
  activeView: {
    type: String,
    default: 'pacientes'
  }
});

// Emits
const emit = defineEmits(['navigate']);

// State
const isOpen = ref(false);
const isCollapsed = ref(false);

// Menu Items
const menuItems = [
  {
    name: 'pacientes',
    view: 'pacientes',
    label: 'Meus Pacientes',
    icon: faUsers
  },
  {
    name: 'convite',
    view: 'convite',
    label: 'Gerar Convite',
    icon: faEnvelope
  },
  {
    name: 'editar',
    view: 'editar-perfil',
    label: 'Editar Perfil',
    icon: faUserPen
  }
];

// Methods
const toggleSidebar = () => {
  isOpen.value = !isOpen.value;
};

const closeSidebar = () => {
  isOpen.value = false;
};

const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value;
};

const handleNavigate = (view) => {
  emit('navigate', view);
  closeSidebar();
};
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
