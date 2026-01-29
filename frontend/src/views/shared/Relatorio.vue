<template>
  <div class="flex-1 p-4 md:p-8 w-full">
    <!-- Botão voltar para profissional -->
    <button
      v-if="userType === TipoUsuario.Profissional && patientId"
      @click="goBack"
      class="mb-6 flex items-center text-sm font-medium text-rose-600 hover:text-rose-800 transition-colors"
      aria-label="Voltar para a lista de pacientes"
    >
      <font-awesome-icon :icon="faArrowLeft" class="mr-2" aria-hidden="true" />
      Voltar para a lista de pacientes
    </button>

    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Relatórios de Bem-Estar</h1>
      <p class="text-gray-600 mt-1">
        Analise suas tendências de humor, sono e energia ao longo do tempo.
      </p>
    </header>

    <!-- Filtros de Período -->
    <div
      class="mb-8 flex justify-center md:justify-start space-x-2"
      role="group"
      aria-label="Filtro de período"
    >
      <button
        v-for="range in timeRanges"
        :key="range.days"
        @click="selectedRange = range.days"
        :aria-pressed="selectedRange === range.days"
        :class="[
          'px-4 py-2 rounded-md font-medium text-sm transition-colors',
          selectedRange === range.days
            ? 'bg-emerald-50 text-emerald-700 shadow-sm'
            : 'bg-white text-gray-600 hover:bg-gray-100 shadow-sm border border-gray-200 hover:text-gray-900',
        ]"
      >
        {{ range.label }}
      </button>
    </div>

    <!-- Cards de Estatísticas -->
    <section class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8" aria-label="Estatísticas resumidas">
      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <div class="flex items-center justify-between mb-2">
          <h3 class="text-gray-500 text-sm font-medium">Média de Sono</h3>
          <Emoji :data="emojiIndex" emoji="sleeping" :size="24" set="facebook" />
        </div>
        <p class="text-3xl font-bold text-blue-600">
          {{ avgSleep }} <span class="text-lg font-medium">horas/noite</span>
        </p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <div class="flex items-center justify-between mb-2">
          <h3 class="text-gray-500 text-sm font-medium">Média de Energia</h3>
          <Emoji :data="emojiIndex" emoji="zap" :size="24" set="facebook" />
        </div>
        <p class="text-3xl font-bold text-amber-600">
          {{ avgEnergy }} <span class="text-lg font-medium">/ 10</span>
        </p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <div class="flex items-center justify-between mb-2">
          <h3 class="text-gray-500 text-sm font-medium">Média de Stress</h3>
          <Emoji :data="emojiIndex" emoji="triumph" :size="24" set="facebook" />
        </div>
        <p class="text-3xl font-bold text-red-600">
          {{ avgStress }} <span class="text-lg font-medium">/ 10</span>
        </p>
      </div>
    </section>

    <!-- Gráficos -->
    <section class="space-y-8" aria-label="Gráficos de tendências">
      <div class="bg-white p-4 md:p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="font-semibold text-lg text-gray-900 mb-4">Horas de Sono</h3>
        <apexchart
          type="area"
          height="350"
          :options="sleepChartOptions"
          :series="sleepSeries"
        ></apexchart>
      </div>
      <div class="bg-white p-4 md:p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="font-semibold text-lg text-gray-900">Nível de Energia</h3>
        <apexchart
          type="area"
          height="350"
          :options="energyChartOptions"
          :series="energySeries"
        ></apexchart>
      </div>
      <div class="bg-white p-4 md:p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="font-semibold text-lg text-gray-900">Nível de Stress</h3>
        <apexchart
          type="area"
          height="350"
          :options="stressChartOptions"
          :series="stressSeries"
        ></apexchart>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/services/api';
import { useToast } from 'vue-toastification';
import { TipoUsuario } from '@/types/usuario.js';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';

