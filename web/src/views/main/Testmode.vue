<template>
	<!-- Timer -->
	<p class="timer">{{ timer }}<icon icon="fa-solid:hourglass-half" :inline="true" class="ms-2" /></p>

	<h1 class="text-center text-success mb-3 mt-5">{{ question_title }} - Testmodus</h1>

	<!-- Progress -->
	<div class="d-flex flex-row justify-content-center">
		<div
			v-for="(bar, id) in questions.length"
			:key="id"
			class="rounded flex-grow-1 m-3"
			style="border: 5px solid"
			:class="{ 'text-success': bar - 1 <= state.step }"
		></div>
	</div>

	<!-- Question -->
	<div class="row">
		<div v-if="!state.finished" class="col-md-12">
			<div class="card mb-2">
				<div class="card-body">
					<h5 class="card-title">{{ questions[state.step].title }}</h5>
					<h6 class="card-subtitle mb-3 text-muted">Frage #{{ state.step + 1 }}</h6>
					<p class="card-text">
						{{ questions[state.step].description }}
					</p>
					<form class="row g-1" @submit.prevent novalidate>
						<div class="input-group" v-for="(choice, aid) in questions[state.step].Answers" :key="aid">
							<div class="input-group-text">
								<input
									class="form-check-input mt-0"
									type="checkbox"
									tabindex="-1"
									aria-label="Auswahl"
									v-model="choice.marked"
								/>
							</div>

							<input
								type="text"
								class="form-control"
								id="answerInput"
								placeholder="Antwort"
								:value="choice.answer"
								disabled
							/>
						</div>
					</form>
				</div>
			</div>
			<div v-if="state.step === 0 && questions.length > 1">
				<div class="d-grid gap-2">
					<button class="btn btn-primary" @click="nextStep">Nächste Frage</button>
				</div>
			</div>
			<div v-else-if="state.step === questions.length - 1">
				<div class="d-grid gap-2">
					<div class="btn-group" role="group">
						<button class="btn btn-info" @click.prevent="previousStep">Vorherige Frage</button>
						<button class="btn btn-success" @click="evaluate">Test beenden</button>
					</div>
				</div>
			</div>
			<div v-else>
				<div class="d-grid gap-2">
					<div class="btn-group" role="group">
						<button class="btn btn-info" @click.prevent="previousStep">Vorherige Frage</button>
						<button class="btn btn-primary" @click="nextStep">Nächste Frage</button>
					</div>
				</div>
			</div>
		</div>
	</div>

	<!-- Result -->
	<div v-if="state.finished">
		<div class="d-flex justify-content-center mt-5">
			<h4>Du hast {{ state.result }} von {{ questions.length }} Fragen richtig beantwortet!</h4>
		</div>
		<div class="d-grid gap-2 mt-5">
			<button class="btn btn-primary" @click="resetTest">Neustart</button>
		</div>
	</div>
</template>

<script setup>
import { mapGetters, mapActions } from "@/store/map-state";
import { computed, reactive } from "vue";

const { username, questions, question_title } = mapGetters();
const { updateUser } = mapActions();

const state = reactive({
	step: 0,
	result: 0,
	elapsedTime: 0,
	finished: false
});

// Timer
const startTimer = setInterval(() => {
	state.elapsedTime += 1000;
}, 1000);
const timer = computed(() => {
	const date = new Date(null);
	date.setSeconds(state.elapsedTime / 1000);
	const utc = date.toUTCString();
	return utc.substr(utc.indexOf(":") - 2, 8);
});

// Evaluation
const evaluate = () => {
	let result = [];
	let request = { username: username.value, useranswers: [] };

	for (const i in questions.value) {
		let real = [];
		let user = [];
		for (const j in questions.value[i].Answers) {
			real.push(questions.value[i].Answers[j].correct);
			user.push(questions.value[i].Answers[j].marked);

			request.useranswers.push({
				answer_id: questions.value[i].Answers[j].id,
				marked: questions.value[i].Answers[j].marked
			});
		}
		var is_same = real.every(function (element, index) {
			return element === user[index];
		});
		result.push(is_same);
	}

	clearInterval(startTimer); // stop timer
	state.result = result.filter(Boolean).length;
	state.finished = true;

	updateUser(request);
};
const previousStep = () => {
	if (state.step === 0) return;
	state.step--;
};
const nextStep = () => {
	if (state.step === questions.length - 1) return;
	state.step++;
};
const resetTest = () => {
	state.step = 0;
	state.result = 0;
	state.elapsedTime = 0;
	state.finished = false;
	for (const i in questions.value) {
		for (const j in questions.value[i].Answers) {
			questions.value[i].Answers[j].marked = false;
		}
	}
};
</script>

<style scoped>
.timer {
	position: absolute;
	top: 100px;
	right: 20px;
	font-size: 18px;
}
</style>
