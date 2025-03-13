import React, { useContext, useState } from 'react';
import { Col, Form, Row } from 'react-bootstrap';

import Button from 'react-bootstrap/Button';
import Modal from 'react-bootstrap/Modal';

import {Context} from "../../index";
import { createDevice } from '../../http/deviceAPI';

const CreateDevice = ({show, onHide}) => {
    const {device} = useContext(Context)
    const [name, setName] = useState("")
    const [price, setPrice] = useState(0)
    const [file, setFile] = useState(null) 
    const [info, setInfo] = useState([]) 

    const addInfo = () => {
        setInfo([...info, {title: '', description: '', number: Date.now()}])
    }
    const removeInfo = (number) => {
        setInfo(info.filter(i => i.number !== number))
    }

    const selectFile = e => {
        setFile(e.target.files[0])
    }

    const addDevice = () => {
        const formData = new FormData()
        formData.append('name', name)
        formData.append('price', `${price}`)
        formData.append('img', file)

        createDevice(formData).then(data => onHide())
    }

    return ( 
        <Modal
        show={show}
        onHide={onHide}
        size="lg"
        centered
    >
      <Modal.Header closeButton>
        <Modal.Title id="contained-modal-title-vcenter">
          Добавить устройство
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form>
            <Form.Control
                value={name}
                onChange={e => setName(e.target.value)} 
                placeholder={"Введите название устройства"}
            />
            <Form.Control
                value={price}
                onChange={e => setPrice(Number(e.target.value))} 
                className="mt-3"
                placeholder={"Введите стоимость устройства"}
                type="number"
            />
            <Form.Control
                className="mt-3"
                type="file"
                onChange={selectFile}
            />
            <hr/>
            <Button
                variant={"outline-dark"}
                onClick={addInfo}
            >
                Добавить новое свойство
            </Button>
            {
                info.map(i =>
                    <Row className="mt-2" key={i.number}>
                        <Col md={4}>
                            <Form.Control
                                placeholder="Введите название устройства"
                            />
                        </Col>
                        <Col md={4}>
                            <Form.Control 
                                placeholder="Введите описание устройства"
                            />
                        </Col>
                        <Col md={4}>
                            <Button 
                            variant={"outline-danger"}
                            onClick={() => removeInfo(i.number)}
                            >
                                Удалить
                            </Button>
                        </Col>
                    </Row>

                )
            }
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="outline-danger" onClick={onHide}>Закрыть</Button>
        <Button variant="outline-success" onClick={addDevice}>Добавить</Button>
      </Modal.Footer>
    </Modal>
     );
}
 
export default CreateDevice;