import {React, useContext, useEffect, useState} from "react"
import { BrowserRouter as Router } from "react-router-dom";
// import "react-bootstrap/dist/react-bootstrap.min.js";

import AppRouter from "./components/AppRouter"
import NavBar from "./components/navbar/NavBar"
import { observer } from "mobx-react-lite";
import { Context } from "./index";

import {check} from "./http/userAPI";
import { Spinner } from "react-bootstrap";

// import "./styles/main.css"

const App = observer(() => {
  const {user} = useContext(Context)
  const [loading, setLoading] = useState(true)
  // #?
  useEffect( () => {
    setTimeout( () => {
      check().then(data => {
        user.setUser(true)
        user.setIsAuth(true)
      }).finally( () => setLoading(false))
    }, 100)
  }, [])

  if (loading) {
    return <Spinner animation={"grow"}/>
  }

  return(
    <Router>
          <NavBar />

          <AppRouter />
    </Router>
  );
});

export default App;
