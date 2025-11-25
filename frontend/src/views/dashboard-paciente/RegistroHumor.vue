<template>
  <div class="flex-1 p-4 md:p-8 w-full">
    <header class="mb-8 text-center md:text-left">
      <h1 class="text-2xl md:text-3xl font-bold text-gray-900">Registro de Humor Diário</h1>
      <p class="text-gray-600 mt-1">Preencha como você se sentiu hoje e registre suas observações.</p>
    </header>

    <form @submit.prevent="submit" class="space-y-8" novalidate>
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Registro para o dia:</h2>
        <div class="flex flex-col md:flex-row md:space-x-4 space-y-2 md:space-y-0">
          <input type="text" :value="currentDate"
            class="flex-1 text-lg font-medium text-gray-700 bg-gray-100 rounded-md p-2 text-center" readonly aria-label="Data atual" />
          <input type="text" :value="currentDay"
            class="flex-1 text-lg font-medium text-gray-700 bg-gray-100 rounded-md p-2 text-center" readonly aria-label="Dia da semana" />
        </div>
      </div>

      <!-- Secao de selecao de humor com emojis -->
      <fieldset class="bg-rose-50 rounded-lg shadow-sm border border-gray-200 p-6">
        <legend class="text-xl font-semibold text-gray-900 mb-4 text-center w-full">Como você se sentiu hoje?</legend>
        <div class="flex flex-wrap gap-6 justify-center" role="radiogroup" aria-label="Selecione seu humor">
          <button 
            v-for="m in moodOptions" 
            :key="m.value" 
            type="button" 
            @click="selectedMood = m.value"
            :aria-pressed="selectedMood === m.value"
            :aria-label="`${m.label} ${m.emoji}`"
            :class="['mood-btn flex flex-col items-center justify-center p-3 rounded-full transition', { 'selected': selectedMood === m.value }]"
          >
            <span class="mood-emoji mb-2" aria-hidden="true">{{ m.emoji }}</span>
            <span class="text-sm">{{ m.label }}</span>
          </button>
        </div>
      </fieldset>

      <!-- Controle de sono com icones de lua -->
      <fieldset class="bg-rose-100 rounded-lg shadow-sm border border-gray-200 p-6">
        <legend class="text-xl font-semibold text-gray-900 mb-6 text-center w-full">Horas de Sono</legend>
        <div class="flex flex-col md:flex-row items-center md:space-x-4">
          <div class="w-full">
            <div class="mb-3">
              <div class="icon-row" role="group" aria-label="Indicadores visuais de sono">
                <button 
                  v-for="step in sleepSteps" 
                  :key="step" 
                  type="button" 
                  class="icon-step"
                  :class="[{ 'active': step <= sleepHours }]" 
                  @click="sleepHours = step"
                  :aria-label="`${step} horas de sono`"
                >
                  <span class="step-emoji hidden md:block" aria-hidden="true">🌜</span>
                </button>
              </div>
            </div>
            <input 
              type="range" 
              :min="sleepConfig.min" 
              :max="sleepConfig.max" 
              step="1" 
              v-model.number="sleepHours" 
              class="w-full slider"
              :aria-label="`Horas de sono: ${sleepLabel}`"
              aria-valuemin="0"
              :aria-valuemax="sleepConfig.max"
              :aria-valuenow="sleepHours"
              :aria-valuetext="sleepLabel"
            />
            <div class="flex justify-between text-sm mt-2 text-gray-600" aria-hidden="true">
              <span v-for="label in sleepConfig.labels" :key="label">{{ label }}</span>
            </div>
          </div>
          <div class="w-full md:w-36 text-center display-area mt-6 md:mt-0" aria-live="polite">
            <div class="display-emoji" aria-hidden="true">🌜</div>
            <div class="font-semibold display-value">{{ sleepLabel }}</div>
          </div>
        </div>
      </fieldset>

      <!-- Controle de energia com icones de bateria -->
      <fieldset class="bg-amber-50 rounded-lg shadow-sm border border-gray-200 p-6">
        <legend class="text-xl font-semibold text-gray-900 mb-6 text-center w-full">Nível de Energia</legend>
        <div class="flex flex-col md:flex-row items-center md:space-x-4">
          <div class="w-full">
            <div class="mb-3">
              <div class="icon-row" role="group" aria-label="Indicadores visuais de energia">
                <button 
                  v-for="step in energySteps" 
                  :key="step" 
                  type="button" 
                  class="icon-step"
                  :class="{ 'active': step < energyLevel }" 
                  @click="energyLevel = step + 1"
                  :aria-label="`Nível ${step + 1} de energia`"
                >
                  <span class="step-emoji hidden md:block" aria-hidden="true">🔋</span>
                </button>
              </div>
            </div>
            <input 
              type="range" 
              :min="energyConfig.min" 
              :max="energyConfig.max" 
              step="1" 
              v-model.number="energyLevel" 
              class="w-full slider"
              :aria-label="`Nível de energia: ${energyLevel} de ${energyConfig.max}`"
              :aria-valuemin="energyConfig.min"
              :aria-valuemax="energyConfig.max"
              :aria-valuenow="energyLevel"
            />
            <div class="flex justify-between text-sm mt-2 text-gray-600" aria-hidden="true">
              <span v-for="label in energyConfig.labels" :key="label">{{ label }}</span>
            </div>
          </div>
          <div class="w-full md:w-36 text-center display-area mt-6 md:mt-0" aria-live="polite">
            <div class="display-emoji" aria-hidden="true">🔋</div>
            <div class="font-semibold display-value">{{ energyLevel }} / {{ energyConfig.max }}</div>
          </div>
        </div>
      </fieldset>

      <!-- Controle de stress com icones de alerta -->
      <fieldset class="bg-yellow-50 rounded-lg shadow-sm border border-gray-200 p-6">
        <legend class="text-xl font-semibold text-gray-900 mb-6 text-center w-full">Nível de Stress</legend>
        <div class="flex flex-col md:flex-row items-center md:space-x-4">
          <div class="w-full">
            <div class="mb-3">
              <div class="icon-row" role="group" aria-label="Indicadores visuais de stress">
                <button 
                  v-for="step in stressSteps" 
                  :key="step" 
                  type="button" 
                  class="icon-step"
                  :class="{ 'active': step < stressLevel }" 
                  @click="stressLevel = step + 1"
                  :aria-label="`Nível ${step + 1} de stress`"
                >
                  <span class="step-emoji hidden md:block" aria-hidden="true">😤</span>
                </button>
              </div>
            </div>
            <input 
              type="range" 
              :min="stressConfig.min" 
              :max="stressConfig.max" 
              step="1" 
              v-model.number="stressLevel" 
              class="w-full slider"
              :aria-label="`Nível de stress: ${stressLevel} de ${stressConfig.max}`"
              :aria-valuemin="stressConfig.min"
              :aria-valuemax="stressConfig.max"
              :aria-valuenow="stressLevel"
            />
            <div class="flex justify-between text-sm mt-2 text-gray-600" aria-hidden="true">
              <span v-for="label in stressConfig.labels" :key="label">{{ label }}</span>
            </div>
          </div>
          <div class="w-full md:w-36 text-center display-area mt-6 md:mt-0" aria-live="polite">
            <div class="display-emoji" aria-hidden="true">😤</div>
            <div class="font-semibold display-value">{{ stressLevel }} / {{ stressConfig.max }}</div>
          </div>
        </div>
      </fieldset>

      <fieldset class="bg-green-200 rounded-lg shadow-sm border border-gray-200 p-6">
        <legend class="text-xl font-semibold text-gray-900 mb-6">Atividades de Autocuidado realizadas hoje</legend>
        <div class="grid grid-cols-2 sm:grid-cols-2 md:grid-cols-3 gap-4">
          <label 
            v-for="activity in selfCareActivities" 
            :key="activity"
            class="flex items-center space-x-3 text-black cursor-pointer"
          >
            <input type="checkbox" :value="activity" v-model="selectedActivities" class="hidden peer" />
            <span class="w-7 h-7 border-2 border-gray-300/70 rounded-lg bg-white inline-flex items-center justify-center transition-all duration-150 peer-checked:bg-emerald-500 peer-checked:border-emerald-500 peer-checked:-translate-y-0.5" aria-hidden="true">
              <svg v-if="selectedActivities.includes(activity)" width="14" height="14" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M20 6L9 17L4 12" stroke="white" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
            </span>
            <span>{{ activity }}</span>
          </label>
          <div class="col-span-2 md:col-span-1">
            <label for="other-activity" class="sr-only">Outra atividade</label>
            <input 
              type="text" 
              id="other-activity"
              v-model="otherActivity" 
              placeholder="Outra atividade..."
              class="w-full px-4 py-2 rounded-lg border border-gray-300 focus:ring-2 focus:ring-white focus:border-emerald-500 outline-none transition-colors text-gray-900 placeholder-gray-500"
            />
          </div>
        </div>
      </fieldset>

      <fieldset class="bg-gray-50 rounded-lg shadow-sm border border-gray-200 p-6">
        <legend class="text-xl font-semibold text-gray-900 mb-6">Anotação Diária (Opcional)</legend>
        <p class="text-sm text-black mb-4">Escreva sobre seus pensamentos, sentimentos ou o que marcou seu dia. Isso ajudará você e seu profissional a entenderem o contexto do seu bem-estar.</p>
        <label for="notes" class="sr-only">Anotação diária</label>
        <textarea 
          id="notes"
          v-model="notes" 
          rows="6" 
          placeholder="Sobre o que gostaria de comentar?"
          class="w-full p-4 rounded-lg border border-gray-300 focus:ring-2 focus:ring-white focus:border-emerald-500 outline-none transition-colors"
        ></textarea>
      </fieldset>

      <div class="flex justify-end">
        <button 
          type="submit"
          :disabled="isSubmitting || !isValid"
          :aria-busy="isSubmitting"
          class="w-full md:w-auto bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-3 px-6 rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 outline-none disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="isSubmitting">Registrando...</span>
          <span v-else>Registrar Humor</span>
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { useMoodForm } from '@/composables/useMoodForm';

// Desestrutura todos os estados, configs e métodos do composable
const {
  // Estado do formulário
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
  submit
} = useMoodForm();
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

/* Novos estilos para os sliders e botoes de humor */
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
  /* Fundo branco removido */
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

/* Emojis e svgs em tamanho ampliado */
.mood-emoji {
  font-size: 34px;
  /* Emojis maiores para facilitar a leitura */
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

/* Aumenta o espacamento e a area de toque dos passos */
.icon-row {
  gap: 8px;
}

.icon-step {
  padding: 10px 6px;
  flex: 0 0 28px;
  /* Largura fixa para evitar overflow */
}

/* Mantem a linha de icones dentro da largura do slider */
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

/* Emojis maiores para passos e display principal */
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



/* Linha de icones representando passos discretos acima do slider */
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
