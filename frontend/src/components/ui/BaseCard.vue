<template>
  <div :class="cardClasses">
    <!-- Header -->
    <div v-if="$slots.header || title" :class="headerClasses">
      <slot name="header">
        <h2 v-if="title" :class="titleClasses">{{ title }}</h2>
        <p v-if="subtitle" class="text-sm text-gray-500 mt-1">{{ subtitle }}</p>
      </slot>
    </div>
    
    <!-- Body -->
    <div :class="bodyClasses">
      <slot></slot>
    </div>
    
    <!-- Footer -->
    <div v-if="$slots.footer" :class="footerClasses">
      <slot name="footer"></slot>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';

/**
 * BaseCard - Componente de card/container reutilizável
 * Suporta header, body, footer, variantes de cor e padding
 */
const props = defineProps({
  /**
   * Título do card
   */
  title: {
    type: String,
    default: null
  },
  /**
   * Subtítulo do card
   */
  subtitle: {
    type: String,
    default: null
  },
  /**
   * Variante de cor de fundo
   */
  variant: {
    type: String,
    default: 'white',
    validator: (v) => ['white', 'gray', 'rose', 'emerald', 'amber', 'yellow', 'green'].includes(v)
  },
  /**
   * Tamanho do padding
   */
  padding: {
    type: String,
    default: 'md',
    validator: (v) => ['none', 'sm', 'md', 'lg'].includes(v)
  },
  /**
   * Card com borda
   */
  bordered: {
    type: Boolean,
    default: true
  },
  /**
   * Card com sombra
   */
  shadow: {
    type: Boolean,
    default: true
  },
  /**
   * Cantos arredondados
   */
  rounded: {
    type: String,
    default: 'lg',
    validator: (v) => ['none', 'sm', 'md', 'lg', 'xl', 'full'].includes(v)
  }
});

// Mapa de variantes de cor
const variantClasses = {
  white: 'bg-white',
  gray: 'bg-gray-50',
  rose: 'bg-rose-50',
  emerald: 'bg-emerald-50',
  amber: 'bg-amber-50',
  yellow: 'bg-yellow-50',
  green: 'bg-green-200'
};

// Mapa de padding
const paddingClasses = {
  none: 'p-0',
  sm: 'p-4',
  md: 'p-6',
  lg: 'p-8'
};

// Mapa de rounded
const roundedClasses = {
  none: 'rounded-none',
  sm: 'rounded-sm',
  md: 'rounded-md',
  lg: 'rounded-lg',
  xl: 'rounded-xl',
  full: 'rounded-full'
};

const cardClasses = computed(() => [
  variantClasses[props.variant],
  roundedClasses[props.rounded],
  props.bordered ? 'border border-gray-200' : '',
  props.shadow ? 'shadow-sm' : ''
]);

const headerClasses = computed(() => [
  paddingClasses[props.padding],
  'border-b border-gray-200'
]);

const bodyClasses = computed(() => [
  paddingClasses[props.padding]
]);

const footerClasses = computed(() => [
  paddingClasses[props.padding],
  'border-t border-gray-200 bg-gray-50',
  props.rounded !== 'none' ? 'rounded-b-lg' : ''
]);

const titleClasses = computed(() => 'text-xl font-semibold text-gray-900');
</script>
