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

//------------- ORDER ----------------------------------------------------------------

export const getAllOrders = async (url: string) => {
    try {
        const response = await AxiosInstance.get(url,
            {
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json',
                    Authorization: 'Bearer ' + localStorage.getItem("JWT")
                }
            })

        return response.data.data
    } catch (error) {
        console.error()
    }
}

export const getAllClient = async (url: string) => {
    try {
        const response = await AxiosInstance.get(url,
            {
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json',
                    Authorization: 'Bearer ' + localStorage.getItem("JWT")
                }
            })

        return response.data.data
    } catch (error) {
        console.error()
    }
}

export const getAllDriver = async (url: string) => {
    try {
        const response = await AxiosInstance.get(url,
            {
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json',
                    Authorization: 'Bearer ' + localStorage.getItem("JWT")
                }
            })

        return response.data.data
    } catch (error) {
        console.error()
    }
}

export const searchOrderByCity = async (url: string, city: string) => {
    try {
        const response = await AxiosInstance.get(url+`/${city}`,
            {
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json',
                    Authorization: 'Bearer ' + localStorage.getItem("JWT")
                }
            })

        return response.data.data
    } catch (error) {
        console.error()
    }
}

export const getOrderId = async (url: string, id: string) => {
    try{
        const response = await AxiosInstance.get(url+`/${id}`,
            {
                headers: {
                    Accept: 'application/json',
                    Authorization: 'Bearer ' + localStorage.getItem("JWT")
                }
            })

        return response.data
    } catch (error) {
        console.error()
    }
}

export const creatOrder = async (url: string, body: IOrder) => {
    try {
        const response = await AxiosInstance.post(
            url,
            body,
            {
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + localStorage.getItem("JWT")
                }
            }
        );
        return response.data
    } catch (error) {
        console.error()
    }
}

export const editOrder = async (url: string, body: IOrder) => {
    try {
        await AxiosInstance.put(url+`/${body.id}`, body,
            {
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + localStorage.getItem("JWT")
                }
            }
        )
    } catch (error) {
        console.error()
    }
}