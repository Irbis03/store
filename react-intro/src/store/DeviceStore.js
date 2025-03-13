import { makeAutoObservable } from "mobx";

export default class DeviceStore {
    constructor() {
        this._devices = [
            // {id: 1, name: "Redmi 10", price: 10000, img: `https://avatars.mds.yandex.net/get-mpic/4219717/img_id8119766774660622192.jpeg/orig`},
        ]
        makeAutoObservable(this)
    }

    setDevices(devices) {
        this._devices = devices
    }

    get devices() {
        return this._devices
    }
}