import {React, useContext} from "react"
import {Routes, Route, Navigate} from "react-router-dom";
import {authRoutes, publicRoutes} from "../routes";
import {Context} from "./../index";

import Catalog from "../pages/Catalog";
import Auth from "../pages/Auth";
import { LOGIN_ROUTE, REGISTRATION_ROUTE, SHOP_ROUTE, DEVICE_ROUTE, ADMIN_ROUTE } from "../utils/consts";
import Product from "../pages/Product";
import Admin from "../pages/Admin";

const AppRouter = () => {
    const {user} = useContext(Context)
    return ( 
        <Routes>
            {/* {
                user.isAuth && authRoutes.map(({path, Component}) => 
                    <Route key={path} path={path} element={Component} />
            )}
            {
                publicRoutes.map(({path, Component}) => 
                    <Route key={path} path={path} element={Component} />
            )} */}
            {user.isAuth && <Route key={ADMIN_ROUTE} path={ADMIN_ROUTE} element={<Admin />} />}

            <Route key={SHOP_ROUTE} path={SHOP_ROUTE} element={<Catalog />}/>
            <Route key={LOGIN_ROUTE} path={LOGIN_ROUTE} element={<Auth />}/>
            <Route key={REGISTRATION_ROUTE} path={REGISTRATION_ROUTE} element={<Auth />}/>
            <Route key={DEVICE_ROUTE + '/:id'} path={DEVICE_ROUTE + '/:id'} element={<Product />}/>

            <Route path="*" element={<Navigate to="/" replace />} />
            {/* <Route key={"/"} path={"/"} element={<Catalog />} /> */}
        </Routes>
     );
};


export default AppRouter;