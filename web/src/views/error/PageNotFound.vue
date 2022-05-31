<template>
    <main class="bg text-center">
        <div class="position-absolute top-50 start-50 translate-middle">
            <img src="/images/logo.png" alt="logo" class="mb-5">
            <h1 class="text-light mb-4">Seite nicht gefunden...</h1>
            <router-link :to="{ name: 'startpage' }" class="btn btn-success">Zurück zur Homepage</router-link>
        </div>
    </main>
</template>

<style scoped>
.bg {
    min-height: 100vh;
    min-width: 100%;
    background-color: #154c29;
}
</style>
