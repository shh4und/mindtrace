<template>
  <div class="flex-1 p-4 md:p-8 w-full">
    <header class="mb-8 text-center md:text-left">
      <h1 class="text-2xl md:text-3xl font-bold text-gray-900">Registro de Humor Diário</h1>
      <p class="text-gray-600 mt-1">Preencha como você se sentiu hoje e registre suas observações.</p>
    </header>

    <form @submit.prevent="submitMood" class="space-y-8">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Registro para o dia:</h2>
        <div class="flex flex-col md:flex-row md:space-x-4 space-y-2 md:space-y-0">
          <input type="text" :value="currentDate"
            class="flex-1 text-lg font-medium text-gray-700 bg-gray-100 rounded-md p-2 text-center" readonly>
          <input type="text" :value="currentDay"
            class="flex-1 text-lg font-medium text-gray-700 bg-gray-100 rounded-md p-2 text-center" readonly>
        </div>
      </div>

      <!-- Mood selection (more generic, with emojis) -->
      <div class="bg-rose-50 rounded-lg shadow-sm border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-4 text-center">Como você se sentiu hoje?</h2>
        <div class="flex flex-wrap gap-6 justify-center">
          <button v-for="m in moodOptions" :key="m.value" type="button" @click="selectedMood = m.value"
            :aria-pressed="selectedMood === m.value"
            :class="['mood-btn flex flex-col items-center justify-center p-3 rounded-full transition', { 'selected': selectedMood === m.value }]">
            <span class="mood-emoji mb-2">{{ m.emoji }}</span>
            <span class="text-sm">{{ m.label }}</span>
          </button>
        </div>
      </div>

      <!-- Sleep slider with discrete moon icons (0..12) -->
      <div class="bg-rose-100 rounded-lg shadow-sm border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-6 text-center">Horas de Sono</h2>
        <div class="flex flex-col md:flex-row items-center md:space-x-4">
          <div class="w-full">
            <div class="mb-3">
              <div class="icon-row">
                <button v-for="step in sleepSteps" :key="step" type="button" class="icon-step "
                  :class="[{ 'active': step <= sleepHours }, '']" @click="sleepHours = step">
                  <span class="step-emoji hidden md:block" aria-hidden="true">🌜</span>
                </button>
              </div>
            </div>
            <input type="range" min="0" :max="sleepMax" step="1" v-model.number="sleepHours" class="w-full slider" />
            <div class="flex justify-between text-sm mt-2 text-gray-600">
              <span>0h</span>
              <span>4h</span>
              <span>8h</span>
              <span>12h+</span>
            </div>
          </div>
          <div class="w-full md:w-36 text-center display-area mt-6 md:mt-0">
            <div class="display-emoji">🌜</div>
            <div class="font-semibold display-value">{{ sleepLabel }}</div>
          </div>
        </div>
      </div>

      <!-- Energy slider with battery icons -->
      <div class="bg-amber-50 rounded-lg shadow-sm border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-6 text-center">Nível de Energia</h2>
        <div class="flex flex-col md:flex-row items-center md:space-x-4">
          <div class="w-full">
            <div class="mb-3">
              <div class="icon-row">
                <button v-for="step in energySteps" :key="step" type="button" class="icon-step"
                  :class="{ 'active': step < energyLevel }" @click="energyLevel = step+1">
                  <span class="step-emoji  hidden md:block" aria-hidden="true">🔋</span>
                </button>
              </div>
            </div>
            <input type="range" min="1" :max="energyMax" step="1" v-model.number="energyLevel" class="w-full slider" />
            <div class="flex justify-between text-sm mt-2 text-gray-600">
              <span>Baixa</span>
              <span>Moderada</span>
              <span>Alta</span>
            </div>
          </div>
          <div class="w-full md:w-36 text-center display-area mt-6 md:mt-0">
            <div class="display-emoji">🔋</div>
            <div class="font-semibold display-value">{{ energyLevel }} / {{ energyMax }}</div>
          </div>
        </div>
      </div>

      <!-- Stress slider with alert icons -->
      <div class="bg-yellow-50 rounded-lg shadow-sm border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-6 text-center">Nível de Stress</h2>
        <div class="flex flex-col md:flex-row items-center md:space-x-4">
          <div class="w-full">
            <div class="mb-3">
              <div class="icon-row">
                <button v-for="step in stressSteps" :key="step" type="button" class="icon-step"
                  :class="{ 'active': step < stressLevel }" @click="stressLevel = step+1">
                  <span class="step-emoji  hidden md:block" aria-hidden="true">😤</span>
                </button>
              </div>
            </div>
            <input type="range" min="1" :max="stressMax" step="1" v-model.number="stressLevel" class="w-full slider" />
            <div class="flex justify-between text-sm mt-2 text-gray-600">
              <span>Calmo</span>
              <span>Moderado</span>
              <span>Alto</span>
            </div>
          </div>
          <div class="w-full md:w-36 text-center display-area mt-6 md:mt-0">
            <div class="display-emoji">😤</div>
            <div class="font-semibold display-value">{{ stressLevel }} / {{ stressMax }}</div>
          </div>
        </div>
      </div>

      <div class="bg-green-200 rounded-lg shadow-sm border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-6">Atividades de Autocuidado realizadas hoje</h2>
        <div class="grid grid-cols-2 sm:grid-cols-2 md:grid-cols-3 gap-4">
          <label v-for="activity in selfCareActivities" :key="activity"
            class="flex items-center space-x-3 text-black custom-checkbox-label">
            <input type="checkbox" :value="activity" v-model="selectedActivities" class="hidden peer" />
            <span class="checkbox-box" :aria-checked="selectedActivities.includes(activity)">
              <svg v-if="selectedActivities.includes(activity)" width="14" height="14" viewBox="0 0 24 24" fill="none"
                xmlns="http://www.w3.org/2000/svg">
                <path d="M20 6L9 17L4 12" stroke="white" stroke-width="2.5" stroke-linecap="round"
                  stroke-linejoin="round" />
              </svg>
            </span>
            <span>{{ activity }}</span>
          </label>
          <div class="col-span-2 md:col-span-1">
            <input type="text" v-model="otherActivity" placeholder="Outra atividade..."
              class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-white focus:border-emerald-500 outline-none transition-colors text-gray-900 placeholder-gray-500">
          </div>
        </div>
      </div>

      <div class="bg-gray-50 rounded-lg shadow-sm border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-6">Anotação Diária (Opcional)</h2>
        <p class="text-sm text-black mb-4">Escreva sobre seus pensamentos, sentimentos ou o que marcou seu dia. Isso
          ajudará você e seu profissional a entenderem o contexto do seu bem-estar.</p>
        <textarea v-model="notes" rows="6" placeholder="Sobre o que gostaria de comentar?"
          class="w-full p-4 rounded-lg border border-gray-300 focus:ring-2 focus:ring-white focus:border-emerald-500 outline-none transition-colors"></textarea>
      </div>

      <div class="flex justify-end">
        <button type="submit"
          class="w-full md:w-auto bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-3 px-6 rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 outline-none">
          Registrar Humor
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useToast } from "vue-toastification";
import api from '../services/api';

