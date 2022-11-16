import React, { useEffect, FC, SyntheticEvent, useState } from 'react'
import { IClient, IOrder, IDriver, IOrderRead } from '../axios/interfaces'
import {Form, Button} from 'react-bootstrap'
import FormContainer from '../components/FormContainer'
import { getAllClient, getAllDriver, getOrderId, editOrder } from '../axios/hooks';
import {useParams} from "react-router-dom";


import Dropdown from 'react-bootstrap/Dropdown';


const OrderEditScreen: FC = () => {
    const [orderRead, setOrderRead] =useState<IOrderRead>()

  const { Id } = useParams()

    const [clients, setClients] = useState<IClient[]>([])
    const [FIOclient, setFIOclient] = useState(" ")

    const [drivers, setDrivers] = useState<IDriver[]>([])
    const [FIOdrivers, setFIOdrivers] = useState(" ")

    const [CliendID, setCliendID] = useState(Number)
    const [DriverID, setDriverID] = useState(Number)
    const [CargoWeight, setCargoWeight] = useState(Number)
    const [Price, setPrice] = useState(Number)
    const [Departure, setDeparture] = useState("")
    const [Destination, setDestination] = useState("")

    const [Submitted, setSubmitted] = useState(false)
    const [OrderId, setOrderId] = useState('')

    useEffect(() => {
      const getClients = async () => {
          const response = await getAllClient("manager-api/client/")
          setClients(response)  
      }

      const getDrivers = async () => {
        const response = await getAllDriver("manager-api/driver/")
        setDrivers(response)
      }

      const getOrderRead = async () => {
          const response = await getOrderId("manager-api/orders", OrderId)
          setOrderRead(response)
      }

      console.log(Id)

      getClients();
      getDrivers();
      getOrderRead();

      if (orderRead) {
        setFIOclient(orderRead.client)
  
        setFIOdrivers(orderRead.driver)
  
        setDeparture(orderRead.departure)
        setDestination(orderRead.destination)
        setCargoWeight(orderRead.cargoWeight)
        setPrice(orderRead.price)
      }
      
    }, [])

    const submitHandler = async (e: SyntheticEvent) => {
      e.preventDefault()
      
      const body: IOrder = {
        id: 1, // как получить конкретный id через ссылку
        clientId: CliendID,
        driverId: DriverID,
        cargoWeight: CargoWeight,
        price: Price,
        departure: Departure,
        destination: Destination,
      }
      await editOrder("manager-api/orders", body)

      setSubmitted(true)
    }

  return (
    <FormContainer>
      {
        Submitted ? (
          <div>
            <h4>Заказ изменен!</h4>
            <h3>номер заказа</h3>
          </div>
        ) : (
            <React.Fragment>
              <h1>Изменение заказа {OrderId} </h1>
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
                            setCliendID(client.id)
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

export default OrderEditScreen