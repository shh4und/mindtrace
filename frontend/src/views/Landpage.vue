<template>
  <div class="bg-gray-50 text-gray-800 font-sans">
    <header class="bg-white shadow-sm sticky top-0 z-50">
      <nav class="container mx-auto px-6 py-4 flex justify-between items-center">
        <a href="#" class="text-2xl font-bold text-blue-600">MindTrace</a>

        <div class="md:hidden">
          <button @click="toggleMenu" class="text-gray-600 focus:outline-none">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7"></path>
            </svg>
          </button>
        </div>

        <ul class="hidden md:flex space-x-6 items-center">
          <li><a href="#hero" class="hover:text-blue-600 transition duration-300">Início</a></li>
          <li><a href="#features" class="hover:text-blue-600 transition duration-300">Recursos</a></li>
          <li><a href="#testimonials" class="hover:text-blue-600 transition duration-300">Depoimentos</a></li>
          <li><a href="#contact" class="hover:text-blue-600 transition duration-300">Contato</a></li>
          <li><router-link to="/login" class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition duration-300">Login</router-link></li>
        </ul>
      </nav>

      <div :class="['md:hidden', { 'hidden': !isMenuOpen }]" class="bg-white shadow-md">
        <ul class="px-6 py-4 space-y-4 text-center">
          <li><a href="#hero" @click="closeMenu" class="block text-gray-800 hover:bg-gray-100 p-2 rounded-lg">Início</a></li>
          <li><a href="#features" @click="closeMenu" class="block text-gray-800 hover:bg-gray-100 p-2 rounded-lg">Recursos</a></li>
          <li><a href="#testimonials" @click="closeMenu" class="block text-gray-800 hover:bg-gray-100 p-2 rounded-lg">Depoimentos</a></li>
          <li><a href="#contact" @click="closeMenu" class="block text-gray-800 hover:bg-gray-100 p-2 rounded-lg">Contato</a></li>
          <li><router-link to="/login" @click="closeMenu" class="block bg-blue-600 text-white p-2 rounded-lg mt-4">Login</router-link></li>
        </ul>
      </div>
    </header>

    <main>
        <section id="hero" class="relative bg-gradient-to-r from-blue-500 to-teal-500 text-white py-20 md:py-32">
            <div class="container mx-auto px-6 text-center">
                <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">
                    Cuidar da sua saúde mental é o primeiro passo para uma vida plena.
                </h1>
                <p class="mt-4 md:mt-6 text-lg md:text-xl max-w-3xl mx-auto">
                    MindTrace é a sua ferramenta para monitorar emoções, conectar-se com profissionais e construir hábitos mais saudáveis.
                </p>
                <a href="#features" class="mt-8 inline-block bg-white text-blue-600 font-bold py-3 px-8 rounded-full shadow-lg hover:bg-gray-100 transition duration-300">
                    Comece Agora
                </a>
            </div>
        </section>

        <section id="features" class="py-20">
            <div class="container mx-auto px-6 text-center">
                <h2 class="text-3xl md:text-4xl font-bold text-gray-900">Nossos Recursos Pensados para Você</h2>
                <p class="mt-4 text-lg text-gray-600 max-w-2xl mx-auto">
                    Ferramentas intuitivas e personalizadas para te ajudar a entender e melhorar seu bem-estar emocional.
                </p>
                
                <div class="mt-12 grid grid-cols-1 md:grid-cols-3 gap-8">
                    <div class="bg-white p-8 rounded-lg shadow-md hover:shadow-xl transition duration-300">
                        <div class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4">
                            <svg xmlns="http://www.w3.org/2000/svg" height="20" width="20" viewBox="0 0 512 512"><path d="M464 256a208 208 0 1 0 -416 0 208 208 0 1 0 416 0zM0 256a256 256 0 1 1 512 0A256 256 0 1 1 0 256zm177.3 63.4C192.3 335 218.4 352 256 352s63.7-17 78.7-32.6c9.2-9.6 24.4-9.9 33.9-.7s9.9 24.4 .7 33.9c-22.1 23-60 47.4-113.3 47.4s-91.2-24.4-113.3-47.4c-9.2-9.6-8.9-24.8 .7-33.9s24.8-8.9 33.9 .7zM144 208a32 32 0 1 1 64 0 32 32 0 1 1 -64 0zm192-32a32 32 0 1 1 0 64 32 32 0 1 1 0-64z"/></svg>
                        </div>
                        <h3 class="text-xl font-semibold text-gray-900">Monitoramento de Humor</h3>
                        <p class="mt-2 text-gray-600">
                            Registre seus sentimentos diariamente e entenda os padrões do seu estado emocional.
                        </p>
                    </div>

                    <div class="bg-white p-8 rounded-lg shadow-md hover:shadow-xl transition duration-300">
                        <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
                            <svg xmlns="http://www.w3.org/2000/svg" height="20" width="20" viewBox="0 0 640 640"><path d="M480 576L192 576C139 576 96 533 96 480L96 160C96 107 139 64 192 64L496 64C522.5 64 544 85.5 544 112L544 400C544 420.9 530.6 438.7 512 445.3L512 512C529.7 512 544 526.3 544 544C544 561.7 529.7 576 512 576L480 576zM192 448C174.3 448 160 462.3 160 480C160 497.7 174.3 512 192 512L448 512L448 448L192 448zM224 216C224 229.3 234.7 240 248 240L424 240C437.3 240 448 229.3 448 216C448 202.7 437.3 192 424 192L248 192C234.7 192 224 202.7 224 216zM248 288C234.7 288 224 298.7 224 312C224 325.3 234.7 336 248 336L424 336C437.3 336 448 325.3 448 312C448 298.7 437.3 288 424 288L248 288z"/></svg>
                        </div>
                        <h3 class="text-xl font-semibold text-gray-900"> Relatório</h3>
                        <p class="mt-2 text-gray-600">
                            Relatórios detalhados, para ajudar na tomada de decisões clínicas. Visão completa e objetiva do bem-estar psicológico, permitindo um acompanhamento personalizado e eficaz.
                        </p>
                    </div>

                    <div class="bg-white p-8 rounded-lg shadow-md hover:shadow-xl transition duration-300">
                        <div class="w-16 h-16 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
                            <svg xmlns="http://www.w3.org/2000/svg" height="14" width="15.75" viewBox="0 0 576 512"><path d="M64 128a112 112 0 1 1 224 0 112 112 0 1 1 -224 0zM0 464c0-97.2 78.8-176 176-176s176 78.8 176 176l0 6c0 23.2-18.8 42-42 42L42 512c-23.2 0-42-18.8-42-42l0-6zM432 64a96 96 0 1 1 0 192 96 96 0 1 1 0-192zm0 240c79.5 0 144 64.5 144 144l0 22.4c0 23-18.6 41.6-41.6 41.6l-144.8 0c6.6-12.5 10.4-26.8 10.4-42l0-6c0-51.5-17.4-98.9-46.5-136.7 22.6-14.7 49.6-23.3 78.5-23.3z"/></svg>
                        </div>
                        <h3 class="text-xl font-semibold text-gray-900">Conexão com Profissionais</h3>
                        <p class="mt-2 text-gray-600">
                            Encontre e conecte-se com terapeutas qualificados que podem te ajudar em sua jornada.
                        </p>
                    </div>
                </div>
            </div>
        </section>
        
        <section id="testimonials" class="bg-gray-100 py-20">
            <div class="container mx-auto px-6 text-center">
                <h2 class="text-3xl md:text-4xl font-bold text-gray-900">O que nossos usuários dizem</h2>
                <p class="mt-4 text-lg text-gray-600 max-w-2xl mx-auto">
                    Histórias reais de pessoas que transformaram suas vidas com MindTrace.
                </p>
                
                <div class="mt-12 grid grid-cols-1 md:grid-cols-2 gap-8">
                    <div class="bg-white p-8 rounded-lg shadow-md">
                        <p class="text-gray-700 italic">"MindTrace me ajudou a entender melhor meus padrões de pensamento e a me conectar com uma terapeuta incrível. Minha ansiedade diminuiu significativamente."</p>
                        <div class="mt-4 flex items-center justify-center">
                            <img src="https://images.pexels.com/photos/774909/pexels-photo-774909.jpeg?auto=compress&cs=tinysrgb&w=100&h=100&fit=crop&crop=face" alt="Avatar Usuário" class="w-12 h-12 rounded-full mr-4">
                            <div class="text-left">
                                <p class="font-semibold text-gray-900">Ana Souza</p>
                                <p class="text-sm text-gray-500">Usuário MindTrace</p>
                            </div>
                        </div>
                    </div>
                    
                    <div class="bg-white p-8 rounded-lg shadow-md">
                        <p class="text-gray-700 italic">"O MindTrace me ajudou a entender melhor como funciona minha mente, agora me sinto mais presente e consciente dos meus sentimentos a cada dia. Recomendo muito!"</p>
                        <div class="mt-4 flex items-center justify-center">
                            <img src="https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=100&h=100&fit=crop&crop=face" alt="Avatar Usuário" class="w-12 h-12 rounded-full mr-4">
                            <div class="text-left">
                                <p class="font-semibold text-gray-900">João Silva</p>
                                <p class="text-sm text-gray-500">Usuário MindTrace</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <section id="contact" class="bg-blue-600 text-white py-20">
            <div class="container mx-auto px-6 text-center">
                <h2 class="text-3xl md:text-4xl font-bold">Pronto para começar sua jornada?</h2>
                <p class="mt-4 text-lg max-w-2xl mx-auto">
                    Junte-se a milhares de pessoas que estão priorizando sua saúde mental. É fácil, seguro e o primeiro passo para uma nova você.
                </p>
                <router-link to="/login" class="mt-8 inline-block bg-white text-blue-600 font-bold py-3 px-8 rounded-full shadow-lg hover:bg-gray-100 transition duration-300">
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

<script>
export default {
  data() {
    return {
      isMenuOpen: false
    };
  },
  methods: {
    toggleMenu() {
      this.isMenuOpen = !this.isMenuOpen;
    },
    closeMenu() {
      this.isMenuOpen = false;
    }
  }
};
</script>

<style>
/* Estilos adicionais para o menu de navegação responsivo */
#mobile-menu {
    /* Adiciona uma transição suave para as propriedades de transformação e opacidade */
    transition: all 0.3s ease-in-out;
}
/* Classe para ocultar o menu mobile, movendo-o para cima e diminuindo a opacidade */
#mobile-menu.hidden {
    transform: translateY(-100%);
    opacity: 0;
}
/* Classe para mostrar o menu mobile, resetando a transformação e opacidade */
#mobile-menu.active {
    transform: translateY(0);
    opacity: 1;
}
/* Habilita a rolagem suave para links âncora na página */
html {
    scroll-behavior: smooth;
}
</style>
