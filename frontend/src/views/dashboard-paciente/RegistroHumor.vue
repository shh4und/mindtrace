<template>
  <div class="max-w-4xl mx-auto p-4 md:p-8">
    <!-- Header com Data Integrada -->
    <header class="mb-10 text-center">
      <div
        class="inline-flex items-center justify-center space-x-2 bg-emerald-50 text-emerald-800 px-4 py-1.5 rounded-full text-sm font-medium mb-4 shadow-sm"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-4 w-4"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
            clip-rule="evenodd"
          />
        </svg>
        <span>{{ currentDay }}, {{ currentDate }}</span>
      </div>
      <h1
        class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
      >
        Como você está hoje?
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Registre seu bem-estar para acompanhar sua jornada.
      </p>
    </header>

    <form @submit.prevent="submit" class="space-y-8" novalidate>
      <!-- Humor - Card Principal -->
      <section
        class="bg-white rounded-3xl shadow-lg border border-gray-100 overflow-hidden relative"
      >
        <div
          class="absolute top-0 left-0 w-full h-2 bg-gradient-to-r from-rose-400 via-purple-400 to-indigo-400"
        ></div>
        <div class="p-6 md:p-10">
          <h2 class="text-xl font-bold text-gray-800 mb-8 text-center">
            Selecione o emoji que melhor representa seu dia
          </h2>

          <div
            class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 gap-4 md:gap-6 justify-items-center"
            role="radiogroup"
            aria-label="Seleção de humor"
          >
            <button
              v-for="m in moodOptions"
              :key="m.value"
              type="button"
              @click="selectedMood = m.value"
              class="group relative flex flex-col items-center justify-center w-full aspect-square rounded-2xl border-2 transition-all duration-300 ease-out focus:outline-none focus:ring-4 focus:ring-emerald-100"
              :class="[
                selectedMood === m.value
                  ? 'border-emerald-500 bg-emerald-50 shadow-emerald-100 scale-105 shadow-lg'
                  : 'border-gray-100 bg-white hover:border-emerald-200 hover:bg-gray-50 hover:-translate-y-1',
              ]"
              :aria-pressed="selectedMood === m.value"
            >
              <div
                class="transform transition-transform duration-300"
                :class="{ 'scale-110': selectedMood === m.value }"
              >
                <Emoji
                  :data="emojiIndex"
                  :emoji="m.emoji"
                  :size="42"
                  set="facebook"
                />
              </div>
              <span
                class="mt-3 text-sm font-semibold transition-colors duration-200"
                :class="
                  selectedMood === m.value
                    ? 'text-emerald-700'
                    : 'text-gray-500 group-hover:text-gray-700'
                "
              >
                {{ m.label }}
              </span>

              <!-- Check icon for selected state -->
              <div
                v-if="selectedMood === m.value"
                class="absolute -top-3 -right-3 bg-emerald-500 text-white rounded-full p-1 shadow-md animate-fade-in-up"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd"
                  />
                </svg>
              </div>
            </button>
          </div>
        </div>
      </section>

      <!-- Métricas (Grid de 2 colunas no desktop) -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Sono -->
        <section
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow"
        >
          <header class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-bold text-gray-800 flex items-center gap-2">
              <span class="p-2 bg-indigo-100 rounded-lg text-indigo-600">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
                  />
                </svg>
              </span>
              Sono
            </h3>
            <span class="text-2xl font-bold text-indigo-600">{{
              sleepLabel
            }}</span>
          </header>

          <div class="space-y-6">
            <div class="flex justify-between px-2">
              <button
                v-for="step in sleepSteps"
                :key="step"
                type="button"
                @click="sleepHours = step"
                class="w-8 h-8 rounded-full flex items-center justify-center text-sm transition-all duration-200"
                :class="
                  step <= sleepHours
                    ? 'bg-indigo-100 text-indigo-700 font-bold scale-110'
                    : 'text-gray-400 hover:bg-gray-100'
                "
              >
                {{ step }}
              </button>
            </div>
            <input
              type="range"
              :min="sleepConfig.min"
              :max="sleepConfig.max"
              v-model.number="sleepHours"
              class="w-full h-3 bg-gray-200 rounded-lg appearance-none cursor-pointer accent-indigo-600 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            />
            <div class="flex justify-center pt-2">
              <Emoji
                :data="emojiIndex"
                :emoji="sleepConfig.emoji"
                :size="32"
                set="facebook"
              />
            </div>
          </div>
        </section>

        <!-- Energia -->
        <section
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow"
        >
          <header class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-bold text-gray-800 flex items-center gap-2">
              <span class="p-2 bg-amber-100 rounded-lg text-amber-600">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13 10V3L4 14h7v7l9-11h-7z"
                  />
                </svg>
              </span>
              Energia
            </h3>
            <span class="text-2xl font-bold text-amber-600"
              >{{ energyLevel }}/{{ energyConfig.max }}</span
            >
          </header>

          <div class="space-y-6">
            <div class="flex justify-between px-2">
              <button
                v-for="(step, index) in energySteps"
                :key="step"
                type="button"
                @click="energyLevel = step + 1"
                class="w-8 h-8 rounded-full flex items-center justify-center transition-all duration-200"
                :class="
                  step < energyLevel
                    ? 'bg-amber-100 text-amber-700 scale-110'
                    : 'bg-gray-100 text-gray-300'
                "
              >
                <div
                  class="w-2 h-2 rounded-full"
                  :class="step < energyLevel ? 'bg-amber-500' : 'bg-gray-300'"
                ></div>
              </button>
            </div>
            <input
              type="range"
              :min="energyConfig.min"
              :max="energyConfig.max"
              v-model.number="energyLevel"
              class="w-full h-3 bg-gray-200 rounded-lg appearance-none cursor-pointer accent-amber-500 focus:outline-none focus:ring-2 focus:ring-amber-500 focus:ring-offset-2"
            />
            <div class="flex justify-center pt-2">
              <Emoji
                :data="emojiIndex"
                :emoji="energyConfig.emoji"
                :size="32"
                set="facebook"
              />
            </div>
          </div>
        </section>

        <!-- Stress (Full Width on mobile, half on large if needed, but let's make it span full if odd) -->
        <section
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 lg:col-span-2 hover:shadow-md transition-shadow"
        >
          <header class="flex items-center justify-between mb-6">
            <h3 class="text-lg font-bold text-gray-800 flex items-center gap-2">
              <span class="p-2 bg-rose-100 rounded-lg text-rose-600">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
                  />
                </svg>
              </span>
              Nível de Stress
            </h3>
            <span class="text-2xl font-bold text-rose-600"
              >{{ stressLevel }}/{{ stressConfig.max }}</span
            >
          </header>

          <div class="space-y-6">
            <div class="flex justify-between px-2">
              <button
                v-for="(step, index) in stressSteps"
                :key="step"
                type="button"
                @click="stressLevel = step + 1"
                class="w-8 h-8 rounded-full flex items-center justify-center transition-all duration-200"
                :class="
                  step < stressLevel
                    ? 'bg-rose-100 text-rose-700 scale-110'
                    : 'bg-gray-100 text-gray-300'
                "
              >
                <div
                  class="w-2 h-2 rounded-full"
                  :class="step < stressLevel ? 'bg-rose-500' : 'bg-gray-300'"
                ></div>
              </button>
            </div>
            <input
              type="range"
              :min="stressConfig.min"
              :max="stressConfig.max"
              v-model.number="stressLevel"
              class="w-full h-3 bg-gray-200 rounded-lg appearance-none cursor-pointer accent-rose-500 focus:outline-none focus:ring-2 focus:ring-rose-500 focus:ring-offset-2"
            />
            <div
              class="flex justify-between text-sm text-gray-500 font-medium px-1"
            >
              <span>Baixo</span>
              <span class="text-rose-500">Alto</span>
            </div>
          </div>
        </section>
      </div>

      <!-- Autocuidado -->
      <section
        class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8"
      >
        <h3
          class="text-xl font-bold text-gray-800 mb-6 flex items-center gap-2"
        >
          <span class="p-2 bg-emerald-100 rounded-lg text-emerald-600">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
              />
            </svg>
          </span>
          O que você fez por você hoje?
        </h3>

        <div class="flex flex-wrap gap-3">
          <label
            v-for="activity in displayedActivities"
            :key="activity"
            class="relative group cursor-pointer"
          >
            <input
              type="checkbox"
              :value="activity"
              v-model="selectedActivities"
              class="peer sr-only"
            />
            <span
              class="inline-flex items-center px-5 py-2.5 rounded-xl text-sm font-medium transition-all duration-200 border-2 select-none peer-checked:bg-emerald-500 peer-checked:border-emerald-500 peer-checked:text-white peer-checked:shadow-md bg-white border-gray-200 text-gray-600 hover:border-emerald-200 hover:bg-emerald-50"
            >
              {{ activity }}
            </span>
          </label>
        </div>

        <div class="mt-6 relative">
          <input
            type="text"
            v-model="otherActivity"
            @keydown.enter.prevent="addCustomActivity"
            placeholder="Digite outra atividade e tecle Enter..."
            class="w-full px-4 py-3 rounded-xl border border-gray-200 focus:ring-2 focus:ring-emerald-500 focus:border-transparent outline-none transition-all placeholder-gray-400 bg-gray-50 hover:bg-white pr-10"
          />
           <button 
            type="button"
            @click="addCustomActivity"
            v-if="otherActivity"
            class="absolute right-3 top-1/2 transform -translate-y-1/2 text-emerald-600 hover:bg-emerald-50 rounded-full p-1 transition-colors"
          >
             <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
          </button>
        </div>
      </section>

      <!-- Notas -->
      <section
        class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8"
      >
        <h3
          class="text-xl font-bold text-gray-800 mb-4 flex items-center gap-2"
        >
          <span class="p-2 bg-gray-100 rounded-lg text-gray-600">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
              />
            </svg>
          </span>
          Diário de Bordo
        </h3>
        <textarea
          v-model="notes"
          rows="4"
          placeholder="Escreva sobre seus pensamentos, gratidão ou desafios do dia..."
          class="w-full p-4 rounded-2xl border border-gray-200 focus:ring-2 focus:ring-emerald-500 focus:border-transparent outline-none transition-all resize-y bg-gray-50 hover:bg-white placeholder-gray-400 text-gray-700 leading-relaxed"
        ></textarea>
      </section>

      <!-- Botão Salvar -->
      <div
        class="fixed bottom-0 left-0 right-0 p-4 bg-white/80 backdrop-blur-md border-t border-gray-200 md:static md:bg-transparent md:border-none md:p-0 z-10 md:flex md:justify-end"
      >
        <button
          type="submit"
          :disabled="isSubmitting || !isValid"
          class="w-full md:w-auto bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 text-white font-bold py-4 px-10 rounded-2xl shadow-lg hover:shadow-xl hover:-translate-y-1 transition-all duration-300 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-lg flex items-center justify-center space-x-2"
        >
          <span v-if="isSubmitting" class="flex items-center">
            <svg
              class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
            >
              <circle
                class="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                stroke-width="4"
              ></circle>
              <path
                class="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
              ></path>
            </svg>
            Salvando...
          </span>
          <span v-else>Salvar Registro</span>
        </button>
      </div>
      <!-- Spacer for mobile fixed bottom button -->
      <div class="h-24 md:h-0"></div>
    </form>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useMoodForm } from "@/composables/useMoodForm";
