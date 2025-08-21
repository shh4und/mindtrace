<template>
  <div class="bg-gray-50 text-gray-800">
    <header class="bg-white shadow-md sticky top-0 z-50">
      <nav class="container mx-auto px-6 py-4 flex justify-between items-center">
  <router-link to="/" class="text-3xl font-bold text-emerald-600">MindTrace</router-link>

        <div class="md:hidden">
          <button @click="toggleMenu" class="text-gray-600 focus:outline-none">
            <svg class="w-7 h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
            </svg>
          </button>
        </div>

        <ul class="hidden md:flex space-x-8 items-center text-lg">
          <li><a href="#hero" class="hover:text-emerald-600 transition duration-300">Início</a></li>
          <li><a href="#features" class="hover:text-emerald-600 transition duration-300">Recursos</a></li>
          <li><a href="#knowledge" class="hover:text-emerald-600 transition duration-300">Conteúdo</a></li>
          <li><a href="#testimonials" class="hover:text-emerald-600 transition duration-300">Depoimentos</a></li>
          <li><a href="#contact" class="hover:text-emerald-600 transition duration-300">Contato</a></li>
          <li><router-link to="/login" class="bg-emerald-600 text-white px-6 py-1 rounded-full hover:bg-emerald-700 transition duration-300 font-semibold shadow-lg">Login</router-link></li>
        </ul>
      </nav>

      <div :class="['md:hidden', { 'hidden': !isMenuOpen, 'active': isMenuOpen }]" class="bg-white shadow-lg">
        <ul class="px-6 py-4 space-y-4 text-center">
          <li><a href="#hero" @click="closeMenu" class="block text-gray-800 hover:bg-blue-50 p-3 rounded-lg">Início</a></li>
          <li><a href="#features" @click="closeMenu" class="block text-gray-800 hover:bg-blue-50 p-3 rounded-lg">Recursos</a></li>
          <li><a href="#knowledge" @click="closeMenu" class="block text-gray-800 hover:bg-blue-50 p-3 rounded-lg">Conteúdo</a></li>
          <li><a href="#testimonials" @click="closeMenu" class="block text-gray-800 hover:bg-blue-50 p-3 rounded-lg">Depoimentos</a></li>
          <li><a href="#contact" @click="closeMenu" class="block text-gray-800 hover:bg-blue-50 p-3 rounded-lg">Contato</a></li>
          <li><router-link to="/login" @click="closeMenu" class="block bg-emerald-600 text-white p-3 rounded-lg mt-4 font-semibold hover:bg-emerald-700">Login</router-link></li>
        </ul>
      </div>
    </header>

    <main>
      <section id="hero" class="relative py-24 md:py-40">
        <div class="hero-carousel">
          <div v-for="(image, index) in heroImages" :key="index" :class="['hero-image', { 'active': index === currentImageIndex }]" :style="{ backgroundImage: 'url(' + image + ')' }"></div>
        </div>

        <div class="hero-content">
          <h1 class="text-4xl md:text-5xl lg:text-6xl font-extrabold leading-tight tracking-wide">
            Cuidar da sua saúde mental é o primeiro passo para uma vida plena.
          </h1>
          <p class="mt-6 md:mt-8 text-lg md:text-xl max-w-4xl mx-auto opacity-90">
            MindTrace é a sua ferramenta para monitorar emoções, conectar-se com profissionais e construir hábitos mais saudáveis.
          </p>
          <a href="#features" class="mt-10 inline-block bg-white text-emerald-600 font-bold py-4 px-10 rounded-full shadow-lg hover:bg-gray-100 transition duration-300 transform hover:scale-105">
            Comece sua Jornada
          </a>
        </div>
      </section>

      <section id="features" class="py-24">
        <div class="container mx-auto px-6 text-center">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">Como o MindTrace te Ajuda</h2>
          <p class="mt-4 text-lg text-gray-600 max-w-2xl mx-auto">
            Ferramentas intuitivas e personalizadas para te ajudar a entender e melhorar seu bem-estar emocional.
          </p>

          <div class="mt-16 grid grid-cols-1 md:grid-cols-3 gap-10">
            <div class="bg-white p-8 rounded-2xl shadow-xl transform transition duration-300 hover:scale-105 hover:shadow-2xl">
              <div class="w-16 h-16 bg-emerald-100 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg class="w-8 h-8 text-emerald-600" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15s-2-1.5-2-3 2-2 2-2 2 .5 2 2-2 3-2 3zm4 0s-2-1.5-2-3 2-2 2-2 2 .5 2 2-2 3-2 3z"></path>
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900">Monitoramento de Humor</h3>
              <p class="mt-2 text-gray-600">
                Registre seus sentimentos diariamente e entenda os padrões do seu estado emocional.
              </p>
            </div>

            <div class="bg-white p-8 rounded-2xl shadow-xl transform transition duration-300 hover:scale-105 hover:shadow-2xl">
              <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg class="w-8 h-8 text-green-600" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zM9 17H7v-7h2v7zm4 0h-2V7h2v10zm4 0h-2v-4h2v4z"></path>
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900">Análises e Relatórios</h3>
              <p class="mt-2 text-gray-600">
                Gráficos e relatórios detalhados para ajudar você e seu profissional a tomar decisões informadas.
              </p>
            </div>

            <div class="bg-white p-8 rounded-2xl shadow-xl transform transition duration-300 hover:scale-105 hover:shadow-2xl">
              <div class="w-16 h-16 bg-rose-200 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg class="w-8 h-8 text-rose-600" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"></path>
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900">Conexão com Profissionais</h3>
              <p class="mt-2 text-gray-600">
                Encontre e conecte-se com terapeutas qualificados que podem te ajudar em sua jornada.
              </p>
            </div>
          </div>
        </div>
      </section>

      <section id="knowledge" class="knowledge-section bg-gray-100">
        <div class="knowledge-container">
          <div class="knowledge-text">
            <h2 class="knowledge-title"> Aqui a prioridade é sua saúde mental!</h2>
            <p class="knowledge-paragraph">Com o MindTrace você não terá somente apoio, mas toda uma estrutura e ambiente que te juda a se tornar mentalmente mais saudável.</p>
          </div>
          <div class="knowledge-image-container">
            <img src="https://images.unsplash.com/photo-1753454116483-417bbc0a975c?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" alt="Ambiente de trabalho moderno" class="knowledge-image">
          </div>
        </div>
      </section>

      <section id="testimonials" class="bg-gray-100 py-24">
        <div class="container mx-auto px-6 text-center">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">O que nossos usuários dizem</h2>
          <p class="mt-4 text-lg text-gray-600 max-w-2xl mx-auto">
            Histórias reais de pessoas que transformaram suas vidas com MindTrace.
          </p>

          <div class="mt-16 grid grid-cols-1 md:grid-cols-2 gap-10">
            <div class="bg-white p-8 rounded-2xl shadow-xl">
              <p class="text-gray-700 italic">"MindTrace me ajudou a entender melhor meus padrões de pensamento e a me conectar com uma terapeuta incrível. Minha ansiedade diminuiu significativamente."</p>
              <div class="mt-6 flex items-center justify-center">
                <img src="https://images.pexels.com/photos/774909/pexels-photo-774909.jpeg?auto=compress&cs=tinysrgb&w=100&h=100&fit=crop&crop=face" alt="Avatar Usuário" class="w-14 h-14 rounded-full mr-4 object-cover">
                <div class="text-left">
                  <p class="font-semibold text-gray-900">Ana Souza</p>
                  <p class="text-sm text-gray-500">Usuário MindTrace</p>
                </div>
              </div>
            </div>

            <div class="bg-white p-8 rounded-2xl shadow-xl">
              <p class="text-gray-700 italic">"O MindTrace me ajudou a entender melhor como funciona minha mente, agora me sinto mais presente e consciente dos meus sentimentos a cada dia. Recomendo muito!"</p>
              <div class="mt-6 flex items-center justify-center">
                <img src="https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=100&h=100&fit=crop&crop=face" alt="Avatar Usuário" class="w-14 h-14 rounded-full mr-4 object-cover">
                <div class="text-left">
                  <p class="font-semibold text-gray-900">João Silva</p>
                  <p class="text-sm text-gray-500">Usuário MindTrace</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

  <section id="contact" class="bg-emerald-600 text-white py-24">
        <div class="container mx-auto px-6 text-center">
          <h2 class="text-3xl md:text-4xl font-bold">Pronto para começar sua jornada?</h2>
          <p class="mt-4 text-lg max-w-2xl mx-auto opacity-90">
            Junte-se a pessoas que estão priorizando sua saúde mental. É fácil, seguro e o primeiro passo para uma nova você.
          </p>
          <router-link to="/cadastro" class="mt-10 inline-block bg-white text-emerald-600 font-bold py-4 px-10 rounded-full shadow-lg hover:bg-gray-100 transition duration-300 transform hover:scale-105">
            Crie sua conta gratuita
          </router-link>
        </div>
      </section>
    </main>

    <footer class="bg-gray-800 text-gray-300 py-8">
      <div class="container mx-auto px-6 text-center">
        <p>&copy; 2025 MindTrace. Todos os direitos reservados.</p>
        <div class="mt-4 space-x-4">
          <a href="#" class="hover:text-white transition duration-300">Política de Privacidade</a>
          <a href="#" class="hover:text-white transition duration-300">Termos de Serviço</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import '../assets/landpage.css';

