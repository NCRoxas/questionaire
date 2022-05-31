import api from "./api";

export const pack = {
	state: {
		info: null,
		packs: null,
		categories: null,
		subcategories: null,
		questions: null,
		question_title: null,
		question_pages: 1,
		answers: null,
		result: {},
		total: 0,
		editmode: false
	},
	actions: {
		/*
            //////////////////////////////////////////////////////////////////////////////////////
            // Pack API
            // getPacks, newPack, newPackBatch, updatePack, deletePack, deletePackBatch
        */
		async getUserPacks({ commit }, data) {
			try {
				let response = await api.get(`/user?username=${data}`);
				commit("PACKS", response.data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async getPacks({ commit }) {
			try {
				let response = await api.get("/pack/batch");
				commit("PACKS", response.data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async newPack({ commit }, data) {
			try {
				let response = await api.post("/pack", data);
				commit("PACKS", response.data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async updatePack({ commit }, data) {
			try {
				let response = await api.patch("/pack", data);
				commit("PACKS", response.data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async deletePack({ commit }, data) {
			try {
				await api.delete("/pack", { data: data });
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		/*
            //////////////////////////////////////////////////////////////////////////////////////
            // Category API
            // newCategory, newCategoryBatch, updateCategory, deleteCategory, deleteCategoryBatch
        */
		async newCategory({ commit }, data) {
			try {
				await api.post("/category", data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async newCategoryBatch({ commit }, data) {
			try {
				await api.post("/category/batch", data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async updateCategory({ commit }, data) {
			try {
				await api.patch("/category", data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async deleteCategory({ commit }, data) {
			try {
				await api.delete("/category", { data: data });
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async deleteCategoryBatch({ commit }, data) {
			try {
				await api.delete("/category/batch", {
					data: { title: data }
				});
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		/*
            //////////////////////////////////////////////////////////////////////////////////////
            // Subcategory API
            // newSubcategory, newSubcategoryBatch, updateSubcategory, deleteSubcategory, 
            // deleteSubcategoryBatch
        */
		async newSubcategory({ commit }, data) {
			try {
				await api.post("/subcategory", data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async newSubcategoryBatch({ commit }, data) {
			try {
				await api.post("/subcategory/batch", data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async updateSubcategory({ commit }, data) {
			try {
				await api.patch("/subcategory", data);
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async deleteSubcategory({ commit }, data) {
			try {
				await api.delete("/subcategory", { data: data });
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		async deleteSubcategoryBatch({ commit }, data) {
			try {
				await api.delete("/subcategory/batch", {
					data: { title: data }
				});
			} catch (error) {
				commit("ERROR_PACKS", error);
			}
		},
		/*
            //////////////////////////////////////////////////////////////////////////////////////
            // Question API
            // newQuestion, newQuestionBatch, updateQuestion, deleteQuestion, deleteQuestionBatch
        */
		async getQuestions({ commit }, data) {
			try {
				let response = await api.get(`/question?subcategory=${data}`);
				commit("QUESTIONS", response.data);
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		async getQuestionList({ commit }, data) {
			try {
				if (data === undefined) {
					data = "";
				}
				let response = await api.get(`/question/batch${data}`);
				commit("QUESTIONS", response.data);
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		async searchQuestions({ commit }, data) {
			try {
				let response = await api.get(`/question/search${data}`);
				commit("QUESTIONS", response.data);
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		async newQuestion({ commit }, data) {
			try {
				await api.post("/question", data);
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		async updateQuestion({ commit }, data) {
			try {
				await api.patch("/question", data);
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		async deleteQuestion({ commit }, data) {
			try {
				await api.delete("/question", { data: data });
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		async newAnswer({ commit }, data) {
			try {
				await api.post("/question/answer", data);
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		async deleteAnswer({ commit }, data) {
			try {
				await api.delete("/question/answer", {
					data: data
				});
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		async evaluateQuestions({ commit }, data) {
			try {
				let response = await api.post(`/question/eval?username=${data}`);
				commit("RESULT", response.data);
			} catch (error) {
				commit("ERROR_QUESTIONS", error);
			}
		},
		/*
            //////////////////////////////////////////////////////////////////////////////////////
            // Utility
        */
		toggleEdit({ commit }, data) {
			commit("EDITMODE", data);
		}
	},
	mutations: {
		PACKS(state, payload) {
			state.packs = payload.packs;
			state.categories = payload.categories;
			state.subcategories = payload.subcategories;
		},
		ERROR_PACKS(state, payload) {
			state.info = payload;
			state.packs = null;
		},
		QUESTIONS(state, payload) {
			state.questions = payload.questions;
			state.question_title = payload.title;
			state.question_pages = payload.pages;
		},
		ERROR_QUESTIONS(state, payload) {
			state.info = payload;
			state.questions = null;
		},
		RESULT(state, payload) {
			state.result = payload.result;
			state.total = payload.total;
			state.subcategories = payload.subcategories;
		},
		EDITMODE(state, payload) {
			state.editmode = payload;
		}
	},
	getters: {
		packs: (state) => state.packs,
		categories: (state) => state.categories,
		subcategories: (state) => state.subcategories,
		questions: (state) => state.questions,
		question_title: (state) => state.question_title,
		question_pages: (state) => state.question_pages,
		result: (state) => state.result,
		editmode: (state) => state.editmode,
		progress: (state) => {
			let done = 0;
			for (const i in state.result) {
				if (state.result[i].correct === true) {
					done++;
				}
			}
			let missing = state.total - done;
			return [done, missing];
		},
		category_list: (state) => (slug) => {
			let title = state.packs.filter((item) => item.slug === slug)[0].title;
			if (state.categories !== null) {
				let data = state.categories.filter((item) => item.pack_slug === slug);
				return { title: title, categories: data };
			} else {
				return { title: title };
			}
		},
		subcategory_list: (state) => (slug) => {
			let title = state.categories.filter((item) => item.slug === slug)[0].title;
			if (state.subcategories !== null) {
				let data = state.subcategories.filter((item) => item.category_slug === slug);
				return { title: title, subcategories: data };
			} else {
				return { title: title };
			}
		}
	}
};
