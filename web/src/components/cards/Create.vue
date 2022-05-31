<template>
	<div v-if="isAdmin" class="col-md-3 p-3">
		<div class="card h-100">
			<div class="text-dark text-center mt-5 mb-5">
				<icon icon="fa-solid:plus" width="100px" />
			</div>
			<div class="card-body mt-4">
				<div class="form-floating addItem">
					<input
						type="text"
						class="form-control"
						id="addItem"
						v-model="title"
						placeholder="item"
						@keydown.enter="add(title)"
					/>
					<label for="addItem">{{ label }} hinzufügen</label>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
import { mapGetters } from "@/store/map-state";
import { ref } from "vue";

defineProps({ label: String });
const emit = defineEmits(["add-item"]);

const { isAdmin } = mapGetters();
const title = ref(null);

const add = () => {
	let newItem = { title: title.value };
	emit("add-item", newItem);
	title.value = null;
};
</script>

<style scoped>
.addItem {
	position: absolute;
	bottom: 20px;
	width: 89%;
}
</style>
