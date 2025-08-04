<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Background Image -->
    <div
      class="absolute inset-0 bg-cover bg-center bg-no-repeat"
      style="background-image: url('/img/bg-office.png')"
    >
      <!-- Overlay -->
      <div class="absolute inset-0 bg-black bg-opacity-20"></div>
    </div>

    <!-- Content -->
    <div class="relative z-10">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
// Layout for office background
</script>
