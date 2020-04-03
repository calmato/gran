import { IStringForm } from '../utils/form'

export interface IBoardListForm {
  name: IStringForm
  color: IStringForm
}

export interface IBoardListFormValidate {
  name: Object
}

export const BoardListForm: IBoardListForm = {
  name: {
    label: 'Name',
    value: '',
  },
  color: {
    label: 'Color',
    value: '',
  },
}

export const BoardListFormValidate: IBoardListFormValidate = {
  name: {
    required: true,
  },
}
