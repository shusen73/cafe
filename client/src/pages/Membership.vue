<template>
  <section class="wrap">
    <div v-if="!isAuthed">
      <p>Please login first.</p>
      <RouterLink :to="{ name: 'login', query: { redirect: '/membership' } }">
        LOGIN
      </RouterLink>
    </div>

    <div v-else-if="isAdmin" class="empty">
      <p>Admin account do not have membership card</p>
      <RouterLink :to="{ name: 'home' }">Home</RouterLink>
    </div>

    <div v-else class="center">
      <MembershipCard :user="user" />
    </div>
  </section>
</template>

<script setup>
import { useAuth } from "@/stores/auth";
import MembershipCard from "@/components/MembershipCard.vue";

// These are computed refs; DO NOT call them as functions in the template.
const { user, isAuthed, isAdmin } = useAuth();
</script>

<style scoped>
.center {
  place-self: center;
}
</style>
