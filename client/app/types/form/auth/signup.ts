import { IStringForm } from '../utils/form'

export interface ISignUpForm {
  email: IStringForm
  password: IStringForm
  passwordConfirmation: IStringForm
}

interface ISignUpFormValidate {
  email: Object
  password: Object
  passwordConfirmation: Object
}

export const SignUpForm: ISignUpForm = {
  email: {
    label: 'Email',
    value: '',
  },
  password: {
    label: 'Password',
    value: '',
  },
  passwordConfirmation: {
    label: 'Password Confirmation',
    value: '',
  },
}

export const SignUpFormValidate: ISignUpFormValidate = {
  email: {
    required: true,
    email: true,
  },
  password: {
    required: true,
  },
  passwordConfirmation: {
    required: true,
    confirmed: SignUpForm.password.label,
  },
}