const toast = useToast();

// Core selections
const selectedMood = ref('');
const sleepHours = ref(0);
const energyLevel = ref(6);
const stressLevel = ref(6);
const selectedEmotions = ref([]);
const otherEmotion = ref('');
const selectedActivities = ref([]);
const otherActivity = ref('');
const notes = ref('');

// Discrete slider configuration (number of steps)
const sleepMax = 12;
const energyMax = 10;
const stressMax = 10;

// Arrays used to render the icon rows (0..max)
const sleepSteps = Array.from({ length: sleepMax + 1 }, (_, i) => i);
const energySteps = Array.from({ length: energyMax }, (_, i) => i);
const stressSteps = Array.from({ length: stressMax }, (_, i) => i);

const moodOptions = [
  { value: 'very_bad', label: 'Muito Mal', emoji: '😖' },
  { value: 'bad', label: 'Aborrecido', emoji: '😕' },
  { value: 'neutral', label: 'Neutro', emoji: '😐' },
  { value: 'cheerful', label: 'Animado', emoji: '😊' },
  { value: 'very_good', label: 'Muito Bem', emoji: '😁' },
];

// More elaborate stress SVG used for the large icon display

const today = new Date();
const currentDate = computed(() => today.toLocaleDateString('pt-BR', { year: 'numeric', month: 'long', day: 'numeric' }));
const currentDay = computed(() => today.toLocaleDateString('pt-BR', { weekday: 'long' }));

const selfCareActivities = ['Leitura', 'Exercício', 'Medicação', 'Yoga', 'Meditação', 'Música', 'Hobbies', 'Nenhuma Atividade'];

const sleepLabel = computed(() => (sleepHours.value >= sleepMax ? '12h+' : sleepHours.value + 'h'));

const moodLevelMapping = {
  'very_bad': 1,
  'bad': 2,
  'neutral': 3,
  'cheerful': 4,
  'very_good': 5
};

