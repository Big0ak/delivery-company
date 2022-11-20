import {useContext} from 'react'
import {Navbar, Nav, Container } from 'react-bootstrap'
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
        <Navbar.Brand href="/">Delivery Company</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          {
            !isLoggedIn ? (
              <Nav className="ms-auto">
                <Nav.Link href="/signup">Зарегистрироваться</Nav.Link>
                <Nav.Link href="/login">Войти</Nav.Link>
              </Nav>
            ) : (
              <Nav className="ms-auto">
                <Nav.Link href="/orders">Список заказов</Nav.Link>
                <Nav.Link href="/creat-order">Создать заказ</Nav.Link> 
                <Nav.Link onClick={logoutHandler} href="/">Выйти</Nav.Link>
              </Nav>
            )
          }
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
} 


export default Header