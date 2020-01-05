export interface IUserStore {
  user: IUser
}

export interface IUser {
  uid: string
  email: string
  creationTime: string
  lastSignInTime: string
}
