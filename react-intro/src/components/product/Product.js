import { NavLink } from "react-router-dom";

// import "./styles.css";

const Product = (props) => {
    return ( 
        <>
            <NavLink to={`/product/${props.ind}`}>
            <li className="product">
                    <img src={props.img} alt="Product image" className="product__img" />
                    <h3 className="product__title">{props.title}</h3>
                    <div className="product__bottom">
                        <h3 className="price">{props.price}</h3>
                        <a href="#" className="price-btn">В корзину</a>
                    </div>
                </li>
            </NavLink>
        </>
     );
}
 
export default Product;