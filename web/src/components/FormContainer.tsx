import React, { FC } from 'react'
import { Container, Row, Col } from 'react-bootstrap'

interface Props {
    children: React.ReactNode
}

const FormContainer: FC<Props> = ({
  children
}) => {
  return (
    <Container className='py-3'>
        <Row className='justify-content-md-center'>
            <Col xs={12} md={8}>
                {children}
            </Col>
        </Row>
    </Container>
  )
}

export default FormContainer
