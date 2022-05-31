<template>
	<div class="container">
		<h1>Neues Passwort</h1>
		<form class="row g-1" @submit.prevent>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="loginInput"
						v-model="password"
						placeholder="name@example.com"
						required
					/>
					<label for="loginInput">Passwort</label>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="loginInput"
						v-model="repeatPassword"
						placeholder="name@example.com"
						required
					/>
					<label for="loginInput">Passwort wiederholen</label>
				</div>
			</div>
			<button @click="submit" type="submit" class="btn btn-primary">Passwort ändern</button>
		</form>
	</div>
</template>

<script setup>
import { ref } from "vue";
import { mapActions } from "@/store/map-state";
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();
const password = ref(null);
const repeatPassword = ref(null);
const { verify } = mapActions();

const submit = () => {
	if (password.value === repeatPassword.value) {
		let request = {
			password: password.value,
			token: route.params.token
		};
		verify(request);
		router.push({ name: "login" });
	}
};
</script>
