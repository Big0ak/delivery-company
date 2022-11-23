import { AxiosInstance } from "./axios";
import { ILoginUser, IOrder, IClient } from "../shared/interfaces";

// регистрация клиента
export const sendPostClient = async (url: string, body: IClient) => {
    try {
        const response = await AxiosInstance.post(url, body);
        return response.data;
    } catch (error) {
        console.error();
    }
}

// вход под клиентом или менеджером
export const sendSignIn =  async (url: string, body: ILoginUser) => {
    try{
        const response = await AxiosInstance.post(url, body);
        return response.data;
    } catch (error) {
        console.error();
    }
}

// --------------------------------------------------------------------------

// получение массива заказов, водителей, клиентом
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
        const response = await AxiosInstance.get(url+`/${id}`);
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

export const deleteOrder = async (url: string, id: string) => {
    try {
        await AxiosInstance.delete(url+`/${id}`);
    } catch (error) {
        console.error()
    }
}

export const getInfoUser = async (url: string) => {
    try {
        const responce = await AxiosInstance.get(url);
        return responce.data
    } catch (error) {
        console.error()
    }
}