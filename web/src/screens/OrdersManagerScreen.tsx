import React, { useEffect } from 'react'
import {FC, SyntheticEvent, useState} from 'react'
import { IOrder } from '../axios/interfaces'
import { getAllOrders } from '../axios/hooks';
import FormContainer from '../components/FormContainer'
import ListGroup from 'react-bootstrap/ListGroup';
import Badge from 'react-bootstrap/Badge';

const OrdersManagerScreen: FC = () => {
    const [orders, setOrders] = useState<IOrder[]>([])


    useEffect(() => {
        const getOrders = async () => {
            const response = await getAllOrders("api/orders/")
            setOrders(response)  
        }
        getOrders();
        console.log(orders)
    }, [])

    
    return (
        <FormContainer>
            <ListGroup>
                {orders.map((order: IOrder, index: number) => (
                    <ListGroup.Item key = {order.id}>
                        <Badge bg="primary" pill>
                            №{order.id}
                        </Badge>
                        <div>
                            <div className="fw-bold"> Номер клиента {order.clientId}</div>
                                Стоимость заказа {order.price}
                        </div>  
                        
                    </ListGroup.Item>
                ))}
            </ListGroup>
        </FormContainer>
    ) 
}

export default OrdersManagerScreen