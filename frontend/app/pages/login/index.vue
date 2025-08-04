<template>
  <AuthLayout title="Login">
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

      <!-- Remember Me Checkbox -->
      <div class="flex w-full justify-center">
        <input
          id="remember"
          type="checkbox"
          v-model="rememberMe"
          class="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded shadow-xl hover:shadow-2xl focus:shadow-2xl transition-shadow duration-200"
        />
        <label for="remember" class="ml-2 text-sm text-gray-700 font-bold">
          Remember Me
        </label>
      </div>

      <!-- Login Button -->
      <button
        type="submit"
        :disabled="isLoading"
        class="w-full bg-blue-600 hover:bg-blue-700 text-white py-3 px-6 rounded-lg font-medium transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        <span v-if="!isLoading">Login</span>
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
          Logging in...
        </span>
      </button>

      <!-- Forgot Password Link -->
      <div class="text-center">
        <NuxtLink to="/forgot" class="text-blue-600 hover:underline text-sm">
          Forgotten account?
        </NuxtLink>
      </div>
    </form>
  </AuthLayout>
</template>

<script setup lang="ts">
definePageMeta({
  layout: "office",
});

const username = ref("");
const password = ref("");
const rememberMe = ref(true);
const isLoading = ref(false);

const handleSubmit = async () => {
  if (!username.value || !password.value) {
    alert("Please fill in all fields");
    return;
  }

  isLoading.value = true;
  try {
    // TODO: Implement login API call
    await new Promise((resolve) => setTimeout(resolve, 2000));
    alert("Login successful!");
  } catch (error) {
    alert("Login failed!");
  } finally {
    isLoading.value = false;
  }
};
</script>
