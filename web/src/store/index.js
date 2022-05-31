import VuexPersistence from "vuex-persist";

import { createStore } from "vuex";
import { auth } from "./auth.module";
import { pack } from "./pack.module";

const vuexStorage = new VuexPersistence({
    key: "questionaire",
    storage: window.localStorage,
});

export default createStore({
    modules: {
        auth,
        pack,
    },
    plugins: [vuexStorage.plugin],
});
