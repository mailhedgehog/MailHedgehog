import { createStore } from 'vuex';

export interface User {
  username: string;
  permissions: Array<string>;
}

export interface State {
  user: User | null
}

const confirmDialog = {
  namespaced: true,
  state: () => ({
    show: false,
    btnCancel: null,
    btnAccept: null,
    modalTitle: null,
    modalText: null,
    promise: null,
  }),
  getters: {
    show(state: any) {
      return state.show;
    },
    btnCancel(state: any) {
      return state.btnCancel;
    },
    btnAccept(state: any) {
      return state.btnCancel;
    },
    modalTitle(state: any) {
      return state.modalTitle;
    },
    modalText(state: any) {
      return state.modalText;
    },
  },
  mutations: {
    DISPLAY(state: any) {
      state.show = true;
    },
    ACCEPT(state: any) {
      state.show = false;
      if (state.promise && state.promise.resolve) {
        state.promise.resolve();
      }
    },
    DECLINE(state: any) {
      state.show = false;
      if (state.promise && state.promise.reject) {
        state.promise.reject();
      }
    },
    SET_CONTENT(state: any, payload: {
      btnCancel: string | null,
      btnAccept: string | null,
      modalTitle: string | null,
      modalText: string | null,
    } | null) {
      state.btnCancel = payload?.btnCancel;
      state.btnAccept = payload?.btnAccept;
      state.modalTitle = payload?.modalTitle;
      state.modalText = payload?.modalText;
    },
    SET_PROMISE(state: any, payload: {
      resolve: any,
      reject: any,
    } | null) {
      state.promise = payload;
    },
  },
  actions: {
    confirm(context: any, payload: {
      btnCancel: string | null,
      btnAccept: string | null,
      modalTitle: string | null,
      modalText: string | null,
    } | null) {
      context.commit('SET_CONTENT', payload);

      context.commit('DISPLAY');

      return new Promise((resolve, reject) => {
        context.commit('SET_PROMISE', {
          resolve,
          reject,
        });
      });
    },
    accept(context: any) {
      context.commit('ACCEPT');
    },
    decline(context: any) {
      context.commit('DECLINE');
    },
  },
};

export function setupStore() {
  return createStore<State>({
    state() {
      return {
        user: null,
      };
    },
    getters: {
      getUser(state): User | null {
        return state.user;
      },
    },
    mutations: {
      SET_USER(state, payload: User | null) {
        state.user = payload;
      },
    },
    actions: {
      setUser(context, payload: User | null) {
        context.commit('SET_USER', payload);
      },
    },
    modules: {
      confirmDialog,
    },
  });
}
