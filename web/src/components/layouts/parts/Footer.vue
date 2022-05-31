<template>
	<footer class="footer pt-3">
		<div class="container">
			<div class="row">
				<div class="col-6 col-md-3 text-white">
					Questionaire GmbH<br />
					<a class="text-white" href="mailto:info@questionaire.com">info@questionaire.com</a>
				</div>

				<div class="col-6 col-md-3">
					<div class="d-grid gap-1">
						<a
							class="text-decoration-none text-white"
							href="#"
							>AGB</a
						>
						<a
							class="text-decoration-none text-white"
							href="#"
							>Datenschutz</a
						>
						<a
							class="text-decoration-none text-white"
							href="#"
							>Impressum</a
						>
						<a
							class="text-decoration-none text-white"
							href="#"
							>Hilfe & FAQ</a
						>
					</div>
				</div>
			</div>
		</div>
		<div class="container-fluid mt-3" style="background-color: #102815">
			<div class="row justify-content-md-center p-3">
				<div class="col-md-auto text-light">© {{ year }} questionaire.com</div>
			</div>
		</div>
	</footer>
</template>

<script setup>
const year = new Date().getFullYear();
</script>

<style scoped>
.footer {
	background-color: #154c29;
}
</style>
