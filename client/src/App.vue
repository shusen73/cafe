<template>
  <header class="header">
    <table>
      <tbody>
        <tr>
          <td class="width-min">
            <h1><RouterLink to="/">Coffee Shop</RouterLink></h1>
          </td>

          <td class="width-auto">
            <nav aria-label="Primary">
              <RouterLink to="/">Home</RouterLink>

              <span v-if="isAdmin">
                | <RouterLink to="/admin">Admin</RouterLink>
              </span>

              <span v-if="isMember">
                | <RouterLink to="/membership">Membership</RouterLink>
              </span>
            </nav>
          </td>

          <td class="width-min">
            <template v-if="isAuthed">
              <div class="">
                <small>
                  Signed in as {{ user && user.email }} ({{
                    user && user.role
                  }})
                </small>
                <button class="" @click="onLogout">Logout</button>
              </div>
            </template>

            <template v-else>
              <nav aria-label="Auth">
                <RouterLink to="/login">Login</RouterLink> /
                <RouterLink to="/register">Register</RouterLink>
              </nav>
            </template>
          </td>
        </tr>
      </tbody>
    </table>
  </header>

  <main>
    <RouterView />
  </main>
</template>

<script setup>
import { computed } from "vue";
import { useRouter } from "vue-router";
import { useAuth } from "@/stores/auth";

const router = useRouter();
const { user, isAuthed, isAdmin, logout } = useAuth();

// Customer = logged in and NOT admin
const isMember = computed(() => isAuthed.value && !isAdmin.value);

function onLogout() {
  logout();
  router.push("/");
}
</script>

<style scoped></style>
