<template>
  <div class="card">
    <div class="card-content">
      <div class="media">
        <div class="media-content">
          <p class="title is-5">🐕 {{ perro.nombre }}</p>
          <p class="subtitle is-6">{{ perro.raza }} · {{ perro.color }}</p>
        </div>
      </div>
      <div class="content">
        <p><strong>Edad:</strong> {{ perro.edad }} años</p>
        <NuxtLink :to="`/perros/${perro._id}`" class="button is-small is-info">
          Ver detalle
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  perro: {
    type: Object,
    required: true,
  },
});
</script>
