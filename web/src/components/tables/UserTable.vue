<template>
	<div class="row">
		<h1 class="fw-light">Nutzertabelle</h1>
		<span class="border-bottom border-dark border-2 mb-4"></span>

		<!-- Entries & Search -->
		<div class="d-flex mb-4">
			<div>
				<select class="form-select" aria-label="Entries" v-model="recordsPerPage" @click="updateTable">
					<option selected>5</option>
					<option>10</option>
					<option>25</option>
					<option>50</option>
					<option>100</option>
				</select>
			</div>
			<div class="ms-auto">
				<input
					type="text"
					class="form-control"
					id="searchInput"
					placeholder="Suche..."
					@input="searchTable($event.target.value)"
				/>
			</div>
		</div>

		<!-- Table -->
		<div class="table-responsive">
			<table class="table table-hover">
				<thead>
					<tr class="table-dark">
						<th scope="col">#</th>
						<th scope="col">Benutzername</th>
						<th scope="col">Vorname</th>
						<th scope="col">Nachname</th>
						<th scope="col">Email</th>
						<th scope="col">Bearbeiten</th>
					</tr>
				</thead>
				<tbody>
					<tr v-for="(user, id) in users" :key="id">
						<th scope="row">{{ user.id }}</th>
						<td>{{ user.username }}</td>
						<td>{{ user.firstname }}</td>
						<td>{{ user.surname }}</td>
						<td>{{ user.email }}</td>
						<td>
							<div class="btn-group" role="group" aria-label="Basic example">
								<button
									type="button"
									class="btn btn-outline-primary btn-sm"
									data-bs-toggle="modal"
									:data-bs-target="'#userModal' + id"
								>
									Update
								</button>
								<button
									type="button"
									class="btn btn-outline-success btn-sm"
									data-bs-toggle="modal"
									:data-bs-target="'#userPack' + id"
								>
									Pakete
								</button>
								<button type="button" class="btn btn-outline-danger btn-sm" @click="remove(user)">
									Löschen
								</button>
							</div>
						</td>
						<UpdateUser :element="user" :target="'userModal' + id" @update-user="update" />
						<UserPack
							:element="user"
							:target="'userPack' + id"
							@update-userpack="addPack"
							@delete-userpack="deletePack"
						/>
					</tr>
				</tbody>
			</table>
		</div>

		<!-- Pagination -->
		<nav class="d-flex" aria-label="Page navigation">
			<ul class="pagination">
				<li v-if="currentPage === 1" class="page-item disabled">
					<a class="page-link" href="#" @click="previousPage">Previous</a>
				</li>
				<li v-else class="page-item">
					<a class="page-link" href="#" @click="previousPage">Previous</a>
				</li>
				<li
					class="page-item"
					v-for="entry in limitation"
					:key="entry"
					v-bind:class="{ active: entry == currentPage }"
				>
					<a class="page-link" href="#" @click="thisPage(entry)">{{ entry }}</a>
				</li>
				<li v-if="currentPage === user_pages" class="page-item disabled">
					<a class="page-link" href="#" @click="nextPage">Next</a>
				</li>
				<li v-else class="page-item">
					<a class="page-link" href="#" @click="nextPage">Next</a>
				</li>
			</ul>
			<div class="ms-auto list-group me-3">
				<button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#newUser">
					<icon icon="fa-solid:user" :inline="true" class="me-2" />
					Neuer Nutzer
				</button>
				<new-user @new-user="create" />
			</div>
		</nav>
	</div>
</template>

<script setup>
import { mapActions, mapGetters } from "@/store/map-state";
import UpdateUser from "@/components/modals/UpdateUser.vue";
import UserPack from "@/components/modals/UserPack.vue";
import NewUser from "@/components/modals/NewUser.vue";

import { ref } from "vue";

const { users, user_pages } = mapGetters();
const { getUserList, searchUsers, register, updateUser, deleteUser, deleteUserRels } = mapActions();

const limitation = ref([]);
const currentPage = ref(1);
const recordsPerPage = ref(5);

const previousPage = () => {
	if (currentPage.value === 0) return;
	currentPage.value--;
	getUserList(`?limit=${recordsPerPage.value}&page=${currentPage.value}`).then(() => {
		setLimit(currentPage.value - 5, currentPage.value + 5);
	});
};
const nextPage = () => {
	if (currentPage.value === user_pages) return;
	currentPage.value++;
	getUserList(`?limit=${recordsPerPage.value}&page=${currentPage.value}`).then(() => {
		setLimit(currentPage.value - 5, currentPage.value + 5);
	});
};
const thisPage = (page) => {
	currentPage.value = page;
	getUserList(`?limit=${recordsPerPage.value}&page=${page}`).then(() => {
		setLimit(currentPage.value - 5, currentPage.value + 5);
	});
};
const updateTable = () => {
	currentPage.value = 1;
	getUserList(`?limit=${recordsPerPage.value}`).then(() => {
		setLimit(currentPage.value - 5, currentPage.value + 5);
	});
};
const searchTable = (data) => {
	if (data !== "") {
		searchUsers(`?limit=100&search=${data}`);
		currentPage.value = 1;
		limitation.value = [1];
	} else {
		getUserList(`?limit=${recordsPerPage.value}`).then(() => {
			setLimit(currentPage.value - 5, currentPage.value + 5);
		});
	}
};
const addPack = (user, change) => {
	if (change.userpacks[0].pack_id > 0) {
		user.userpacks.push(change.userpacks[0]);

		let formatted = new Date(change.userpacks[0].expiry).toISOString();
		change.userpacks[0].expiry = formatted;
		updateUser(change).then(() => getUserList(`?limit=${recordsPerPage.value}`));
	}
};
const deletePack = (user, id) => {
	const index = user.userpacks.findIndex((i) => {
		return i.ID === id;
	});
	!!user.userpacks.splice(index, 1);
	let request = { username: user.username, userpacks: [{ pack_id: id }] };
	deleteUserRels(request).then(() => getUserList(`?limit=${recordsPerPage.value}`));
};
const create = (data) => {
	register(data).then(() => getUserList(`?limit=${recordsPerPage.value}`));
};
const update = (data) => {
	updateUser(data).then(() => getUserList(`?limit=${recordsPerPage.value}`));
};
const remove = (data) => {
	const index = users.value.findIndex((i) => {
		return i.id === data.id;
	});
	!!users.value.splice(index, 1);
	deleteUser(data);
};
const setLimit = (start, stop) => {
	let rStart = start < 1 ? 1 : start;
	let rStop = stop > user_pages.value ? user_pages.value + 1 : stop;
	limitation.value = new Array(rStop - rStart).fill(rStart).map((n, i) => n + i);
};
getUserList(`?limit=${recordsPerPage.value}`).then(() => {
	setLimit(currentPage.value - 5, currentPage.value + 5);
});
</script>
