import Admin from "./pages/Admin"
import Auth from "./pages/Auth"
import Catalog from "./pages/Catalog"
import Product from "./pages/Product"
import {ADMIN_ROUTE, LOGIN_ROUTE, REGISTRATION_ROUTE, SHOP_ROUTE, DEVICE_ROUTE} from "./utils/consts"

export const authRoutes = [
    {
        path: ADMIN_ROUTE,
        Component: Admin
    }
]

export const publicRoutes = [
    {
        path: SHOP_ROUTE,
        Component: Catalog 
    },
    {
        path: LOGIN_ROUTE,
        Component: Auth 
    },
    {
        path: REGISTRATION_ROUTE,
        Component: Auth
    },
    {
        path: DEVICE_ROUTE + '/:id',
        Component: Product 
    },
]