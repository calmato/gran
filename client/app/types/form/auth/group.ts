import { IStringForm } from '../utils/form'

export interface IGroupForm {
  name: IStringForm
  description: IStringForm
}

interface IGroupFormValidate {
  name: Object
  description: Object
}

export const GroupNewForm: IGroupForm = {
  name: {
    label: 'グループ名',
    value: ''
  },
  description: {
    label: 'グループの説明',
    value: ''
  }
}

export const GroupNewFormValidate: IGroupFormValidate = {
  name: {
    required: true
  },
  description: {
    required: true
  }
}
