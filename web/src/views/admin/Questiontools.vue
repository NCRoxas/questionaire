<template>
	<div class="container mb-4 mt-4">
		<div class="row">
			<question-table class="mb-3" />

			<h1 class="fw-light">Fragen hinzufügen</h1>
			<span class="border-bottom border-dark border-2 mb-4"></span>

			<div class="col-md-12">
				<form class="row g-1 mb-5" @submit.prevent novalidate>
					<div class="form-floating mb-2">
						<select
							class="form-select"
							id="floatingSelect"
							aria-label="Floating label select example"
							v-model="question.subcategory_slug"
						>
							<option v-for="(item, id) in subcategories" :key="id" :value="item.slug">
								{{ item.title }}
							</option>
						</select>
						<label for="floatingSelect">Unterkategorie wählen</label>
					</div>
					<div class="form-floating">
						<input
							type="text"
							class="form-control"
							id="titleInput"
							placeholder="title"
							v-model="question.title"
						/>
						<label for="titleInput">Titel</label>
					</div>
					<div class="form-floating">
						<textarea
							class="form-control"
							placeholder="Leave a description here"
							id="descriptionInput"
							style="height: 200px"
							v-model="question.description"
						></textarea>
						<label for="descriptionInput" class="form-label">Beschreibung</label>
					</div>
					<div class="input-group" v-for="(input, id) in question.answers" :key="id">
						<div class="input-group-text">
							<input
								class="form-check-input mt-0"
								type="checkbox"
								tabindex="-1"
								aria-label="Checkbox if true or false"
								v-model="input.correct"
							/>
						</div>
						<input
							type="text"
							class="form-control"
							id="answerInput"
							placeholder="Antwort"
							v-model="input.answer"
						/>
						<button
							class="btn btn-outline-primary"
							type="button"
							tabindex="-1"
							@click="addAnswer(id, question.answers)"
						>
							<icon icon="fa-solid:plus" :inline="true" />
						</button>
						<button
							class="btn btn-outline-danger"
							type="button"
							tabindex="-1"
							@click="removeAnswer(id, question.answers)"
						>
							<icon icon="fa-solid:minus" :inline="true" />
						</button>
					</div>
					<button @click="addQuestion" type="submit" class="btn btn-primary mt-3">Frage hinzufügen</button>
				</form>
			</div>

			<question-sort />
		</div>
	</div>
</template>

<script setup>
import { mapActions, mapGetters } from "@/store/map-state";
import QuestionSort from "@/components/sortables/QuestionSort.vue";
import QuestionTable from "@/components/tables/QuestionTable.vue";
import { reactive } from "vue";

const { subcategories } = mapGetters();
const { getQuestionList, newQuestion } = mapActions();

const question = reactive({
	title: "",
	description: "",
	answers: [{ answer: "", correct: false }],
	subcategory_slug: ""
});

const addQuestion = () => {
	var valid = true;
	for (const i in question.answers) {
		if (question.title === "") {
			valid = false;
		} else if (question.description === "") {
			valid = false;
		} else if (question.answers[i].answer === "") {
			valid = false;
		}
	}

	if (valid) {
		newQuestion(question).then(() => getQuestionList(""));
	}
};
const addAnswer = (value, fieldType) => {
	fieldType.push({ answer: "", correct: false });
};
const removeAnswer = (index, fieldType) => {
	if (index !== 0) {
		fieldType.splice(index, 1);
	}
};
</script>
