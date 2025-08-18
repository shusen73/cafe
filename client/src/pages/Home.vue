<template>
  <header class="header"></header>

  <section>
    <h1>Menu</h1>

    <table v-if="items.length">
      <thead>
        <tr>
          <th class="width-min">Item</th>
          <th class="width-auto">Description</th>
          <th class="width-auto">Price</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id">
          <td>{{ item.name }}</td>
          <td>{{ item.description }}</td>
          <td>{{ formatPrice(item.priceCents) }}</td>
        </tr>
      </tbody>
    </table>

    <p v-else>No items yet.</p>
  </section>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { fetchMenu } from "@/services/menuAdmin";

const items = ref([]);

onMounted(async () => {
  const data = await fetchMenu();
  // keep only active items
  items.value = data.filter((item) => item.active);
});

function formatPrice(cents) {
  const v = (cents ?? 0) / 100;
  return v.toLocaleString(undefined, { style: "currency", currency: "USD" });
}
</script>
