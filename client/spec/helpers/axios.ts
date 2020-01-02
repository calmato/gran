import response from '~~/spec/helpers/response'

// Error を返したいときだけ false にする
let isSafetyMode: boolean = true

export default {
  setSafetyMode: (flag: boolean) => (isSafetyMode = flag),

  get: (key: any): Promise<{ data: any }> =>
    isSafetyMode
      ? Promise.resolve({ data: response['get'][key] })
      : Promise.reject(Error('some error')),

  post: (key: any): Promise<{ data: any }> =>
    isSafetyMode
      ? Promise.resolve({ data: response['post'][key] })
      : Promise.reject(Error('some error')),

  patch: (key: any): Promise<{ data: any }> =>
    isSafetyMode
      ? Promise.resolve({ data: response['patch'][key] })
      : Promise.reject(Error('some error')),

  put: (key: any): Promise<{ data: any }> =>
    isSafetyMode
      ? Promise.resolve({ data: response['put'][key] })
      : Promise.reject(Error('some error')),

  delete: (key: any): Promise<{ data: any }> =>
    isSafetyMode
      ? Promise.resolve({ data: response['delete'][key] })
      : Promise.reject(Error('some error'))
}
