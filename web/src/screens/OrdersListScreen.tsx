import { useEffect } from 'react'
import {FC, useState} from 'react'
import { IOrderRead } from '../shared/interfaces'
import { getRequest, searchOrderByCity, getOrderId} from '../axios/services';
import FormContainer from '../components/FormContainer'
import { roles } from '../shared/constants'

import Col from 'react-bootstrap/Col';
import ListGroup from 'react-bootstrap/ListGroup';
import Row from 'react-bootstrap/Row';
import Tab from 'react-bootstrap/Tab';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import Nav from 'react-bootstrap/Nav';

const OrdersListScreen: FC = () => {
    const [role] = useState(localStorage.getItem("role"))
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
        if (searchCity !== ''){
            const response = await searchOrderByCity(`${role}-api/orders/search`, searchCity)
            if (response && response !== null){
                setOrders(response)
            }
        }
    }

    const getActiveOrders = async () => {
        const response = await getRequest(`${role}-api/orders/active`)
        setOrders(response)  
    }
    
    const getCompletedOrders = async () => {
        const response = await getRequest(`${role}-api/orders/completed`)
        setOrders(response) 
    }

    return (
        <FormContainer>
            <InputGroup className="mb-3">
                <Form.Control
                    placeholder="?????????? ???? ??????????????"
                    aria-label="Search-by-city"
                    aria-describedby="basic-addon"
                    value ={searchCity}
                    onChange={e => setSearchCity(e.target.value)}  
                />
                <Button variant="outline-secondary" id="button-search" onClick={searchByCity}>
                    ??????????
                </Button>
            </InputGroup>

            <Nav variant="tabs">
                <Nav.Item>
                    <Nav.Link onClick = {getActiveOrders}>????????????????</Nav.Link>
                </Nav.Item>
                <Nav.Item>
                    <Nav.Link onClick = {getCompletedOrders}>??????????????????????</Nav.Link>
                </Nav.Item>
            </Nav>

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
                                        ??? {order.id}
                                    </Badge>
                                    <div>
                                        <div className="fw-bold"> ??????????????: {order.departure} - {order.destination}</div>
                                        ????????????: {order.client}
                                    </div> 
                                        <Button 
                                            className="btn btn-info btn-sm"
                                            onClick= {() => selectOrder(String(order.id)) }
                                        >??????????????????
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
                                        <label>  ?????????? ????????????: </label> {currentOrder.id}
                                    </div>
                                    <div>
                                        <label> ????????????: </label> {currentOrder.client}
                                    </div>
                                    <div>
                                        <label> ????????????????: </label> {currentOrder.manager.Valid ? currentOrder.manager.String: "---"}
                                    </div>
                                    <div>
                                        <label> ????????????????: </label> {currentOrder.driver}
                                    </div>
                                    <div>
                                        <label> ??????:  </label> {currentOrder.cargoWeight} ??.
                                    </div>
                                    <div>
                                        <label> ????????: </label> {currentOrder.price} ??.
                                    </div>
                                    <div>
                                        <label> ???????? ????????????????: </label> {String(currentOrder.date).split('T')[0]}
                                    </div>
                                    <div>
                                        <label> ???????? ???????????????????? ????????????: </label> {String(currentOrder.deliveryDate).split('T')[0]}
                                    </div>
                                    {role === roles.manager && (
                                            <Button
                                                className="btn btn-outline-warning btn-link"
                                                href={`/order/${currentOrder.id}`}
                                            >
                                                ????????????????
                                            </Button>
                                    )}
                                </div>
                            ) : (
                                <div>
                                    ???????????????? ?????????? ?????? ??????????????????...
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