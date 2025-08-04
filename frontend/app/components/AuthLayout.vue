<template>
  <div class="min-h-screen flex items-center justify-center p-4">
    <div class="bg-white rounded-2xl shadow-2xl p-8 w-full max-w-md">
      <!-- BlueStone Logo and Header -->
      <div class="text-center mb-8">
        <div class="mb-6">
          <img
            src="/img/logo2.png"
            alt="BlueStone Logo"
            class="mx-auto w-auto"
          />
        </div>

        <h2 class="text-xl font-semibold text-gray-800 mt-2">
          {{ title }}
        </h2>
      </div>

      <!-- Form Content Slot -->
      <slot />

      <!-- Divider -->
      <div class="mt-6">
        <div class="relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm"></div>
        </div>
      </div>

      <span
        class="flex w-full px-2 bg-white text-gray-500 text-sm justify-center text-center"
        >Connect Us</span
      >

      <!-- Social Login Buttons -->
      <div class="flex justify-center space-x-8">
        <a
          href="https://www.facebook.com/Bluestone.co.th/"
          target="_blank"
          class="rounded-full text-white transition-colors h-16 w-16 hover:scale-110 transition-transform duration-200"
        >
          <img src="/img/icons-facebook.png" alt="Facebook Logo" />
        </a>
        <a
          href="https://www.instagram.com/bluestonethailand/"
          target="_blank"
          class="rounded-full text-white transition-colors h-16 w-16 hover:scale-110 transition-transform duration-200 mt-2"
        >
          <img src="/img/icons-ig.png" alt="Instagram Logo" />
        </a>
        <a
          href="https://line.me/ti/p/~@bluestonethailand"
          target="_blank"
          class="rounded-full text-white transition-colors h-16 w-16 hover:scale-110 transition-transform duration-200"
        >
          <img src="/img/icons-line.png" alt="Line Logo" />
        </a>
        <a
          href="https://www.youtube.com/channel/UCQ3mRpetmm5Ek-LLdTjwaNQ"
          target="_blank"
          class="rounded-full text-white transition-colors h-16 w-16 hover:scale-110 transition-transform duration-200"
        >
          <img src="/img/icons-youtube.png" alt="YouTube Logo" />
        </a>
      </div>

      <!-- Footer Text -->
      <div class="mt-6 text-center">
        <a
          href="https://www.bluestone.co.th"
          target="_blank"
          class="text-xs transition-colors duration-200 text-blue-800 hover:text-blue-600 underline"
          >www.bluestone.co.th</a
        >
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title: string;
}

defineProps<Props>();
</script>
