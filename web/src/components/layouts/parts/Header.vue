<template>
	<nav class="navbar navbar-expand-lg navbar-dark">
		<div class="container-fluid">
			<router-link class="navbar-brand p-2" :to="{ name: 'startpage' }">
				<h3 class="fst-italic fw-bold">Logo</h3>
			</router-link>

			<button
				class="navbar-toggler"
				type="button"
				data-bs-toggle="collapse"
				data-bs-target="#navbarSupportedContent"
				aria-controls="navbarSupportedContent"
				aria-expanded="false"
				aria-label="Toggle navigation"
			>
				<span class="navbar-toggler-icon"></span>
			</button>

			<div class="collapse navbar-collapse" id="navbarSupportedContent">
				<ul v-if="isLoggedIn" class="navbar-nav me-auto">
					<li class="nav-item">
						<router-link class="nav-link active" :to="{ name: 'home' }">
							<icon icon="fa-solid:home" :inline="true" class="me-2" />
							Übersicht
						</router-link>
					</li>
					<li class="nav-item">
						<router-link class="nav-link active" :to="{ name: 'results' }">
							<icon icon="fa-solid:poll" :inline="true" class="me-2" />
							Testergebnisse
						</router-link>
					</li>
					<li class="nav-item">
						<router-link class="nav-link active" :to="{ name: 'profile' }">
							<icon icon="fa-solid:cog" :inline="true" class="me-2" />
							Einstellungen
						</router-link>
					</li>
					<li class="nav-item ms-3" v-if="isAdmin">
						<router-link class="nav-link btn btn-warning text-dark" :to="{ name: 'usertools' }">
							<icon icon="fa-solid:lock" :inline="true" class="me-2" />
							Admin Modus
						</router-link>
					</li>
					<li class="nav-item ms-3" v-if="isAdmin">
						<button
							class="nav-link btn text-dark"
							@click="toggleEdit(!editmode)"
							v-bind:class="editmode ? 'btn-danger' : 'btn-info'"
						>
							<icon icon="fa-solid:pen" :inline="true" class="me-2" /> Bearbeiten
						</button>
					</li>
				</ul>
				<ul v-if="!isLoggedIn" class="navbar-nav ms-auto">
					<li class="nav-item me-3">
						<router-link :to="{ name: 'register' }" class="nav-link text-light">Registrieren</router-link>
					</li>
					<li class="nav-item me-3">
						<router-link :to="{ name: 'login' }" class="nav-link text-light">Login</router-link>
					</li>
				</ul>
				<ul v-if="isLoggedIn" class="d-flex navbar-nav">
					<li class="nav-item me-3">
						<router-link :to="{ name: 'login' }" class="nav-link text-light" @click="logout"
							><icon icon="fa-solid:sign-out-alt" :inline="true" />
						</router-link>
					</li>
				</ul>
			</div>
		</div>
	</nav>
</template>

<script setup>
import { mapActions, mapGetters } from "@/store/map-state";

const { isLoggedIn, isAdmin, editmode } = mapGetters();
const { logout, toggleEdit } = mapActions();
</script>

<style scoped>
nav {
	background-color: #154c29;
}
.pkg-dropdown {
	background-color: #1b773d;
}
.dropdown-item {
	color: white;
}
</style>
