import axios from "axios";
import Vue from "vue";
export const userList = {
  state: [
    {
      trap_id: "",
      is_admin: false
    }
  ],
  mutations: {
    setUserList(state, newState) {
      state.splice(0);
      newState.forEach((element, index) => {
        Vue.set(state, index, element);
      });
    }
  },
  actions: {
    async getUserList({ commit }) {
      try {
        const response = await axios.get("/api/users");
        commit("setUserList", response.data);
      } catch (err) {
        console.log(err);
        if (err.response.status === 401) {
          try {
            commit("toggleLoginDialog");
          } catch (err) {
            console.log(err);
          }
        }
      }
    }
  },
  getters: {
    trap_ids: state => {
      const trap_ids = state.map(data => data.trap_id);
      return trap_ids;
    },
    adminList: state => {
      let admin = [];
      state.forEach(user => {
        if (user.is_admin) admin.push(user.trap_id);
      });
      return admin;
    },
    notAdminList: state => {
      let notAdmin = [];
      state.forEach(user => {
        if (!user.is_admin) notAdmin.push(user.trap_id);
      });
      return notAdmin;
    }
  }
};
