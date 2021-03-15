import { observable, action, makeObservable } from 'mobx'
import {ICreateAccountStore, ICreateAccountRequest, ICreateAccountResponse} from '..'
import { IAccountCreateAPI } from '../../api'

const stub: ICreateAccountResponse = {id: 0}

export class CreateAccountStore implements ICreateAccountStore {
  @observable response: ICreateAccountResponse = stub
  @observable loading: boolean = false
  @observable error: string = ''

  constructor(private readonly api: IAccountCreateAPI) {
    makeObservable(this)
  }

  @action
  create =  async (req: ICreateAccountRequest): Promise<void>  => {
    try {      
      this.loading = true
      this.error = ''
      const response = await this.api.createAccount({
          customerID: req.customerID,
          initialCredit: req.initialCredit,
      })
        if (response instanceof Error) {
          console.log(response)
          this.error = response.message
          this.loading = false
          return
        }
  
        this.response = {
          id: response.id
        }
        this.loading = false
    } catch (e) {
      console.log("err in catch", e)
        if (e instanceof Error) {
          this.error = e.message
        } else {
          this.error =  "Unknown error: " + e
        }
        this.loading = false
    }
  }

  @action
  reset = (): void => {
    this.loading = false
    this.response = stub
    this.error = ''
  }
}