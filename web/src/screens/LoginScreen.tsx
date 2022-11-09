import {FC, SyntheticEvent, useState} from 'react'
import {Form, Button} from 'react-bootstrap'
import FormContainer from '../components/FormContainer'
import { ILoginUser } from '../axios/interfaces';
import { sendSignInManager } from '../axios/hooks';
import { roles } from '../constants';
import { useNavigate } from 'react-router-dom';

const LoginScreen: FC = () => {

  const [login, setLogin] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate()
  
  const submitHandler = async(e: SyntheticEvent) => { // тип события
    e.preventDefault()
    const body: ILoginUser = {
      login: login,
      password: password, 
    }
    const token = await sendSignInManager("auth/sign-in", body)
    if (token !== null){  
      localStorage.setItem('JWT', token.token)
      localStorage.setItem('role', roles.manager)
    }

    navigate('/')
  }

  return (
    <FormContainer>
      <h1>Вход в личный кабинет</h1>
      <Form onSubmit={submitHandler}>

        <Form.Group className="mb-3" controlId="Login">
          <Form.Label>Логин</Form.Label>
          <Form.Control 
            type="login" 
            placeholder="введите логин"
            value={login}
            onChange={e => setLogin(e.target.value)}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="Password">
          <Form.Label>Пароль</Form.Label>
          <Form.Control 
            type="password" 
            placeholder="пароль" 
            value={password}
            onChange={e => setPassword(e.target.value)}
          />
        </Form.Group>
        
        <Button variant="primary" type="submit">
          Войти
        </Button>
      </Form>
    </FormContainer>
  )
}

export default LoginScreen