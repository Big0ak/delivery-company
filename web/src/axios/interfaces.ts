export interface IManager {
    name: string;
    surname: string;
    login: string;
    password: string;
}

export interface ILoginUser {
    login: string;
    password: string;
}

export interface IOrder {
    id: number;
    clientId: number;
    routeId: number;
    driverid: number;
    managerId?: number;
    cargoWeight: number;
    price: number;
    date: string;
}