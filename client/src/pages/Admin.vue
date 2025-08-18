<template>
  <header class="header">
    <h1>Admin — Menu</h1>
  </header>

  <section v-if="!isAdmin">
    <p>
      You must log in as an admin to manage items.
      <a href="/login">Go to Login</a>
    </p>
  </section>

  <section v-else>
    <h2>Create Item</h2>
    <form @submit.prevent="create">
      <div class="grid">
        <label>
          Name
          <input v-model="form.name" required />
        </label>
        <label>
          Price (cents)
          <input
            v-model.number="form.priceCents"
            type="number"
            min="0"
            required
          />
        </label>
      </div>
      <label>
        Description
        <input v-model="form.description" />
      </label>
      <div class="grid">
        <button type="submit">Create</button>
      </div>
    </form>

    <h2>Items</h2>
    <table v-if="items.length">
      <thead>
        <tr>
          <th class="width-min">Item</th>
          <th class="width-auto">Description</th>
          <th class="width-auto">Price</th>
          <th class="width-min">Active</th>
          <th class="width-auto">Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.id">
          <template v-if="editingId === item.id">
            <td><input v-model="edit.name" /></td>
            <td><input v-model="edit.description" /></td>
            <td>
              <input v-model.number="edit.priceCents" type="number" min="0" />
            </td>
            <td>
              <input
                type="checkbox"
                v-model="edit.active"
                :aria-label="`Active ${edit.name}`"
              />
            </td>
            <td>
              <div class="grid">
                <button @click="save(item.id)">Save</button>
                <button @click="cancel">Cancel</button>
              </div>
            </td>
          </template>
          <template v-else>
            <td>{{ item.name }}</td>
            <td>{{ item.description }}</td>
            <td>{{ formatPrice(item.priceCents) }}</td>
            <td>{{ item.active ? "Yes" : "No" }}</td>
            <td>
              <div class="grid">
                <button @click="startEdit(item)">Edit</button>
                <button @click="remove(item.id)">Delete</button>
              </div>
            </td>
          </template>
        </tr>
      </tbody>
    </table>
    <p v-else>No items.</p>
  </section>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { fetchMenu } from "@/services/menuAdmin";
import {
  createMenuItem,
  updateMenuItem,
  deleteMenuItem,
} from "@/services/menuAdmin";
import { useAuth } from "@/stores/auth";

const { isAdmin } = useAuth();

const items = ref([]);
const form = ref({ name: "", priceCents: 0, description: "" });
const editingId = ref(null);
const edit = ref({ name: "", priceCents: 0, description: "", active: true });

onMounted(load);

async function load() {
  items.value = await fetchMenu();
}

function formatPrice(cents) {
  const v = (cents ?? 0) / 100;
  return v.toLocaleString(undefined, { style: "currency", currency: "USD" });
}

async function create() {
  const payload = {
    name: form.value.name.trim(),
    priceCents: Number(form.value.priceCents),
    description: form.value.description.trim(),
    active: true,
  };
  await createMenuItem(payload);
  form.value = { name: "", priceCents: 0, description: "" };
  await load();
}

function startEdit(item) {
  editingId.value = item.id;
  edit.value = {
    name: item.name,
    priceCents: item.priceCents,
    description: item.description,
    active: !!item.active,
  };
}

async function save(id) {
  await updateMenuItem(id, { ...edit.value });
  editingId.value = null;
  await load();
}

function cancel() {
  editingId.value = null;
}

async function remove(id) {
  await deleteMenuItem(id);
  await load();
}
</script>
