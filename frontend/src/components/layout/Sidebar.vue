<template>
  <!-- Sidebar principal - Componente unificado para Paciente e Profissional -->
  <aside 
    :class="[
      'fixed md:static inset-y-0 left-0 transform transition-transform duration-300 ease-in-out md:translate-x-0',
      isOpen ? 'translate-x-0' : '-translate-x-full',
      isCollapsed ? 'md:w-20' : 'md:w-64',
      'w-64 bg-white border-r border-gray-200 flex flex-col shadow-lg md:shadow-none',
      'z-30'
    ]"
    role="navigation"
    :aria-label="`Menu de navegação do ${variantLabel}`"
  >
    <!-- Botao desktop para recolher -->
    <div class="hidden md:flex justify-end p-2 border-b border-gray-200">
      <button 
        @click="toggleCollapse"
        class="p-2 rounded-lg hover:bg-gray-100 transition-colors text-gray-500 hover:text-gray-700"
        :aria-label="isCollapsed ? 'Expandir menu' : 'Recolher menu'"
        :aria-expanded="!isCollapsed"
      >
        <font-awesome-icon :icon="isCollapsed ? faChevronRight : faChevronLeft" class="w-4 h-4" />
      </button>
    </div>

    <!-- Lista de navegacao principal -->
    <nav class="flex-1 px-3 pb-3 overflow-y-auto pt-16 md:pt-3">
      <ul class="space-y-1" role="menubar">
        <li v-for="item in menuItems" :key="item.name" role="none">
          <button
            @click="handleNavigate(item.view)"
            role="menuitem"
            :aria-current="activeView === item.view ? 'page' : undefined"
            :class="[
              'w-full flex items-center px-3 py-3 rounded-lg font-medium transition-all duration-200',
              activeView === item.view 
                ? activeItemClasses 
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
              aria-hidden="true"
            />
            <transition name="fade">
              <span v-if="!isCollapsed" class="truncate">{{ item.label }}</span>
            </transition>
            <font-awesome-icon 
              v-if="!isCollapsed && activeView === item.view"
              :icon="faCheck"
              :class="['w-4 h-4 ml-auto', checkIconClass]"
              aria-hidden="true"
            />
          </button>
        </li>
      </ul>
    </nav>

    <!-- Informacoes no rodape -->
    <div 
      v-if="!isCollapsed"
      class="p-4 border-t border-gray-200 bg-gray-50"
    >
      <div class="text-xs text-gray-500 text-center">
        <p class="font-medium">🧠 MindTrace</p>
        <p>{{ variantLabel }}</p>
      </div>
    </div>
  </aside>

  <!-- Sobreposicao mobile -->
  <Teleport to="body">
    <Transition name="overlay">
      <div 
        v-if="isOpen"
        @click="closeSidebar"
        class="fixed inset-0 bg-black/50 z-20 md:hidden"
        aria-hidden="true"
      ></div>
    </Transition>
  </Teleport>

  <!-- Botao mobile para abrir -->
  <button
    @click="toggleSidebar"
    :class="[
      'fixed bottom-4 left-4 z-40 md:hidden text-white p-4 rounded-full shadow-lg transition-colors',
      mobileButtonClasses
    ]"
    :aria-label="isOpen ? 'Fechar menu' : 'Abrir menu'"
    :aria-expanded="isOpen"
  >
    <font-awesome-icon :icon="isOpen ? faTimes : faBars" class="w-5 h-5" />
  </button>
</template>

<script setup>
import { ref, computed } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faChevronLeft,
  faChevronRight,
  faCheck,
  faBars,
  faTimes
} from '@fortawesome/free-solid-svg-icons';

/**
 * Sidebar unificada para Paciente e Profissional
 * Recebe menuItems como prop para customização por tipo de usuário
 */

const props = defineProps({
  /**
   * Items do menu - array de objetos com name, view, label, icon
   */
  menuItems: {
    type: Array,
    required: true,
    validator: (items) => items.every(item => 
      item.name && item.view && item.label && item.icon
    )
  },
  /**
   * View atualmente ativa
   */
  activeView: {
    type: String,
    default: ''
  },
  /**
   * Variante visual: 'paciente' (emerald) ou 'profissional' (rose)
   */
  variant: {
    type: String,
    default: 'paciente',
    validator: (v) => ['paciente', 'profissional'].includes(v)
  }
});

const emit = defineEmits(['navigate']);

// Estado do drawer lateral
const isOpen = ref(false);
const isCollapsed = ref(false);

// Classes dinâmicas baseadas na variante
const activeItemClasses = computed(() => {
  return props.variant === 'paciente'
    ? 'bg-emerald-50 text-emerald-700 shadow-sm'
    : 'bg-rose-50 text-rose-700 shadow-sm';
});

const checkIconClass = computed(() => {
  return props.variant === 'paciente' ? 'text-emerald-600' : 'text-rose-600';
});

const mobileButtonClasses = computed(() => {
  return props.variant === 'paciente'
    ? 'bg-emerald-600 hover:bg-emerald-700'
    : 'bg-rose-500 hover:bg-rose-600';
});

const variantLabel = computed(() => {
  return props.variant === 'paciente' ? 'Paciente' : 'Profissional';
});

// Alterna visibilidade da sidebar em telas pequenas
const toggleSidebar = () => {
  isOpen.value = !isOpen.value;
};

// Fecha a sidebar apos interacao
const closeSidebar = () => {
  isOpen.value = false;
};

// Controla modo compacto da sidebar em telas largas
const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value;
};

// Faz navegacao para a rota selecionada
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

.overlay-enter-active,
.overlay-leave-active {
  transition: opacity 0.3s ease;
}

.overlay-enter-from,
.overlay-leave-to {
  opacity: 0;
}
</style>
