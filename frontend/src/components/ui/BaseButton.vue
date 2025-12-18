<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    :class="buttonClasses"
    :aria-busy="loading"
    v-bind="$attrs"
  >
    <!-- Loading spinner -->
    <span v-if="loading" class="inline-flex items-center">
      <svg 
        class="animate-spin -ml-1 mr-2 h-4 w-4" 
        xmlns="http://www.w3.org/2000/svg" 
        fill="none" 
        viewBox="0 0 24 24"
        aria-hidden="true"
      >
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
      <span>{{ loadingText || 'Carregando...' }}</span>
    </span>
    
    <!-- Icon (left) -->
    <font-awesome-icon 
      v-if="icon && iconPosition === 'left' && !loading" 
      :icon="icon" 
      :class="['w-4 h-4', $slots.default ? 'mr-2' : '']"
      aria-hidden="true"
    />
    
    <!-- Slot content -->
    <span v-if="!loading">
      <slot></slot>
    </span>
    
    <!-- Icon (right) -->
    <font-awesome-icon 
      v-if="icon && iconPosition === 'right' && !loading" 
      :icon="icon" 
      :class="['w-4 h-4', $slots.default ? 'ml-2' : '']"
      aria-hidden="true"
    />
  </button>
</template>

<script setup>
import { computed } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

/**
 * BaseButton - Componente de botão reutilizável
 * Suporta variantes, tamanhos, estados de loading e ícones
 */
const props = defineProps({
  /**
   * Tipo do botão HTML
   */
  type: {
    type: String,
    default: 'button',
    validator: (v) => ['button', 'submit', 'reset'].includes(v)
  },
  /**
   * Variante visual do botão
   * - primary: emerald (padrão)
   * - emerald: tema paciente explícito
   * - rose: tema profissional
   * - secondary, danger, ghost, outline: utilitárias
   */
  variant: {
    type: String,
    default: 'primary',
    validator: (v) => ['primary', 'emerald', 'rose', 'secondary', 'danger', 'ghost', 'outline'].includes(v)
  },
  /**
   * Tamanho do botão
   */
  size: {
    type: String,
    default: 'md',
    validator: (v) => ['sm', 'md', 'lg'].includes(v)
  },
  /**
   * Botão desabilitado
   */
  disabled: {
    type: Boolean,
    default: false
  },
  /**
   * Estado de carregamento
   */
  loading: {
    type: Boolean,
    default: false
  },
  /**
   * Texto exibido durante loading
   */
  loadingText: {
    type: String,
    default: null
  },
  /**
   * Ícone FontAwesome
   */
  icon: {
    type: Object,
    default: null
  },
  /**
   * Posição do ícone
   */
  iconPosition: {
    type: String,
    default: 'left',
    validator: (v) => ['left', 'right'].includes(v)
  },
  /**
   * Botão de largura total
   */
  fullWidth: {
    type: Boolean,
    default: false
  }
});

// Classes base compartilhadas
const baseClasses = 'inline-flex items-center justify-center font-medium rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-offset-2 outline-none disabled:opacity-50 disabled:cursor-not-allowed';

// Mapa de variantes - Design System MindTrace
const variantClasses = {
  // Tema Paciente (Emerald)
  primary: 'bg-emerald-600 hover:bg-emerald-700 text-white focus:ring-emerald-500',
  emerald: 'bg-emerald-600 hover:bg-emerald-700 text-white focus:ring-emerald-500',
  
  // Tema Profissional (Rose)
  rose: 'bg-rose-500 hover:bg-rose-600 text-white focus:ring-rose-400',
  
  // Variantes utilitárias
  secondary: 'bg-gray-100 hover:bg-gray-200 text-gray-700 focus:ring-gray-500',
  danger: 'bg-red-600 hover:bg-red-700 text-white focus:ring-red-500',
  ghost: 'bg-transparent hover:bg-gray-100 text-gray-700 focus:ring-gray-500',
  outline: 'border-2 border-emerald-600 text-emerald-600 hover:bg-emerald-50 focus:ring-emerald-500'
};

// Mapa de tamanhos
const sizeClasses = {
  sm: 'py-2 px-3 text-sm',
  md: 'py-3 px-4 text-base',
  lg: 'py-4 px-6 text-lg'
};

const buttonClasses = computed(() => [
  baseClasses,
  variantClasses[props.variant],
  sizeClasses[props.size],
  props.fullWidth ? 'w-full' : ''
]);
</script>
