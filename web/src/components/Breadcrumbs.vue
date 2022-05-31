<template>
	<div class="container" aria-label="breadcrumb">
		<ol class="breadcrumb">
			<li class="breadcrumb-item" v-for="(breadcrumb, idx) in breadcrumbs" :key="idx">
				<router-link :to="breadcrumb.to" class="text-decoration-none text-capitalize badge bg-success">{{
					breadcrumb.text
				}}</router-link>
			</li>
		</ol>
	</div>
</template>

<script setup>
import { computed } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const breadcrumbs = computed(() => {
	if (typeof route.meta.breadcrumb === "function") {
		return route.meta.breadcrumb.call(this, route);
	}
	return route.meta.breadcrumb;
});
</script>

<style scoped>
.breadcrumb {
	margin: 10px 0 0 20px;
}
</style>
