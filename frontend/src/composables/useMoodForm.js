import { ref, computed } from 'vue';
import { useToast } from 'vue-toastification';
import api from '@/services/api';
import {
  MOOD_OPTIONS,
  MOOD_LEVEL_MAPPING,
  SLIDER_CONFIG,
  SELF_CARE_ACTIVITIES,
  generateSteps,
  formatSleepLabel,
  clampValue
} from '@/constants/mood.js';

/**
 * Composable para gerenciamento do formulário de registro de humor
 * Extrai lógica complexa do RegistroHumor.vue
 */
export function useMoodForm() {
  const toast = useToast();

  // Estado do formulário
  const selectedMood = ref('');
  const sleepHours = ref(SLIDER_CONFIG.sleep.default);
  const energyLevel = ref(SLIDER_CONFIG.energy.default);
  const stressLevel = ref(SLIDER_CONFIG.stress.default);
  const selectedActivities = ref([]);
  const otherActivity = ref('');
  const notes = ref('');
  const isSubmitting = ref(false);

  // Configurações expostas
  const moodOptions = MOOD_OPTIONS;
  const selfCareActivities = SELF_CARE_ACTIVITIES;
  const sleepConfig = SLIDER_CONFIG.sleep;
  const energyConfig = SLIDER_CONFIG.energy;
  const stressConfig = SLIDER_CONFIG.stress;

  // Steps para os sliders
  const sleepSteps = generateSteps(sleepConfig.max, sleepConfig.min);
  const energySteps = generateSteps(energyConfig.max - 1, 0); // 0-indexed para icons
  const stressSteps = generateSteps(stressConfig.max - 1, 0);

  // Computed
  const sleepLabel = computed(() => formatSleepLabel(sleepHours.value, sleepConfig.max));
  
  const today = new Date();
  const currentDate = computed(() => 
    today.toLocaleDateString('pt-BR', { year: 'numeric', month: 'long', day: 'numeric' })
  );
  const currentDay = computed(() => 
    today.toLocaleDateString('pt-BR', { weekday: 'long' })
  );

  const isValid = computed(() => !!selectedMood.value);

  /**
   * Processa atividades de autocuidado
   * @returns {string[]}
   */
  function getProcessedActivities() {
    let activities = [...selectedActivities.value];
    
    if (otherActivity.value.trim()) {
      activities.push(otherActivity.value.trim());
    }
    
    // Se "Nenhuma Atividade" foi selecionada, retorna apenas ela
    if (activities.includes('Nenhuma Atividade')) {
      return [];
    }
    
    return activities;
  }

  /**
   * Monta o objeto de submissão para o backend
   * @returns {Object}
   */
  function buildSubmission() {
    return {
      nivel_humor: MOOD_LEVEL_MAPPING[selectedMood.value],
      horas_sono: clampValue(sleepHours.value, sleepConfig.min, sleepConfig.max),
      nivel_energia: clampValue(energyLevel.value, energyConfig.min, energyConfig.max),
      nivel_stress: clampValue(stressLevel.value, stressConfig.min, stressConfig.max),
      auto_cuidado: getProcessedActivities(),
      observacoes: notes.value,
      data_hora_registro: new Date().toISOString()
    };
  }

  /**
   * Valida o formulário
   * @returns {{ valid: boolean, message?: string }}
   */
  function validate() {
    if (!selectedMood.value) {
      return {
        valid: false,
        message: 'Por favor, selecione seu humor principal antes de registrar.'
      };
    }
    return { valid: true };
  }

  /**
   * Submete o registro de humor
   * @returns {Promise<boolean>}
   */
  async function submit() {
    const validation = validate();
    
    if (!validation.valid) {
      toast.warning(validation.message);
      return false;
    }

    isSubmitting.value = true;

    try {
      const submission = buildSubmission();
      await api.registrarHumor(submission);
      toast.success('Seu humor foi registrado com sucesso!');
      reset();
      return true;
    } catch (error) {
      toast.error('Houve um erro ao registrar seu humor.');
      console.error('Erro ao registrar humor:', error);
      return false;
    } finally {
      isSubmitting.value = false;
    }
  }

  /**
   * Reseta o formulário para o estado inicial
   */
  function reset() {
    selectedMood.value = '';
    sleepHours.value = SLIDER_CONFIG.sleep.default;
    energyLevel.value = SLIDER_CONFIG.energy.default;
    stressLevel.value = SLIDER_CONFIG.stress.default;
    selectedActivities.value = [];
    otherActivity.value = '';
    notes.value = '';
  }

  return {
    // Estado
    selectedMood,
    sleepHours,
    energyLevel,
    stressLevel,
    selectedActivities,
    otherActivity,
    notes,
    isSubmitting,

    // Configurações
    moodOptions,
    selfCareActivities,
    sleepConfig,
    energyConfig,
    stressConfig,
    sleepSteps,
    energySteps,
    stressSteps,

    // Computed
    sleepLabel,
    currentDate,
    currentDay,
    isValid,

    // Métodos
    validate,
    submit,
    reset
  };
}
