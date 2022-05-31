<template>
	<div class="container mb-4 mt-4">
		<h1 class="mb-3">Registrierung</h1>
		<form class="row g-1 needs-validation" @submit.prevent novalidate>
			<div class="col-md-6">
				<div class="form-floating">
					<select
						class="form-select"
						id="titleSelect"
						aria-label="title"
						v-model="v$.gender.$model"
						v-bind:class="{
							'is-invalid': v$.gender.$dirty && v$.gender.$invalid
						}"
					>
						<option value="">-</option>
						<option value="Herr">Herr</option>
						<option value="Frau">Frau</option>
					</select>
					<label for="titleSelect">Anrede</label>
					<div class="invalid-feedback">Bitte auswählen!</div>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="date"
						class="form-control"
						id="birthdayInput"
						placeholder="Birthday"
						v-model="v$.birthday.$model"
						v-bind:class="{
							'is-invalid': v$.birthday.$dirty && v$.birthday.$invalid
						}"
					/>
					<label for="birthdayInput">Geburtstag</label>
					<div class="invalid-feedback">Geburtstag wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="usernameInput"
						placeholder="Username"
						v-model.trim="v$.username.$model"
						v-bind:class="{
							'is-invalid': v$.username.$dirty && v$.username.$invalid
						}"
					/>
					<label for="usernameInput">Benutzername</label>
					<div class="invalid-feedback">Benutzername wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="firstnameInput"
						placeholder="Vorname"
						v-model="v$.firstname.$model"
						v-bind:class="{
							'is-invalid': v$.firstname.$dirty && v$.firstname.$invalid
						}"
					/>
					<label for="firstnameInput">Vorname</label>
					<div class="invalid-feedback">Vorname wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="surnameInput"
						placeholder="Nachname"
						v-model="v$.surname.$model"
						v-bind:class="{
							'is-invalid': v$.surname.$dirty && v$.surname.$invalid
						}"
					/>
					<label for="surnameInput">Nachname</label>
					<div class="invalid-feedback">Nachname wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="streetInput"
						placeholder="Straße"
						v-model="v$.street.$model"
						v-bind:class="{
							'is-invalid': v$.street.$dirty && v$.street.$invalid
						}"
					/>
					<label for="streetInput">Straße</label>
					<div class="invalid-feedback">Straße wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="zipInput"
						placeholder="PLZ"
						v-model.trim="v$.zip.$model"
						v-bind:class="{
							'is-invalid': v$.zip.$dirty && v$.zip.$invalid
						}"
					/>
					<label for="zipInput">PLZ</label>
					<div class="invalid-feedback">PLZ wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="cityInput"
						placeholder="Ort"
						v-model="v$.city.$model"
						v-bind:class="{
							'is-invalid': v$.city.$dirty && v$.city.$invalid
						}"
					/>
					<label for="cityInput">Ort</label>
					<div class="invalid-feedback">Ort wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="stateInput"
						placeholder="State"
						v-model="v$.state.$model"
						v-bind:class="{
							'is-invalid': v$.state.$dirty && v$.state.$invalid
						}"
					/>
					<label for="stateInput">Land</label>
					<div class="invalid-feedback">Land wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-6">
				<div class="form-floating mb-3">
					<input
						type="email"
						class="form-control"
						id="emailInput"
						placeholder="name@example.com"
						v-model.trim="v$.email.$model"
						v-bind:class="{
							'is-invalid': v$.email.$dirty && v$.email.$invalid
						}"
					/>
					<label for="emailInput">Email</label>
					<div class="invalid-feedback">Email wird benötigt!</div>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-hint">
					<h3>Passwort</h3>
					<p>Mindestens 8 Zeichen - am besten eine Mischung aus Buchstaben, Symbolen und Zahlen verwenden.</p>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="password"
						class="form-control"
						id="passwordInput"
						placeholder="Passwort"
						v-model="v$.password.$model"
						v-bind:class="{
							'is-invalid': v$.password.$dirty && v$.password.$invalid
						}"
					/>
					<label for="passwordInput">Passwort</label>
					<div class="invalid-feedback" v-if="v$.password.required.$invalid">Passwort wird benötigt!</div>
					<div class="invalid-feedback" v-if="v$.password.minLength.$invalid">
						Passwort muss aus mindestens 8 Zeichen bestehen!
					</div>
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="password"
						class="form-control"
						id="repasswordInput"
						placeholder="RePasswort"
						v-model="v$.repeatPassword.$model"
						v-bind:class="{
							'is-invalid': v$.repeatPassword.$dirty && v$.repeatPassword.$invalid
						}"
					/>
					<label for="repasswordInput">Passwort wiederholen</label>
					<div class="invalid-feedback" v-if="v$.repeatPassword.sameAsPassword.$invalid">
						Passwörter sind nicht gleich!
					</div>
				</div>
			</div>

			<div class="form-check">
				<input
					class="form-check-input"
					type="checkbox"
					value=""
					id="dataprotection"
					v-model="v$.dataprotection.$model"
					v-bind:class="{
						'is-invalid': v$.dataprotection.$dirty && v$.dataprotection.$invalid
					}"
				/>
				<label class="form-check-label" for="dataprotection">
					Ich stimme zu, dass meine oben genannten personenbezogenen Daten zum Zweck der Bearbeitung und
					Verwaltung der Anfrage gespeichert und verarbeitet werden. (<a
						href="#"
						target="_blank"
						>Datenschutzerklärung lesen</a
					>) *
				</label>
				<div class="invalid-feedback">Bitte akzeptieren!</div>
			</div>
			<div class="form-check">
				<input
					class="form-check-input"
					type="checkbox"
					value=""
					id="agb"
					v-model="v$.agb.$model"
					v-bind:class="{
						'is-invalid': v$.agb.$dirty && v$.agb.$invalid
					}"
				/>
				<label class="form-check-label" for="agb">
					Ich habe die
					<a href="#">AGB</a>
					gelesen und akzeptiert! *
				</label>
				<div class="invalid-feedback">Bitte akzeptieren!</div>
			</div>
			<div class="form-check">
				<input class="form-check-input" type="checkbox" value="" id="newsletter" v-model="form.newsletter" />
				<label class="form-check-label" for="newsletter">
					Ich möchte den Newsletter von Questionaire erhalten und immer up to date sein!
				</label>
			</div>

			<button @click="submit" type="submit" class="btn btn-primary mt-3">Registrieren</button>
		</form>

		<div class="toast-container position-fixed top-0 end-0 p-3">
			<div
				class="toast align-items-center text-white bg-danger border-0"
				role="alert"
				aria-live="assertive"
				aria-atomic="true"
				ref="toast"
			>
				<div class="d-flex">
					<div class="toast-body">Benutzer existiert bereits!</div>
					<button
						type="button"
						class="btn-close btn-close-white me-2 m-auto"
						data-bs-dismiss="toast"
						aria-label="Close"
					></button>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
import useVuelidate from "@vuelidate/core";
import { required, email, sameAs, minLength } from "@vuelidate/validators";
import { mapActions } from "@/store/map-state";
import { computed, reactive } from "vue";
import { useRouter } from "vue-router";

const { register } = mapActions();
const router = useRouter();

const form = reactive({});
const passRef = computed(() => form.password);
const rules = {
	username: { required },
	gender: { required },
	birthday: { required },
	firstname: { required },
	surname: { required },
	street: { required },
	zip: { required },
	city: { required },
	state: { required },
	email: { required, email },
	password: { required, minLength: minLength(8) },
	repeatPassword: { required, sameAsPassword: sameAs(passRef) },
	dataprotection: { checkMe: sameAs(true) },
	agb: { checkMe: sameAs(true) }
};
const v$ = useVuelidate(rules, form);

const submit = async () => {
	if (!v$.value.$invalid) {
		form.birthday = new Date(form.birthday).toISOString();
		await register(form);
		router.push({ name: "login" });
	}
};
</script>
