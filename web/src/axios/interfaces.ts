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
    id?: number | null;
    clientId: number;
    routeId: number;
    driverid: number;
    managerId?: number | null;
    cargoWeight: number;
    price: number;
    date?: string | null;
}