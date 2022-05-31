<template>
	<div v-if="['testmode', 'helpmode'].includes(route.name)">
		<test>
			<router-view v-slot="{ Component }">
				<component :is="Component" />
			</router-view>
		</test>
	</div>
	<div v-else-if="route.path.startsWith('/admin')">
		<admin>
			<router-view v-slot="{ Component }">
				<transition name="fade" mode="out-in">
					<component v-if="!isLoading" :is="Component" />
				</transition>
				<spinner v-if="isLoading" />
			</router-view>
		</admin>
	</div>
	<div v-else>
		<basic>
			<router-view v-slot="{ Component }">
				<transition name="fade" mode="out-in">
					<component v-if="!isLoading" :is="Component" />
				</transition>
				<spinner v-if="isLoading" />
			</router-view>
		</basic>
	</div>
</template>

<script setup>
import basic from "@/components/layouts/Basiclayout.vue";
import admin from "@/components/layouts/Adminlayout.vue";
import test from "@/components/layouts/Testlayout.vue";
import spinner from "@/components/Spinner.vue";
import { mapGetters } from "@/store/map-state";
import { useRoute } from "vue-router";

const route = useRoute();
const { isLoading } = mapGetters();
</script>

<style scoped>
.fade-enter-from,
.fade-leave-to {
	transform: translateX(20px);
	opacity: 0;
}

.fade-enter-to,
.fade-leave-from {
	transition: transform 0.3s, opacity 0.5s;
}
</style>
