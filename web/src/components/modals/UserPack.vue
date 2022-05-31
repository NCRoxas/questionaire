<template>
	<div class="modal fade" :id="target" tabindex="-1" aria-labelledby="userModal" aria-hidden="true">
		<div class="modal-dialog modal-xl">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title">
						<span class="text-capitalize">{{ change.username }}</span> - Paket hinzufügen oder löschen
					</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<form class="row g-1" @submit.prevent>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<select
									class="form-select"
									aria-label="selectPack"
									v-model="pack"
									@change="addPackInfo()"
								>
									<option selected value="">-</option>
									<option v-for="(pack, id) in packs" :key="id" :value="pack">
										{{ pack.title }}
									</option>
								</select>
								<label>Paket hinzufügen</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<input
									type="text"
									class="form-control"
									placeholder="Rechnungsnummer"
									v-model="change.userpacks[0].billnumber"
								/>
								<label>Rechnungsnummer</label>
							</div>
						</div>
						<div class="col-md-6">
							<div class="form-floating mb-3">
								<input
									type="text"
									class="form-control"
									placeholder="Betrag"
									v-model="change.userpacks[0].amount"
								/>
								<label>Betrag</label>
							</div>
						</div>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<input
									type="text"
									class="form-control"
									placeholder="Anmerkung"
									v-model="change.userpacks[0].info"
								/>
								<label>Anmerkung</label>
							</div>
						</div>
						<div class="col-md-12">
							<div class="form-floating mb-3">
								<input
									type="date"
									class="form-control"
									placeholder="expiry"
									v-model="change.userpacks[0].expiry"
								/>
								<label>Paket Ablaufdatum</label>
							</div>
						</div>
					</form>
					<ul class="list-group">
						<li class="list-group-item active mb-1" aria-current="true">Aktive Pakete:</li>
					</ul>
					<div v-for="(pack, id) in user.userpacks" :key="id" class="input-group mb-1">
						<input
							type="text"
							class="form-control"
							placeholder="Paket Titel"
							aria-label="pack title"
							:value="pack.pack_title"
							disabled
						/>
						<input
							type="text"
							class="form-control"
							placeholder="Rechnungsnummer"
							aria-label="billnumber"
							:value="'Rechnungsnummer: ' + pack.billnumber"
							disabled
						/>
						<input
							type="text"
							class="form-control"
							placeholder="Betrag"
							aria-label="amount"
							:value="'Betrag: ' + pack.amount"
							disabled
						/>
						<input
							type="text"
							class="form-control"
							placeholder="Ablaufdatum"
							aria-label="expiry"
							:value="'Ablauf: ' + dateToLocal(pack.expiry)"
							disabled
						/>
						<button
							class="btn btn-outline-danger"
							type="button"
							@click="$emit('delete-userpack', user, pack.pack_id)"
						>
							<icon icon="fa-solid:times" :inline="true" />
						</button>
					</div>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Schließen</button>
					<button type="button" class="btn btn-primary" @click="$emit('update-userpack', user, change)">
						Speichern
					</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
import { mapGetters } from "@/store/map-state";
import { ref } from "vue";

const props = defineProps({ element: Object, target: String });

const { packs } = mapGetters();

const user = ref(props.element);
const pack = ref({});
const change = ref({
	username: props.element.username,
	userpacks: [
		{
			pack_id: null,
			pack_title: null,
			billnumber: null,
			amount: null,
			info: null,
			expiry: null
		}
	]
});
const dateToLocal = (data) => {
	if (data !== undefined) {
		return new Date(data).toLocaleDateString("de-DE");
	}
};
const addPackInfo = () => {
	change.value.userpacks[0].pack_id = pack.value.id;
	change.value.userpacks[0].pack_title = pack.value.title;
};
</script>
