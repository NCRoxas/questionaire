<template>
	<div class="container mb-4 mt-4">
		<h2 v-if="subcategories === null" class="text-center">Keine Pakete gekauft!</h2>
		<apexchart
			v-if="progress[1] > 0"
			width="550"
			type="donut"
			:options="chartOptions"
			:series="progress"
		></apexchart>
		<ul class="list-group list-group-numbered mb-5">
			<li
				v-for="(item, id) in subcategories"
				:key="id"
				class="list-group-item d-flex justify-content-between align-items-start"
			>
				<div class="ms-2 me-auto">
					<div class="fw-bold mb-2">{{ item.title }}</div>
					<div v-for="(item, id) in item.Questions" :key="id">
						<div class="d-grid gap-2">
							<button
								class="btn position-relative mb-1"
								data-bs-toggle="modal"
								:data-bs-target="'#qModal' + id"
								v-bind:class="check(item.title) == true ? 'btn-success' : 'btn-light'"
							>
								<icon v-show="check(item.title)" icon="fa-solid:check" :inline="true" class="me-2" />
								{{ item.title }}
							</button>
							<TestQuestion :element="item" :target="'qModal' + id" />
						</div>
					</div>
				</div>
				<button class="btn btn-danger btn-sm me-2" @click="deleteResults(item.Questions)">
					Ergebnisse löschen
				</button>
			</li>
		</ul>
	</div>
</template>

<script setup>
import TestQuestion from "@/components/modals/TestQuestion.vue";
import { mapGetters, mapActions } from "@/store/map-state";

const { username, result, progress, subcategories } = mapGetters();
const { evaluateQuestions, deleteUserRels } = mapActions();

const chartOptions = {
	chart: {
		id: "total-progress-chart"
	},
	labels: ["Richtig beantwortete Fragen", "Falsche / Nicht beantwortete Fragen"],
	plotOptions: {
		pie: {
			startAngle: -90,
			endAngle: 90,
			offsetY: 10
		}
	},
	grid: {
		padding: {
			bottom: -90
		}
	},
	responsive: [
		{
			breakpoint: 480,
			options: {
				chart: {
					width: 200
				},
				legend: {
					position: "bottom"
				}
			}
		}
	]
};

const check = (data) => {
	for (const i in result.value) {
		if (data === result.value[i].question_title && result.value[i].correct === true) {
			return true;
		}
	}
};
const deleteResults = (data) => {
	let request = { username: username.value, useranswers: [] };
	for (const i in data) {
		let answers = data[i].Answers;
		for (const j in answers) {
			request.useranswers.push({ answer_id: answers[j].id });
		}
	}
	deleteUserRels(request).then(() => evaluateQuestions(username.value));
};
</script>
