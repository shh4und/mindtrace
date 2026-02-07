<template>
  <div class="max-w-4xl mx-auto p-4 md:p-8">
    <header class="mb-10 text-center">
      <div
        class="inline-flex items-center justify-center space-x-2 bg-emerald-50 text-emerald-800 px-4 py-1.5 rounded-full text-sm font-medium mb-4 shadow-sm"
      >
        <font-awesome-icon :icon="faLink" class="h-4 w-4" />
        <span>Conexão</span>
      </div>
      <h1
        class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
      >
        Vincular a um Profissional
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Insira o token de convite fornecido pelo seu profissional ou clique no link recebido por e-mail.
      </p>
    </header>

    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 max-w-lg mx-auto">
      <div class="space-y-6">
        <BaseInput
          v-model="token"
          label="Código do Convite"
          placeholder="Cole o token aqui..."
          :disabled="loading"
          @input="verificarTokenManualmente"
        />

        <div
          v-if="dadosConvite"
          class="bg-emerald-50 p-5 rounded-2xl border border-emerald-200 flex items-center gap-4"
        >
          <div class="p-3 bg-emerald-100 rounded-xl text-emerald-600 shrink-0">
            <font-awesome-icon :icon="faUserDoctor" class="w-5 h-5" />
          </div>
          <div>
            <p class="text-sm text-emerald-700 font-medium">
              Convite encontrado:
            </p>
            <p class="text-lg font-bold text-emerald-900">
              {{ dadosConvite.nome_profissional }}
            </p>
            <p class="text-sm text-emerald-600">
              {{ dadosConvite.especialidade }}
            </p>
          </div>
        </div>

        <p v-if="erro" class="text-red-500 text-sm font-medium">{{ erro }}</p>

        <BaseButton
          variant="emerald"
          full-width
          :loading="loadingVinculo"
          :disabled="!dadosConvite"
          @click="bindWithToken"
        >
          Confirmar Vínculo
        </BaseButton>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import api from "@/services/api";
import { useToast } from "vue-toastification";
import { BaseInput, BaseButton } from "@/components/ui";
import { debounce } from "lodash";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faUserDoctor, faLink } from "@fortawesome/free-solid-svg-icons";

const route = useRoute();
const router = useRouter();
const toast = useToast();

const token = ref("");
const dadosConvite = ref(null);
const erro = ref("");
const loading = ref(false);
const loadingVinculo = ref(false);

onMounted(() => {
  const tokenUrl = route.query.token;
  if (tokenUrl) {
    token.value = tokenUrl;
  }
});

watch(token, () => {
  verificarTokenManualmente();
});

// Função que consulta a API para ver de quem é o convite
const buscarInfoToken = async (tokenValue) => {
  if (!tokenValue || tokenValue.length < 10) return;

  loading.value = true;
  erro.value = "";
  dadosConvite.value = null;

  try {
    const response = await api.consultarToken(tokenValue);
    dadosConvite.value = response.data;
  } catch (err) {
    erro.value = "Código inválido ou expirado.";
    dadosConvite.value = null;
  } finally {
    loading.value = false;
  }
};

const verificarTokenManualmente = debounce(() => {
  buscarInfoToken(token.value);
}, 500);

const bindWithToken = async () => {
  if (!token.value.trim()) {
    return;
  }

  loadingVinculo.value = true;
  try {
    await api.vincularComToken(token.value.trim());
    toast.success(
      `Você agora está vinculado a ${dadosConvite.value.nome_profissional}`
    );
    router.push({ name: "paciente-profissionais" });
  } catch (error) {
    const errorMessage = error.response?.data?.erro || "Erro ao vincular.";
    toast.error(errorMessage);
    console.error("Erro ao vincular com token:", error);
  } finally {
    loadingVinculo.value = false;
  }
};
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
