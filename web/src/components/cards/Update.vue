<template>
	<OnClickOutside @trigger="editing = false">
		<div v-if="isAdmin" class="card-text">
			<span @click="focusOnInput" v-show="!editing">
				{{ title }}
			</span>
			<span v-show="editing">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="titleValue"
						placeholder="title"
						v-model="title"
						@keydown.enter="
							$emit('update-item', element, title);
							editing = false;
						"
						ref="editInput"
					/>
					<label for="titleValue">Titel</label>
				</div>
			</span>
		</div>
		<div v-else>
			<p class="card-text">{{ title }}</p>
		</div>
	</OnClickOutside>
</template>

<script setup>
import { OnClickOutside } from "@vueuse/components";
import { mapGetters } from "@/store/map-state";
import { ref } from "vue";

const props = defineProps({ element: Object });

const { isAdmin } = mapGetters();
const title = ref(props.element.title);
const editing = ref(false);

const focusOnInput = () => {
	editing.value = true;
};
</script>

<style scoped>
.card-text {
	font-size: 16px;
	font-weight: 600;
}
</style>
