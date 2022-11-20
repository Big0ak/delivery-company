import { AxiosInstance } from "./axios";
import { IManager, ILoginUser, IOrder, IOrderRead } from "../shared/interfaces";

export const sendPostManager = async (url: string, body: IManager) => {
    try {
        const response = await AxiosInstance.post(url, body);
        return response.data;
    } catch (error) {
        console.error();
    }
}

export const sendSignInManager =  async (url: string, body: ILoginUser) => {
    try{
        const response = await AxiosInstance.post(url, body);
        return response.data;
    } catch (error) {
        console.error();
    }
}

// --------------------------------------------------------------------------

export const getRequest = async (url: string) => {
    try {
        const response = await AxiosInstance.get(url);
        return response.data.data;
    } catch (error) {
        console.error();
    }
}

//------------- ORDER ----------------------------------------------------------------

export const searchOrderByCity = async (url: string, city: string) => {
    try {
        const response = await AxiosInstance.get(url+`/${city}`);
        return response.data.data;
    } catch (error) {
        console.error();
    }
}

export const getOrderId = async (url: string, id: string) => {
    try{
        const response = await AxiosInstance.get<IOrderRead>(url+`/${id}`);
        return response.data;
    } catch (error) {
        console.error();
    }
}

export const creatOrder = async (url: string, body: IOrder) => {
    try {
        const response = await AxiosInstance.post(url, body);
        return response.data;
    } catch (error) {
        console.error();
    }
}

export const editOrder = async (url: string, body: IOrder) => {
    try {
        await AxiosInstance.put(url+`/${body.id}`, body);
    } catch (error) {
        console.error();
    }
}