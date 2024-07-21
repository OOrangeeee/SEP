import Vue from 'vue';
import vuex from 'vuex';
import createPersistedState from 'vuex-persistedstate';

Vue.use(vuex);
import {getUserInfo} from "@/api";

export default  new vuex.Store({
  state: {
    userInfo:{},
    recordList:[]
  },
  mutations: {
    // ...
    SET_USERINFO(state,info){
      state.userInfo = info.userInfo;
    },
    SET_RECORD_LIST(state,list) {
      state.recordList = list
    }
  },
  actions: {
    // ...
    async setUserInfo({ commit}) {
        const res = await getUserInfo()
        commit('SET_USERINFO',res)
    },
    async setRecordList ({commit},list) {

      commit('SET_RECORD_LIST',list)
    }
  },
  plugins: [createPersistedState({
    storage: window.localStorage, // 或者 localStorage
  })]
});