import data from 'emoji-mart-vue-fast/data/all.json';
import { EmojiIndex, Emoji } from 'emoji-mart-vue-fast/src';
import 'emoji-mart-vue-fast/css/emoji-mart.css';

const emojiIndex = new EmojiIndex(data);

const props = defineProps({
  patientId: {
    type: [Number, String],
    default: null, // ID do paciente, usado pelo profissional
  },
  userType: {
    type: String,
    default: TipoUsuario.Paciente,
    validator: (value) => [TipoUsuario.Paciente, TipoUsuario.Profissional].includes(value),
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

const timeRanges = [
  { label: 'Últimos 7 dias', days: 7 },
  { label: 'Últimos 30 dias', days: 30 },
  { label: 'Últimos 90 dias', days: 90 },
];

const moodOptions = [
  { label: 'Muito Mal', emoji: 'confounded' },
  { label: 'Aborrecido', emoji: 'confused' },
  { label: 'Neutro', emoji: 'neutral_face' },
  { label: 'Animado', emoji: 'blush' },
  { label: 'Muito Bem', emoji: 'grin' },
];

// Helper para obter URL da imagem do emoji do Facebook
const getFacebookEmojiUrl = (id) => {
  if (!id) return '';
  const emoji = emojiIndex.findEmoji(id);
  if (!emoji || !emoji.unified) return '';
  // Constrói a URL usando um CDN confiável para emoji-datasource-apple
  // emoji.unified geralmente é uppercase, ex: '1F600'
  return `https://cdnjs.cloudflare.com/ajax/libs/emoji-datasource-facebook/15.0.1/img/facebook/64/${emoji.unified.toLowerCase()}.png`;
};

// Navegação
const goBack = () => {
  router.push({ name: 'profissional-pacientes' });
};

// --- DADOS PROCESSADOS PARA OS GRÁFICOS ---
const chartData = computed(() => {
  // No futuro, a API poderia retornar os dados já filtrados.
  // Por agora, filtramos no frontend.
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
      report = (
        await api.buscarRelatorioPacienteDoProfissional(selectedRange.value, props.patientId)
      ).data;
    }

    // Transforma os dados da API para o formato que os gráficos precisam
    const formattedData = report.grafico_sono.map((_, i) => ({
      date: report.grafico_sono[i].data,
      valor_sono: report.grafico_sono[i].valor,
      valor_energia: report.grafico_energia[i].valor,
      valor_stress: report.grafico_stress[i].valor,
      humor: report.grafico_sono[i].humor, // Assumindo que o humor é o mesmo para todos os pontos do dia
      anotacao: report.grafico_sono[i].anotacao,
    }));
    allData.value = formattedData;

    avgSleep.value = (report.media_sono || 0).toFixed(1);
    avgEnergy.value = (report.media_energia || 0).toFixed(1);
    avgStress.value = (report.media_stress || 0).toFixed(1);
  } catch (error) {
    toast.error('Não foi possível carregar os dados do relatório.');
    console.error('Erro ao buscar relatório:', error);
  } finally {
    isLoading.value = false;
  }
};

onMounted(fetchReportData);
watch(selectedRange, fetchReportData);

const sortedChartData = computed(() =>
  [...chartData.value].sort((a, b) => new Date(a.date) - new Date(b.date))
);

// --- OPÇÕES DOS GRÁFICOS (adaptado para reatividade com ref) ---
const getChartOptions = (title, color, dataKey) => ({
  chart: {
    type: 'area',
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
      easing: 'easeinout',
      speed: 800,
    },
  },
  colors: [color],
  fill: {
    type: 'gradient',
    gradient: {
      shadeIntensity: 1,
      opacityFrom: 0.45,
      opacityTo: 0.05,
      stops: [0, 100],
    },
  },
  dataLabels: { enabled: false },
  stroke: { curve: 'smooth', width: 3 },
  xaxis: {
    type: 'datetime',
    categories: sortedChartData.value.map((d) => d.date),
    labels: {
      show: true,
      style: { colors: '#9CA3AF', fontSize: '12px' },
      datetimeFormatter: {
        year: 'yyyy',
        month: 'MMM',
        day: 'dd',
        hour: 'HH:mm',
      },
    },
    title: {
      text: `Tempo (${timeRanges.find((r) => r.days === selectedRange.value)?.label})`,
      style: { fontSize: '14px', fontWeight: 500, color: '#6B7280' },
      offsetY: 80,
    },
    axisBorder: { show: false },
    axisTicks: { show: false },
    tooltip: { enabled: false },
  },
  yaxis: {
    title: { text: title, style: { color: '#6B7280', fontWeight: 500 } },
    labels: { style: { colors: '#9CA3AF' } },
  },
  markers: {
    size: 4,
    colors: [color],
    strokeColors: '#fff',
    strokeWidth: 2,
    hover: { size: 6 },
  },
  grid: {
    borderColor: '#F3F4F6',
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
      // image: {
      //   path: getFacebookEmojiUrl(moodOptions[point.humor - 1]?.emoji),
      //   width: 24,
      //   height: 24,
      //   offsetY: -15,
      // },
    })),
  },
  tooltip: {
    theme: 'light',
    custom: function ({ series, seriesIndex, dataPointIndex, w }) {
      const pointData = sortedChartData.value[dataPointIndex];
      if (!pointData) return '';
      const seriesName = w.globals.seriesNames[seriesIndex];
      const color = w.globals.colors[seriesIndex];

      const moodId = moodOptions[pointData.humor - 1]?.emoji;
      const moodLabel = moodOptions[pointData.humor - 1]?.label;
      const moodImageUrl = getFacebookEmojiUrl(moodId);

      return `
          <div class="p-3 bg-white border border-gray-100 shadow-lg rounded-lg text-sm">
            <div class="font-bold text-gray-800 mb-2 border-b border-gray-50 pb-1">
              ${new Date(pointData.date).toLocaleDateString('pt-BR', { weekday: 'short', day: 'numeric', month: 'short' })}
            </div>
            <div class="flex items-center gap-2 mb-1">
              <span class="w-2 h-2 rounded-full" style="background-color: ${color}"></span>
              <span class="text-gray-600">${seriesName}: <span class="font-bold text-gray-900">${series[seriesIndex][dataPointIndex]}</span></span>
            </div>
            <div class="flex items-center gap-2 mb-1">
              <span class="w-2 h-2"></span>
              <span class="text-gray-600 flex items-center">Humor: <img src="${moodImageUrl}" width="20" height="20" style="margin: 0 4px; vertical-align: text-bottom;" /> <span class="font-bold text-gray-900">${moodLabel}</span></span>
            </div>
            ${
              pointData.anotacao
                ? `
              <div class="mt-2 pt-2 border-top border-gray-50 max-w-[200px] text-xs text-gray-500 italic">
                "${pointData.anotacao.length > 50 ? pointData.anotacao.substring(0, 50) + '...' : pointData.anotacao}"
              </div>
            `
                : ''
            }
          </div>
        `;
    },
  },
});

const sleepChartOptions = computed(() => getChartOptions('Horas', '#3B82F6', 'valor_sono'));
const sleepSeries = computed(() => [
  { name: 'Horas de Sono', data: sortedChartData.value.map((d) => d.valor_sono) },
]);

const energyChartOptions = computed(() =>
  getChartOptions('Nível (0-10)', '#F59E0B', 'valor_energia')
);
const energySeries = computed(() => [
  { name: 'Nível de Energia', data: sortedChartData.value.map((d) => d.valor_energia) },
]);

const stressChartOptions = computed(() =>
  getChartOptions('Nível (0-10)', '#EF4444', 'valor_stress')
);
const stressSeries = computed(() => [
  { name: 'Nível de Stress', data: sortedChartData.value.map((d) => d.valor_stress) },
]);
</script>
