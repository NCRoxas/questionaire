import { createRouter, createWebHistory } from "vue-router";
import store from "@/store";

import Startpage from "@/views/Startpage.vue";

const Home = () => import("@/views/main/Home.vue");
const Category = () => import("@/views/main/Category.vue");
const Subcategory = () => import("@/views/main/Subcategory.vue");
const Testmode = () => import("@/views/main/Testmode.vue");
const Testresults = () => import("@/views/main/Testresults.vue");
const Helpmode = () => import("@/views/main/Helpmode.vue");
const Profile = () => import("@/views/main/Profile.vue");

const Register = () => import("@/views/auth/Register.vue");
const Login = () => import("@/views/auth/Login.vue");
const Forgot = () => import("@/views/auth/Forgot.vue");
const Reset = () => import("@/views/auth/Reset.vue");

const Usertools = () => import("@/views/admin/Usertools.vue");
const Batchtools = () => import("@/views/admin/Batchtools.vue");
const Questiontools = () => import("@/views/admin/Questiontools.vue");
const Dbtools = () => import("@/views/admin/Dbtools.vue");

const PageNotFound = () => import("@/views/error/PageNotFound.vue");

const routes = [
	{
		path: "/",
		name: "startpage",
		component: Startpage
	},
	{
		path: "/home",
		name: "home",
		component: Home,
		meta: {
			title: "Home - Questionaire",
			requiresAuth: true
		},
		beforeEnter: () => {
			if (store.getters.isAdmin) {
				store.dispatch("getPacks");
			} else {
				store.dispatch("getUserPacks", store.getters.username);
			}
		}
	},
	{
		path: "/results",
		name: "results",
		component: Testresults,
		meta: {
			title: "Testergebnisse - Questionaire",
			requiresAuth: true
		},
		beforeEnter: () => {
			store.dispatch("evaluateQuestions", store.getters.username);
		}
	},
	{
		path: "/profile",
		name: "profile",
		component: Profile,
		meta: {
			title: "Profil - Questionaire",
			requiresAuth: true
		},
		beforeEnter: () => {
			store.dispatch("getUser");
		}
	},
	{
		path: "/:pack",
		name: "category",
		component: Category,
		meta: {
			title: "Kategorien - Questionaire",
			requiresAuth: true,
			breadcrumb(route) {
				return [
					{
						text: "Home",
						to: { name: "home" }
					},
					{
						text: store.getters.category_list(route.params.pack).title,
						to: {
							name: "category",
							params: {
								pack: route.params.pack
							}
						}
					}
				];
			}
		}
	},
	{
		path: "/:pack/:category",
		name: "subcategory",
		component: Subcategory,
		meta: {
			title: "Unterkategorien - Questionaire",
			requiresAuth: true,
			breadcrumb(route) {
				return [
					{
						text: "Home",
						to: { name: "home" }
					},
					{
						text: store.getters.category_list(route.params.pack).title,
						to: {
							name: "category",
							params: {
								pack: route.params.pack
							}
						}
					},
					{
						text: store.getters.subcategory_list(route.params.category).title,
						to: {
							name: "subcategory",
							params: {
								category: route.params.category
							}
						}
					}
				];
			}
		}
	},
	{
		path: "/:pack/:category/:subcategory/testmode",
		name: "testmode",
		component: Testmode,
		meta: {
			title: "Testmodus - Questionaire",
			requiresAuth: true
		},
		beforeEnter: (to) => {
			store.dispatch("getQuestions", to.params.subcategory);
		}
	},
	{
		path: "/:pack/:category/:subcategory/helpmode",
		name: "helpmode",
		component: Helpmode,
		meta: {
			title: "Lernmodus - Questionaire",
			requiresAuth: true
		},
		beforeEnter: (to) => {
			store.dispatch("getQuestions", to.params.subcategory);
		}
	},
	// Authentication section --------------------------------------
	{
		path: "/register",
		name: "register",
		component: Register,
		meta: {
			title: "Registrieren - Questionaire"
		}
	},
	{
		path: "/login",
		name: "login",
		component: Login,
		meta: {
			title: "Login - Questionaire"
		},
		beforeEnter: () => {
			if (store.getters.isLoggedIn) {
				router.push({ name: "home" });
			}
		}
	},
	{
		path: "/forgot",
		name: "forgot",
		component: Forgot,
		meta: {
			title: "Passwort vergessen - Questionaire"
		}
	},
	{
		path: "/reset/:token",
		name: "reset",
		component: Reset,
		meta: {
			title: "Passwort zurücksetzen - Questionaire"
		}
	},
	// Admin section -----------------------------------------------
	{
		path: "/admin/usertools",
		name: "usertools",
		component: Usertools,
		meta: {
			title: "Tools - Questionaire",
			requiresAuth: true,
			requiresAdmin: true
		}
	},
	{
		path: "/admin/batchtools",
		name: "batchtools",
		component: Batchtools,
		meta: {
			title: "Tools - Questionaire",
			requiresAuth: true,
			requiresAdmin: true
		}
	},
	{
		path: "/admin/qtools",
		name: "qtools",
		component: Questiontools,
		meta: {
			title: "Tools - Questionaire",
			requiresAuth: true,
			requiresAdmin: true
		}
	},
	{
		path: "/admin/dbtools",
		name: "dbtools",
		component: Dbtools,
		meta: {
			title: "Tools - Questionaire",
			requiresAuth: true,
			requiresAdmin: true
		}
	},
	// Utility section ---------------------------------------------
	{
		path: "/:catchAll(.*)*",
		name: "pagenotfound",
		component: PageNotFound,
		meta: {
			title: "404"
		}
	}
];

const router = createRouter({
	history: createWebHistory(),
	base: import.meta.env.BASE_URL,
	routes
});

// Redirect user to login for protected pages
router.beforeEach(async (to, from, next) => {
	await store.restored;
	document.title = to.meta.title ? to.meta.title : "Questionaire";

	if (to.matched.some((record) => record.meta.requiresAdmin)) {
		if (!store.getters.isAdmin) {
			next("/home");
			return;
		}
	}

	if (to.matched.some((record) => record.meta.requiresAuth)) {
		if (!store.getters.isLoggedIn) {
			next("/login");
			return;
		}
		next();
	} else {
		next();
	}
});

// Clear userdata
// router.afterEach(() => {
// 	store.dispatch("clear");
// });

export default router;
