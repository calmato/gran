export const state = () => ({
  board: {},
})

export const getters = {
  board: (state) => state.board,
  lists: (state) => state.board.lists,
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
        commit('addColumn', {
          ...newColumn,
          id: res.data.id,
          tasks: [],
        })
      })
      .catch((err) => {
        return Promise.reject(err)
      })
  },

  // ColumnにTaskを追加
  addNewTask({ commit, state }, formData): Promise<void> {
    const groupId: string = state.board.groupId
    const boardId: string = state.board.id
    const listId: string = state.board.lists[formData.index].id

    const newTask = {
      index: formData.index,
      task: {
        name: formData.value,
        listId,
        labels: [],
      },
    }

    return this.$axios
      .post(`/v1/groups/${groupId}/boards/${boardId}/tasks`, newTask.task)
      .then((res) => {
        commit('addTask', {
          index: formData.index,
          task: res.data,
        })
      })
      .catch((err) => {
        return Promise.reject(err)
      })
  },

  // kanban更新の関数
  updateKanban({ _commit, state }): Promise<void> {
    const groupId: string = state.board.groupId
    const boardId: string = state.board.id

    return this.$axios
      .patch(`/v1/groups/${groupId}/boards/${boardId}/kanban`, { lists: getters.lists(state) })
      .then((_res) => {})
      .catch((err) => {
        return Promise.reject(err)
      })
  },
}
