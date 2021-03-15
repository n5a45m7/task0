export interface IAPIConfig {
    baseURL: string
}

interface ITx {
    id: number
    amount: number
}

interface IAccount {
    id: number
    balance: number
    txs: ITx[]
}

export interface IGetUserInfoResponse {
    id: number
    name: string
    surname: string
    accounts: IAccount[]
}


export interface IUserInfoAPI {
    getUserInfo(id: number): Promise<IGetUserInfoResponse | Error>
}

export interface ICreateAccountRequest {
    customerID: number
    initialCredit: number
}

export interface ICreateAccountResponse {
    id: number
}

export interface IAccountCreateAPI {
    createAccount(req: ICreateAccountRequest): Promise<ICreateAccountResponse | Error>
}