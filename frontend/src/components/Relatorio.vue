<template>
  <div class="flex-1 p-4 md:p-8 w-full">
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Relatórios de Bem-Estar</h1>
      <p class="text-gray-600 mt-1">Analise suas tendências de humor, sono e energia ao longo do tempo.</p>
    </header>

    <!-- Filtros de Período -->
    <div class="mb-8 flex justify-center md:justify-start space-x-2">
      <button v-for="range in timeRanges" :key="range.days" @click="selectedRange = range.days"
        :class="['px-4 py-2 rounded-md font-medium text-sm transition-colors', selectedRange === range.days ? 'bg-emerald-600 text-white shadow' : 'bg-white text-gray-700 hover:bg-gray-100']">
        {{ range.label }}
      </button>
    </div>

    <!-- Cards de Estatísticas -->
    <section class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="text-gray-500 text-sm font-medium mb-2">Média de Sono</h3>
        <p class="text-3xl font-bold text-blue-600">{{ avgSleep }} <span class="text-lg font-medium">horas/noite</span>
        </p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="text-gray-500 text-sm font-medium mb-2">Média de Energia</h3>
        <p class="text-3xl font-bold text-amber-600">{{ avgEnergy }} <span class="text-lg font-medium">/ 10</span></p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="text-gray-500 text-sm font-medium mb-2">Média de Stress</h3>
        <p class="text-3xl font-bold text-red-600">{{ avgStress }} <span class="text-lg font-medium">/ 10</span></p>
      </div>
    </section>

    <!-- Gráficos -->
    <section class="space-y-8">
      <div class="bg-white p-4 md:p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="font-semibold text-lg text-gray-900 mb-4">Horas de Sono</h3>
        <apexchart type="area" height="350" :options="sleepChartOptions" :series="sleepSeries"></apexchart>
      </div>
      <div class="bg-white p-4 md:p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="font-semibold text-lg text-gray-900">Nível de Energia</h3>
        <apexchart type="area" height="350" :options="energyChartOptions" :series="energySeries"></apexchart>
      </div>
      <div class="bg-white p-4 md:p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 class="font-semibold text-lg text-gray-900">Nível de Stress</h3>
        <apexchart type="area" height="350" :options="stressChartOptions" :series="stressSeries"></apexchart>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
import api from '../services/api';
import { useToast } from 'vue-toastification';

const props = defineProps({
  patientId: {
    type: Number,
    default: null, // ID do paciente, usado pelo profissional
  }
});

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
  {label: 'Muito Mal', emoji: '😖' },
  {label: 'Aborrecido', emoji: '😕' },
  {label: 'Neutro', emoji: '😐' },
  {label: 'Animado', emoji: '😊' },
  {label: 'Muito Bem', emoji: '😁' },
];

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
    // TODO: A API de relatório precisa aceitar um ID de paciente para o profissional
    const response = await api.buscarRelatorio(selectedRange.value);
    const report = response.data;

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

// --- OPÇÕES DOS GRÁFICOS (adaptado para reatividade com ref) ---
const getChartOptions = (title) => ({
  chart: {
    type: 'area',
    height: 350,
    zoom: { enabled: false },
    toolbar: { show: true, tools: { download: true, selection: false, zoom: false, zoomin: false, zoomout: false, pan: false, reset: true } },
  },
  dataLabels: { enabled: false },
  stroke: { curve: 'smooth', width: 3 },
  xaxis: {
    type: 'datetime',
    categories: sortedChartData.value.map(d => d.date),
    labels: { show: false },
    title: { text: `Tempo (${timeRanges.find(r => r.days === selectedRange.value)?.label})`, style: { fontSize: '14px', fontWeight: 400, color: '#6B7280' } },
    tooltip: { enabled: false },
  },
  yaxis: { title: { text: title } },
  markers: { size: 5, hover: { size: 7 } },
  grid: { borderColor: '#e7e7e7', row: { colors: ['#f3f3f3', 'transparent'], opacity: 0.5 } },
  tooltip: {
    custom: function ({ series, seriesIndex, dataPointIndex, w }) {
        const pointData = sortedChartData.value[dataPointIndex];
        if (!pointData) return '';
        const seriesName = w.globals.seriesNames[seriesIndex];
        return `
          <div style="padding: 10px 14px; background: #FFF; border: 1px solid #DDD; box-shadow: 0 3px 8px rgba(0,0,0,0.15); border-radius: 6px;">
            <div style="font-weight: 600; color: #333; margin-bottom: 6px;">
              ${new Date(pointData.date).toLocaleDateString('pt-BR', { weekday: 'long', day: 'numeric', month: 'long' })}
            </div>
            <div style="font-size: 13px; color: #555;">
              <span style="display: inline-block; width: 10px; height: 10px; margin-right: 6px; border-radius: 50%; background-color: ${w.globals.colors[seriesIndex]};"></span>
              <span>${seriesName}: <strong>${series[seriesIndex][dataPointIndex]}</strong></span>
            </div>
            <div style="font-size: 13px; color: #555; margin-top: 5px;">
              <span style="display: inline-block; width: 10px; height: 10px; margin-right: 6px;"></span>
              <span>Humor: <strong>${moodOptions[pointData.humor-1].label} - ${moodOptions[pointData.humor-1].emoji}</strong></span>
            </div>
          </div>
        `;
    }
  },
});

const sleepChartOptions = computed(() => ({ ...getChartOptions('Horas'), colors: ['#3B82F6'] }));
const sleepSeries = computed(() => [{ name: 'Horas de Sono', data: sortedChartData.value.map(d => d.valor_sono) }]);

const energyChartOptions = computed(() => ({ ...getChartOptions('Nível (0-10)'), colors: ['#F59E0B'] }));
const energySeries = computed(() => [{ name: 'Nível de Energia', data: sortedChartData.value.map(d => d.valor_energia) }]);

const stressChartOptions = computed(() => ({ ...getChartOptions('Nível (0-10)'), colors: ['#EF4444'] }));
const stressSeries = computed(() => [{ name: 'Nível de Stress', data: sortedChartData.value.map(d => d.valor_stress) }]);

</script>
