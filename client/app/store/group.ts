export const actions = {
  create({ _ }, form) {
    return new Promise((resolve, reject) => {
      this.$axios
        .post('/v1/groups', form)
        .then(() => resolve())
        .catch((err) => reject(new Error(err)))
    })
  },
  groupAll({ _ }) {
    return new Promise((resolve, reject) => {
      this.$axios
        .get('/v1/groups')
        .then((res: any) => {
          console.log('res', res)
          resolve()
        })
        .catch((err: any) => reject(err))
    })
  }
}
