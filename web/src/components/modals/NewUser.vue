<template>
	<div class="modal fade" id="newUser" tabindex="-1" aria-labelledby="newUser" aria-hidden="true">
		<div class="modal-dialog modal-xl">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title">Neuen Benutzer anlegen</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>

				<div class="modal-body">
					<form class="row g-1" @submit.prevent novalidate @keydown.enter="$emit('update-user', user)">
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<input
									type="text"
									class="form-control"
									placeholder="Username"
									v-model="user.username"
									v-bind:class="{
										'is-invalid': v$.username.$dirty && v$.username.$invalid
									}"
								/>
								<label>Benutzername</label>
								<div class="invalid-feedback">Benutzername wird benötigt!</div>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<input
									type="text"
									class="form-control"
									placeholder="Vorname"
									v-model="user.firstname"
								/>
								<label>Vorname</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<input type="text" class="form-control" placeholder="Nachname" v-model="user.surname" />
								<label>Nachname</label>
							</div>
						</div>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<input type="text" class="form-control" placeholder="Straße" v-model="user.street" />
								<label>Straße</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<input type="text" class="form-control" placeholder="PLZ" v-model="user.zip" />
								<label>PLZ</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<input type="text" class="form-control" placeholder="Ort" v-model="user.city" />
								<label>Ort</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<input type="text" class="form-control" placeholder="State" v-model="user.state" />
								<label>Land</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<input
									type="email"
									class="form-control"
									placeholder="name@example.com"
									v-model="user.email"
									v-bind:class="{
										'is-invalid': v$.email.$dirty && v$.email.$invalid
									}"
								/>
								<label>Email</label>
								<div class="invalid-feedback">Email wird benötigt!</div>
							</div>
						</div>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<input
									type="date"
									class="form-control"
									placeholder="Birthday"
									:value="dateToPicker(user.birthday)"
									@change="updateBirthday"
								/>
								<label>Geburtstag</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<select class="form-select" v-model="user.role">
									<option>admin</option>
									<option>user</option>
								</select>
								<label>Rolle</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<select class="form-select" v-model="user.gender">
									<option>Frau</option>
									<option>Herr</option>
								</select>
								<label>Anrede</label>
							</div>
						</div>
						<div class="col-md-12 mb-3">
							<div class="form-check form-switch">
								<input class="form-check-input" type="checkbox" v-model="user.active" />
								<label class="form-check-label"
									>Aktiv
									<span class="fw-light fst-italic"
										>Hindert den Nutzer daran sich einzuloggen wenn deaktiviert!</span
									>
								</label>
							</div>
						</div>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<input type="password" class="form-control" placeholder="Passwort" v-model="password" />
								<label>Neues Passwort</label>
							</div>
						</div>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<input
									type="password"
									class="form-control"
									placeholder="RePasswort"
									v-model="repeatPassword"
									v-bind:class="{
										'is-invalid': pwCheck
									}"
								/>
								<label>Passwort wiederholen</label>
								<div class="invalid-feedback">Passwörter sind nicht gleich!</div>
							</div>
						</div>
					</form>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Schließen</button>
					<button type="button" class="btn btn-primary" @click="submit">Speichern</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
import useVuelidate from "@vuelidate/core";
import { required, email } from "@vuelidate/validators";
import { computed, reactive, ref } from "vue";

const emit = defineEmits(["new-user"]);

const user = reactive({});
const password = ref(null);
const repeatPassword = ref(null);

const rules = {
	username: { required },
	email: { required, email }
};
const v$ = useVuelidate(rules, user, { $autoDirty: true });

const pwCheck = computed(() => {
	if (password.value !== repeatPassword.value) {
		return true;
	}
});
const dateToPicker = (data) => {
	if (data !== undefined) {
		return new Date(data).toISOString().substr(0, 10);
	}
};
const updateBirthday = (event) => {
	user.birthday = new Date(event.target.value).toISOString();
};
const submit = async () => {
	if (!v$.value.$invalid && !pwCheck.value) {
		let request = { ...user, password: password.value };
		emit("new-user", request);
	}
};
</script>
