import React from 'react'
import {Navbar, Nav, Container } from 'react-bootstrap'

const Header = () => {

  const logoutHandler = () => {
    localStorage.clear()
  }

  return (
    <Navbar bg="dark" variant='dark' expand="lg" collapseOnSelect>
      <Container>
        <Navbar.Brand href="/">Delivery Company</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          {
            localStorage.getItem("JWT") === null ? (
              <Nav className="ms-auto">
                <Nav.Link href="/signup">Зарегистрироваться</Nav.Link>
                <Nav.Link href="/login">Войти</Nav.Link>
              </Nav>
            ) : (
              <Nav className="ms-auto">
                <Nav.Link href="/orders">Список заказов</Nav.Link>
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