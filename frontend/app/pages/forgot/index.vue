<template>
  <AuthLayout title="Reset Password">
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
        <span v-if="!isResetting">Reset Password</span>
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
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 714 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            ></path>
          </svg>
          Resetting...
        </span>
      </button>
    </form>
  </AuthLayout>
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
