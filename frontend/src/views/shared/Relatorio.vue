<template>
  <div class="max-w-6xl mx-auto p-4 md:p-8">
    <!-- Botão voltar para profissional -->
    <div v-if="userType === TipoUsuario.Profissional && patientId" class="mb-6">
      <button
        @click="goBack"
        class="p-2.5 rounded-xl hover:bg-gray-100 transition-colors"
        aria-label="Voltar para a lista de pacientes"
      >
        <font-awesome-icon :icon="faArrowLeft" class="w-5 h-5 text-gray-600" />
      </button>
    </div>

    <!-- Header -->
    <header class="mb-10">
      <h1
        class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
      >
        {{
          patientName
            ? `Relatório de ${patientName}`
            : "Relatórios de Bem-Estar"
        }}
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        {{
          patientName
            ? "Acompanhamento detalhado do paciente."
            : "Analise suas tendências de humor, sono e energia ao longo do tempo."
        }}
      </p>
    </header>

    <!-- Filtros de Período -->
    <div
      class="mb-8 flex flex-wrap gap-2"
      role="group"
      aria-label="Filtro de período"
    >
      <button
        v-for="range in timeRanges"
        :key="range.days"
        @click="selectedRange = range.days"
        :aria-pressed="selectedRange === range.days"
        :class="[
          'px-5 py-2.5 rounded-xl font-bold text-sm transition-all',
          selectedRange === range.days
            ? 'bg-indigo-600 text-white shadow-md'
            : 'bg-white text-gray-600 hover:bg-gray-50 shadow-sm border border-gray-100 hover:text-gray-900',
        ]"
      >
        {{ range.label }}
      </button>
    </div>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-20">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"
      ></div>
    </div>

    <template v-else>
      <!-- Cards de Estatísticas -->
      <section
        class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8"
        aria-label="Estatísticas resumidas"
      >
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between mb-3">
            <p class="text-xs text-gray-500 uppercase tracking-wide font-bold">
              Média de Sono
            </p>
            <span class="p-2 bg-blue-100 rounded-lg text-blue-600">
              <font-awesome-icon :icon="faBed" class="w-5 h-5" />
            </span>
          </div>
          <p class="text-4xl font-extrabold text-blue-600">
            {{ avgSleep }}
            <span class="text-base font-bold text-gray-400">horas/noite</span>
          </p>
        </div>
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between mb-3">
            <p class="text-xs text-gray-500 uppercase tracking-wide font-bold">
              Média de Energia
            </p>
            <span class="p-2 bg-amber-100 rounded-lg text-amber-600">
              <font-awesome-icon :icon="faBolt" class="w-5 h-5" />
            </span>
          </div>
          <p class="text-4xl font-extrabold text-amber-600">
            {{ avgEnergy }}
            <span class="text-base font-bold text-gray-400">/ 10</span>
          </p>
        </div>
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between mb-3">
            <p class="text-xs text-gray-500 uppercase tracking-wide font-bold">
              Média de Stress
            </p>
            <span class="p-2 bg-red-100 rounded-lg text-red-600">
              <font-awesome-icon :icon="faHeartPulse" class="w-5 h-5" />
            </span>
          </div>
          <p class="text-4xl font-extrabold text-red-600">
            {{ avgStress }}
            <span class="text-base font-bold text-gray-400">/ 10</span>
          </p>
        </div>
      </section>

      <!-- Gráficos -->
      <section class="space-y-6" aria-label="Gráficos de tendências">
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow"
        >
          <h3
            class="font-bold text-lg text-gray-800 mb-4 flex items-center gap-2"
          >
            <span class="p-2 bg-blue-100 rounded-lg text-blue-600">
              <font-awesome-icon :icon="faBed" class="w-5 h-5" />
            </span>
            Horas de Sono
          </h3>
          <apexchart
            type="area"
            height="350"
            :options="sleepChartOptions"
            :series="sleepSeries"
          ></apexchart>
        </div>
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow"
        >
          <h3
            class="font-bold text-lg text-gray-800 mb-4 flex items-center gap-2"
          >
            <span class="p-2 bg-amber-100 rounded-lg text-amber-600">
              <font-awesome-icon :icon="faBolt" class="w-5 h-5" />
            </span>
            Nível de Energia
          </h3>
          <apexchart
            type="area"
            height="350"
            :options="energyChartOptions"
            :series="energySeries"
          ></apexchart>
        </div>
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow"
        >
          <h3
            class="font-bold text-lg text-gray-800 mb-4 flex items-center gap-2"
          >
            <span class="p-2 bg-red-100 rounded-lg text-red-600">
              <font-awesome-icon :icon="faHeartPulse" class="w-5 h-5" />
            </span>
            Nível de Stress
          </h3>
          <apexchart
            type="area"
            height="350"
            :options="stressChartOptions"
            :series="stressSeries"
          ></apexchart>
        </div>
      </section>
    </template>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from "vue";