import data from "emoji-mart-vue-fast/data/all.json";
import { EmojiIndex, Emoji } from "emoji-mart-vue-fast/src";
import "emoji-mart-vue-fast/css/emoji-mart.css";

const emojiIndex = new EmojiIndex(data);

const {
  selectedMood,
  sleepHours,
  energyLevel,
  stressLevel,
  selectedActivities,
  otherActivity,
  notes,
  isSubmitting,
  moodOptions,
  selfCareActivities: initialSelfCareActivities,
  sleepConfig,
  energyConfig,
  stressConfig,
  sleepSteps,
  energySteps,
  stressSteps,
  sleepLabel,
  currentDate,
  currentDay,
  isValid,
  submit,
} = useMoodForm();

// Estado local para a lista de atividades (para permitir adição dinâmica)
const displayedActivities = ref([...initialSelfCareActivities]);

const addCustomActivity = () => {
  const val = otherActivity.value.trim();
  if (val) {
    // Se não existir na lista, adiciona
    if (!displayedActivities.value.includes(val)) {
      displayedActivities.value.push(val);
    }
    // Seleciona automaticamente (se já não estiver selecionado)
    if (!selectedActivities.value.includes(val)) {
      selectedActivities.value.push(val);
    }
    // Limpa o campo
    otherActivity.value = "";
  }
};
</script>

