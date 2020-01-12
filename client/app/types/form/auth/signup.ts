export interface ISignUpForm {
  email: any
  password: any
  passwordConfirmation: any
}

interface ISignUpFormValidate {
  email: Object
  password: Object
  passwordConfirmation: Object
}

export const SignUpForm: ISignUpForm = {
  email: {
    label: 'Email',
    value: ''
  },
  password: {
    label: 'Password',
    value: ''
  },
  passwordConfirmation: {
    label: 'Password Confirmation',
    value: ''
  }
}

export const SignUpFormValidate: ISignUpFormValidate = {
  email: {
    required: true,
    email: true
  },
  password: {
    required: true
  },
  passwordConfirmation: {
    required: true,
    confirmed: SignUpForm.password.label
  }
}
