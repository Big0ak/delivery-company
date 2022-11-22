import './App.css';
import Header from './components/Header';
import Footer from './components/Footer';
import {Container } from 'react-bootstrap'
import {BrowserRouter, Routes, Route} from "react-router-dom"
import HomeScreen from './screens/HomeScreen';
import LoginScreen from './screens/LoginScreen';
import SignupScreen from './screens/SignupScreen';
import OrdersListScreen from './screens/OrdersListScreen';
import OrderCreationScreen from './screens/OrderCreationScreen';
import OrderEditScreen from './screens/OrderEditScreen';
import React, { useState } from 'react';
import { AppContext } from './shared/Context';
import { roles } from './shared/constants';


function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(
    sessionStorage.getItem("isLoggedIn") === "true" ? true : false
  );

  return (
    <AppContext.Provider 
      value = {{
        isLoggedIn,
        setIsLoggedIn
      }}
    >
      <BrowserRouter>
        <Header />
          <main>
          <Container>
            <Routes>
              <Route path='/' element={<HomeScreen />} />
              <Route path='/signup' element={<SignupScreen />} />
              <Route path='/login' element={<LoginScreen />} />
              {localStorage.getItem("role") === roles.manager ? (
                  <React.Fragment>
                    <Route path='/orders' element={<OrdersListScreen /> }/>
                    <Route path='/creat-order' element={<OrderCreationScreen /> }/>
                    <Route path='/order/:id' element={<OrderEditScreen /> }/>
                  </React.Fragment>
              ) : (
                <React.Fragment>
                    <Route path='/client-orders' element={<OrdersListScreen /> }/>
                  </React.Fragment>
              )}
            </Routes>
          </Container>
          </main>
      </BrowserRouter>
    </AppContext.Provider>
    
  );
}

export default App;
