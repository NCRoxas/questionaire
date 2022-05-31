<template>
	<div class="fit-page">
		<HeaderAdmin />
		<main class="flex-shrink-0 mt-3 mb-3">
			<div class="container">
				<slot />
			</div>
		</main>
		<Footer />
	</div>
</template>

<script setup>
import HeaderAdmin from "@/components/layouts/parts/HeaderAdmin.vue";
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
