<template>
  <article>
    <header>
      <h3>Coffee Club</h3>
      <span>Member</span>
    </header>

    <section>
      <div>
        <span>Name</span>
        <span>{{ displayName }}</span>
      </div>
      <div>
        <span>Member ID</span>
        <span>{{ String(user.id) }}</span>
      </div>
      <div>
        <span>Role</span>
        <span>{{ user?.role || "customer" }}</span>
      </div>
    </section>

    <footer>
      <!-- purely decorative stripe to feel like a card -->
      <div></div>
    </footer>
  </article>
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  user: { type: Object, required: true },
});

const displayName = computed(() => {
  const u = props.user || {};
  if (u.name && String(u.name).trim()) return u.name;
  if (u.email && typeof u.email === "string") return u.email.split("@")[0];
  return "Member";
});
</script>

<style scoped></style>
