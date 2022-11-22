import { useEffect } from 'react'
import {FC, useState} from 'react'
import { IOrderRead } from '../shared/interfaces'
import { getRequest, searchOrderByCity, getOrderId} from '../axios/services';
import FormContainer from '../components/FormContainer'

import Col from 'react-bootstrap/Col';
import ListGroup from 'react-bootstrap/ListGroup';
import Row from 'react-bootstrap/Row';
import Tab from 'react-bootstrap/Tab';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';

const OrdersListScreen: FC = () => {
    const [role, useRole] = useState(localStorage.getItem("role"))
    const [orders, setOrders] = useState<IOrderRead[]>([])
    const [currentOrder, setCurrentOrder] = useState<IOrderRead>()
    const [searchCity, setSearchCity] = useState("")

    useEffect(() => {
        const getOrders = async () => {
            const response = await getRequest(`${role}-api/orders/`)
            setOrders(response)  
        }

        getOrders()
    }, [])

    const selectOrder = async (id: string) => {
        if (id) {
            const order = await getOrderId(`${role}-api/orders`, id)
            setCurrentOrder(order)
        }
    }

    const searchByCity = async () => {
        const response = await searchOrderByCity(`${role}-api/orders/search`, searchCity)
        if (response && response !== null){
            setOrders(response)
        }
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
                <Button variant="outline-secondary" id="button-search" onClick={searchByCity}>
                    Поиск
                </Button>
            </InputGroup>

            <Tab.Container id="list-group-tabs-example">
                <Row>
                    <Col sm={6}>
                        <ListGroup>
                            {orders.map((order: IOrderRead) => (
                                <ListGroup.Item 
                                    key = {order.id}
                                    onClick= {() => selectOrder(String(order.id)) }
                                >
                                    <Badge bg="badge bg-info" pill>
                                        № {order.id}
                                    </Badge>
                                    <div>
                                        <div className="fw-bold"> Маршрут: {order.departure} - {order.destination}</div>
                                        Клиент: {order.client}
                                    </div> 
                                        <Button 
                                            className="btn btn-info btn-sm"
                                            onClick= {() => selectOrder(String(order.id)) }
                                        >Подробнее
                                        </Button>
                                </ListGroup.Item>
                            ))}
                        </ListGroup>
                    </Col>

                    <Col sm={6}>
                        {
                            currentOrder ? (
                                <div>
                                    <div>
                                        <label>  Номер заказа: </label> {currentOrder.id}
                                    </div>
                                    <div>
                                        <label> Клинет: </label> {currentOrder.client}
                                    </div>
                                    <div>
                                        <label> Менеджер: </label> {currentOrder.manager.Valid ? currentOrder.manager.String: "---"}
                                    </div>
                                    <div>
                                        <label> Водитель: </label> {currentOrder.driver}
                                    </div>
                                    <div>
                                        <label> Вес:  </label> {currentOrder.cargoWeight} т.
                                    </div>
                                    <div>
                                        <label> Цена: </label> {currentOrder.price} р.
                                    </div>
                                    <div>
                                        <label> Дата: </label> {String(currentOrder.date).split('T')[0]}
                                    </div>
                                    <Button
                                        className="btn btn-outline-warning btn-link"
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

export default OrdersListScreen