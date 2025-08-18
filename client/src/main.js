import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

import "./assets/styles/reset.css"; // uses your global CSS (keep path consistent)
import "./assets/styles/global.css"; // uses your global CSS (keep path consistent)

createApp(App).use(router).mount("#app");