const submitMood = async () => {
  if (!selectedMood.value) {
    toast.warning('Por favor, selecione seu humor principal antes de registrar.');
    return;
  }

  let finalActivities = [...selectedActivities.value];
  if (otherActivity.value) finalActivities.push(otherActivity.value);
  if (finalActivities.includes('Nenhuma Atividade')) finalActivities = ['Nenhuma Atividade'];

  // Objeto padronizado para corresponder ao DTO do backend
  const submission = {
    nivel_humor: moodLevelMapping[selectedMood.value],
    horas_sono: sleepHours.value > sleepMax ? sleepMax : sleepHours.value,
    nivel_energia: energyLevel.value > energyMax ? energyMax : energyLevel.value,
    nivel_stress: stressLevel.value > stressMax ? stressMax : stressLevel.value,
    auto_cuidado: finalActivities,
    observacoes: notes.value,
    data_hora_registro: new Date().toISOString(),
  };

  try {
    await api.registrarHumor(submission);
    toast.success('Seu humor foi registrado com sucesso!');
    // Opcional: limpar o formulário ou navegar para outra página
  } catch (error) {
    toast.error('Houve um erro ao registrar seu humor.');
    console.error("Erro ao registrar humor:", error);
  }
};
</script>

<style scoped>
.mood-icon {
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
  border-radius: 9999px;
  box-shadow: 0 0 0 3px transparent;
}

.mood-icon:hover {
  transform: scale(1.1);
}

.mood-icon.selected {
  box-shadow: 0 0 0 3px rgba(255, 104, 122, 0.95);
  transform: scale(1.1);
}

.sleep-icon {
  cursor: pointer;
  transition: transform 0.2s, opacity 0.2s;
  font-size: 2rem;
  opacity: 0.5;
}

.sleep-icon.selected {
  opacity: 1;
  transform: scale(1.1);
  filter: drop-shadow(0 0 4px rgba(255, 104, 122, 0.35));
}

/* New styles for sliders and mood buttons */
.slider {
  appearance: none;
  -webkit-appearance: none;
  height: 6px;
  background: linear-gradient(90deg, rgba(255, 104, 122, 0.9), rgba(102, 206, 131, 0.9));
  border-radius: 999px;
  outline: none;
}

.slider::-webkit-slider-thumb {
  appearance: none;
  -webkit-appearance: none;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: white;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.2);
  border: 3px solid rgba(255, 104, 122, 0.95);
}

.mood-btn {
  width: 110px;
  height: 110px;
  background: transparent;
  /* removed white background */
  border: 1px solid transparent;
}

.mood-btn:hover {
  border-color: rgba(16, 185, 129, 0.12);
  background: rgba(102, 206, 131, 0.04);
}

.mood-btn.selected {
  border-color: rgba(102, 206, 131, 0.9);
  background: linear-gradient(180deg, rgba(102, 206, 131, 0.08), rgba(102, 206, 131, 0.03));
  box-shadow: 0 6px 18px rgba(16, 185, 129, 0.08);
}

/* Larger emoji and svg icon sizes */
.mood-emoji {
  font-size: 34px;
  /* larger emojis */
  line-height: 1;
}

.icon-svg {
  display: inline-block;
  width: 48px;
  height: 48px;
}

.step-svg {
  width: 22px;
  height: 22px;
  display: inline-block;
}

.display-icon {
  width: 48px;
  height: 48px;
  margin: 0 auto 4px;
}

/* increase spacing and tappable area for steps */
.icon-row {
  gap: 8px;
}

.icon-step {
  padding: 10px 6px;
  flex: 0 0 28px;
  /* fixed step width to avoid overflow */
}

/* make the icon row constrained to the slider width */
.w-full .icon-row {
  max-width: 100%;
  justify-content: space-between;
}

.display-area .display-emoji {
  font-size: 44px;
}

.display-area .display-value {
  font-size: 18px;
}

/* Larger emoji steps and display */
.step-emoji {
  font-size: 22px;
  line-height: 1;
}

.display-emoji {
  font-size: 36px;
  line-height: 1;
  margin-bottom: 4px;
}

.icon-step.active .step-emoji {
  transform: translateY(-6px) scale(1.14);
}

/* Custom checkbox visuals */
.custom-checkbox-label {
  align-items: center;
}

.checkbox-box {
  width: 28px;
  height: 28px;
  border-radius: 8px;
  border: 2px solid rgba(15, 23, 42, 0.06);
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background: white;
  transition: background-color 0.12s ease, border-color 0.12s ease, transform 0.08s ease;
}

.peer:checked+.checkbox-box,
.checkbox-box[aria-checked="true"] {
  background: #10B981;
  border-color: #10B981;
  transform: translateY(-2px);
}

.checkbox-box svg {
  display: block;
}

/* Icon row (discrete steps above the slider) */
.icon-row {
  display: flex;
  gap: 6px;
  align-items: center;
  width: 100%;
}

.icon-step {
  flex: 1 1 0;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  border: none;
  cursor: pointer;
  padding: 6px 4px;
  border-radius: 6px;
  opacity: 0.45;
  transition: transform 0.12s ease, opacity 0.12s ease, background-color 0.12s ease;
}

.icon-step.active {
  opacity: 1;
  transform: translateY(-4px) scale(1.08);
}
</style>
