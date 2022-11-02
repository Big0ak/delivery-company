import {SyntheticEvent, useState} from 'react'
import {Form, Button} from 'react-bootstrap'
import FormContainer from '../components/FormContainer'
import { useNavigate } from "react-router-dom";


const SignupScreen = () => {

  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [login, setLogin] = useState('')
  const [password, setPassword] = useState('')

  const submitHandler = async(e: SyntheticEvent) => {
    e.preventDefault()

    //связь с backend
    await fetch('http://localhost:8000/auth/sign-up', {
      mode: 'no-cors',
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: firstName,
        surname: lastName,
        login,
        password
      })
    })
  }

  return (
    <FormContainer>
        <h1>Регистрация</h1>
      <Form onSubmit={submitHandler}>
        <Form.Group className="mb-3" controlId="firstName">
          <Form.Label>Имя</Form.Label>
          <Form.Control 
            type="firtsName" 
            placeholder="введите имя" 
            value={firstName}
            onChange={e => setFirstName(e.target.value)}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="lastName">
          <Form.Label>Фамилия</Form.Label>
          <Form.Control 
            type="lastName"
            placeholder="введите фамилию"
            value={lastName}
            onChange={e => setLastName(e.target.value)}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="Login">
          <Form.Label>Логин</Form.Label>
          <Form.Control 
            type="login"
            placeholder="введите логин"
            value={login}
            onChange={e => setLogin(e.target.value)} 
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicPassword">
          <Form.Label>Пароль</Form.Label>
          <Form.Control 
            type="password" 
            placeholder="пароль" 
            value={password}
            onChange={e => setPassword(e.target.value)}
          />
        </Form.Group>
        
        <Button variant="primary" type="submit">
          Зарегистрироваться
        </Button>
      </Form>
    </FormContainer>
  )
}

export default SignupScreen