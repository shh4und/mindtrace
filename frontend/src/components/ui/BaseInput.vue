<template>
  <div class="w-full">
    <!-- Label -->
    <label 
      v-if="label" 
      :for="inputId" 
      :class="labelClasses"
    >
      {{ label }}
      <span v-if="required" class="text-red-500 ml-1" aria-hidden="true">*</span>
    </label>
    
    <!-- Input container -->
    <div class="relative">
      <!-- Prefix icon -->
      <div 
        v-if="prefixIcon" 
        class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none"
      >
        <font-awesome-icon :icon="prefixIcon" class="w-5 h-5" aria-hidden="true" />
      </div>
      
      <!-- Input field -->
      <input
        :id="inputId"
        ref="inputRef"
        :type="computedType"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readonly"
        :required="required"
        :autocomplete="autocomplete"
        :aria-invalid="hasError"
        :aria-describedby="errorId"
        :class="inputClasses"
        v-bind="$attrs"
        @input="handleInput"
        @blur="$emit('blur', $event)"
        @focus="$emit('focus', $event)"
      />
      
      <!-- Password toggle -->
      <button
        v-if="type === 'password'"
        type="button"
        class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
        :aria-label="showPassword ? 'Ocultar senha' : 'Mostrar senha'"
        @click="togglePasswordVisibility"
      >
        <font-awesome-icon 
          :icon="showPassword ? faEyeSlash : faEye" 
          class="w-5 h-5"
          aria-hidden="true"
        />
      </button>
      
      <!-- Suffix icon (not for password) -->
      <div 
        v-if="suffixIcon && type !== 'password'" 
        class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none"
      >
        <font-awesome-icon :icon="suffixIcon" class="w-5 h-5" aria-hidden="true" />
      </div>
    </div>
    
    <!-- Error message -->
    <p 
      v-if="error" 
      :id="errorId" 
      class="mt-1 text-sm text-red-600"
      role="alert"
    >
      {{ error }}
    </p>
    
    <!-- Hint text -->
    <p 
      v-else-if="hint" 
      class="mt-1 text-sm text-gray-500"
    >
      {{ hint }}
    </p>
  </div>
</template>

<script setup>
import { ref, computed, useId } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons';

/**
 * BaseInput - Componente de input reutilizável
 * Suporta validação, ícones, toggle de senha e acessibilidade
 */
const props = defineProps({
  /**
   * Valor do input (v-model)
   */
  modelValue: {
    type: [String, Number],
    default: ''
  },
  /**
   * Label do campo
   */
  label: {
    type: String,
    default: null
  },
  /**
   * Tipo do input
   */
  type: {
    type: String,
    default: 'text'
  },
  /**
   * Placeholder
   */
  placeholder: {
    type: String,
    default: ''
  },
  /**
   * Campo obrigatório
   */
  required: {
    type: Boolean,
    default: false
  },
  /**
   * Campo desabilitado
   */
  disabled: {
    type: Boolean,
    default: false
  },
  /**
   * Campo somente leitura
   */
  readonly: {
    type: Boolean,
    default: false
  },
  /**
   * Mensagem de erro
   */
  error: {
    type: String,
    default: null
  },
  /**
   * Texto de ajuda
   */
  hint: {
    type: String,
    default: null
  },
  /**
   * Ícone no início
   */
  prefixIcon: {
    type: Object,
    default: null
  },
  /**
   * Ícone no final
   */
  suffixIcon: {
    type: Object,
    default: null
  },
  /**
   * Autocomplete
   */
  autocomplete: {
    type: String,
    default: 'off'
  },
  /**
   * Tamanho do input
   */
  size: {
    type: String,
    default: 'md',
    validator: (v) => ['sm', 'md', 'lg'].includes(v)
  }
});

const emit = defineEmits(['update:modelValue', 'blur', 'focus']);

// Gera IDs únicos para acessibilidade
const uniqueId = useId();
const inputId = computed(() => `input-${uniqueId}`);
const errorId = computed(() => `error-${uniqueId}`);

// Ref para o elemento input
const inputRef = ref(null);

// Estado para toggle de senha
const showPassword = ref(false);

// Computed
const hasError = computed(() => !!props.error);

const computedType = computed(() => {
  if (props.type === 'password') {
    return showPassword.value ? 'text' : 'password';
  }
  return props.type;
});

const labelClasses = computed(() => [
  'block font-medium text-gray-700 mb-2',
  props.size === 'sm' ? 'text-sm' : props.size === 'lg' ? 'text-lg' : 'text-base'
]);

const sizeClasses = {
  sm: 'py-2 px-3 text-sm',
  md: 'py-3 px-4 text-base',
  lg: 'py-4 px-5 text-lg'
};

const inputClasses = computed(() => [
  'w-full rounded-lg border outline-none transition-colors',
  sizeClasses[props.size],
  // Padding para ícones
  props.prefixIcon ? 'pl-12' : '',
  props.suffixIcon || props.type === 'password' ? 'pr-12' : '',
  // Estados
  props.disabled ? 'bg-gray-100 cursor-not-allowed' : 'bg-white',
  props.readonly ? 'bg-gray-50' : '',
  // Borda e foco
  hasError.value
    ? 'border-red-300 focus:ring-2 focus:ring-red-500 focus:border-red-500'
    : 'border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500',
  'text-gray-900 placeholder-gray-500'
]);

// Methods
function handleInput(event) {
  emit('update:modelValue', event.target.value);
}

function togglePasswordVisibility() {
  showPassword.value = !showPassword.value;
}

// Expõe métodos para componente pai
defineExpose({
  focus: () => inputRef.value?.focus(),
  blur: () => inputRef.value?.blur()
});
</script>
