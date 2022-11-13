import './App.css';
import Header from './components/Header';
import Footer from './components/Footer';
import {Container } from 'react-bootstrap'
import {BrowserRouter, Routes, Route} from "react-router-dom"
import HomeScreen from './screens/HomeScreen';
import LoginScreen from './screens/LoginScreen';
import SignupScreen from './screens/SignupScreen';
import OrdersManagerScreen from './screens/OrdersManagerScreen';
import OrderCreationScreen from './screens/OrderCreationScreen';
import React from 'react';

function App() {
  return (
    <BrowserRouter>
      <Header />
        <main>
        <Container>
          <Routes>
            <Route path='/' element={<HomeScreen />} />
            <Route path='/signup' element={<SignupScreen />} />
            <Route path='/login' element={<LoginScreen />} />
            {localStorage.getItem("JWT") && (
                <React.Fragment>
                  <Route path='/orders' element={<OrdersManagerScreen /> }/>
                  <Route path='/creat-order' element={<OrderCreationScreen /> }/>
                </React.Fragment>
                //<Route path='/orders/:id' element={<OrderItemScreen /> }/>
            )}
          </Routes>
        </Container>
        </main>
      <Footer/>
    </BrowserRouter>
  );
}

export default App;
