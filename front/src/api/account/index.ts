import axios from "axios"
import { IAPIConfig, IAccountCreateAPI, ICreateAccountResponse, ICreateAccountRequest } from ".."

export class AccountCreateAPI implements IAccountCreateAPI {
    constructor(private readonly config: IAPIConfig) {        
    }

    async createAccount(req: ICreateAccountRequest): Promise<ICreateAccountResponse | Error> {
        return axios
            .post(
                `/account`, 
                {
                    customerID: req.customerID,
                    initialCredit: req.initialCredit,
                },
                {baseURL: this.config.baseURL},
            )
            .then(response => response.data)
            .catch(err => err)
    }
}