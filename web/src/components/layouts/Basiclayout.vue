<template>
	<div class="fit-page">
		<Header />
		<main class="flex-shrink-0">
			<slot />
		</main>
		<Footer />
	</div>
</template>

<script setup>
import Header from "@/components/layouts/parts/Header.vue";
import Footer from "@/components/layouts/parts/Footer.vue";
</script>

<style scoped>
.fit-page {
	display: flex;
	min-height: 100vh;
	flex-direction: column;
	justify-content: space-between;
}
</style>
