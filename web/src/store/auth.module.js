import Cookies from "js-cookie";
import api from "./api";

function parseJwt(token) {
	var base64Url = token.split(".")[1];
	var base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
	var jsonPayload = decodeURIComponent(
		atob(base64)
			.split("")
			.map(function (c) {
				return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
			})
			.join("")
	);

	return JSON.parse(jsonPayload);
}

export const auth = {
	state: {
		user: null,
		username: null,
		user_answers: null,
		role: null,
		users: null,
		user_pages: 1,
		loggedIn: false,
		status: {
			user: {
				loading: false,
				error: null
			}
		}
	},
	actions: {
		/*
            //////////////////////////////////////////////////////////////////////////////////////
            // User API
            // getUser, updateUser, login, register, reset
        */
		async getUser({ commit, state }) {
			try {
				let response = await api.get(`/user?username=${state.username}`);
				commit("UPDATE", response.data);
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async getUserList({ commit }, data) {
			try {
				if (data === undefined) {
					data = "";
				}
				let response = await api.get(`/user/batch${data}`);
				commit("ADMIN_LIST", response.data);
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async searchUsers({ commit }, data) {
			try {
				let response = await api.get(`/user/search${data}`);
				commit("ADMIN_LIST", response.data);
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async updateUser({ commit }, data) {
			try {
				let response = await api.patch("/user", data);
				commit("UPDATE", response.data);
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async deleteUser({ commit }, data) {
			try {
				await api.delete("/user", { data: data });
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async deleteUserRels({ commit }, data) {
			try {
				let response = await api.delete("/user/rels", { data: data });
				commit("UPDATE", response.data);
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async login({ commit }, data) {
			try {
				commit("STATUS", { data: "user", loading: true });
				let response = await api.post("/user/auth", data);
				commit("STATUS", { data: "user", loading: false });
				commit("LOGIN", response.data);
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async register({ commit }, data) {
			try {
				await api.post("/user", data);
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async reset({ commit }, data) {
			try {
				await api.post("/user/reset", { login: data });
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		async verify({ commit }, data) {
			try {
				await api.post("/user/verify", data);
			} catch (error) {
				commit("STATUS", { data: "user", error: error });
			}
		},
		clear({ commit }) {
			commit("CLEAR");
		},
		logout({ commit }) {
			commit("LOGOUT");
		}
	},
	mutations: {
		LOGIN(state, payload) {
			let jwt = parseJwt(payload.token);
			state.loggedIn = true;
			state.role = payload.role;
			state.username = jwt.iss;
			var expiry = (new Date(jwt.exp * 1000).getTime() - new Date().getTime()) / (1000 * 3600 * 24);
			Cookies.set("token", payload.token, { expires: Math.ceil(expiry) });
		},
		LOGOUT(state) {
			state.username = null;
			state.user = null;
			state.role = null;
			state.loggedIn = false;
			Cookies.remove("token");
		},
		UPDATE(state, payload) {
			state.user = payload;
			state.user_answers = payload.user_answers;
		},
		ADMIN_LIST(state, payload) {
			state.users = payload.users;
			state.user_pages = payload.pages;
		},
		STATUS(state, payload) {
			state.status[payload.data].loading = payload.loading !== undefined ? payload.loading : false;
			state.status[payload.data].error = payload.error !== undefined ? payload.error : null;
		},
		CLEAR(state) {
			state.user = null;
		}
	},
	getters: {
		isLoading: (state) => state.status.user.loading,
		isLoggedIn: (state) => state.loggedIn,
		isAdmin: (state) => state.role === "admin",
		username: (state) => state.username,
		user_answers: (state) => state.user_answers,
		userInfo: (state) => state.user,
		users: (state) => state.users,
		user_pages: (state) => state.user_pages
	}
};
