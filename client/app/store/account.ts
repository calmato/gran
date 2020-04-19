export const actions = {
  edit({ _ }, form) {
    const params: any = {
      name: form.name.value as string,
      displayName: form.userName.value as string,
      email: form.email.value as string,
      phoneNumber: form.phone.value as string,
      biography: form.profile.value as string,
    }

    return new Promise((resolve, reject) => {
      this.$axios
        .patch('/v1/users/profile', params)
        .then(() => resolve())
        .catch((err) => reject(new Error(err)))
    })
  },
}
