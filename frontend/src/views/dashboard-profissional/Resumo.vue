<template>
  <div>
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Resumo</h1>
      <p class="text-gray-600 mt-1">
        Bem-vindo(a) de volta! Aqui está uma visão geral do seu trabalho.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-16">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-rose-500"></div>
    </div>

    <template v-else>
      <!-- Cards de Estatísticas -->
      <section class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div 
          v-for="stat in estatisticas" 
          :key="stat.id"
          class="bg-white p-6 rounded-xl shadow-sm border border-gray-200 hover:shadow-md transition-shadow cursor-pointer"
          @click="stat.route && navigateTo(stat.route)"
        >
          <div class="flex items-center justify-between mb-3">
            <div :class="[stat.bgColor, 'p-3 rounded-lg']">
              <font-awesome-icon :icon="stat.icon" :class="[stat.iconColor, 'w-6 h-6']" />
            </div>
            <span 
              v-if="stat.trend" 
              :class="[
                'text-xs font-medium px-2 py-1 rounded-full',
                stat.trend > 0 ? 'bg-green-100 text-green-700' : 'bg-gray-100 text-gray-600'
              ]"
            >
              {{ stat.trend > 0 ? '+' : '' }}{{ stat.trend }}%
            </span>
          </div>
          <p class="text-3xl font-bold text-gray-900">{{ stat.value }}</p>
          <p class="text-sm text-gray-500 mt-1">{{ stat.label }}</p>
        </div>
      </section>

      <!-- Acesso Rápido -->
      <section class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <!-- Ações Rápidas -->
        <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
            <font-awesome-icon :icon="faBolt" class="w-5 h-5 mr-2 text-amber-500" />
            Ações Rápidas
          </h2>
          <div class="grid grid-cols-2 gap-3">
            <button
              v-for="action in acoesRapidas"
              :key="action.id"
              @click="navigateTo(action.route)"
              :class="[
                'flex flex-col items-center justify-center p-4 rounded-lg border-2 border-dashed transition-all',
                'hover:border-rose-400 hover:bg-rose-50 border-gray-200'
              ]"
            >
              <font-awesome-icon :icon="action.icon" class="w-6 h-6 text-rose-500 mb-2" />
              <span class="text-sm font-medium text-gray-700">{{ action.label }}</span>
            </button>
          </div>
        </div>

        <!-- Últimos Pacientes Atendidos -->
        <div class="bg-white p-6 rounded-xl shadow-sm border border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
            <font-awesome-icon :icon="faUsers" class="w-5 h-5 mr-2 text-rose-500" />
            Últimos Pacientes
          </h2>
          
          <div v-if="ultimosPacientes.length === 0" class="text-center py-6">
            <font-awesome-icon :icon="faUserPlus" class="w-10 h-10 text-gray-300 mb-2" />
            <p class="text-gray-500 text-sm">Nenhum paciente vinculado ainda.</p>
            <button 
              @click="navigateTo('profissional-convite')"
              class="text-rose-600 hover:text-rose-700 text-sm font-medium mt-2"
            >
              Gerar convite →
            </button>
          </div>

          <ul v-else class="space-y-3">
            <li 
              v-for="(paciente, index) in ultimosPacientes" 
              :key="paciente.id"
              @click="viewPatientReport(paciente.id)"
              class="flex items-center space-x-3 p-3 rounded-lg hover:bg-gray-50 cursor-pointer transition-colors"
            >
              <div :class="[getAvatarColor(index), 'w-10 h-10 rounded-full flex items-center justify-center']">
                <font-awesome-icon :icon="faUser" class="w-5 h-5 text-white" />
              </div>
              <div class="flex-1 min-w-0">
                <p class="font-medium text-gray-900 truncate">{{ paciente.name }}</p>
                <p class="text-xs text-gray-500">{{ paciente.age }}</p>
              </div>
              <font-awesome-icon :icon="faChevronRight" class="w-4 h-4 text-gray-400" />
            </li>
          </ul>

          <button 
            v-if="ultimosPacientes.length > 0"
            @click="navigateTo('profissional-pacientes')"
            class="w-full mt-4 text-center text-sm text-rose-600 hover:text-rose-700 font-medium"
          >
            Ver todos os pacientes →
          </button>
        </div>
      </section>

      <!-- Questionários Pendentes -->
      <section class="bg-white p-6 rounded-xl shadow-sm border border-gray-200">
        <h2 class="text-lg font-semibold text-gray-900 mb-4 flex items-center">
          <font-awesome-icon :icon="faClipboardList" class="w-5 h-5 mr-2 text-rose-500" />
          Questionários Atribuídos Recentemente
        </h2>

        <div v-if="questionariosRecentes.length === 0" class="text-center py-8">
          <font-awesome-icon :icon="faClipboardList" class="w-12 h-12 text-gray-300 mb-3" />
          <p class="text-gray-500">Nenhum questionário atribuído recentemente.</p>
        </div>

        <div v-else class="overflow-x-auto">
          <table class="w-full">
            <thead>
              <tr class="text-left text-xs text-gray-500 uppercase tracking-wider border-b">
                <th class="pb-3 font-medium">Paciente</th>
                <th class="pb-3 font-medium">Questionário</th>
                <th class="pb-3 font-medium">Status</th>
                <th class="pb-3 font-medium">Data</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-100">
              <tr 
                v-for="q in questionariosRecentes" 
                :key="q.id"
                class="hover:bg-gray-50 cursor-pointer"
                @click="q.respondido && viewResponses(q.id)"
              >
                <td class="py-3">
                  <span class="font-medium text-gray-900">{{ q.pacienteNome }}</span>
                </td>
                <td class="py-3 text-gray-600">{{ q.questionarioNome }}</td>
                <td class="py-3">
                  <span 
                    :class="[
                      'px-2 py-1 text-xs font-medium rounded-full',
                      q.respondido ? 'bg-green-100 text-green-700' : 'bg-amber-100 text-amber-700'
                    ]"
                  >
                    {{ q.respondido ? 'Respondido' : 'Pendente' }}
                  </span>
                </td>
                <td class="py-3 text-gray-500 text-sm">{{ q.dataFormatada }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <button 
          v-if="questionariosRecentes.length > 0"
          @click="navigateTo('profissional-questionarios-atribuidos')"
          class="w-full mt-4 text-center text-sm text-rose-600 hover:text-rose-700 font-medium"
        >
          Ver todos os questionários →
        </button>
      </section>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import api from '@/services/api';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faUsers,
  faUser,
  faUserPlus,
  faClipboardList,
  faEnvelope,
  faChartLine,
  faBolt,
  faChevronRight,
  faCalendarCheck,
  faUserCheck
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const toast = useToast();

const isLoading = ref(true);
const estatisticas = ref([]);
const ultimosPacientes = ref([]);
const questionariosRecentes = ref([]);

// Cores para avatares
const avatarColors = [
  'bg-blue-500',
  'bg-green-500',
  'bg-purple-500',
  'bg-red-500',
  'bg-yellow-500',
  'bg-indigo-500',
  'bg-pink-500',
  'bg-teal-500',
];

const getAvatarColor = (index) => avatarColors[index % avatarColors.length];

// Ações rápidas
const acoesRapidas = [
  { id: 'pacientes', label: 'Ver Pacientes', icon: faUsers, route: 'profissional-pacientes' },
  { id: 'convite', label: 'Gerar Convite', icon: faEnvelope, route: 'profissional-convite' },
  { id: 'questionarios', label: 'Questionários', icon: faClipboardList, route: 'profissional-questionarios-atribuidos' },
  { id: 'relatorios', label: 'Relatórios', icon: faChartLine, route: 'profissional-relatorios' }
];

// Navegação
const navigateTo = (routeName) => {
  router.push({ name: routeName });
};

const viewPatientReport = (patientId) => {
  router.push({ name: 'profissional-paciente-relatorio', params: { patientId } });
};

const viewResponses = (atribuicaoId) => {
  router.push({ name: 'profissional-visualizar-respostas', params: { atribuicaoId } });
};

// Calcular idade
const calculateAge = (birthdate) => {
  if (!birthdate) return '';
  const birthDate = new Date(birthdate);
  const today = new Date();
  let age = today.getFullYear() - birthDate.getFullYear();
  const m = today.getMonth() - birthDate.getMonth();
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
};

// Formatar data
const formatDate = (dateString) => {
  if (!dateString) return '';
  return new Date(dateString).toLocaleDateString('pt-BR', {
    day: '2-digit',
    month: 'short'
  });
};

// Buscar dados do resumo
const fetchResumoData = async () => {
  isLoading.value = true;
  try {
    // Buscar pacientes
    let totalPacientes = 0;
    try {
      const pacientesResponse = await api.listarPacientesDoProfissional();
      const pacientes = pacientesResponse.data || [];
      totalPacientes = pacientes.length;
      
      // Pegar últimos 3 pacientes
      ultimosPacientes.value = pacientes.slice(0, 3).map(p => ({
        id: p.id,
        name: p.usuario?.nome || 'Paciente',
        age: `${calculateAge(p.data_nascimento)} anos`
      }));
    } catch (e) {
      console.log('Erro ao buscar pacientes:', e);
    }

    // Buscar questionários atribuídos
    let totalQuestionarios = 0;
    let pendentes = 0;
    try {
      const questionariosResponse = await api.listarAtribuicoesProfissional();
      const questionarios = questionariosResponse.data || [];
      totalQuestionarios = questionarios.length;
      pendentes = questionarios.filter(q => !q.respondido_em).length;

      // Pegar últimos 5 questionários
      questionariosRecentes.value = questionarios
        .sort((a, b) => new Date(b.atribuido_em) - new Date(a.atribuido_em))
        .slice(0, 5)
        .map(q => ({
          id: q.id,
          pacienteNome: q.paciente?.usuario?.nome || 'Paciente',
          questionarioNome: q.instrumento?.nome || 'Questionário',
          respondido: !!q.respondido_em,
          dataFormatada: formatDate(q.atribuido_em)
        }));
    } catch (e) {
      console.log('Erro ao buscar questionários:', e);
    }

    // Montar estatísticas
    estatisticas.value = [
      {
        id: 'pacientes',
        label: 'Pacientes Vinculados',
        value: totalPacientes,
        icon: faUsers,
        bgColor: 'bg-blue-100',
        iconColor: 'text-blue-600',
        route: 'profissional-pacientes'
      },
      {
        id: 'questionarios',
        label: 'Questionários Atribuídos',
        value: totalQuestionarios,
        icon: faClipboardList,
        bgColor: 'bg-purple-100',
        iconColor: 'text-purple-600',
        route: 'profissional-questionarios-atribuidos'
      },
      {
        id: 'pendentes',
        label: 'Pendentes de Resposta',
        value: pendentes,
        icon: faCalendarCheck,
        bgColor: 'bg-amber-100',
        iconColor: 'text-amber-600',
        route: 'profissional-questionarios-atribuidos'
      },
      {
        id: 'respondidos',
        label: 'Respondidos',
        value: totalQuestionarios - pendentes,
        icon: faUserCheck,
        bgColor: 'bg-green-100',
        iconColor: 'text-green-600',
        route: 'profissional-questionarios-atribuidos'
      }
    ];

  } catch (error) {
    console.error('Erro ao carregar resumo:', error);
    toast.error('Erro ao carregar dados do resumo.');
  } finally {
    isLoading.value = false;
  }
};

onMounted(fetchResumoData);
</script>
