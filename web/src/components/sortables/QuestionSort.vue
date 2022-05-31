<template>
	<div class="row">
		<h1 class="fw-light">Fragen umsortieren</h1>
		<span class="border-bottom border-dark border-2 mb-4"></span>

		<div class="col-md-6">
			<div class="form-floating mb-2">
				<select
					class="form-select"
					id="floatingSelect"
					aria-label="Floating label select example"
					v-model="subcategory1"
				>
					<option v-for="(item, id) in subcategories" :key="id" :value="item">
						{{ item.title }}
					</option>
				</select>
				<label for="floatingSelect">Unterkategorie wählen</label>
				<span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-primary">
					{{ subcategory1.Questions.length }}
				</span>
			</div>
		</div>
		<div class="col-md-6">
			<div class="form-floating mb-2">
				<select
					class="form-select"
					id="floatingSelect"
					aria-label="Floating label select example"
					v-model="subcategory2"
				>
					<option v-for="(item, id) in subcategories" :key="id" :value="item">
						{{ item.title }}
					</option>
				</select>
				<label for="floatingSelect">Unterkategorie wählen</label>
				<span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-primary">
					{{ subcategory2.Questions.length }}
				</span>
			</div>
		</div>
		<div class="col-md-6">
			<draggable
				class="list-group"
				:list="subcategory1.Questions"
				v-bind="dragOptions"
				@change="addList1"
				item-key="id"
			>
				<template #item="{ element, index }">
					<div class="input-group mb-1">
						<input type="text" class="form-control" placeholder="title" :value="element.title" disabled />
						<button
							class="btn btn-outline-primary"
							type="button"
							data-bs-toggle="modal"
							:data-bs-target="'#qModal' + index"
						>
							<icon icon="fa-solid:pen" :inline="true" />
						</button>
						<button class="btn btn-outline-danger" type="button" @click="remove(element)">
							<icon icon="fa-solid:times" :inline="true" />
						</button>
						<UpdateQuestion :element="element" :target="'qModal' + index" @update-question="update" />
					</div>
				</template>
			</draggable>
		</div>
		<div class="col-md-6">
			<draggable
				class="list-group"
				:list="subcategory2.Questions"
				v-bind="dragOptions"
				@change="addList2"
				item-key="id"
			>
				<template #item="{ element, index }">
					<div class="input-group mb-1">
						<input type="text" class="form-control" placeholder="title" :value="element.title" disabled />
						<button
							class="btn btn-outline-primary"
							type="button"
							data-bs-toggle="modal"
							:data-bs-target="'#qModal' + index"
						>
							<icon icon="fa-solid:pen" :inline="true" />
						</button>
						<button class="btn btn-outline-danger" type="button" @click="remove(element)">
							<icon icon="fa-solid:times" :inline="true" />
						</button>
						<UpdateQuestion :element="element" :target="'qModal' + index" @update-question="update" />
					</div>
				</template>
			</draggable>
		</div>
	</div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import draggable from "vuedraggable";
import UpdateQuestion from "@/components/modals/UpdateQuestion.vue";

export default {
	components: { UpdateQuestion, draggable },
	data() {
		return {
			subcategory1: { Questions: [] },
			subcategory2: { Questions: [] },
			clone: false
		};
	},
	computed: {
		...mapGetters(["subcategories", "questions"]),
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
		...mapActions(["getPacks", "newQuestion", "updateQuestion", "deleteQuestion"]),
		addList1: function (evt) {
			if (evt.added) {
				let element = evt.added.element;
				element.subcategory_slug = this.subcategory1.slug;

				if (this.clone) {
					element.slug.replace(this.subcategory2.slug, this.subcategory1.slug);
					this.subcategory1.Questions.push(element);
					this.newQuestion(element);
				} else {
					this.updateQuestion(element);
				}
			}
		},
		addList2: function (evt) {
			if (evt.added) {
				let element = evt.added.element;
				element.subcategory_slug = this.subcategory2.slug;

				if (this.clone) {
					element.slug.replace(this.subcategory1.slug, this.subcategory2.slug);
					this.subcategory2.Questions.push(element);
					this.newQuestion(element);
				} else {
					this.updateQuestion(element);
				}
			}
		},
		update: function (data) {
			this.updateQuestion(data).then(() => this.getPacks());
		},
		remove: function (data) {
			this.deleteQuestion(data).then(() => this.getPacks());
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
