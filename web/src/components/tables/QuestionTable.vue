<template>
	<div class="row">
		<h1 class="fw-light">Fragentabelle</h1>
		<span class="border-bottom border-dark border-2 mb-4"></span>

		<!-- Entries & Search -->
		<div class="d-flex mb-4">
			<div>
				<select class="form-select" aria-label="Entries" v-model="recordsPerPage" @click="updateTable">
					<option selected>5</option>
					<option>10</option>
					<option>25</option>
					<option>50</option>
					<option>100</option>
				</select>
			</div>
			<div class="ms-auto">
				<input
					type="text"
					class="form-control"
					id="searchInput"
					placeholder="Suche..."
					@input="searchTable($event.target.value)"
				/>
			</div>
		</div>

		<!-- Table -->
		<div class="table-responsive">
			<table class="table table-hover">
				<thead>
					<tr class="table-dark">
						<th scope="col">#</th>
						<th scope="col">Titel</th>
						<th scope="col">Beschreibung</th>
						<th scope="col">Bearbeiten</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="(item, id) in questions" :key="id">
						<th scope="row">{{ item.id }}</th>
						<td>{{ item.title }}</td>
						<td>{{ item.description }}</td>
						<td>
							<div class="btn-group" role="group" aria-label="Basic example">
								<button
									type="button"
									class="btn btn-outline-primary btn-sm"
									data-bs-toggle="modal"
									:data-bs-target="'#qModal' + id"
								>
									Update
								</button>
								<button type="button" class="btn btn-outline-danger btn-sm" @click="remove(item)">
									Löschen
								</button>
							</div>
						</td>
						<UpdateQuestion :element="item" :target="'qModal' + id" @update-question="update" />
					</tr>
				</tbody>
			</table>
		</div>

		<!-- Pagination -->
		<nav aria-label="Page navigation">
			<ul class="pagination">
				<li v-if="currentPage === 1" class="page-item disabled">
					<a class="page-link" href="#" @click="previousPage">Previous</a>
				</li>
				<li v-else class="page-item">
					<a class="page-link" href="#" @click="previousPage">Previous</a>
				</li>
				<li
					class="page-item"
					v-for="entry in limitation"
					:key="entry"
					v-bind:class="{ active: entry == currentPage }"
				>
					<a class="page-link" href="#" @click="thisPage(entry)">{{ entry }}</a>
				</li>
				<li v-if="currentPage === question_pages" class="page-item disabled">
					<a class="page-link" href="#" @click="nextPage">Next</a>
				</li>
				<li v-else class="page-item">
					<a class="page-link" href="#" @click="nextPage">Next</a>
				</li>
			</ul>
		</nav>
	</div>
</template>

<script setup>
import UpdateQuestion from "@/components/modals/UpdateQuestion.vue";
import { mapActions, mapGetters } from "@/store/map-state";
import { ref } from "vue";

const { questions, question_pages } = mapGetters();
const { searchQuestions, getQuestionList, updateQuestion, newAnswer, deleteQuestion } = mapActions();

const limitation = ref([]);
const currentPage = ref(1);
const recordsPerPage = ref(5);

const previousPage = () => {
	if (currentPage.value === 0) return;
	currentPage.value--;
	getQuestionList(`?limit=${recordsPerPage.value}&page=${currentPage.value}`).then(() => {
		setLimit(currentPage.value - 5, currentPage.value + 5);
	});
};
const nextPage = () => {
	if (currentPage.value === question_pages) return;
	currentPage.value++;
	getQuestionList(`?limit=${recordsPerPage.value}&page=${currentPage.value}`).then(() => {
		setLimit(currentPage.value - 5, currentPage.value + 5);
	});
};
const thisPage = (page) => {
	currentPage.value = page;
	getQuestionList(`?limit=${recordsPerPage.value}&page=${page}`).then(() => {
		setLimit(currentPage.value - 5, currentPage.value + 5);
	});
};
const updateTable = () => {
	currentPage.value = 1;
	getQuestionList(`?limit=${recordsPerPage.value}`).then(() => {
		setLimit(currentPage.value - 5, currentPage.value + 5);
	});
};
const searchTable = (data) => {
	if (data !== "") {
		searchQuestions(`?limit=100&search=${data}`);
		currentPage.value = 1;
		limitation.value = [1];
	} else {
		getQuestionList(`?limit=${recordsPerPage.value}`).then(() => {
			setLimit(currentPage.value - 5, currentPage.value + 5);
		});
	}
};
const update = (data) => {
	for (const i in data.Answers) {
		if (!data.Answers[i].id) {
			let request = data.Answers[i];
			request.question_slug = data.slug;
			newAnswer(request);
		}
	}
	updateQuestion(data).then(() => getQuestionList(`?limit=${recordsPerPage.value}`));
};
const remove = (data) => {
	deleteQuestion(data).then(() => getQuestionList(`?limit=${recordsPerPage.value}`));
};
const setLimit = (start, stop) => {
	let rStart = start < 1 ? 1 : start;
	let rStop = stop > question_pages.value ? question_pages.value + 1 : stop;
	limitation.value = new Array(rStop - rStart).fill(rStart).map((n, i) => n + i);
};
getQuestionList(`?limit=${recordsPerPage.value}`).then(() => {
	setLimit(currentPage.value - 5, currentPage.value + 5);
});
</script>
