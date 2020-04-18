import { IStringForm } from '../utils/form'

export interface IAccountForm {
  name: IStringForm
  userName: IStringForm
  profile: IStringForm
  email: IStringForm
  phone: IStringForm
}

interface IAccountFormValidate {
  name: Object
  userName: Object
}

export const AccountForm: IAccountForm = {
  name: {
    label: 'Name',
    value: '',
  },
  userName: {
    label: 'Username',
    value: '',
  },
  profile: {
    label: 'Bio',
    value: '',
  },
  email: {
    label: 'Email',
    value: '',
  },
  phone: {
    label: 'Phone',
    value: '',
  },
}

export const AccountFormValidate: IAccountFormValidate = {
  name: {
    required: true,
    max: '20',
  },
  userName: {
    required: true,
    max: '20',
  },
}
