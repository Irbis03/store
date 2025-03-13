import { Button } from "react-bootstrap";
import React, { useEffect, useState } from "react";

import {Col, Container, Image, Row} from "react-bootstrap";
import { Card } from "react-bootstrap";

import {useParams} from "react-router-dom";
import { fetchOneDevice } from "../http/deviceAPI";

const Product = () => {
    const [device, setDevice] = useState({})
    const {id} = useParams()
    useEffect(() => {
        fetchOneDevice(id).then(data => setDevice(data))
    }, [])

    // const device = {id: 1, name: "Redmi 10", price: 10000, img: `https://avatars.mds.yandex.net/get-mpic/4219717/img_id8119766774660622192.jpeg/orig`}
    return (
       <Container className="mt-3">
            <Row>
            <Col md={4}>
                <Image width={300} height={300} src={process.env.REACT_APP_API_URL + "static/" + device.img} />
            </Col>
            <Col md={4}>
                <Row className="d-flex flex-column align-items-center">
                    <h2>{device.name}</h2>
                </Row>
            </Col>
            <Col md={4}>
                <Card
                className="d-flex flex-column align-items-center justify-content-around"
                style={{width: 300, height: 300, fontsize:32, border: '5px solid lightgray' }}
                >
                    <h3>{device.price} руб.</h3>
                    <Button variant={"outline-dark"}>Добавить в корзину</Button>
                </Card>
                
            </Col>
            </Row>
       </Container>
    );
}

export default Product;
