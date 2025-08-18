import { authHeader } from "./auth";

export async function fetchMenu() {
  const res = await fetch("/api/menu", {
    headers: { Accept: "application/json" },
  });
  if (!res.ok) throw new Error("Failed to fetch menu");
  const data = await res.json();
  return data.items ?? [];
}

export async function createMenuItem(payload) {
  const res = await fetch("/api/menu", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
      ...authHeader(),
    },
    body: JSON.stringify(payload),
  });
  if (!res.ok) throw new Error("Create failed");
  return (await res.json()).item;
}

export async function updateMenuItem(id, payload) {
  const res = await fetch(`/api/menu/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
      ...authHeader(),
    },
    body: JSON.stringify(payload),
  });
  if (!res.ok) throw new Error("Update failed");
  return (await res.json()).item;
}

export async function deleteMenuItem(id) {
  const res = await fetch(`/api/menu/${id}`, {
    method: "DELETE",
    headers: { ...authHeader() },
  });
  if (!res.ok && res.status !== 204) throw new Error("Delete failed");
}
