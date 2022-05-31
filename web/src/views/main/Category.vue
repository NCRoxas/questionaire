<template>
	<div class="container mt-4">
		<breadcrumbs />

		<h1 class="text-center text-success mb-5 mt-3">{{ title }}</h1>
		<p class="text-center fs-4 mb-3">Wählen Sie eine Unterkategorie:</p>
		<div class="border-top border-success border-2"></div>

		<div class="row">
			<div class="col-md-3 p-3" v-for="(element, id) in data" :key="id">
				<div class="card">
					<Delete v-if="editmode" :item="element" @remove-item="remove" />

					<router-link
						:to="{ name: 'subcategory', params: { category: element.slug } }"
						:aria-label="element.slug"
					>
						<img
							v-if="!editmode"
							:src="`/images/vendor/${element.slug}.png?cache=${cacheKey}`"
							class="card-img-top border border-success border-2"
							alt="Category Image"
							@error="$event.target.src = '/images/not-found.png'"
						/>
						<img
							v-else
							:src="`/images/vendor/${element.slug}.png?cache=${cacheKey}`"
							class="card-img-top border border-danger border-3"
							alt="Category Image"
							@error="$event.target.src = '/images/not-found.png'"
							@dragover.prevent="dragover"
							@dragleave.prevent="dragleave"
							@drop.prevent="drop($event, element)"
						/>
					</router-link>

					<div class="card-body text-center">
						<Update :element="element" @update-item="update" />
					</div>
				</div>
			</div>
			<Create v-if="editmode" label="Kategorie" @add-item="add" />
		</div>
	</div>
</template>

<script setup>
import { useRoute } from "vue-router";
import { mapGetters, mapActions } from "@/store/map-state";
import Breadcrumbs from "@/components/Breadcrumbs.vue";
import Create from "@/components/cards/Create.vue";
import Update from "@/components/cards/Update.vue";
import Delete from "@/components/cards/Delete.vue";
import api from "@/store/api";
import { computed, ref } from "vue";

const route = useRoute();
const { packs, categories, editmode } = mapGetters();
const { getPacks, newCategory, updateCategory, deleteCategory } = mapActions();

const cacheKey = ref(0);
const title = packs.value.filter((item) => item.slug === route.params.pack)[0].title;
const data = computed(() => {
	if (categories.value === null){
		return {}
	}
	return categories.value.filter((item) => item.pack_slug === route.params.pack);
});

const add = (data) => {
	let request = { pack_slug: route.params.pack, title: data.title };
	newCategory(request).then(() => getPacks());
};
const remove = (data) => {
	deleteCategory(data).then(() => getPacks());
};
const update = (data, title) => {
	data.title = title.trim();
	updateCategory(data).then(() => getPacks());
};

// Drag and drop upload handler
const dragover = (event) => {
	event.currentTarget.classList.add("opacity-50");
};
const dragleave = (event) => {
	event.currentTarget.classList.remove("opacity-50");
};
const drop = async (event, element) => {
	event.target.classList.remove("opacity-50");

	let file = event.dataTransfer.files[0];
	if (file.type.indexOf("image/") >= 0) {
		let newFilename = element.slug + "." + file.name.split(".")[1];

		const formData = new FormData();
		formData.append("file", file, newFilename);
		const headers = { "Content-Type": "multipart/form-data" };
		await api.post("/upload", formData, { headers });
		cacheKey.value++;
	}
};
</script>
