<template>
  <div>
    <h1 class="title">Registrar Nuevo Perro</h1>

    <div v-if="success" class="notification is-success">
      <button class="delete" @click="success = false"></button>
      ¡Perro registrado exitosamente!
      <NuxtLink :to="`/perros/${nuevoId}`" class="has-text-weight-bold ml-2"
        >Ver perro →</NuxtLink
      >
    </div>

    <div v-if="error" class="notification is-danger">
      <button class="delete" @click="error = null"></button>
      {{ error }}
    </div>

    <form @submit.prevent="crearPerro" class="box">
      <div class="field">
        <label class="label">Nombre</label>
        <div class="control">
          <input
            v-model="form.nombre"
            class="input"
            type="text"
            placeholder="Nombre del perro"
            required
          />
        </div>
      </div>

      <div class="field">
        <label class="label">Raza</label>
        <div class="control">
          <input
            v-model="form.raza"
            class="input"
            type="text"
            placeholder="Raza"
            required
          />
        </div>
      </div>

      <div class="field">
        <label class="label">Color</label>
        <div class="control">
          <input
            v-model="form.color"
            class="input"
            type="text"
            placeholder="Color"
            required
          />
        </div>
      </div>

      <div class="columns">
        <div class="column">
          <div class="field">
            <label class="label">Edad</label>
            <div class="control">
              <input
                v-model.number="form.edad"
                class="input"
                type="number"
                min="0"
                placeholder="Edad"
                required
              />
            </div>
          </div>
        </div>
        <div class="column">
          <div class="field">
            <label class="label">ID Dueño</label>
            <div class="control">
              <input
                v-model.number="form.id_dueno"
                class="input"
                type="number"
                min="1"
                placeholder="ID del dueño"
                required
              />
            </div>
          </div>
        </div>
      </div>

      <div class="field">
        <div class="control">
          <button class="button is-primary" type="submit" :disabled="enviando">
            {{ enviando ? "Guardando..." : "Registrar Perro" }}
          </button>
        </div>
      </div>
    </form>
  </div>
</template>

<script setup>
const { $api } = useNuxtApp();

const form = ref({
  nombre: "",
  raza: "",
  color: "",
  edad: 0,
  id_dueno: 1,
});

const enviando = ref(false);
const success = ref(false);
const error = ref(null);
const nuevoId = ref(null);

const crearPerro = async () => {
  enviando.value = true;
  error.value = null;
  success.value = false;

  try {
    const { data } = await $api.post("/perros", form.value);
    nuevoId.value = data._id;
    success.value = true;
    form.value = { nombre: "", raza: "", color: "", edad: 0, id_dueno: 1 };
  } catch (e) {
    error.value = e.response?.data?.error || e.message;
  } finally {
    enviando.value = false;
  }
};
</script>
