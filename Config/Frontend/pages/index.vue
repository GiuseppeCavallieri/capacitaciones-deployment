<template>
  <div>
    <h1 class="title is-2 has-text-centered">🐾 Veterinaria Citiaps</h1>
    <p class="subtitle has-text-centered">Sistema de gestión veterinaria</p>

    <div class="columns is-centered mt-5">
      <div class="column is-4">
        <div class="card">
          <div class="card-content has-text-centered">
            <span class="icon is-large"><i class="fas fa-dog"></i></span>
            <p class="title is-4">🐕 Perros</p>
            <p class="subtitle is-6">Gestiona los perros registrados</p>
            <NuxtLink to="/perros" class="button is-primary"
              >Ver Perros</NuxtLink
            >
          </div>
        </div>
      </div>

      <div class="column is-4">
        <div class="card">
          <div class="card-content has-text-centered">
            <p class="title is-4">➕ Nuevo Perro</p>
            <p class="subtitle is-6">Registra un nuevo perro</p>
            <NuxtLink to="/nuevo" class="button is-success"
              >Agregar Perro</NuxtLink
            >
          </div>
        </div>
      </div>

      <div class="column is-4">
        <div class="card">
          <div class="card-content has-text-centered">
            <p class="title is-4">💉 Vacunas</p>
            <p class="subtitle is-6">Consulta vacunas por perro</p>
            <NuxtLink to="/perros" class="button is-info">Consultar</NuxtLink>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
