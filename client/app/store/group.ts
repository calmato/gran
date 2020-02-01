export const actions = {
  create({ _ }, form) {
    return new Promise((resolve, reject) => {
      this.$axios
        .post('/v1/groups', form)
        .then(() => resolve())
        .catch((err) => reject(new Error(err)))
    })
  }
}
