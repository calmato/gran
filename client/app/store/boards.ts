export const state = () => ({
  board: {},
})

export const getters = {
  board: (state) => state.board,
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
  addTask(state, payload) {
    state.board.lists[payload.index].tasks.push(payload.task)
  },
}

export const actions = {
  getBoardById({ commit }, payload) {
    return new Promise((resolve, reject) => {
      this.$axios
        .get(`/v1/groups/${payload.groupId}/boards/${payload.boardId}`)
        .then((res: any) => {
          commit('setBoard', res.data)
          resolve()
        })
        .catch((err: any) => reject(err))
    })
  },

  // ボードのColumnを追加
  addNewColumn({ commit, state }, formData): Promise<void> {
    const newColumn = {
      name: formData.name.value,
      color: formData.color.value,
    }

    const groupId: string = state.board.groupId
    const boardId: string = state.board.id

    return this.$axios
      .post(`/v1/groups/${groupId}/boards/${boardId}/lists`, newColumn)
      .then((res) => {
        console.log(res.data)
        // todo: レスポンスで作成したIDを受け取ってそれを使ってstateの書き換えをしたい
        commit('addColumn', {
          ...newColumn,
          id: res.data.id,
          tasks: [],
        })
      })
      .catch((err) => {
        console.log(err)
      })
  },

  // ColumnにTaskを追加
  addNewTask({ commit, state }, formData) {
    const groupId: string = state.board.groupId
    const boardId: string = state.board.id
    const listId: string = state.board.lists[formData.index].id

    const newTask = {
      index: formData.index,
      task: {
        id: Date.now(),
        name: formData.value,
        listId,
        labels: [],
      },
    }

    this.$axios
      .post(`/v1/groups/${groupId}/boards/${boardId}/tasks`, newTask.task)
      .then((res) => {
        console.log(res)
        // todo: レスポンスで作成したIDを受け取ってそれを使ってstateの書き換えをしたい
        commit('addTask', { index: formData.index, task: res.data })
      })
      .catch((err) => {
        console.log(err)
      })
  },
}
