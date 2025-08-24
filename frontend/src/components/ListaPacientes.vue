<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Meus Pacientes</h1>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="patient in patients" 
        :key="patient.id"
        @click="$emit('view-patient', patient.id)"
        class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 cursor-pointer hover:shadow-lg hover:border-indigo-400 transition-all duration-200"
      >
        <div class="flex items-center space-x-4 mb-4">
          <img :src="patient.avatar" :alt="patient.name" class="w-14 h-14 rounded-full object-cover">
          <div>
            <h3 class="font-semibold text-gray-900 text-lg">{{ patient.name }}</h3>
            <p class="text-sm text-gray-500">{{ patient.age }}</p>
          </div>
        </div>
        <div class="text-sm text-gray-600">
          <p class="font-medium">Foco do tratamento:</p>
          <p>{{ patient.focus }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

defineEmits(['view-patient']);

// Dados fictícios que viriam da API
const patients = ref([
    {
        id: 1,
        name: "Marina",
        age: "30 anos",
        focus: "Ansiedade e estresse crônico.",
        avatar: "https://images.pexels.com/photos/774909/pexels-photo-774909.jpeg?auto=compress&cs=tinysrgb&w=150&h=150&fit=crop&crop=face",
    },
    {
        id: 2,
        name: "Gilberto",
        age: "25 anos",
        focus: "Síndrome de burnout e exaustão emocional.",
        avatar: "https://images.pexels.com/photos/1043474/pexels-photo-1043474.jpeg?auto=compress&cs=tinysrgb&w=150&h=150&fit=crop&crop=face",
    },
    {
        id: 3,
        name: "Ana Clara",
        age: "42 anos",
        focus: "Terapia de casal e comunicação não-violenta.",
        avatar: "https://images.pexels.com/photos/1239291/pexels-photo-1239291.jpeg?auto=compress&cs=tinysrgb&w=150&h=150&fit=crop&crop=face",
    }
]);
</script>
