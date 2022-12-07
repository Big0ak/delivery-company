export interface IManager {
    name: string;
    surname: string;
    login: string;
    password: string;
}

export interface IClient {
    id?: number;
	login: string;
	password: string;
	name: string;
	surname: string;
	phone: string;
	registrationDate?: string | null;
}

export interface IDriver {
    id: number; 
	name: string;
	surname: string;
}

export interface ILoginUser {
    login: string;
    password: string;
}

export interface IOrder {
    id?: number | null;
    clientId: number;
    driverId: number;
    managerId?: number | null;
    cargoWeight: number;
    price: number;
    departure: string;
	destination: string;
    date?: string | null;
    deliveryDate: string;
}

interface IManagerRead {
    String : string;
    Valid : boolean;
} 

export interface IOrderRead {
    id?: number | null;
    client: string;
    driver: string;
    manager: IManagerRead;
    cargoWeight: number;
    price: number;
    departure: string;
	destination: string;
    date?: string | null;
    deliveryDate: string;
}