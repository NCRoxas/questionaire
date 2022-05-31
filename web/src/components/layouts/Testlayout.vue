<template>
	<HeaderTest />
	<main class="flex-shrink-0 mt-4 mb-3">
		<div class="container">
			<slot />
		</div>
	</main>
</template>

<script setup>
import HeaderTest from "@/components/layouts/parts/HeaderTest.vue";
</script>

<style scoped>
nav {
	background-color: #154c29;
}
</style>
