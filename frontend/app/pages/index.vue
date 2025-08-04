<template>
  <AuthLayout title="Register">
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <!-- Username Field -->
      <div>
        <input
          type="text"
          v-model="username"
          required
          class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all duration-200"
          placeholder="Username"
        />
      </div>

      <!-- Email Field -->
      <div>
        <input
          type="email"
          v-model="email"
          required
          class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all duration-200"
          placeholder="Email Address"
        />
      </div>

      <!-- Password Field -->
      <div>
        <input
          type="password"
          v-model="password"
          required
          class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all duration-200"
          placeholder="Password"
        />
      </div>

      <!-- Register Button -->
      <button
        type="submit"
        :disabled="isLoading"
        class="w-full bg-blue-800 hover:bg-blue-700 text-white py-3 px-6 rounded-lg font-medium transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <span v-if="!isLoading">Register</span>
        <span v-else class="flex items-center justify-center">
          <svg
            class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              class="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              stroke-width="4"
            ></circle>
            <path
              class="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            ></path>
          </svg>
          Registering...
        </span>
      </button>
    </form>
  </AuthLayout>
</template>

<script setup lang="ts">
import type { RegisterRequest } from "~/models/auth.model";
import services from "~/services";

definePageMeta({
  layout: "office",
});

const username = ref("");
const password = ref("");
const email = ref("");
const isLoading = ref(false);

const handleSubmit = async () => {
  if (!username.value || !password.value || !email.value) {
    alert("Please fill in all fields");
    return;
  }

  isLoading.value = true;
  const body: RegisterRequest = {
    username: username.value,
    password: password.value,
    email: email.value,
  };

  await services.auth
    .Register(body)
    .then((resp: any) => {
      console.log("Register response:", resp);

      if (resp.status !== 200) {
        alert("Registration failed. Please try again.");
        return;
      }
      alert("Register Success");
      navigateTo("/login");
    })
    .catch((error) => {
      console.error("Registration failed:", error);
      alert("Registration failed. Please try again.");
    })
    .finally(() => {
      isLoading.value = false;
    });
};
</script>
