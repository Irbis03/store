package internal

// import "errors"
import "database/sql"

// user database

var Db *sql.DB

type UserStorageDb struct {
	db *sql.DB
}

func (u *UserStorageDb) PutUser(email, password, role string) (User, error) {
	user := User{
		Email: email,
		Password: password,
		Role: role,
	}

	err := u.db.QueryRow("insert into users (email, password, role) values ($1, $2, $3) returning id;", user.Email, user.Password, user.Role).Scan(&user.Id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *UserStorageDb) GetUserByEmail(email string) (User, error) {
	var user User
	err := u.db.QueryRow("select * from users where email = $1;", email).Scan(&user.Id, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UserStorageDb) DeleteUser(user_id int) error {
	_, err := u.db.Exec("delete from users where id = $1;", user_id)
	return err
}

func (u *UserStorageDb) UpdateUser(user User) error {
	_, err := u.db.Exec("update users set email = $1, password = $2, role = $3 where id = $4;",
	user.Email, user.Password, user.Role, user.Id)
	return err
}

func (u *UserStorageDb) GetAllUser() ([]User, error) {
	rows, err := u.db.Query("select * from users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var id int
		var email string
		var password string
		var role string

		err = rows.Scan(&id, &email, &password, &role)
		if err != nil {
			return nil, err
		}

		user := User{
			Id: id,
			Email: email,
			Password: password,
			Role: role, 
		}

		users = append(users, user)
	}
	
	return users, nil
}

// create user storage db

func CreateUserStorageDb() *UserStorageDb {
	userStorageDb := &UserStorageDb{
		db: Db,
	}

	return userStorageDb
}

// // user storage in memory implementation

// type UserStorageImpl struct {
// 	users map[int]User
// 	nextId int
// }

// func (u *UserStorageImpl) PutUser(user User) error {
// 	user.Id = u.nextId
// 	u.users[user.Id] = user
// 	return nil
// }

// func (u *UserStorageImpl) GetUserByEmail(email string) (User, error) {
// 	for _, user := range u.users {
// 		if user.Email == email {
// 			return user, nil
// 		}
// 	}
	
// 	return User{}, errors.New("this user doesn't exist")
// }

// func (u *UserStorageImpl) DeleteUser(user_id int) error {
// 	_, ok := u.users[user_id]
// 	if !ok {
// 		return errors.New("this user doesn't exist")
// 	}

// 	delete(u.users, user_id)
// 	return nil
// }

// func (u *UserStorageImpl) GetAllUser() ([]User, error) {
// 	users := make([]User, 0)

// 	for _, user := range u.users {
// 		users = append(users, user)
// 	}
	
// 	return users, nil
// }

// // create storage

// func CreateUserStorage() *UserStorageImpl {
// 	userStorage := &UserStorageImpl{
// 		users: map[int]User{1: User{
// 			Id: 1,
// 			Email: "test@mail.ru",
// 			Password: "$2a$10$caaeDWfqMRlyrwS4P5MCZ.d0EsR6gYcUo/VOb986BPKRKCyzQKTQG", // "123"
// 			Role: "user"},
// 		},
// 		nextId: 0,
// 	}

// 	return userStorage
// }