const isMenuOpen = ref(false);

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};

const closeMenu = () => {
  isMenuOpen.value = false;
};

const heroImages = [
  'https://images.unsplash.com/photo-1754215683705-ee0a731d7288?w=600&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHx0b3BpYy1mZWVkfDEyfDZzTVZqVExTa2VRfHxlbnwwfHx8fHw%3D',
  'https://images.unsplash.com/photo-1497373637916-e47a55e22d0a?w=600&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTB8fENhbmFkJUMzJUExfGVufDB8fDB8fHww',
  'https://images.unsplash.com/photo-1532274402911-5a369e4c4bb5?w=600&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8N3x8bGFuZHNjYXBlfGVufDB8fDB8fHww',
  'https://images.unsplash.com/photo-1617718295766-0f839c2853e7?w=600&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTR8fERvbWluaWNhbiUyMFJlcHVibGljfGVufDB8MHwwfHx8MA%3D%3D',
  'https://images.unsplash.com/photo-1708598660454-4011df6397f0?q=80&w=1174&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
];

const currentImageIndex = ref(0);
let intervalId = null;

const nextImage = () => {
  currentImageIndex.value = (currentImageIndex.value + 1) % heroImages.length;
};

onMounted(() => {
  intervalId = setInterval(nextImage, 5000);
});

onUnmounted(() => {
  clearInterval(intervalId);
});

</script>
