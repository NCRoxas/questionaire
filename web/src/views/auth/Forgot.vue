<template>
	<div class="container">
		<h1>Passwort zurücksetzen</h1>
		<form class="row g-1" @submit.prevent>
			<div class="col-md-12">
				<div class="form-floating mb-3">
					<input
						type="text"
						class="form-control"
						id="loginInput"
						v-model="login"
						placeholder="name@example.com"
						required
					/>
					<label for="loginInput">Benutzername oder Email</label>
				</div>
			</div>
			<button @click="reset(login)" type="submit" class="btn btn-primary">Anfrage schicken</button>
		</form>
	</div>
</template>

<script setup>
import { ref } from "vue";
import { mapActions } from "@/store/map-state";

const login = ref(null);
const { reset } = mapActions();
</script>
