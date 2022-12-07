import React, { useEffect, FC, SyntheticEvent, useState } from 'react'
import { IClient, IOrder, IDriver } from '../shared/interfaces'
import {Form, Button} from 'react-bootstrap'
import FormContainer from '../components/FormContainer'
import { creatOrder, getRequest } from '../axios/services';

import Dropdown from 'react-bootstrap/Dropdown';

const OrderCreationScreen: FC = () => {
    const [clients, setClients] = useState<IClient[]>([])
    const [FIOclient, setFIOclient] = useState("")

    const [drivers, setDrivers] = useState<IDriver[]>([])
    const [FIOdrivers, setFIOdrivers] = useState("")

    const [CliendID, setCliendID] = useState(Number)
    const [DriverID, setDriverID] = useState(Number)
    const [CargoWeight, setCargoWeight] = useState(Number)
    const [Price, setPrice] = useState(Number)
    const [Departure, setDeparture] = useState("")
    const [Destination, setDestination] = useState("")
    const [DeliveryDate, setDeliveryDate] = useState("")

    const [Submitted, setSubmitted] = useState(false)
    const [OrderId, setOrderId] = useState()

    useEffect(() => {
      const getClients = async () => {
          const response = await getRequest("manager-api/client/")
          setClients(response)  
      }

      const getDrivers = async () => {
        const response = await getRequest("manager-api/driver/")
        setDrivers(response)
      }
      getClients();
      getDrivers();
    }, [])

    const submitHandler = async (e: SyntheticEvent) => {
      e.preventDefault()
      
      const body: IOrder = {
        clientId: CliendID,
        driverId: DriverID,
        cargoWeight: CargoWeight,
        price: Price,
        departure: Departure,
        destination: Destination,
        deliveryDate: DeliveryDate,
      }
      const response = await creatOrder("manager-api/orders/", body)

      setOrderId(response.OrderId)
      setSubmitted(true)
    }
    
    const newOrder = () => {
      setSubmitted(false)
      setFIOclient("")
      setCliendID(0)
      setFIOdrivers("")
      setDriverID(0)
      setDeparture("")
      setDestination("")
      setCargoWeight(0)
      setDeliveryDate("")
      setPrice(0)
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
                <Form.Group className="mb-3" controlId="Client">

                  <Form.Label>Клиент</Form.Label>
                  <Dropdown>
                    <Dropdown.Toggle variant="success" id="dropdown-basic">
                      {FIOclient}
                    </Dropdown.Toggle>

                    <Dropdown.Menu>
                      {clients.map((client: IClient, index: number) => (
                        <Dropdown.Item
                          key = {client.id}
                          onClick={() => {
                            setCliendID(client.id ? client.id : 0)
                            setFIOclient(client.name + " " + client.surname)
                          }}
                        >
                          {client.name} {client.surname}
                        </Dropdown.Item>

                      ))}
                    </Dropdown.Menu>
                  </Dropdown>
                </Form.Group>

                <Form.Group className="mb-3" controlId="route">
                  <Form.Label>Откуда</Form.Label>
                  <Form.Control 
                    type="text"
                    placeholder=""
                    value={Departure}
                    onChange={e => setDeparture(e.target.value)}
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="route">
                  <Form.Label>Куда</Form.Label>
                  <Form.Control 
                    type="text"
                    placeholder=""
                    value={Destination}
                    onChange={e => setDestination(e.target.value)}
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="Driver">
                  <Form.Label>Водитель</Form.Label>
                  <Dropdown>
                    <Dropdown.Toggle variant="success" id="dropdown-basic">
                      {FIOdrivers}
                    </Dropdown.Toggle>

                    <Dropdown.Menu>
                      {drivers.map((driver: IDriver, index: number) => (
                        <Dropdown.Item
                          key = {driver.id}
                          onClick={() => {
                            setDriverID(driver.id)
                            setFIOdrivers(driver.name + " " + driver.surname)
                          }}
                        >
                          {driver.name} {driver.surname}
                        </Dropdown.Item>

                      ))}
                    </Dropdown.Menu>
                  </Dropdown>
                </Form.Group>

                <Form.Group className="mb-3" controlId="cargoWeight">
                  <Form.Label>Вес (тонн)</Form.Label>
                  <Form.Control 
                    type="text" 
                    placeholder="" 
                    value={CargoWeight}
                    onChange={e => setCargoWeight(Number(e.target.value))}
                  />
                </Form.Group>
                
                <Form.Group className="mb-3" controlId="deliveryDate">
                  <Form.Label>Дата посавки</Form.Label>
                  <input
                    type="date"
                    className="form-control"
                    value={DeliveryDate}
                    onChange={e => setDeliveryDate(e.target.value)}
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="price">
                  <Form.Label>Цена (р)</Form.Label>
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