<template>
	<div class="container mb-4 mt-4">
		<div class="card mb-4">
			<div class="card-header">Datenbank exportieren</div>
			<div class="card-body">
				<p class="card-text">Datenbank wird als .json exportiert.</p>
				<button class="btn btn-primary" @click="exportDb">Download</button>
			</div>
		</div>
		<div class="card">
			<div class="card-header">Datenbank importieren</div>
			<div class="card-body">
				<p class="card-text">Es wird nur das .json Format akzeptiert.</p>
				<input class="form-control mb-3" type="file" id="databaseFile" ref="file" />
				<button class="btn btn-primary" @click="importDb">Upload</button>
			</div>
		</div>
	</div>
</template>

<script setup>
import { ref } from "vue";
import api from "@/store/api";

const file = ref(null);

const exportDb = async () => {
	await api.get("/backup", { responseType: "blob" }).then((response) => {
		const type = response.headers["content-type"];
		const blob = new Blob([response.data], { type: type, encoding: "UTF-8" });
		const link = document.createElement("a");
		link.href = window.URL.createObjectURL(blob);
		link.download = "db.json";
		link.click();
	});
};
const importDb = async () => {
	const formData = new FormData();
	formData.append("file", file.value.files[0]);
	const headers = { "Content-Type": "multipart/form-data" };
	await api.post("/backup", formData, { headers }).then((res) => {
		res.data.files;
		res.status;
	});
};
</script>
