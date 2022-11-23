import {FC, useState, useEffect} from 'react'
import { Col, Form, Row } from 'react-bootstrap'
import { getInfoUser } from '../axios/services'
import FormContainer from '../components/FormContainer'
import { roles } from '../shared/constants'
import { IClient, IManager } from '../shared/interfaces'

const CabinetScreen: FC = () => {
    const [role] = useState(localStorage.getItem("role"))
    //const [Client, setClient] = useState<IClient | IManager>()

    const [firstName, setFirstName] = useState('')
    const [lastName, setLastName] = useState('')
    const [login, setLogin] = useState('')
    const [password, setPassword] = useState('')
    const [phone, setPhone] = useState('')

    useEffect(() => {
        const getUserInfo = async () => {
            const response = await getInfoUser(`${role}-api/cabinet/`)
            let user 
            if (role === roles.manager){
                user = response as IManager
            } else {
                user = response as IClient
                setPhone(user.phone)
            }
            setFirstName(user.login)
            setFirstName(user.name)
            setLastName(user.surname)
            setLogin(user.login)
        }

        getUserInfo()
    }, [])

    return (
        <FormContainer>
            <h1>Личный кабинет</h1>
            <Form>
            <Row>
                    <Form.Group as={Col} controlId="formGridfirstName">
                        <Form.Label>Имя</Form.Label>
                        <Form.Control 
                            type="firtsName" 
                            placeholder="ваше имя" 
                            value={firstName}
                            onChange={e => setFirstName(e.target.value)}
                        />
                    </Form.Group>
                    <Form.Group as={Col} controlId="formGridLastName">
                        <Form.Label>Фамилия</Form.Label>
                        <Form.Control 
                            type="lastName" 
                            placeholder="ваша фамилия" 
                            value={lastName}
                            onChange={e => setLastName(e.target.value)}
                        />
                    </Form.Group> 
                </Row> <br/>
                
                <Row>
                    <Form.Group as={Col} xs={3} controlId="formGridLogin">
                        <Form.Label>Логин</Form.Label>
                        <Form.Control
                            type="text"
                            placeholder={login}
                            aria-label="Disabled input"
                            disabled
                            readOnly
                        />
                    </Form.Group>
                    <Form.Group as={Col} controlId="formGridPassword">
                        <Form.Label>Пароль</Form.Label>
                        <Form.Control 
                            placeholder="пароль" 
                            value={password}
                            onChange={e => setPassword(e.target.value)}
                        />
                    </Form.Group>
                    {role === roles.client && (
                        <Form.Group as={Col} controlId="Phone">
                            <Form.Label>Телефон</Form.Label>
                            <Form.Control
                                type="phone"
                                placeholder="телефон"
                                value={phone}
                                onChange={e => setPhone(e.target.value)}
                            />
                        </Form.Group>
                    )}
                </Row>
            </Form>
        </FormContainer>

    )
}

export default CabinetScreen