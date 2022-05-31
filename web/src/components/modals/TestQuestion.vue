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
								<textarea
									class="form-control"
									placeholder="Leave a description here"
									style="height: 200px"
									:value="question.description"
									readonly
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
									v-model="input.marked"
									disabled
								/>
							</div>
							<input
								type="text"
								class="form-control"
								placeholder="Antwort"
								:value="input.answer"
								readonly
							/>
						</div>
					</form>
				</div>
				<div class="modal-footer">
					<button
						class="btn btn-warning me-auto"
						@click="
							show = !show;
							showAnswer(question.Answers);
						"
					>
						<icon icon="fa-solid:question" :inline="true" />
					</button>
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Schließen</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
import { ref } from "vue";

const props = defineProps({ element: Object, target: String });

const question = ref(props.element);
const show = ref(false);

const showAnswer = (data) => {
	if (show.value) {
		for (const i in data) {
			if (data[i].correct === true) {
				data[i].marked = true;
			}
		}
	} else {
		for (const i in data) {
			data[i].marked = false;
		}
	}
};
</script>
