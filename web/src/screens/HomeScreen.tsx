import React from 'react'
import FormContainer from '../components/FormContainer'
import {Row, Col} from 'react-bootstrap'
const HomeScreen = () => {
  return (
    <Row>
      <Col>
        <h1>Добро пожаловать</h1>
        <p className="mb-0" > Предлагаем широкий ассортимент услуг, удовлетворяющих ваши потребности в перевозке грузов, работая в более чем 220 странах и регионах по всему миру.</p>
      </Col>
      <Col xs={8} md={13}>
        <img src="https://lorrycraft.ru/images/blog/multi-country-consolidation.jpg" className="img-fluid" alt="Responsive image"/>
      </Col>
    </Row>
    
  )
}

export default HomeScreen