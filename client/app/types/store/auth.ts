export interface IUserStore {
  emailVerified: boolean
  token: string
  user: IUser
}

export interface IUser {
  uid: string
  email: string
  creationTime: string
  lastSignInTime: string
}
