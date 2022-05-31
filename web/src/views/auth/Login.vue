<template>
	<div class="container">
		<h1>Login</h1>

		<p>Wenn Sie sich bereits auf questionaire.com registriert haben können Sie sich hier einloggen.</p>

		<form class="row g-1 needs-validation" @submit.prevent novalidate>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="loginInput"
						v-model="v$.login.$model"
						v-bind:class="{
							'is-invalid': v$.login.$dirty && v$.login.$invalid
						}"
						placeholder="name@example.com"
					/>
					<label for="loginInput">Benutzername oder Email</label>
					<div class="invalid-feedback" v-if="v$.login.required.$invalid">
						Benutzername/Email wird benötigt!
					</div>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="password"
						class="form-control"
						id="passwordInput"
						v-model="v$.password.$model"
						v-bind:class="{
							'is-invalid': v$.password.$dirty && v$.password.$invalid
						}"
						placeholder="Password"
					/>
					<label for="passwordInput">Password</label>
					<div class="invalid-feedback" v-if="v$.password.required.$invalid">Password wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-3">
				<div class="form-check mb-3">
					<input type="checkbox" class="form-check-input" id="remember" v-model="form.remember" />
					<label class="form-check-label" for="remember">Eingeloggt bleiben</label>
				</div>
			</div>
			<div class="col-md-2">
				<router-link :to="{ name: 'forgot' }">Passwort vergessen?</router-link>
			</div>
			<button @click="submit" type="submit" class="btn btn-primary">Login</button>
		</form>
	</div>
</template>

<script setup>
import { mapActions } from "@/store/map-state";
import useVuelidate from "@vuelidate/core";
import { required } from "@vuelidate/validators";
import { reactive } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const { login } = mapActions();

const form = reactive({});
const rules = {
	login: { required },
	password: { required }
};
const v$ = useVuelidate(rules, form);

const submit = async () => {
	if (!v$.value.$invalid) {
		await login(form);
		router.push({ name: "home" });
	}
};
</script>
