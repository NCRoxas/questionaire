<template>
	<nav class="navbar navbar-expand-lg navbar-dark">
		<div class="container-fluid">
			<router-link class="navbar-brand p-3" to="/">
				<img src="/images/logo.png" alt="logo" width="120" />
			</router-link>

			<ul class="navbar-nav ms-auto">
				<li class="nav-item me-3">
					<router-link
						:to="{
							name: 'subcategory',
							params: closeLink
						}"
						class="nav-link text-light"
						><icon icon="fa-solid:times" :inline="true"
					/></router-link>
				</li>
			</ul>
		</div>
	</nav>
</template>

<script setup>
import { useRoute } from "vue-router";
const route = useRoute();
const closeLink = { category: route.params.category };
</script>
