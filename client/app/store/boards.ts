export const state = () => ({
  board: {}
})

export const getters = {
  board: (state) => state.board
}

export const mutations = {
  setBoard(state, board) {
    state.board = board
  },
  updateBoardLists(state, payload) {
    state.board.lists = payload
  },
  updateBoardTasks(state, payload) {
    state.board.lists[payload.index].tasks = payload.value
  },
  addColumn(state, payload) {
    state.board.lists.push(payload)
  },
  initBoard(state) {
    state.board = {
      id: '36de3d32-4d67-4959-9056-23b6763299db',
      name: 'test board with vuex',
      closed: false,
      thumbnailUrl: '',
      backgroundColor: '#E6EE9C',
      labels: null,
      groupId: 'be02c252-d1c3-4e53-9f38-6bdef8e03e3f',
      lists: [
        {
          id: '0ai5jtQZIBtQKrNCT3Bf',
          name: 'To Do',
          color: '#009688',
          tasks: [
            {
              id: '65bv2GQawTYMY2Toatao',
              name: 'テストタスク',
              labels: [
                { name: 'Client', color: '#2196F3' },
                { name: 'API', color: '#F44336' },
                { name: 'Infra', color: '#FFEB3B' }
              ],
              assignedUserIds: null,
              deadlinedAt: '0001-01-01T00:00:00Z'
            },
            {
              id: '67bv2GQawTYMY2Toatao',
              name: 'UI設計',
              labels: [],
              assignedUserIds: null,
              deadlinedAt: '0001-01-01T00:00:00Z'
            }
          ]
        },
        {
          id: '1ai5jtQZIBtQKrNCT3Bf',
          name: 'In Prgress',
          color: '#D81B60',
          tasks: [
            {
              id: '63bv2GQawTYMY2Toatao',
              name: 'Figma',
              labels: [],
              assignedUserIds: null,
              deadlinedAt: '0001-01-01T00:00:00Z'
            },
            {
              id: '61bv2GQawTYMY2Toatao',
              name: '本番環境',
              labels: [],
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
  },
  addNewColumn({ commit }, formData) {
    const newColumn = {
      id: Date.now(),
      name: formData.name.value,
      color: formData.color.value,
      tasks: []
    }
    commit('addColumn', newColumn)
  }
}
