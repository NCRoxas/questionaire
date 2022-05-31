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
				<ul v-if="isLoggedIn" class="navbar-nav me-auto mb-2 mb-lg-0">
					<li class="me-3 nav-item">
						<router-link class="nav-link text-dark btn btn-info" :to="{ name: 'home' }">
							<icon icon="fa-solid:arrow-left" :inline="true" class="me-2" />
							Zurück
						</router-link>
					</li>
					<li class="nav-item">
						<router-link class="nav-link active" :to="{ name: 'usertools' }">
							<icon icon="fa-solid:users" :inline="true" class="me-2" />
							Nutzerverwaltung
						</router-link>
					</li>
					<li class="nav-item">
						<router-link class="nav-link active" :to="{ name: 'batchtools' }">
							<icon icon="fa-solid:toolbox" :inline="true" class="me-2" />
							Massentools
						</router-link>
					</li>
					<li class="nav-item">
						<router-link class="nav-link active" :to="{ name: 'qtools' }">
							<icon icon="fa-solid:question" :inline="true" class="me-2" />
							Fragenverwaltung
						</router-link>
					</li>
					<li class="nav-item">
						<router-link class="nav-link active" :to="{ name: 'dbtools' }">
							<icon icon="fa-solid:database" :inline="true" class="me-2" />
							Datenbank
						</router-link>
					</li>
				</ul>
				<ul v-if="isLoggedIn" class="d-flex navbar-nav">
					<li class="nav-item me-3">
						<router-link :to="{ name: 'login' }" class="nav-link text-light" @click="logout"
							><icon icon="fa-solid:sign-out-alt" :inline="true" class="me-2"
						/></router-link>
					</li>
				</ul>
			</div>
		</div>
	</nav>
</template>

<script setup>
import { mapActions, mapGetters } from "@/store/map-state";

const { isLoggedIn } = mapGetters();
const { logout } = mapActions();
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