import { useRouter } from "vue-router";
import api from "@/services/api";
import { useToast } from "vue-toastification";
import { TipoUsuario } from "@/types/usuario.js";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faArrowLeft,
  faBed,
  faBolt,
  faHeartPulse,
} from "@fortawesome/free-solid-svg-icons";

import data from "emoji-mart-vue-fast/data/all.json";
import { EmojiIndex, Emoji } from "emoji-mart-vue-fast/src";
import "emoji-mart-vue-fast/css/emoji-mart.css";

const emojiIndex = new EmojiIndex(data);

const props = defineProps({
  patientId: {
    type: [Number, String],
    default: null,
  },
  userType: {
    type: String,
    default: TipoUsuario.Paciente,
    validator: (value) =>
      [TipoUsuario.Paciente, TipoUsuario.Profissional].includes(value),
  },
});

const router = useRouter();
const toast = useToast();

// --- ESTADO DO COMPONENTE ---
const allData = ref([]);
const selectedRange = ref(30);
const isLoading = ref(true);
const avgSleep = ref(0);
const avgEnergy = ref(0);
const avgStress = ref(0);
const patientName = ref("");

const timeRanges = [
  { label: "Últimos 7 dias", days: 7 },
  { label: "Últimos 30 dias", days: 30 },
  { label: "Últimos 90 dias", days: 90 },
];

const moodOptions = [
  { label: "Muito Mal", emoji: "confounded" },
  { label: "Aborrecido", emoji: "confused" },
  { label: "Neutro", emoji: "neutral_face" },
  { label: "Animado", emoji: "blush" },
  { label: "Muito Bem", emoji: "grin" },
];

const getFacebookEmojiUrl = (id) => {
  if (!id) return "";
  const emoji = emojiIndex.findEmoji(id);
  if (!emoji || !emoji.unified) return "";
  return `https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-facebook/15.0.1/img/facebook/64/${emoji.unified.toLowerCase()}.png`;
};

const goBack = () => {
  router.push({ name: "profissional-pacientes" });
};

// --- DADOS PROCESSADOS PARA OS GRÁFICOS ---
const chartData = computed(() => {
  return allData.value.slice(-selectedRange.value);
});

// --- LÓGICA DE BUSCA DE DADOS ---
const fetchReportData = async () => {
  isLoading.value = true;
  try {
    let report;
    if (props.userType === TipoUsuario.Paciente) {
      report = (await api.buscarRelatorio(selectedRange.value)).data;
    } else {
      if (!patientName.value) {
        try {
          const patients = (await api.listarPacientesDoProfissional()).data;
          const patient = patients.find(
            (p) => p.id === Number(props.patientId)
          );
          if (patient) {
            patientName.value = patient.usuario?.nome || patient.nome;
          }
        } catch (e) {
          console.error("Erro ao buscar nome do paciente", e);
        }
      }

      report = (
        await api.buscarRelatorioPacienteDoProfissional(
          selectedRange.value,
          props.patientId
        )
      ).data;
    }

    const formattedData = report.grafico_sono.map((_, i) => ({
      date: report.grafico_sono[i].data,
      valor_sono: report.grafico_sono[i].valor,
      valor_energia: report.grafico_energia[i].valor,
      valor_stress: report.grafico_stress[i].valor,
      humor: report.grafico_sono[i].humor,
      anotacao: report.grafico_sono[i].anotacao,
    }));
    allData.value = formattedData;

    avgSleep.value = (report.media_sono || 0).toFixed(1);
    avgEnergy.value = (report.media_energia || 0).toFixed(1);
    avgStress.value = (report.media_stress || 0).toFixed(1);
  } catch (error) {
    toast.error("Não foi possível carregar os dados do relatório.");
    console.error("Erro ao buscar relatório:", error);
  } finally {
    isLoading.value = false;
  }
};

onMounted(fetchReportData);
watch(selectedRange, fetchReportData);

const sortedChartData = computed(() =>
  [...chartData.value].sort((a, b) => new Date(a.date) - new Date(b.date))
);

