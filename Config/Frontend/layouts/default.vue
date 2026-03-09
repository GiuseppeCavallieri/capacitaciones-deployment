<template>
  <div>
    <nav class="navbar is-dark" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <NuxtLink class="navbar-item has-text-weight-bold" to="/">
          🐾 Veterinaria Citiaps
        </NuxtLink>
        <a
          role="button"
          class="navbar-burger"
          aria-label="menu"
          aria-expanded="false"
          @click="menuOpen = !menuOpen"
          :class="{ 'is-active': menuOpen }"
        >
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>
      <div class="navbar-menu" :class="{ 'is-active': menuOpen }">
        <div class="navbar-start">
          <NuxtLink class="navbar-item" to="/">Inicio</NuxtLink>
          <NuxtLink class="navbar-item" to="/perros">Perros</NuxtLink>
          <NuxtLink class="navbar-item" to="/nuevo">Nuevo Perro</NuxtLink>
        </div>
      </div>
    </nav>

    <section class="section">
      <div class="container">
        <slot />
      </div>
    </section>

    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          <strong>Veterinaria Citiaps</strong> — Capacitación Deployment CITI
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup>
const menuOpen = ref(false);
</script>
