<template>
  <div>
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Resumo</h1>
      <p class="text-gray-600 mt-1">
        Bem-vindo(a) de volta! Aqui está um resumo do seu bem-estar.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-16">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-emerald-600"></div>
    </div>

    <template v-else>
      <section class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <!-- Card: Último Registro de Humor -->
        <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-200 hover:shadow-md transition-shadow">
          <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
            <font-awesome-icon :icon="faFaceSmile" class="w-5 h-5 mr-2 text-emerald-600" />
            Último Registro de Humor
          </h2>
          <div v-if="userLastData.humor" class="flex items-center space-x-4">
            <span class="text-5xl">{{ userLastData.emoji }}</span>
            <div>
              <p class="text-2xl font-bold text-emerald-600">
                {{ userLastData.label }}
              </p>
              <p class="text-gray-500 text-sm">
                {{ formatDate(userLastData.data) }}
              </p>
              <p class="text-gray-600 text-sm mt-1">
                {{ userLastData.anotacao || "Nenhuma anotação." }}
              </p>
            </div>
          </div>
          <div v-else class="text-center py-4">
            <p class="text-gray-500">Nenhum registro de humor ainda.</p>
            <router-link 
              to="/dashboard-paciente/humor" 
              class="text-emerald-600 hover:text-emerald-700 text-sm font-medium mt-2 inline-block"
            >
              Registrar agora →
            </router-link>
          </div>
        </div>

        <!-- Card: Últimas Notificações (placeholder) -->
        <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-200 hover:shadow-md transition-shadow">
          <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
            <font-awesome-icon :icon="faBell" class="w-5 h-5 mr-2 text-amber-500" />
            Últimas Notificações
          </h2>
          <div class="text-center py-6">
            <font-awesome-icon :icon="faBellSlash" class="w-12 h-12 text-gray-300 mb-3" />
            <p class="text-gray-500 text-sm">Nenhuma notificação no momento.</p>
            <p class="text-gray-400 text-xs mt-1">Em breve você receberá alertas e lembretes aqui.</p>
          </div>
        </div>
      </section>

      <!-- Card: Atividades Recentes -->
      <section class="bg-white p-6 rounded-xl shadow-sm border border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
          <font-awesome-icon :icon="faClock" class="w-5 h-5 mr-2 text-emerald-600" />
          Atividades Recentes
        </h2>
        
        <div v-if="atividadesRecentes.length === 0" class="text-center py-8">
          <font-awesome-icon :icon="faListCheck" class="w-12 h-12 text-gray-300 mb-3" />
          <p class="text-gray-500">Nenhuma atividade recente.</p>
          <p class="text-gray-400 text-sm mt-1">Registre seu humor ou responda um questionário para começar.</p>
        </div>

        <ul v-else class="space-y-4">
          <li 
            v-for="(atividade, index) in atividadesRecentes" 
            :key="index"
            class="flex items-start space-x-4 p-3 rounded-lg hover:bg-gray-50 transition-colors"
          >
            <div 
              :class="[
                'shrink-0 w-10 h-10 rounded-full flex items-center justify-center',
                atividade.tipo === 'humor' ? 'bg-emerald-100 text-emerald-600' : 'bg-blue-100 text-blue-600'
              ]"
            >
              <font-awesome-icon 
                :icon="atividade.tipo === 'humor' ? faFaceSmile : faClipboardList" 
                class="w-5 h-5"
              />
            </div>
            <div class="flex-1 min-w-0">
              <p class="font-medium text-gray-900">{{ atividade.titulo }}</p>
              <p class="text-sm text-gray-500">{{ atividade.descricao }}</p>
              <p class="text-xs text-gray-400 mt-1">{{ atividade.dataFormatada }}</p>
            </div>
            <span 
              v-if="atividade.tipo === 'humor'" 
              class="text-2xl"
            >
              {{ atividade.emoji }}
            </span>
          </li>
        </ul>
      </section>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import api from "@/services/api";
import { useToast } from "vue-toastification";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { 
  faFaceSmile, 
  faBell, 
  faBellSlash, 
  faClock, 
  faClipboardList,
  faListCheck
} from "@fortawesome/free-solid-svg-icons";

const toast = useToast();

// --- ESTADO DO COMPONENTE ---
const isLoading = ref(true);
const userLastData = ref({});
const atividadesRecentes = ref([]);

const moodOptions = [
  { label: "Muito Mal", emoji: "😖" },
  { label: "Aborrecido", emoji: "😕" },
  { label: "Neutro", emoji: "😐" },
  { label: "Animado", emoji: "😊" },
  { label: "Muito Bem", emoji: "😁" },
];

// --- FUNÇÕES AUXILIARES ---
const formatDate = (dateString) => {
  if (!dateString) return '';
  return new Date(dateString).toLocaleDateString("pt-BR", {
    weekday: "long",
    day: "numeric",
    month: "long",
  });
};

const formatDateShort = (dateString) => {
  if (!dateString) return '';
  return new Date(dateString).toLocaleDateString("pt-BR", {
    day: "numeric",
    month: "short",
    hour: "2-digit",
    minute: "2-digit"
  });
};

// --- LÓGICA DE BUSCA DE DADOS ---
const fetchSummaryData = async () => {
  isLoading.value = true;
  try {
    // Buscar resumo (último registro de humor)
    const summaryResponse = await api.buscarResumo();
    const summary = summaryResponse.data;
    
    if (summary && summary.humor) {
      summary.emoji = moodOptions[summary.humor - 1]?.emoji || "😐";
      summary.label = moodOptions[summary.humor - 1]?.label || "Neutro";
      userLastData.value = summary;
    }

    // Buscar atividades recentes (mock + dados reais combinados)
    await fetchAtividadesRecentes(summary);

  } catch (error) {
    toast.error("Não foi possível carregar os dados do resumo.");
    console.error(error);
  } finally {
    isLoading.value = false;
  }
};

const fetchAtividadesRecentes = async (summaryData) => {
  const atividades = [];

  // Adicionar último registro de humor como atividade
  if (summaryData && summaryData.humor) {
    atividades.push({
      tipo: 'humor',
      titulo: 'Registro de Humor',
      descricao: `Você registrou: ${moodOptions[summaryData.humor - 1]?.label || 'Humor'}`,
      data: summaryData.data,
      dataFormatada: formatDateShort(summaryData.data),
      emoji: moodOptions[summaryData.humor - 1]?.emoji || '😐'
    });
  }

  // Tentar buscar questionários respondidos
  try {
    const atribuicoesResponse = await api.listarAtribuicoesPaciente();
    const atribuicoes = atribuicoesResponse.data || [];
    
    // Filtrar apenas respondidos e pegar os mais recentes
    const respondidos = atribuicoes
      .filter(a => a.respondido_em)
      .sort((a, b) => new Date(b.respondido_em) - new Date(a.respondido_em))
      .slice(0, 2);

    respondidos.forEach(q => {
      atividades.push({
        tipo: 'questionario',
        titulo: `Questionário "${q.instrumento?.nome || 'Avaliação'}" respondido`,
        descricao: q.instrumento?.descricao || 'Questionário completado',
        data: q.respondido_em,
        dataFormatada: formatDateShort(q.respondido_em)
      });
    });
  } catch (error) {
    console.log('Questionários não disponíveis:', error);
  }

  // Ordenar por data e limitar a 3
  atividadesRecentes.value = atividades
    .sort((a, b) => new Date(b.data) - new Date(a.data))
    .slice(0, 3);
};

onMounted(fetchSummaryData);
</script>
