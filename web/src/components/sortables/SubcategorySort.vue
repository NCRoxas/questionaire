<template>
	<div class="row">
		<h1 class="fw-light">Unterkategorien umsortieren</h1>
		<span class="border-bottom border-dark border-2 mb-4"></span>

		<div class="col-md-6">
			<div class="form-floating mb-2">
				<select
					class="form-select"
					id="floatingSelect"
					aria-label="Floating label select example"
					v-model="category1"
				>
					<option v-for="(item, id) in categories" :key="id" :value="item">
						{{ item.title }}
					</option>
				</select>
				<label for="floatingSelect">Kategorie wählen</label>
				<span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-primary">
					{{ category1.Subcategories.length }}
				</span>
			</div>
		</div>
		<div class="col-md-6">
			<div class="form-floating mb-2">
				<select
					class="form-select"
					id="floatingSelect"
					aria-label="Floating label select example"
					v-model="category2"
				>
					<option v-for="(item, id) in categories" :key="id" :value="item">
						{{ item.title }}
					</option>
				</select>
				<label for="floatingSelect">Kategorie wählen</label>
				<span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-primary">
					{{ category2.Subcategories.length }}
				</span>
			</div>
		</div>
		<div class="col-md-6">
			<draggable
				class="list-group"
				:list="category1.Subcategories"
				v-bind="dragOptions"
				@change="addList1"
				item-key="id"
			>
				<template #item="{ element }">
					<div class="input-group mb-1">
						<input type="text" class="form-control" placeholder="title" :value="element.title" disabled />
						<button class="btn btn-outline-danger" type="button" @click="remove(element)">
							<icon icon="fa-solid:times" :inline="true" />
						</button>
					</div>
				</template>
			</draggable>
		</div>
		<div class="col-md-6">
			<draggable
				class="list-group"
				:list="category2.Subcategories"
				v-bind="dragOptions"
				@change="addList2"
				item-key="id"
			>
				<template #item="{ element }">
					<div class="input-group mb-1">
						<input type="text" class="form-control" placeholder="title" :value="element.title" disabled />
						<button class="btn btn-outline-danger" type="button" @click="remove(element)">
							<icon icon="fa-solid:times" :inline="true" />
						</button>
					</div>
				</template>
			</draggable>
		</div>
	</div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import draggable from "vuedraggable";

export default {
	components: { draggable },
	data() {
		return {
			category1: { Subcategories: [] },
			category2: { Subcategories: [] },
			clone: false
		};
	},
	computed: {
		...mapGetters(["categories"]),
		dragOptions() {
			return {
				animation: 200,
				group: "description",
				disabled: false,
				ghostClass: "ghost"
			};
		}
	},
	methods: {
		...mapActions(["getPacks", "newSubcategory", "updateSubcategory", "deleteSubcategory"]),
		addList1: function (evt) {
			if (evt.added) {
				let element = evt.added.element;
				element.category_slug = this.category1.slug;

				if (this.clone) {
					element.slug.replace(this.category2.slug, this.category1.slug);
					this.category2.Subcategories.push(element);
					this.newSubcategory(element);
				} else {
					this.updateSubcategory(element);
				}
			}
		},
		addList2: function (evt) {
			if (evt.added) {
				let element = evt.added.element;
				element.category_slug = this.category2.slug;

				if (this.clone) {
					element.slug.replace(this.category1.slug, this.category2.slug);
					this.category1.Subcategories.push(element);
					this.newSubcategory(element);
				} else {
					this.updateSubcategory(element);
				}
			}
		},
		remove: function (data) {
			this.deleteCategory(data);
		}
	},
	mounted() {
		this.getPacks();
	}
};
</script>

<style scoped>
input[disabled] {
	background-color: white;
}
</style>
