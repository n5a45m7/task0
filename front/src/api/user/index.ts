import axios from "axios"
import { IAPIConfig, IGetUserInfoResponse, IUserInfoAPI } from ".."

export class UserInfoAPI implements IUserInfoAPI {
    constructor(private readonly config: IAPIConfig) {        
    }

    async getUserInfo(id: number): Promise<IGetUserInfoResponse | Error> {
        return axios
            .get(`/userinfo?userID=${id}`, {baseURL: this.config.baseURL})
            .then(response => response.data)
            .catch(err => err)
    }
}