<style scoped>
/* Animação suave para entrada de elementos */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translate3d(0, 10px, 0);
  }
  to {
    opacity: 1;
    transform: translate3d(0, 0, 0);
  }
}

.animate-fade-in-up {
  animation: fadeInUp 0.3s ease-out forwards;
}

/* Custom Range Slider Styling */
input[type="range"] {
  -webkit-appearance: none;
  background: transparent;
}

input[type="range"]::-webkit-slider-thumb {
  -webkit-appearance: none;
  height: 24px;
  width: 24px;
  border-radius: 50%;
  background: white;
  border: 2px solid currentColor;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
  margin-top: -8px;
}

/* Specific thumb colors per section handled by text-color utility in parent or currentColor if mapped properly, 
   but simplistic approach: */
input[type="range"].accent-indigo-600::-webkit-slider-thumb {
  border-color: #4f46e5; /* indigo-600 */
}
input[type="range"].accent-amber-500::-webkit-slider-thumb {
  border-color: #f59e0b; /* amber-500 */
}
input[type="range"].accent-rose-500::-webkit-slider-thumb {
  border-color: #f43f5e; /* rose-500 */
}

input[type="range"]::-moz-range-thumb {
  height: 24px;
  width: 24px;
  border-radius: 50%;
  background: white;
  border: 2px solid currentColor;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
}

input[type="range"]::-webkit-slider-runnable-track {
  width: 100%;
  height: 8px;
  cursor: pointer;
  background: #e5e7eb; /* gray-200 */
  border-radius: 999px;
}
</style>