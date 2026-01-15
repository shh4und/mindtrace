<template>
  <div
    :class="[
      'bg-white rounded-xl shadow-sm border border-gray-200 p-5 cursor-pointer transition-all duration-200',
      hoverClass,
      'hover:shadow-lg'
    ]"
    role="button"
    :aria-label="ariaLabel"
    tabindex="0"
    @click="$emit('click')"
    @keydown.enter="$emit('click')"
    @keydown.space.prevent="$emit('click')"
  >
    <!-- Header com Avatar e Info -->
    <div class="flex items-center space-x-4 mb-4">
      <div 
        :class="[avatarColor, 'w-14 h-14 rounded-full flex items-center justify-center shrink-0']"
      >
        <font-awesome-icon 
          :icon="avatarIcon" 
          class="w-7 h-7 text-white" 
          aria-hidden="true" 
        />
      </div>
      <div class="flex-1 min-w-0">
        <h3 class="font-semibold text-gray-900 text-lg truncate">{{ title }}</h3>
        <p class="text-sm text-gray-500 truncate">{{ subtitle }}</p>
      </div>
    </div>

    <!-- Ações -->
    <div class="flex flex-wrap gap-2">
      <button 
        v-for="(action, index) in actions"
        :key="index"
        @click.stop="$emit('action', action.id)"
        :class="[
          'flex items-center text-sm font-medium transition-colors px-3 py-1.5 rounded-lg',
          actionClass
        ]"
      >
        <font-awesome-icon 
          v-if="action.icon" 
          :icon="action.icon" 
          class="mr-2 w-4 h-4" 
          aria-hidden="true" 
        />
        {{ action.label }}
      </button>
    </div>
  </div>
</template>

<script setup>
/**
 * CardListaUsuario - Card reutilizável para listas de usuários
 * Usado em ListaPacientes (profissional) e ListaProfissionais (paciente)
 */
import { computed } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faUser, faUserDoctor } from '@fortawesome/free-solid-svg-icons';

const props = defineProps({
  /**
   * Título principal do card (nome do usuário)
   */
  title: {
    type: String,
    required: true
  },
  /**
   * Subtítulo (idade, especialidade, etc.)
   */
  subtitle: {
    type: String,
    default: ''
  },
  /**
   * Variante do card: 'paciente' ou 'profissional'
   */
  variant: {
    type: String,
    default: 'paciente',
    validator: (v) => ['paciente', 'profissional'].includes(v)
  },
  /**
   * Cor do avatar (classe Tailwind)
   */
  avatarColor: {
    type: String,
    default: 'bg-blue-500'
  },
  /**
   * Array de ações: [{ id: 'view', label: 'Ver', icon: faIcon }]
   */
  actions: {
    type: Array,
    default: () => []
  },
  /**
   * Label para acessibilidade
   */
  ariaLabel: {
    type: String,
    default: 'Ver detalhes'
  }
});

defineEmits(['click', 'action']);

// Ícone baseado na variante
const avatarIcon = computed(() => {
  return props.variant === 'profissional' ? faUserDoctor : faUser;
});

// Classes de hover baseadas na variante
const hoverClass = computed(() => {
  return props.variant === 'profissional' 
    ? 'hover:border-rose-400' 
    : 'hover:border-emerald-400';
});

// Classes de ação baseadas na variante
const actionClass = computed(() => {
  return props.variant === 'profissional'
    ? 'text-rose-600 hover:text-rose-800 hover:bg-rose-50'
    : 'text-emerald-600 hover:text-emerald-800 hover:bg-emerald-50';
});
</script>
