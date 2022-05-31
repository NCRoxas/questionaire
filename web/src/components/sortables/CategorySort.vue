<template>
	<div class="row">
		<h1 class="fw-light">Kategorien umsortieren</h1>
		<span class="border-bottom border-dark border-2 mb-4"></span>

		<div class="col-md-6">
			<div class="form-floating mb-2">
				<select
					class="form-select"
					id="floatingSelect"
					aria-label="Floating label select example"
					v-model="pack1"
				>
					<option v-for="(item, id) in packs" :key="id" :value="item">
						{{ item.title }}
					</option>
				</select>
				<label for="floatingSelect">Paket wählen</label>
				<span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-primary">
					{{ pack1.categories.length }}
				</span>
			</div>
		</div>
		<div class="col-md-6">
			<div class="form-floating mb-2">
				<select
					class="form-select"
					id="floatingSelect"
					aria-label="Floating label select example"
					v-model="pack2"
				>
					<option v-for="(item, id) in packs" :key="id" :value="item">
						{{ item.title }}
					</option>
				</select>
				<label for="floatingSelect">Paket wählen</label>
				<span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-primary">
					{{ pack2.categories.length }}
				</span>
			</div>
		</div>
		<div class="col-md-6">
			<draggable
				class="list-group"
				:list="pack1.categories"
				v-bind="dragOptions"
				@change="addList1"
				@start="start"
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
				:list="pack2.categories"
				v-bind="dragOptions"
				@change="addList2"
				@start="start"
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
			pack1: { categories: [] },
			pack2: { categories: [] },
			clone: false
		};
	},
	computed: {
		...mapGetters(["packs"]),
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
		...mapActions(["getPacks", "newCategory", "updateCategory", "deleteCategory"]),
		addList1: function (evt) {
			if (evt.added) {
				let element = evt.added.element;
				element.pack_slug = this.pack1.slug;

				if (this.clone) {
					element.slug.replace(this.pack2.slug, this.pack1.slug);
					this.pack2.categories.push(element);
					this.newCategory(element);
				} else {
					this.updateCategory(element);
				}
			}
		},
		addList2: function (evt) {
			if (evt.added) {
				let element = evt.added.element;
				element.pack_slug = this.pack2.slug;

				if (this.clone) {
					element.slug = element.slug.replace(this.pack1.slug, this.pack2.slug);
					this.pack1.categories.push(element);
					this.newCategory(element);
				} else {
					this.updateCategory(element);
				}
			}
		},
		remove: function (data) {
			this.deleteCategory(data);
		},
		start({ originalEvent }) {
			this.clone = originalEvent.ctrlKey;
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
