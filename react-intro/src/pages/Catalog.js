import { useContext, useEffect } from "react";
import { Col, Container } from "react-bootstrap";
import Product from "../components/product/Product";

import { Row } from "react-bootstrap";
import DeviceList from "../components/DeviceList";
import {Context} from "../index";
import { observer } from "mobx-react-lite";
import {fetchDevices} from "../http/deviceAPI";

const Catalog = observer(() => {
    const {device} = useContext(Context)

    useEffect(() => {
        fetchDevices().then(data => device.setDevices(data))
    })

    return ( 
        <Container>
            <Row className="mt-2">
                <Col md={3}>
                </Col>
                <Col md={9}>
                <DeviceList />
                </Col>
            </Row>
        </Container>    
     );
})
 
export default Catalog;