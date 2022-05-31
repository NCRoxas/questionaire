<template>
	<div class="container mb-4 mt-4">
		<h1 class="mb-3">Einstellungen</h1>

		<form class="row g-1 needs-validation" @submit.prevent novalidate>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="usernameInput"
						placeholder="Username"
						v-model="form.username"
					/>
					<label for="usernameInput">Benutzername</label>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="firstnameInput"
						placeholder="Vorname"
						v-model="form.firstname"
					/>
					<label for="firstnameInput">Vorname</label>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="surnameInput"
						placeholder="Nachname"
						v-model="form.surname"
					/>
					<label for="surnameInput">Nachname</label>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="streetInput"
						placeholder="Street"
						v-model="form.street"
					/>
					<label for="streetInput">Straße</label>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="postalInput"
						placeholder="PLZ"
						v-model.trim="form.zip"
					/>
					<label for="postalInput">PLZ</label>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input type="text" class="form-control" id="cityInput" placeholder="Ort" v-model="form.city" />
					<label for="cityInput">Ort</label>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input type="text" class="form-control" id="stateInput" placeholder="State" v-model="form.state" />
					<label for="stateInput">Land</label>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="email"
						class="form-control"
						id="emailInput"
						placeholder="name@example.com"
						v-model.trim="form.email"
					/>
					<label for="emailInput">Email</label>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="date"
						class="form-control"
						id="birthdayInput"
						placeholder="Birthday"
						:value="dateToPicker(form.birthday)"
						@change="updateBirthday"
					/>
					<label for="birthdayInput">Geburtstag</label>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-hint">
					<h3>Passwort</h3>
					<p>
						Bitte nur ausfüllen wenn Sie das Passwort ändern wollen, ansonsten leer lassen! Mindestens 8
						Zeichen.
					</p>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="password"
						class="form-control"
						id="passwordInput"
						placeholder="Passwort"
						v-model="form.password"
					/>
					<label for="passwordInput">Passwort</label>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="password"
						class="form-control"
						id="repasswordInput"
						placeholder="RePasswort"
						v-model="form.repeatPassword"
						v-bind:class="{
							'is-invalid': pwCheck
						}"
					/>
					<label for="repasswordInput">Passwort wiederholen</label>
					<div class="invalid-feedback">Passwörter sind nicht gleich!</div>
				</div>
			</div>
			<div class="form-check">
				<input class="form-check-input" type="checkbox" id="newsletter" v-model="form.newsletter" />
				<label class="form-check-label" for="newsletter">
					Ich möchte den Newsletter von Questionaire erhalten und immer up to date sein!
				</label>
			</div>

			<button @click="submit" type="submit" class="btn btn-primary mt-3">Nutzerdaten ändern</button>
		</form>
	</div>
</template>

<script setup>
import useVuelidate from "@vuelidate/core";
import { email } from "@vuelidate/validators";
import { mapGetters, mapActions } from "@/store/map-state";
import { computed, reactive } from "vue";

const { userInfo } = mapGetters();
const { updateUser } = mapActions();

const form = reactive(userInfo.value);
const rules = {
	email: { email }
};
const v$ = useVuelidate(rules, form);

const pwCheck = computed(() => {
	if (form.password !== form.repeatPassword) {
		return true;
	}
});
const dateToPicker = (data) => {
	if (data !== undefined) {
		return new Date(data).toISOString().substr(0, 10);
	}
};
const updateBirthday = (event) => {
	form.birthday = new Date(event.target.value).toISOString();
};
const submit = async () => {
	if (!v$.value.$invalid && !pwCheck.value) {
		form.birthday = new Date(form.birthday).toISOString();
		await updateUser(form);
	}
};
</script>
