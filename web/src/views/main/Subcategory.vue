<template>
	<div class="container mt-4">
		<breadcrumbs />

		<h1 class="text-center text-success mb-5 mt-3">{{ title }}</h1>

		<div class="text-center mb-5">
			<p class="fw-bold">Wissen überprüfen oder vertiefen - so einfach geht's:</p>
			<p>
				Wähle hier dein gewünschtes Wissensgebiet aus und vertiefe dein Wissen in der jeweiligen Kategorie oder
				überprüfe dein bereits vorhandenes Wissen in einem Gesamttest, der alle Wissenskategorien dieser Seite
				beinhaltet.
			</p>
			<p class="text-decoration-underline">Bitte wähle den entsprechenden Modus.</p>
		</div>

		<p class="text-center fs-4 mb-3">Wählen Sie eine Unterkategorie:</p>
		<div class="border-top border-success border-2"></div>

		<div class="row">
			<div class="col-md-3 p-3" v-for="(element, id) in data" :key="id">
				<div class="card">
					<Delete v-if="editmode" :item="element" @remove-item="remove" />

					<router-link
						:to="{ name: 'testmode', params: { subcategory: element.slug } }"
						:aria-label="element.slug"
					>
						<img
							v-if="!editmode"
							:src="`/images/vendor/${element.slug}.png?cache=${cacheKey}`"
							class="card-img-top border border-success border-2"
							alt="Subcategory Image"
							@error="$event.target.src = '/images/not-found.png'"
						/>
						<img
							v-else
							:src="`/images/vendor/${element.slug}.png?cache=${cacheKey}`"
							class="card-img-top border border-danger border-3"
							alt="Subcategory Image"
							@error="$event.target.src = '/images/not-found.png'"
							@dragover.prevent="dragover"
							@dragleave.prevent="dragleave"
							@drop.prevent="drop($event, element)"
						/>
					</router-link>

					<div class="card-body text-center">
						<Update :element="element" @update-item="update" />
						<div class="d-grid gap-2 mt-3">
							<router-link
								:to="{ name: 'testmode', params: { subcategory: element.slug } }"
								class="btn btn-primary btn-sm"
								>Test starten</router-link
							>
							<router-link
								:to="{ name: 'helpmode', params: { subcategory: element.slug } }"
								class="btn btn-primary btn-sm"
								>Lernmodus starten</router-link
							>
						</div>
					</div>
				</div>
			</div>
			<Create v-if="editmode" label="Unterkategorie" @add-item="add" />
		</div>

		<div class="d-grid gap-2 mt-3 mb-5">
			<button class="btn btn-success">Gesamttest für diese Kategorie starten</button>
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
const { categories, subcategories, editmode } = mapGetters();
const { getPacks, newSubcategory, updateSubcategory, deleteSubcategory } = mapActions();

const cacheKey = ref(0);
const title = categories.value.filter((item) => item.slug === route.params.category)[0].title;
const data = computed(() => {
	if (subcategories.value === null){
		return {}
	}
	return subcategories.value.filter((item) => item.category_slug === route.params.category);
});

const add = (data) => {
	let request = { category_slug: route.params.category, title: data.title };
	newSubcategory(request).then(() => getPacks());
};
const remove = (data) => {
	deleteSubcategory(data).then(() => getPacks());
};
const update = (data, title) => {
	data.title = title.trim();
	updateSubcategory(data).then(() => getPacks());
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
