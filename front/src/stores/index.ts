// --
// user

export interface IUser {
  id: number
}

export interface IAccountTransaction {
    id: number
    amount: number
}

export interface IUserAccount {
    id: number
    balance: number
    txs: IAccountTransaction[]
}

export interface IUserStore {
  id: number
  name: string
  surname: string
  accounts: IUserAccount[]

  load: (user: IUser) => void

  loading: boolean
  error: string
}

// --
// account

export interface ICreateAccountRequest {
  customerID: number
  initialCredit: number
}

export interface ICreateAccountResponse {
  id: number
}

export interface ICreateAccountStore {
  create: (req: ICreateAccountRequest) => void
  reset: () => void

  response: ICreateAccountResponse
  loading: boolean
  error: string
}