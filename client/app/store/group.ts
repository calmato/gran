export const state = () => ({
  groups: [],
})

export const getters = {
  groups: (state) => state.groups,
}

export const mutations = {
  setGroups(state, groups) {
    state.groups = groups
  },
}

export const actions = {
  create({ _ }, form) {
    return new Promise((resolve, reject) => {
      this.$axios
        .post('/v1/groups/', form)
        .then(() => resolve())
        .catch((err) => reject(new Error(err)))
    })
  },

  groupAll({ commit }) {
    return new Promise((resolve, reject) => {
      this.$axios
        .get('/v1/groups/')
        .then((res: any) => {
          commit('setGroups', res.data.groups)
          resolve()
        })
        .catch((err: any) => reject(err))
    })
  },
}
