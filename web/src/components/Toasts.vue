<template>
	<div class="toast-container position-fixed top-0 end-0 p-3">
		<div
			class="toast align-items-center bg-success text-white border-0"
			role="alert"
			aria-live="assertive"
			aria-atomic="true"
			ref="toast"
			v-bind:class="classObject"
		>
			<div class="d-flex">
				<div class="toast-body">{{ authStatusMsg }}</div>
				<button
					type="button"
					class="btn-close btn-close-white me-2 m-auto"
					data-bs-dismiss="toast"
					aria-label="Close"
				></button>
			</div>
		</div>
	</div>
</template>

<script setup>
import { Toast } from "bootstrap";
import { mapGetters } from "@/store/map-state";

const { authStatus, authStatusMsg } = mapGetters();

const show = () => {
	new Toast(this.$refs.toast, Toast.options).show();
};
const classObject = () => {
	return {
		"bg-danger": authStatus === "danger",
		"bg-success": authStatus === "success"
	};
};
</script>
