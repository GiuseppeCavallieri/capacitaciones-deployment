<template>
  <div>
    <NuxtLink to="/perros" class="button is-light mb-4">← Volver</NuxtLink>

    <div v-if="loading" class="has-text-centered">
      <p>Cargando...</p>
    </div>

    <div v-else-if="error" class="notification is-danger">
      <p>{{ error }}</p>
    </div>

    <div v-else>
      <!-- Info del Perro -->
      <div class="card mb-5">
        <div class="card-content">
          <h1 class="title">🐕 {{ perro.nombre }}</h1>
          <div class="columns">
            <div class="column">
              <p><strong>Raza:</strong> {{ perro.raza }}</p>
              <p><strong>Color:</strong> {{ perro.color }}</p>
              <p><strong>Edad:</strong> {{ perro.edad }} años</p>
            </div>
            <div class="column">
              <h2 class="subtitle">Dueño</h2>
              <div v-if="dueno">
                <p><strong>Nombre:</strong> {{ dueno.nombre }}</p>
                <p><strong>Edad:</strong> {{ dueno.edad }} años</p>
                <p><strong>Sexo:</strong> {{ dueno.sexo }}</p>
              </div>
              <p v-else class="has-text-grey">Sin dueño asignado</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Vacunas -->
      <div class="card mb-5">
        <div class="card-content">
          <h2 class="title is-4">💉 Vacunas</h2>
          <table
            v-if="vacunas.length > 0"
            class="table is-fullwidth is-striped"
          >
            <thead>
              <tr>
                <th>Nombre</th>
                <th>Fecha</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="vacuna in vacunas" :key="vacuna._id">
                <td>{{ vacuna.nombrevacuna }}</td>
                <td>{{ vacuna.fecha }}</td>
              </tr>
            </tbody>
          </table>
          <p v-else class="has-text-grey">No hay vacunas registradas.</p>
        </div>
      </div>

      <!-- Formulario Agregar Vacuna -->
      <VacunaForm
        :id-perro="Number(route.params.id)"
        @vacuna-creada="cargarVacunas"
      />
    </div>
  </div>
</template>

<script setup>
const { $api } = useNuxtApp();
const route = useRoute();

const perro = ref({});
const dueno = ref(null);
const vacunas = ref([]);
const loading = ref(true);
const error = ref(null);

const cargarVacunas = async () => {
  try {
    const { data } = await $api.get(`/perros/${route.params.id}/vacunas`);
    vacunas.value = data || [];
  } catch (e) {
    console.error("Error cargando vacunas:", e);
  }
};

onMounted(async () => {
  try {
    const [perroRes, duenoRes, vacunasRes] = await Promise.allSettled([
      $api.get(`/perros/${route.params.id}`),
      $api.get(`/perros/${route.params.id}/dueno`),
      $api.get(`/perros/${route.params.id}/vacunas`),
    ]);

    if (perroRes.status === "fulfilled") {
      perro.value = perroRes.value.data;
    } else {
      throw new Error("Perro no encontrado");
    }

    if (duenoRes.status === "fulfilled") {
      dueno.value = duenoRes.value.data;
    }

    if (vacunasRes.status === "fulfilled") {
      vacunas.value = vacunasRes.value.data || [];
    }
  } catch (e) {
    error.value = e.message;
  } finally {
    loading.value = false;
  }
});
</script>
