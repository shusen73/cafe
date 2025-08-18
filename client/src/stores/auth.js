// client/src/stores/auth.js
import { ref, computed } from "vue";
import {
  getAuth as readPersisted,
  setAuth as persist,
  clearAuth as clearPersisted,
} from "@/services/auth";

const state = ref(readPersisted()); // { token, user, exp } | null

function setState(data) {
  state.value = data;
  if (data) persist(data);
  else clearPersisted();
}

export function useAuth() {
  const isAuthed = computed(() => !!state.value?.token);
  const user = computed(() => state.value?.user || null);
  const token = computed(() => state.value?.token || "");
  const isAdmin = computed(() => user.value?.role === "admin");
  const isCustomer = computed(() => user.value?.role === "customer");
  const exp = computed(() => state.value?.exp || 0);

  function logout() {
    setState(null);
  }

  function logoutIfExpired() {
    // assumes exp is seconds-since-epoch; adjust if ms
    const now = Math.floor(Date.now() / 1000);
    if (exp.value && now >= exp.value) logout();
  }

  async function login(email, password) {
    const res = await fetch("/api/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({ email, password }),
    });
    if (!res.ok) throw new Error("Invalid credentials");
    const { token, user, exp } = await res.json();
    setState({ token, user, exp });
    return user;
  }

  async function register(email, password) {
    const res = await fetch("/api/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify({ email, password }),
    });
    if (!res.ok) throw new Error("Registration failed");
    return (await res.json()).user;
  }

  return {
    state,
    user,
    token,
    isAuthed,
    isAdmin,
    isCustomer,
    login,
    register,
    logout,
    logoutIfExpired,
  };
}