// --- OPÇÕES DOS GRÁFICOS ---
const getChartOptions = (title, color, dataKey) => ({
  chart: {
    type: "area",
    height: 350,
    zoom: { enabled: false },
    toolbar: {
      show: true,
      tools: {
        download: true,
        selection: false,
        zoom: false,
        zoomin: false,
        zoomout: false,
        pan: false,
        reset: true,
      },
    },
    animations: {
      enabled: true,
      easing: "easeinout",
      speed: 800,
    },
  },
  colors: [color],
  fill: {
    type: "gradient",
    gradient: {
      shadeIntensity: 1,
      opacityFrom: 0.45,
      opacityTo: 0.05,
      stops: [0, 100],
    },
  },
  dataLabels: { enabled: false },
  stroke: { curve: "smooth", width: 3 },
  xaxis: {
    type: "datetime",
    categories: sortedChartData.value.map((d) => d.date),
    labels: {
      show: true,
      style: { colors: "#9CA3AF", fontSize: "12px" },
      datetimeFormatter: {
        year: "yyyy",
        month: "MMM",
        day: "dd",
        hour: "HH:mm",
      },
    },
    title: {
      text: `Tempo (${
        timeRanges.find((r) => r.days === selectedRange.value)?.label
      })`,
      style: { fontSize: "14px", fontWeight: 500, color: "#6B7280" },
      offsetY: 80,
    },
    axisBorder: { show: false },
    axisTicks: { show: false },
    tooltip: { enabled: false },
  },
  yaxis: {
    title: { text: title, style: { color: "#6B7280", fontWeight: 500 } },
    labels: { style: { colors: "#9CA3AF" } },
  },
  markers: {
    size: 4,
    colors: [color],
    strokeColors: "#fff",
    strokeWidth: 2,
    hover: { size: 6 },
  },
  grid: {
    borderColor: "#F3F4F6",
    strokeDashArray: 4,
    xaxis: { lines: { show: true } },
    yaxis: { lines: { show: true } },
    padding: { bottom: 10 },
  },
  annotations: {
    points: sortedChartData.value.map((point) => ({
      x: new Date(point.date).getTime(),
      y: point[dataKey],
      marker: {
        size: 0,
      },
    })),
  },
  tooltip: {
    theme: "light",
    custom: function ({ series, seriesIndex, dataPointIndex, w }) {
      const pointData = sortedChartData.value[dataPointIndex];
      if (!pointData) return "";
      const seriesName = w.globals.seriesNames[seriesIndex];
      const color = w.globals.colors[seriesIndex];

      const moodId = moodOptions[pointData.humor - 1]?.emoji;
      const moodLabel = moodOptions[pointData.humor - 1]?.label;
      const moodImageUrl = getFacebookEmojiUrl(moodId);

      return `
          <div class="p-3 bg-white border border-gray-100 shadow-lg rounded-xl text-sm">
            <div class="font-bold text-gray-800 mb-2 border-b border-gray-50 pb-1">
              ${new Date(pointData.date).toLocaleDateString("pt-BR", {
                weekday: "short",
                day: "numeric",
                month: "short",
              })}
            </div>
            <div class="flex items-center gap-2 mb-1">
              <span class="w-2 h-2 rounded-full" style="background-color: ${color}"></span>
              <span class="text-gray-600">${seriesName}: <span class="font-bold text-gray-900">${
        series[seriesIndex][dataPointIndex]
      }</span></span>
            </div>
            <div class="flex items-center gap-2 mb-1">
              <span class="w-2 h-2"></span>
              <span class="text-gray-600 flex items-center">Humor: <img src="${moodImageUrl}" width="20" height="20" style="margin: 0 4px; vertical-align: text-bottom;" /> <span class="font-bold text-gray-900">${moodLabel}</span></span>
            </div>
            ${
              pointData.anotacao
                ? `
              <div class="mt-2 pt-2 border-top border-gray-50 max-w-[200px] text-xs text-gray-500 italic">
                "${
                  pointData.anotacao.length > 50
                    ? pointData.anotacao.substring(0, 50) + "..."
                    : pointData.anotacao
                }"
              </div>
            `
                : ""
            }
          </div>
        `;
    },
  },
});

const sleepChartOptions = computed(() =>
  getChartOptions("Horas", "#3B82F6", "valor_sono")
);
const sleepSeries = computed(() => [
  {
    name: "Horas de Sono",
    data: sortedChartData.value.map((d) => d.valor_sono),
  },
]);

const energyChartOptions = computed(() =>
  getChartOptions("Nível (0-10)", "#F59E0B", "valor_energia")
);
const energySeries = computed(() => [
  {
    name: "Nível de Energia",
    data: sortedChartData.value.map((d) => d.valor_energia),
  },
]);

const stressChartOptions = computed(() =>
  getChartOptions("Nível (0-10)", "#EF4444", "valor_stress")
);
const stressSeries = computed(() => [
  {
    name: "Nível de Stress",
    data: sortedChartData.value.map((d) => d.valor_stress),
  },
]);
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
