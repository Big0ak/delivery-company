import {useContext} from 'react'
import {Navbar, Nav, Container } from 'react-bootstrap'
import { roles } from '../shared/constants';
import { AppContext } from '../shared/Context';

const Header = () => {
  const { isLoggedIn, setIsLoggedIn } = useContext(AppContext);

  const logoutHandler = () => {
    sessionStorage.setItem("isLoggedIn", "false");
    localStorage.clear()
    setIsLoggedIn(false)
  }

  return (
    <Navbar bg="dark" variant='dark' expand="lg" collapseOnSelect>
      <Container>
        <Navbar.Brand href="/">Max Speed</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          {
            !isLoggedIn ? (
              <Nav className="ms-auto">
                <Nav.Link href="/signup">Зарегистрироваться</Nav.Link>
                <Nav.Link href="/login">Войти</Nav.Link>
              </Nav>
            ) : (
                localStorage.getItem("role") === roles.manager ? (
                  <Nav className="ms-auto">
                    <Nav.Link href="/orders">Список заказов</Nav.Link>
                    <Nav.Link href="/creat-order">Оформить заказ</Nav.Link>
                    <Nav.Link href="/manager-cabinet">Личный кабинет</Nav.Link> 
                    <Nav.Link onClick={logoutHandler} href="/">Выйти</Nav.Link>
                  </Nav>
                ) : (
                  <Nav className="ms-auto">
                    <Nav.Link href="/client-orders">Список заказов</Nav.Link>
                    <Nav.Link href="/client-cabinet">Личный кабинет</Nav.Link>
                    <Nav.Link onClick={logoutHandler} href="/">Выйти</Nav.Link>
                  </Nav>
                )   
            )
          }
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
} 


export default Header