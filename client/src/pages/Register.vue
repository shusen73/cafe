<template>
  <header class="header">
    <h1>Register</h1>
  </header>

  <section>
    <form @submit.prevent="onSubmit">
      <label>
        Email
        <input
          v-model="email"
          type="email"
          required
          placeholder="you@example.com"
        />
      </label>
      <br />
      <label>
        Password
        <input
          v-model="password"
          type="password"
          required
          minlength="6"
          placeholder="min 6 chars"
        />
      </label>
      <br />
      <div class="grid">
        <button type="submit">Create Account</button>
      </div>
    </form>

    <p v-if="error"><strong>Error:</strong> {{ error }}</p>
  </section>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuth } from "@/stores/auth";

const router = useRouter();
const { register, login } = useAuth();

const email = ref("");
const password = ref("");
const error = ref("");

async function onSubmit() {
  error.value = "";
  try {
    await register(email.value, password.value);
    // auto-login for convenience
    await login(email.value, password.value);
    router.push("/");
  } catch (e) {
    error.value = e.message || "Registration failed";
  }
}
</script>
