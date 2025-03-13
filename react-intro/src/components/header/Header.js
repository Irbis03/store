import { NavLink } from "react-router-dom";

// import "./styles.css"

const Header = () => {
    return ( 
        <>
            <header className="header">
                <div className="header-wrapper">
                    <h1 className="header__title">CS - интернет-магазин по продаже смартфонов</h1>
                    <NavLink to="catalog" className="btn">Перейти в каталог</NavLink>
                </div>
            </header>
        </>
     );
}
 
export default Header;