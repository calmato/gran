export const state = () => ({
  board: {}
})

export const mutations = {
  setBoard(state, board) {
    state.board = board
  },
  initBoard(state) {
    state.board = {
      id: '36de3d32-4d67-4959-9056-23b6763299db',
      name: 'test board with vuex',
      closed: false,
      thumbnailUrl: '',
      backgroundColor: '#A7FFEB',
      labels: null,
      groupId: 'be02c252-d1c3-4e53-9f38-6bdef8e03e3f',
      lists: [
        {
          id: '0ai5jtQZIBtQKrNCT3Bf',
          name: 'テスト',
          color: 'blue',
          tasks: [
            {
              id: '65bv2GQawTYMY2Toatao',
              name: 'テストタスク',
              labels: ['Client', 'API', 'Infra'],
              assignedUserIds: null,
              deadlinedAt: '0001-01-01T00:00:00Z'
            }
          ]
        }
      ],
      createdAt: '2020-02-03T12:08:30.540047Z',
      updatedAt: '2020-02-03T12:08:30.540047Z'
    }
  }
}

export const actions = {
  init({ commit }) {
    commit('initBoard')
  }
}
