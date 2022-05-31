<template>
	<div class="modal fade" :id="target" tabindex="-1" aria-labelledby="qModal" aria-hidden="true">
		<div class="modal-dialog modal-xl">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title text-capitalize">{{ question.title }}</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<form class="row g-1" @submit.prevent novalidate>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<input type="text" class="form-control" placeholder="Titel" v-model="question.title" />
								<label>Titel</label>
							</div>
						</div>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<textarea
									class="form-control"
									placeholder="Leave a description here"
									style="height: 200px"
									v-model="question.description"
								></textarea>
								<label class="form-label">Beschreibung</label>
							</div>
						</div>
						<div class="input-group" v-for="(input, id) in question.Answers" :key="id">
							<div class="input-group-text">
								<input
									class="form-check-input mt-0"
									type="checkbox"
									tabindex="-1"
									aria-label="Checkbox if true or false"
									v-model="input.correct"
								/>
							</div>
							<input type="text" class="form-control" placeholder="Antwort" v-model="input.answer" />
							<button
								class="btn btn-outline-primary"
								type="button"
								tabindex="-1"
								@click="addAnswer(question.Answers)"
							>
								<icon icon="fa-solid:plus" :inline="true" />
							</button>
							<button
								class="btn btn-outline-danger"
								type="button"
								tabindex="-1"
								@click="removeAnswer(id, question.Answers, input.id)"
							>
								<icon icon="fa-solid:minus" :inline="true" />
							</button>
						</div>
					</form>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Schließen</button>
					<button
						type="button"
						class="btn btn-primary"
						data-bs-dismiss="modal"
						@click="$emit('update-question', question)"
					>
						Speichern
					</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
import { reactive } from "vue";
import { mapActions } from "@/store/map-state";

const props = defineProps({ element: Object, target: String });

const question = reactive(props.element);
const { getPacks, deleteAnswer } = mapActions();

const addAnswer = (field) => {
	field.push({ answer: "", correct: false });
};
const removeAnswer = (index, fieldType, id) => {
	if (index !== 0) {
		fieldType.splice(index, 1);
	}
	let request = { id: id, question_slug: question.slug };
	deleteAnswer(request).then(() => getPacks());
};
</script>
