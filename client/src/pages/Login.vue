<template>
  <header class="header">
    <h1>Login</h1>
  </header>

  <section>
    <form @submit.prevent="onSubmit">
      <label>
        Email
        <input
          v-model="email"
          type="email"
          required
          placeholder="admin@example.com"
        />
      </label>
      <br />
      <label>
        Password
        <input
          v-model="password"
          type="password"
          required
          placeholder="••••••••"
        />
      </label>
      <br />
      <div class="grid">
        <button type="submit">Sign In</button>
      </div>
    </form>

    <p><a href="/register">Create an account</a></p>
    <p v-if="error"><strong>Error:</strong> {{ error }}</p>
  </section>
</template>

<script setup>
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useAuth } from "@/stores/auth";

const router = useRouter();
const route = useRoute();
const { login } = useAuth();

const email = ref("");
const password = ref("");
const error = ref("");

async function onSubmit() {
  error.value = "";
  try {
    const user = await login(email.value, password.value);
    const redirect =
      route.query.redirect || (user.role === "admin" ? "/admin" : "/");
    router.push(String(redirect));
  } catch (e) {
    error.value = e.message || "Login failed";
  }
}
</script>

<style scoped></style>
