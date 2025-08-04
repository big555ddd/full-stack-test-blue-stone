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
          {{ isReset ? "Reset Password" : "Forgot Password" }}
        </h2>
      </div>

      <!-- Email Form (Step 1) -->
      <form @submit.prevent="handleEmailSubmit" class="space-y-4">
        <div>
          <input
            type="email"
            v-model="email"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all duration-200"
            placeholder="Email"
          />
        </div>

        <button
          type="submit"
          :disabled="isLoading"
          class="w-full bg-blue-800 hover:bg-blue-600 text-white py-3 px-6 rounded-lg font-medium transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="!isLoading">Send Reset Token</span>
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
            Sending...
          </span>
        </button>
      </form>

      <!-- Reset Password Form (Step 2) -->
      <form
        v-if="isReset"
        @submit.prevent="handleResetSubmit"
        class="space-y-4 mt-4"
      >
        <div>
          <input
            type="text"
            v-model="otp"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all duration-200"
            placeholder="Verification Code"
          />
        </div>
        <div>
          <input
            type="password"
            v-model="password"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all duration-200"
            placeholder="New Password"
          />
        </div>
        <div>
          <input
            type="password"
            v-model="confirmPassword"
            required
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all duration-200"
            placeholder="Confirm Password"
          />
        </div>
        <button
          type="submit"
          :disabled="isResetting"
          class="w-full bg-blue-800 hover:bg-blue-600 text-white py-3 px-6 rounded-lg font-medium transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <span v-if="!isResetting">Submit</span>
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
            Resetting...
          </span>
        </button>
      </form>

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
          href="https://www.instagram.com/bluestone.co.th"
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
import type {
  ForgotRequest,
  ForgotResponse,
  LoginRequest,
  VerifyOtpRequest,
} from "~/models/auth.model";
import services from "~/services";

definePageMeta({
  layout: "office",
});

const isLoading = ref(false);
const isResetting = ref(false);
const isReset = ref(false);
const email = ref("");
const otp = ref("");
const password = ref("");
const confirmPassword = ref("");
const forgotResponse = ref<ForgotResponse>({
  code: 0,
  message: "",
  data: {
    id: "",
    userId: "",
    otp: "",
    expiresAt: 0,
    used: false,
    createdAt: 0,
  },
});

// Handle email submission (Step 1)
const handleEmailSubmit = async () => {
  if (!email.value) {
    alert("Please fill in your email");
    return;
  }

  isLoading.value = true;
  const body: ForgotRequest = {
    email: email.value,
  };

  await services.auth
    .ForgotPassword(body)
    .then((resp: any) => {
      if (resp.status !== 200) {
        throw new Error(resp.data.message || "email not found");
      }
      console.log(resp);
      if (resp.status === 200) {
        forgotResponse.value = resp.data;
        isReset.value = true; // Move to step 2
        alert("Verification code sent to your email!");
      }
    })
    .catch((error) => {
      alert(error.message || "email not found");
    })
    .finally(() => {
      isLoading.value = false;
    });
};

// Handle password reset submission (Step 2)
const handleResetSubmit = async () => {
  if (!otp.value || !password.value || !confirmPassword.value) {
    alert("Please fill in all fields");
    return;
  }

  if (password.value !== confirmPassword.value) {
    alert("Passwords do not match");
    return;
  }

  isResetting.value = true;

  const body: VerifyOtpRequest = {
    id: forgotResponse.value.data.id,
    otp: otp.value,
    newPassword: password.value,
    confirmPassword: confirmPassword.value,
  };

  await services.auth
    .VerifyOtp(body)
    .then((resp: any) => {
      if (resp.status !== 200) {
        throw new Error(resp.data.message || "Verification failed");
      }
      alert("Password reset successful!");
      window.location.href = "/login";
    })
    .catch((error) => {
      alert(error.message || "Verification failed");
    })
    .finally(() => {
      isResetting.value = false;
    });
};
</script>
