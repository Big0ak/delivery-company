import React, { useEffect } from 'react'
import {FC, SyntheticEvent, useState} from 'react'
import { IOrderRead } from '../axios/interfaces'
import { getAllOrders, searchOrderByCity, getOrderId} from '../axios/hooks';
import FormContainer from '../components/FormContainer'

import Col from 'react-bootstrap/Col';
import ListGroup from 'react-bootstrap/ListGroup';
import Row from 'react-bootstrap/Row';
import Tab from 'react-bootstrap/Tab';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';

const OrdersManagerScreen: FC = () => {
    const [orders, setOrders] = useState<IOrderRead[]>([])
    const [currentOrder, setCurrentOrder] = useState<IOrderRead>()
    const [OrderID, setOrderID] = useState(String)
    const [searchCity, setSearchCity] = useState('')

    useEffect(() => {
        const getOrders = async () => {
            const response = await getAllOrders("manager-api/orders/")
            setOrders(response)  
        }
        getOrders();
    }, [])

    const selectOrder = async (id: string) => {
        setOrderID(id)
        if (OrderID) {
            const order = await getOrderId("manager-api/orders", OrderID)
            setCurrentOrder(order)
            console.log(currentOrder)
        }
    }

    const SearchByCity = async () => {
        const response = await searchOrderByCity("manager-api/orders/search", searchCity)
        if (response && response !== null){
            setOrders(response)
        }
        console.log(response)
    }
    
    return (
        <FormContainer>
            <InputGroup className="mb-3">
                <Form.Control
                    placeholder="Поиск по городам"
                    aria-label="Search-by-city"
                    aria-describedby="basic-addon"
                    value ={searchCity}
                    onChange={e => setSearchCity(e.target.value)}  
                />
                <Button variant="outline-secondary" id="button-search" onClick={SearchByCity}>
                    Поиск
                </Button>
            </InputGroup>

            <Tab.Container id="list-group-tabs-example">
                <Row>
                    <Col sm={6}>
                        <ListGroup>
                            {orders.map((order: IOrderRead, index: number) => (
                                <ListGroup.Item 
                                    key = {order.id}
                                >
                                    <Badge bg="primary" pill>
                                        № {order.id}
                                    </Badge>
                                    <div>
                                        <div className="fw-bold"> Маршрут: {order.departure} - {order.destination}</div>
                                            Клиент: {order.client}
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
                                        <label> <strong> Номер заказа: </strong> </label> {currentOrder.id}
                                    </div>
                                    <div>
                                        <label> <strong> Клинет : </strong> </label> {currentOrder.client}
                                    </div>
                                    <div>
                                        <label> <strong> Водитель : </strong> </label> {currentOrder.driver}
                                    </div>
                                    <div>
                                        <label> <strong> Вес: </strong> </label> {currentOrder.cargoWeight} т.
                                    </div>
                                    <div>
                                        <label> <strong> Цена: </strong> </label> {currentOrder.price} р.
                                    </div>
                                    <Button
                                        href={`/order/${currentOrder.id}`}
                                    >
                                        Изменить
                                    </Button>
                                </div>
                            ) : (
                                <div>
                                    Выберете заказ для просмотра...
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