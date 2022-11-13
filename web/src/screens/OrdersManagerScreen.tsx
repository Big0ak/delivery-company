import React, { useEffect } from 'react'
import {FC, SyntheticEvent, useState} from 'react'
import { IOrder } from '../axios/interfaces'
import { getAllOrders } from '../axios/hooks';
import FormContainer from '../components/FormContainer'
import {getOrderId} from '../axios/hooks';

import Col from 'react-bootstrap/Col';
import ListGroup from 'react-bootstrap/ListGroup';
import Row from 'react-bootstrap/Row';
import Tab from 'react-bootstrap/Tab';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';

const OrdersManagerScreen: FC = () => {
    const [orders, setOrders] = useState<IOrder[]>([])
    const [currentOrder, setCurrentOrder] = useState<IOrder>()
    const [OrderID, setOrderID] = useState(String)

    useEffect(() => {
        const getOrders = async () => {
            const response = await getAllOrders("api/orders/")
            setOrders(response)  
        }
        getOrders();
    }, [])

    const selectOrder = async (id: string) => {
        setOrderID(id)
        if (OrderID) {
            const order = await getOrderId("api/orders", OrderID)
            setCurrentOrder(order)
            console.log(OrderID)
        }
    }
    
    return (
        <FormContainer>
            <Tab.Container id="list-group-tabs-example">
                <Row>
                    <Col sm={6}>
                        <ListGroup>
                            {orders.map((order: IOrder, index: number) => (
                                <ListGroup.Item 
                                    key = {order.id}
                                >
                                    <Badge bg="primary" pill>
                                        № {order.id}
                                    </Badge>
                                    <div>
                                        <div className="fw-bold"> Номер клиента {order.clientId}</div>
                                            Стоимость заказа {order.price}
                                    </div> 
                                    <Button onClick= {() => selectOrder(String(order.id)) }>Info</Button>{' '}
                                </ListGroup.Item>
                            ))}
                        </ListGroup>
                    </Col>

                    <Col sm={6}>
                        {
                            currentOrder ? (
                                <div>
                                    <div>
                                        Вес заказа {currentOrder.cargoWeight}
                                    </div>
                                    <Button
                                        href={"/order/" + currentOrder.id}
                                    >
                                        Edit
                                    </Button>
                                </div>
                            ) : (
                                <div>
                                    gege
                                </div>
                            )
                        }
                    </Col>
                </Row>
            </Tab.Container>
        </FormContainer>
    ) 
}

export default OrdersManagerScreen