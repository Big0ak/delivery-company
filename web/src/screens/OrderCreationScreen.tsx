import React, { FC, SyntheticEvent, useState } from 'react'
import { IOrder } from '../axios/interfaces'
import {Form, Button} from 'react-bootstrap'
import FormContainer from '../components/FormContainer'
import { creatOrder } from '../axios/hooks';

const OrderCreationScreen: FC = () => {

    const [CliendID, setCliendID] = useState(Number)
    const [RouteID, setRouteID] = useState(Number)
    const [DriverID, setDriverID] = useState(Number)
    const [CargoWeight, setCargoWeight] = useState(Number)
    const [Price, setPrice] = useState(Number)
    const [Submitted, setSubmitted] = useState(false)
    const [OrderId, setOrderId] = useState()

    const submitHandler = async (e: SyntheticEvent) => {
      e.preventDefault()
      
      const body: IOrder = {
        clientId: CliendID,
        routeId: RouteID,
        driverid: DriverID,
        cargoWeight: CargoWeight,
        price: Price,
      }
      const response = await creatOrder("api/orders/", body)

      setOrderId(response.OrderId)
      setSubmitted(true)
    }
    
    const newOrder = () => {
      setSubmitted(false)
      setCliendID(Number)
      setRouteID(Number)
      setDriverID(Number)
      setCargoWeight(Number)
      setPrice(Number)
    }

  return (
    <FormContainer>
      {
        Submitted ? (
          <div>
            <h4>Заказ создан!</h4>
            <h3>номер заказа #{OrderId}</h3>
            <Button variant="primary" onClick={newOrder}>
              Создать заказ
            </Button>
          </div>
        ) : (
            <React.Fragment>
              <h1>Создание заказа</h1>
              <Form onSubmit={submitHandler}>

                <Form.Group className="mb-3" controlId="clientId">
                  <Form.Label>Клиент</Form.Label>
                  <Form.Control 
                    type="number" 
                    placeholder="" 
                    value={CliendID}
                    onChange={e => setCliendID(Number(e.target.value))}
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="route">
                  <Form.Label>Маршрут</Form.Label>
                  <Form.Control 
                    type="number"
                    placeholder=""
                    value={RouteID}
                    onChange={e => setRouteID(Number(e.target.value))}
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="driver">
                  <Form.Label>Водитель</Form.Label>
                  <Form.Control 
                    type="text"
                    placeholder=""
                    value={DriverID}
                    onChange={e => setDriverID(Number(e.target.value))} 
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="cargoWeight">
                  <Form.Label>Вес</Form.Label>
                  <Form.Control 
                    type="text" 
                    placeholder="" 
                    value={CargoWeight}
                    onChange={e => setCargoWeight(Number(e.target.value))}
                  />
                </Form.Group>
                
                <Form.Group className="mb-3" controlId="price">
                  <Form.Label>Цена</Form.Label>
                  <Form.Control 
                    type="text" 
                    placeholder="" 
                    value={Price}
                    onChange={e => setPrice(Number(e.target.value))}
                  />
                </Form.Group>

                <Button variant="primary" type="submit">
                  Сохранить
                </Button>
              </Form>
            </React.Fragment>
        )}
    </FormContainer>
  )
}

export default OrderCreationScreen