/**
 * Constantes e configurações para o registro de humor
 * Centraliza magic strings, mapeamentos e opções do formulário
 */

/**
 * Opções de humor com emoji e label
 */
export const MOOD_OPTIONS = [
  { value: 'very_bad', label: 'Muito Mal', emoji: '😖' },
  { value: 'bad', label: 'Aborrecido', emoji: '😕' },
  { value: 'neutral', label: 'Neutro', emoji: '😐' },
  { value: 'cheerful', label: 'Animado', emoji: '😊' },
  { value: 'very_good', label: 'Muito Bem', emoji: '😁' },
];

/**
 * Mapeamento de valor de humor para nível numérico (usado no backend)
 */
export const MOOD_LEVEL_MAPPING = {
  'very_bad': 1,
  'bad': 2,
  'neutral': 3,
  'cheerful': 4,
  'very_good': 5
};

/**
 * Mapeamento reverso: nível numérico para valor de humor
 */
export const MOOD_LEVEL_REVERSE = {
  1: 'very_bad',
  2: 'bad',
  3: 'neutral',
  4: 'cheerful',
  5: 'very_good'
};

/**
 * Configuração dos sliders
 */
export const SLIDER_CONFIG = {
  sleep: {
    min: 0,
    max: 12,
    default: 0,
    emoji: '🌜',
    labels: ['0h', '4h', '8h', '12h+']
  },
  energy: {
    min: 1,
    max: 10,
    default: 6,
    emoji: '🔋',
    labels: ['Baixa', 'Moderada', 'Alta']
  },
  stress: {
    min: 1,
    max: 10,
    default: 6,
    emoji: '😤',
    labels: ['Calmo', 'Moderado', 'Alto']
  }
};

/**
 * Atividades de autocuidado pré-definidas
 */
export const SELF_CARE_ACTIVITIES = [
  'Leitura',
  'Exercício',
  'Medicação',
  'Yoga',
  'Meditação',
  'Música',
  'Hobbies',
  'Nenhuma Atividade'
];

/**
 * Helper para gerar array de steps para sliders
 * @param {number} max - Valor máximo
 * @param {number} min - Valor mínimo (default 0)
 * @returns {number[]} - Array de valores
 */
export function generateSteps(max, min = 0) {
  return Array.from({ length: max - min + 1 }, (_, i) => i + min);
}

/**
 * Helper para formatar label de horas de sono
 * @param {number} hours - Horas de sono
 * @param {number} max - Valor máximo
 * @returns {string} - Label formatado
 */
export function formatSleepLabel(hours, max = SLIDER_CONFIG.sleep.max) {
  return hours >= max ? `${max}h+` : `${hours}h`;
}

/**
 * Valida e limita valor ao range do slider
 * @param {number} value - Valor a validar
 * @param {number} min - Valor mínimo
 * @param {number} max - Valor máximo
 * @returns {number} - Valor limitado ao range
 */
export function clampValue(value, min, max) {
  return Math.max(min, Math.min(max, value));
}
