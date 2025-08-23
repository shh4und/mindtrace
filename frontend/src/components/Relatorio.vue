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
import { ref, computed } from 'vue';

// --- DADOS FICTÍCIOS ---
const moodMapping = { 1: '😖 Muito Mal', 2: '😕 Mal', 3: '😐 Neutro', 4: '😊 Bem', 5: '😁 Muito Bem' };

const generateMockData = (days) => {
  const data = [];
  for (let i = 0; i < days; i++) {
    const date = new Date();
    date.setDate(date.getDate() - i);
    data.push({
      date: date.toISOString().split('T')[0],
      sleep:  Math.floor(Math.random() * 9) + 3, // entre 5 e 9 horas
      energy: Math.floor(Math.random() * 8) + 2, // entre 2 e 9
      stress: Math.floor(Math.random() * 8) + 1, // entre 1 e 8
      mood: moodMapping[Math.floor(Math.random() * 5) + 1],
    });
  }
  return data.reverse(); // Ordem cronológica
};

const allData = ref(generateMockData(90));
const selectedRange = ref(30);

const timeRanges = [
  { label: 'Últimos 7 dias', days: 7 },
  { label: 'Últimos 30 dias', days: 30 },
  { label: 'Últimos 90 dias', days: 90 },
];

const chartData = computed(() => {
  return allData.value.slice(-selectedRange.value);
});

// --- ESTATÍSTICAS ---
const calculateAverage = (key) => {
  if (chartData.value.length === 0) return 'N/A';
  const total = chartData.value.reduce((acc, curr) => acc + curr[key], 0);
  return (total / chartData.value.length).toFixed(1);
};

const avgSleep = computed(() => calculateAverage('sleep'));
const avgEnergy = computed(() => calculateAverage('energy'));
const avgStress = computed(() => calculateAverage('stress'));


// --- OPÇÕES DOS GRÁFICOS ---
const xAxisTitle = computed(() => {
  const rangeInfo = timeRanges.find(r => r.days === selectedRange.value);
  return rangeInfo ? `Tempo (${rangeInfo.label})` : 'Tempo';
});

const baseChartOptions = computed(() => ({
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
    categories: chartData.value.map(d => d.date),
    labels: {
      show: false, // Esconde as labels individuais do eixo X
    },
    title: {
      text: xAxisTitle.value, // Adiciona o título dinâmico
      style: {
        fontSize: '14px',
        fontWeight: 400,
        color: '#6B7280'
      }
    },
    tooltip: {
      enabled: false, // Desativa o tooltip do eixo X para não ser redundante
    }
  },
  tooltip: {
    custom: function ({ series, seriesIndex, dataPointIndex, w }) {
      const pointData = chartData.value[dataPointIndex];
      if (!pointData) return '';

      const value = series[seriesIndex][dataPointIndex];
      const seriesName = w.globals.seriesNames[seriesIndex];

      // Retorna um bloco HTML autossuficiente com estilos inline
      return `
        <div style="padding: 10px 14px; background: #FFF; border: 1px solid #DDD; box-shadow: 0 3px 8px rgba(0,0,0,0.15); border-radius: 6px;">
          <div style="font-weight: 600; color: #333; margin-bottom: 6px;">
            ${new Date(pointData.date).toLocaleDateString('pt-BR', { weekday: 'long', day: 'numeric', month: 'long' })}
          </div>
          <div style="font-size: 13px; color: #555;">
            <span style="display: inline-block; width: 10px; height: 10px; margin-right: 6px; border-radius: 50%; background-color: ${w.globals.colors[seriesIndex]};"></span>
            <span>${seriesName}: <strong>${value}</strong></span>
          </div>
          <div style="font-size: 13px; color: #555; margin-top: 5px;">
            <span style="display: inline-block; width: 10px; height: 10px; margin-right: 6px;"></span>
            <span>Humor: <strong>${pointData.mood}</strong></span>
          </div>
        </div>
      `;
    }
  },
  markers: {
    size: 5,
    hover: { size: 7 }
  },
  grid: {
    borderColor: '#e7e7e7',
    row: { colors: ['#f3f3f3', 'transparent'], opacity: 0.5 },
  },
}));

// Opções e Séries para cada gráfico
const sleepChartOptions = computed(() => ({ ...baseChartOptions.value, yaxis: { title: { text: 'Horas' } }, colors: ['#3B82F6'] }));
const sleepSeries = computed(() => [{ name: 'Horas de Sono', data: chartData.value.map(d => d.sleep) }]);

const energyChartOptions = computed(() => ({ ...baseChartOptions.value, yaxis: { title: { text: 'Nível (0-10)' } }, colors: ['#F59E0B'] }));
const energySeries = computed(() => [{ name: 'Nível de Energia', data: chartData.value.map(d => d.energy) }]);

const stressChartOptions = computed(() => ({ ...baseChartOptions.value, yaxis: { title: { text: 'Nível (0-10)' } }, colors: ['#EF4444'] }));
const stressSeries = computed(() => [{ name: 'Nível de Stress', data: chartData.value.map(d => d.stress) }]);

</script>
