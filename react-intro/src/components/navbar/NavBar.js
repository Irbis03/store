import { NavLink, useNavigate } from "react-router-dom";
import {observer} from "mobx-react-lite";

import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import Container from "react-bootstrap/Container";
import Button from "react-bootstrap/Button";


import { useContext } from "react";
import { Context } from "../..";
import { ADMIN_ROUTE, LOGIN_ROUTE, SHOP_ROUTE } from "../../utils/consts";

// import "./styles.css";

const NavBar = observer(() => {
    const {user} = useContext(Context)
    const navigate = useNavigate()

    const logout = () => {
        user.setUser({})
        user.setIsAuth(false)
    }

    return ( 
        <Navbar bg="dark" data-bs-theme="dark">
            <Container>
            <NavLink style={{color:'white'}} to={SHOP_ROUTE}>CS</NavLink>
            {
                user.isAuth ?
                <Nav className="ml-auto" style={{color: 'white'}}>
                    <Button 
                        variant={"outline-light"}
                        onClick={() => navigate(ADMIN_ROUTE)}
                    >
                        Админ панель
                    </Button>
                    <Button 
                        variant={"outline-light"}
                        onClick={logout}
                        className="ml-2"
                    >
                        Выйти
                    </Button>
                </Nav>
                : // ? onClick
                <Nav className="ml-auto" style={{color: 'white'}}>
                    <Button 
                        variant={"outline-light"}
                        onClick={() => navigate(LOGIN_ROUTE)}
                    >
                        Авторизация
                    </Button>
                </Nav>
            }
            </Container>
       </Navbar>
     );
});
 
export default NavBar;