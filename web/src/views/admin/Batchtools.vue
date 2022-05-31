<template>
	<div class="container mb-4 mt-4">
		<div class="row">
			<h1 class="fw-light">
				Verschiedene Verwaltungstools (experimentell)
				<icon icon="fa-solid:exclamation-triangle" :inline="true" />
			</h1>
			<span class="border-bottom border-dark border-2 mb-4"></span>

			<div class="col-md-6">
				<form class="row g-1 mb-3" @submit.prevent novalidate>
					<div class="form-floating">
						<select
							class="form-select"
							id="floatingSelect"
							aria-label="Floating label select example"
							v-model="state.select_pack"
						>
							<option v-for="item in packs" :key="item.id" :value="item.slug">
								{{ item.title }}
							</option>
						</select>
						<label for="floatingSelect">Paket wählen</label>
					</div>
					<div class="form-floating">
						<textarea
							class="form-control"
							placeholder="Leave a comment here"
							id="exampleFormControlTextarea1"
							style="height: 200px"
							v-model="state.create_cats"
						></textarea>
						<label for="exampleFormControlTextarea1" class="form-label">Kategorien</label>
					</div>
					<button @click="createCategories" type="submit" class="btn btn-primary">
						Batch insert categories
					</button>
				</form>
			</div>
			<div class="col-md-6">
				<form class="row g-1 mb-3" @submit.prevent novalidate>
					<div class="form-floating">
						<textarea
							class="form-control"
							placeholder="Leave a comment here"
							id="exampleFormControlTextarea1"
							style="height: 200px"
							v-model="state.del_cats"
						></textarea>
						<label for="exampleFormControlTextarea1" class="form-label">Kategorien</label>
					</div>
					<button @click="deleteCategories" type="submit" class="btn btn-danger">
						Batch delete categories
					</button>
				</form>
			</div>
			<div class="col-md-6">
				<form class="row g-1 mb-3" @submit.prevent novalidate>
					<div class="form-floating">
						<select
							class="form-select"
							id="floatingSelect"
							aria-label="Floating label select example"
							v-model="state.select_category"
						>
							<option v-for="item in categories" :key="item.id" :value="item.slug">
								{{ item.title }}
							</option>
						</select>
						<label for="floatingSelect">Kategorie wählen</label>
					</div>
					<div class="form-floating">
						<textarea
							class="form-control"
							placeholder="Leave a comment here"
							id="exampleFormControlTextarea1"
							style="height: 200px"
							v-model="state.create_subs"
						></textarea>
						<label for="exampleFormControlTextarea1" class="form-label">Unterkategorien</label>
					</div>
					<button @click="createSubcategories" type="submit" class="btn btn-primary">
						Batch insert subcategories
					</button>
				</form>
			</div>
			<div class="col-md-6">
				<form class="row g-1 mb-3" @submit.prevent novalidate>
					<div class="form-floating">
						<textarea
							class="form-control"
							placeholder="Leave a comment here"
							id="exampleFormControlTextarea1"
							style="height: 200px"
							v-model="state.del_subs"
						></textarea>
						<label for="exampleFormControlTextarea1" class="form-label">Unterkategorien</label>
					</div>
					<button @click="deleteSubcategories" type="submit" class="btn btn-danger">
						Batch delete subcategories
					</button>
				</form>
			</div>

			<div class="col-md-4">
				<ul class="list-group mb-1">
					<li class="list-group-item active" aria-current="true">Alle Pakete</li>
				</ul>
				<div v-for="(item, id) in packs" :key="id" class="input-group mb-1">
					<input
						type="text"
						class="form-control"
						placeholder="Recipient's username"
						aria-label="Recipient's username"
						aria-describedby="button-addon2"
						:value="item.title"
						disabled
					/>
					<button class="btn btn-outline-danger" type="button" @click="deleteP(item)">
						<icon icon="fa-solid:times" :inline="true" />
					</button>
				</div>
			</div>

			<div class="col-md-4">
				<ul class="list-group mb-1">
					<li class="list-group-item active" aria-current="true">Alle Kategorien</li>
				</ul>
				<div v-for="(item, id) in categories" :key="id" class="input-group mb-1">
					<input
						type="text"
						class="form-control"
						placeholder="Recipient's username"
						aria-label="Recipient's username"
						aria-describedby="button-addon2"
						:value="item.title"
						disabled
					/>
					<button class="btn btn-outline-danger" type="button" @click="deleteC(item)">
						<icon icon="fa-solid:times" :inline="true" />
					</button>
				</div>
			</div>

			<div class="col-md-4">
				<ul class="list-group mb-1">
					<li class="list-group-item active" aria-current="true">Alle Unterkategorien</li>
				</ul>
				<div v-for="(item, id) in subcategories" :key="id" class="input-group mb-1">
					<input
						type="text"
						class="form-control"
						placeholder="Recipient's username"
						aria-label="Recipient's username"
						aria-describedby="button-addon2"
						:value="item.title"
						disabled
					/>
					<button class="btn btn-outline-danger" type="button" @click="deleteS(item)">
						<icon icon="fa-solid:times" :inline="true" />
					</button>
				</div>
			</div>
		</div>
		<category-sort class="mt-4" />
		<subcategory-sort class="mt-4" />
	</div>
</template>

<script setup>
import { mapActions, mapGetters } from "@/store/map-state";
import CategorySort from "@/components/sortables/CategorySort.vue";
import SubcategorySort from "@/components/sortables/SubcategorySort.vue";
import { reactive } from "vue";

const { packs, categories, subcategories } = mapGetters();
const {
	newCategoryBatch,
	deleteCategoryBatch,
	newSubcategoryBatch,
	deleteSubcategoryBatch,
	getPacks,
	deletePack,
	deleteCategory,
	deleteSubcategory
} = mapActions();

const state = reactive({
	select_pack: null,
	select_category: null,
	create_cats: null,
	create_subs: null,
	del_cats: null,
	del_subs: null
});
const createCategories = () => {
	if (state.select_pack != "") {
		let cats = state.create_cats.split(/\r?\n/).map((element) => element.trim());
		let request = { pack: state.select_pack, title: cats };
		newCategoryBatch(request).then(() => getPacks());
	}
	state.create_cats = "";
	state.select_pack = "";
};
const deleteCategories = () => {
	let request = state.del_cats.split(/\r?\n/).map((element) => element.trim());
	deleteCategoryBatch(request).then(() => getPacks());
	state.del_cats = "";
};
const createSubcategories = () => {
	if (state.select_category != "") {
		let subs = state.create_subs.split(/\r?\n/).map((element) => element.trim());
		let request = { category: state.select_category, title: subs };
		newSubcategoryBatch(request).then(() => getPacks());
	}
	state.create_subs = "";
	state.select_category = "";
};
const deleteSubcategories = () => {
	let request = state.del_subs.split(/\r?\n/).map((element) => element.trim());
	deleteSubcategoryBatch(request).then(() => getPacks());
	state.del_subs = "";
};
const deleteP = (data) => {
	deletePack(data).then(() => getPacks());
};
const deleteC = (data) => {
	deleteCategory(data).then(() => getPacks());
};
const deleteS = (data) => {
	deleteSubcategory(data).then(() => getPacks());
};
</script>
