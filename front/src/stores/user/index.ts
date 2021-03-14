import { observable, action, makeObservable } from 'mobx'
import {IUserStore, IUserAccount, IUser} from '..'
import { IUserInfoAPI } from '../../api'

export class UserStore implements IUserStore {
  @observable id: number = 0
  @observable name: string = ''
  @observable surname: string = ''
  @observable accounts: IUserAccount[] = []

  @observable loading: boolean = false
  @observable error: string = ''

  constructor(private readonly api: IUserInfoAPI) {
    makeObservable(this)
  }

  @action
  load =  async ({id}: IUser): Promise<void> => {
    try {      
      this.loading = true
      this.error = ''
      this.id = id
      const response = await this.api.getUserInfo(this.id)
        if (response instanceof Error) {
          console.log(response)
          this.error = response.message
          this.loading = false
          return
        }
  
        this.name = response.name
        this.surname = response.surname
        this.accounts = (response.accounts || []).map(acc => {
          return {
            id: acc.id,
            balance: acc.balance,
            txs: (acc.txs || []).map(tx => {
              return {
                id: tx.id,
                amount: tx.amount,
              }
            })
          }
        })
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
}