import { createStore } from 'vuex';

export interface User {
  username: string;
}

export interface State {
  user: User | null
}

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
  });
}
