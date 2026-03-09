<template>
  <div class="card">
    <div class="card-content">
      <h3 class="title is-5">Agregar Vacuna</h3>

      <div v-if="success" class="notification is-success">
        <button class="delete" @click="success = false"></button>
        Vacuna registrada exitosamente.
      </div>

      <div v-if="error" class="notification is-danger">
        <button class="delete" @click="error = null"></button>
        {{ error }}
      </div>

      <form @submit.prevent="crearVacuna">
        <div class="columns">
          <div class="column">
            <div class="field">
              <label class="label">Nombre de la Vacuna</label>
              <div class="control">
                <input
                  v-model="form.nombrevacuna"
                  class="input"
                  type="text"
                  placeholder="Ej: Antirrábica"
                  required
                />
              </div>
            </div>
          </div>
          <div class="column">
            <div class="field">
              <label class="label">Fecha</label>
              <div class="control">
                <input
                  v-model="form.fecha"
                  class="input"
                  type="date"
                  required
                />
              </div>
            </div>
          </div>
        </div>

        <div class="field">
          <div class="control">
            <button
              class="button is-success"
              type="submit"
              :disabled="enviando"
            >
              {{ enviando ? "Guardando..." : "Agregar Vacuna" }}
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
const { $api } = useNuxtApp();

const props = defineProps({
  idPerro: {
    type: Number,
    required: true,
  },
});

const emit = defineEmits(["vacuna-creada"]);

const form = ref({
  nombrevacuna: "",
  fecha: "",
});

const enviando = ref(false);
const success = ref(false);
const error = ref(null);

const crearVacuna = async () => {
  enviando.value = true;
  error.value = null;
  success.value = false;

  try {
    await $api.post("/vacunas", {
      ...form.value,
      id_perro: props.idPerro,
    });
    success.value = true;
    form.value = { nombrevacuna: "", fecha: "" };
    emit("vacuna-creada");
  } catch (e) {
    error.value = e.response?.data?.error || e.message;
  } finally {
    enviando.value = false;
  }
};
</script>
