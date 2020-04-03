import { IStringForm } from '../utils/form'

export interface ITaskForm {
  name: IStringForm
}

export interface ITaskFormValidate {
  name: Object
}

export const TaskForm: ITaskForm = {
  name: {
    label: 'New Task Name',
    value: '',
  },
}

export const TaskFormValidate: ITaskFormValidate = {
  name: {
    required: true,
  },
}
