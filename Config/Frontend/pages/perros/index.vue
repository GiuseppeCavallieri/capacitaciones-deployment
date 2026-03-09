<template>
  <div>
    <h1 class="title">Lista de Perros</h1>

    <div v-if="loading" class="has-text-centered">
      <p>Cargando perros...</p>
    </div>

    <div v-else-if="error" class="notification is-danger">
      <p>Error al cargar los perros: {{ error }}</p>
    </div>

    <div v-else class="columns is-multiline">
      <div v-for="perro in perros" :key="perro.id" class="column is-4">
        <PerroCard :perro="perro" />
      </div>
    </div>

    <div
      v-if="!loading && perros.length === 0 && !error"
      class="notification is-warning"
    >
      No hay perros registrados.
    </div>
  </div>
</template>

<script setup>
const { $api } = useNuxtApp();

const perros = ref([]);
const loading = ref(true);
const error = ref(null);

onMounted(async () => {
  try {
    const { data } = await $api.get("/perros");
    perros.value = data;
  } catch (e) {
    error.value = e.message;
  } finally {
    loading.value = false;
  }
});
</script>
