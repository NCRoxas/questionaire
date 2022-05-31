<template>
	<div v-if="isAdmin" class="position-absolute top-0 start-100 translate-middle">
		<button class="btn btn-danger btn-rounded btn-sm" @click="$emit('remove-item', item)" aria-label="deleteButton">
			<icon icon="fa-solid:times" :inline="true" />
		</button>
	</div>
</template>

<script setup>
import { mapGetters } from "@/store/map-state";

defineProps({ item: Object });
const { isAdmin } = mapGetters();
</script>

<style scoped>
.btn-rounded {
	width: 30px;
	height: 30px;
	padding: 0 0;
	border-radius: 15px;
	font-size: 16px;
	border: solid 3px rgb(243, 154, 154);
}
</style>
