import { createStore } from 'vuex';

export function setupStore() {
  return createStore({
    state() {
      return {
        username: 0,
      };
    },
    getters: {
      getUsername(state) {
        return state.username;
      },
    },
    mutations: {
      SET_USERNAME(state, payload) {
        state.username += payload;
      },
    },
    actions: {
      setUsername(context, payload) {
        context.commit('SET_USERNAME', payload);
      },
    },
  });
}
