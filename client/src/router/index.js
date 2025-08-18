// client/src/router/index.js
import { createRouter, createWebHistory } from "vue-router";
import { useAuth } from "@/stores/auth";

const Home = () => import("@/pages/Home.vue");
const Login = () => import("@/pages/Login.vue");
const Register = () => import("@/pages/Register.vue");
const Admin = () => import("@/pages/Admin.vue");
const Membership = () => import("@/pages/Membership.vue");

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "home", component: Home },
    {
      path: "/login",
      name: "login",
      component: Login,
      meta: { guestOnly: true },
    },
    {
      path: "/register",
      name: "register",
      component: Register,
      meta: { guestOnly: true },
    },
    // roles: ['admin'] implies requiresAuth
    {
      path: "/admin",
      name: "admin",
      component: Admin,
      meta: { roles: ["admin"] },
    },
    // roles: ['customer'] for customer-only page
    {
      path: "/membership",
      name: "membership",
      component: Membership,
      meta: { roles: ["customer"] },
    },
  ],
  scrollBehavior: () => ({ top: 0 }),
});

router.beforeEach((to) => {
  const { isAuthed, user, logoutIfExpired } = useAuth();

  // Optional: expire check; clears state if token is stale
  logoutIfExpired?.();

  const isLoggedIn = isAuthed.value === true;
  const role = user.value?.role; // 'admin' | 'customer' | undefined

  // Block authed users from visiting login/register
  if (to.meta?.guestOnly && isLoggedIn) return { name: "home" };

  // If route declares roles, require auth
  const requiredRoles = (to.matched || []).flatMap((r) => r.meta?.roles || []);
  if (requiredRoles.length > 0) {
    if (!isLoggedIn) return { name: "login", query: { redirect: to.fullPath } };
    if (!requiredRoles.includes(role)) return { name: "home" }; // role mismatch
  }

  // If a route explicitly says requiresAuth (even without roles)
  const requiresAuth = (to.matched || []).some((r) => r.meta?.requiresAuth);
  if (requiresAuth && !isLoggedIn) {
    return { name: "login", query: { redirect: to.fullPath } };
  }

  return true;
});

export default router;
