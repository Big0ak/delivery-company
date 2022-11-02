import {SyntheticEvent, useState} from 'react'
import {Form, Button} from 'react-bootstrap'
import FormContainer from '../components/FormContainer'

const LoginScreen = () => {

  const [login, setLogin] = useState('')
  const [password, setPassword] = useState('')

  const submitHandler = async(e: SyntheticEvent) => { // тип события
    e.preventDefault()
    
    await fetch('http://localhost:8000/auth/sign-in', {
      mode: 'no-cors',
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify({
        login,
        password
      })
    }).then(response => {
      const result = response.json()
      return result;
    }).then(data => {
      console.log(data);
    }).catch(err => {
      console.log(err)
    })
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