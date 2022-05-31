import App from "./App.vue";
import router from "./router";
import store from "./store";
import VueApexCharts from "vue3-apexcharts";
import { createApp } from "vue";
import { Icon } from "@iconify/vue";
import { registerSW } from "virtual:pwa-register";
import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";

const intervalMS = 60 * 60 * 1000;
registerSW({
	onRegistered(r) {
		r &&
			setInterval(() => {
				r.update();
			}, intervalMS);
	}
});

const app = createApp(App);
app.use(router);
app.use(store);
app.use(VueApexCharts);
app.component("icon", Icon);
app.mount("#app");
