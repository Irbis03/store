import React from 'react';
import {Col, Image} from "react-bootstrap";
import { Card } from 'react-bootstrap';
import {useNavigate} from "react-router-dom";
import { DEVICE_ROUTE } from '../utils/consts';

const DeviceItem = ({device}) => {
    const navigate = useNavigate()
    return ( 
        <Col md={3} onClick={() => navigate(DEVICE_ROUTE + '/' + device.id)}>
            <Card style={{width: 150, cursor: 'pointer'}} border={"light"}>
                <Image width={150} height={150} src={process.env.REACT_APP_API_URL + "static/" + device.img} />
                <div>{device.name}</div>
            </Card>
        </Col>
     );
}
 
export default DeviceItem;