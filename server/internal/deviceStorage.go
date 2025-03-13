package internal

// import "errors"
import "database/sql"

// user database

type DeviceStorageDb struct {
	db *sql.DB
}

func (d *DeviceStorageDb) Create(name string, price int, img string) (Device, error) {
	var id int

	err := d.db.QueryRow("insert into device (name, price, img) values ($1, $2, $3) returning id;", name, price, img).Scan(&id)
	if err != nil {
		return Device{}, err
	}

	device := Device{
		Id: id,
		Name: name,
		Price: price,
		Img: img,
	}
	return device, nil
}

func (d *DeviceStorageDb) GetAll() ([]Device, error) {
	rows, err := d.db.Query("select * from device;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := make([]Device, 0)

	for rows.Next() {
		device := Device{}
		err = rows.Scan(&device.Id, &device.Name, &device.Price, &device.Img)
		if err != nil {
			return nil, err
		}

		devices = append(devices, device)
	}
	
	return devices, nil
}

func (d *DeviceStorageDb) GetOne(id int) (Device, error) {
	var device Device
	err := d.db.QueryRow("select * from device where id = $1;", id).Scan(&device.Id, &device.Name, &device.Price, &device.Img)
	if err != nil {
		return Device{}, err
	}

	return device, nil
}

// func (u *UserStorageDb) DeleteUser(user_id int) error {
// 	_, err := u.db.Exec("delete from users where id = $1;", user_id)
// 	return err
// }

// func (u *UserStorageDb) UpdateUser(user User) error {
// 	_, err := u.db.Exec("update users set email = $1, password = $2, role = $3 where id = $4;",
// 	user.Email, user.Password, user.Role, user.Id)
// 	return err
// }

// create storage db

func CreateDeviceStorageDb() *DeviceStorageDb {
	stDb := &DeviceStorageDb{
		db: Db,
	}

	return stDb
}