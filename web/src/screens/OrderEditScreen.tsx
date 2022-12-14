import React, { useEffect, FC, SyntheticEvent, useState } from 'react'
import { IClient, IOrder, IDriver, IOrderRead } from '../shared/interfaces'
import {Form, Button} from 'react-bootstrap'
import FormContainer from '../components/FormContainer'
import { getRequest, getOrderId, editOrder, deleteOrder } from '../axios/services';
import {useParams} from "react-router-dom";

import Dropdown from 'react-bootstrap/Dropdown';

const OrderEditScreen: FC = () => {
    const [orderRead, setOrderRead] =useState<IOrderRead>();

    const { id } = useParams();

    const [clients, setClients] = useState<IClient[]>([]);
    const [FIOclient, setFIOclient] = useState("");

    const [drivers, setDrivers] = useState<IDriver[]>([]);
    const [FIOdrivers, setFIOdrivers] = useState("");

    const [CliendID, setCliendID] = useState(0);
    const [DriverID, setDriverID] = useState(0);
    const [CargoWeight, setCargoWeight] = useState(0);
    const [Price, setPrice] = useState(0);
    const [Departure, setDeparture] = useState("");
    const [Destination, setDestination] = useState("");
    const [DeliveryDate, setDeliveryDate] = useState("");

    const [Submitted, setSubmitted] = useState(false)

    useEffect(() => {
      const getClients = async () => {
          const response = await getRequest("manager-api/client/")
          setClients(response)  
      }

      const getDrivers = async () => {
        const response = await getRequest("manager-api/driver/")
        setDrivers(response)
      }

      const getOrderRead = async () => {
          if (id) {
            const response = await getOrderId("manager-api/orders", id)
            setOrderRead(response)
          }
      }

      getClients();
      getDrivers();
      getOrderRead();
    
    }, [])

    useEffect(() => {
      if (orderRead) {
        let findIdClient: number = 0
        let strFIO = orderRead.client.split(' ')

        clients.map((elem) => {
          if (elem.name === strFIO[0] && elem.surname === strFIO[1]){
            findIdClient = Number(elem.id)
          }
        })
        setCliendID(findIdClient);

        let findIdDriver = 0
        strFIO = orderRead.driver.split(' ')
        drivers.map((elem) => {
          if (elem.name == strFIO[0] && elem.surname === strFIO[1]){
            findIdDriver = elem.id
          }
        })
        setDriverID(findIdDriver)

        // ?????????????????????? ???????????? ???????????? ?? ?????? ????????????, ?????????? ???????????????????? Id ?????????????? ?? ????????????????
        // ???? ???????????? ?????????????? ?? ?????????????? ????????, ???????????????????? ?????????????????????????? ????????????????
        if (findIdClient && findIdDriver) {
          setFIOclient(orderRead.client)
          setFIOdrivers(orderRead.driver)
          setDeparture(orderRead.departure)
          setDestination(orderRead.destination)
          setCargoWeight(orderRead.cargoWeight)
          setDeliveryDate(orderRead.deliveryDate.split('T')[0])
          setPrice(orderRead.price)
        }
      }
    }, [orderRead])

    const submitHandler = async (e: SyntheticEvent) => {
      e.preventDefault()
      const body: IOrder = {
        id: Number(id),
        clientId: CliendID,
        driverId: DriverID,
        cargoWeight: CargoWeight,
        price: Price,
        departure: Departure,
        destination: Destination,
        deliveryDate: DeliveryDate
      }
      await editOrder("manager-api/orders", body)

      setSubmitted(true)
    }

    const submitDeleteOrder = async (e: SyntheticEvent) => {
      e.preventDefault()
      await deleteOrder("manager-api/orders", String(id))

      setSubmitted(true)
    }

  return (
    <FormContainer>
      {
        Submitted ? (
          <div>
            <h2>???????????????????????? ???????????????? ?????? ?????????????? ???{id}!</h2>
            <Button
                className="btn btn-link btn-light"
                href={`/orders`}
            >
              ?? ???????????? ??????????????
            </Button>
          </div>
        ) : (
            <React.Fragment>
              <h1>?????????????????? ???????????? ???{id} </h1>
              <Form onSubmit={submitHandler}>

                <Form.Group className="mb-3" controlId="Client">
                  <Form.Label>????????????</Form.Label>
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
                  <Form.Label>????????????</Form.Label>
                  <Form.Control 
                    type="text"
                    placeholder=""
                    value={Departure}
                    onChange={e => setDeparture(e.target.value)}
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="route">
                  <Form.Label>????????</Form.Label>
                  <Form.Control 
                    type="text"
                    placeholder=""
                    value={Destination}
                    onChange={e => setDestination(e.target.value)}
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="Driver">
                  <Form.Label>????????????????</Form.Label>
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
                  <Form.Label>?????? (????????)</Form.Label>
                  <Form.Control 
                    type="text" 
                    placeholder="" 
                    value={CargoWeight}
                    onChange={e => setCargoWeight(Number(e.target.value))}
                  />
                </Form.Group>
                
                <Form.Group className="mb-3" controlId="deliveryDate">
                  <Form.Label>???????? ??????????????</Form.Label>
                  <input
                    type="date"
                    className="form-control"
                    value={DeliveryDate}
                    onChange={e => setDeliveryDate(e.target.value)}
                  />
                </Form.Group>

                <Form.Group className="mb-3" controlId="price">
                  <Form.Label>???????? (??)</Form.Label>
                  <Form.Control 
                    type="text" 
                    placeholder="" 
                    value={Price}
                    onChange={e => setPrice(Number(e.target.value))}
                  />
                </Form.Group>
                
                
                <Button variant="primary" type="submit">
                  ??????????????????
                </Button>
                 {' '}
                <Button variant="primary" className="btn btn-danger" onClick={submitDeleteOrder}>
                  ??????????????
                </Button>
              </Form>
            </React.Fragment>
        )}
    </FormContainer>
  )
}

export default OrderEditScreen