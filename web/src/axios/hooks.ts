import { AxiosInstance } from "./axios";
import { IManager, ILoginUser, IOrder } from "./interfaces";

export const sendPostManager = async (url: string, body: IManager) => {
    try {
        const response = await AxiosInstance.post(
                url,
                body
            );
        return response.data
    } catch (error) {
        console.error()
    }
}

export const sendSignInManager =  async (url: string, body: ILoginUser) => {
    try{
        const response = await AxiosInstance.post(
                url,
                body,
                {
                    headers: {
                        'Content-Type': 'application/json',
                        Accept: 'application/json',
                    }
                }
               
            );

        return response.data
    } catch (error) {
        console.error()
    }
}