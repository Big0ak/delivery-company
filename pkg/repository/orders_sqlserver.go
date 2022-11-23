package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Big0ak/delivery-company/models"
)

type OrderDB struct {
	db *sql.DB
}

func NewOrderDB(db *sql.DB) *OrderDB {
	return &OrderDB{db: db}
}

func (r *OrderDB) CreateManagerDB(managerId int, order models.Orders) (int, error) {
	var id int
	creatOrder := fmt.Sprintf("INSERT INTO %s (ClientID, DriverID, ManagerID, CargoWeight, Price, Departure, Destination) OUTPUT Inserted.OrderID VALUES('%d', '%d', '%d', '%d','%d', '%s', '%s')",
		ordersTable, order.ClientID, order.DriverID, managerId, order.CargoWeight, order.Price, order.Departure, order.Destination)
	row := r.db.QueryRow(creatOrder)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrderDB) GetAllDB(managerId int) ([]models.OrdersRead, error) {
	var orders []models.OrdersRead
	query := fmt.Sprintf(`SELECT o.OrderID, c.Name + ' ' + c.Surname as 'Client', ` +
							`d.Name + ' ' + d.Surname as 'Driver', ` +
							`m.Name + ' ' + m.Surname as 'Manager', ` +
							`o.CargoWeight, o.Price, o.Departure, o.Destination, o.Date FROM %s o ` +
								`JOIN %s c ON o.ClientID = c.ClientID ` +
								`JOIN %s d ON o.DriverID = d.DriverID ` +
								`LEFT JOIN %s m ON o.ManagerID = m.ManagerID `,
								 	ordersTable, clientTable, driverTable, managerTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := models.OrdersRead{}
		var price []uint8
		err := rows.Scan(
			&o.OrderID, &o.Client, &o.Driver, &o.Manager, &o.CargoWeight, &price, &o.Departure, &o.Destination, &o.Date)
		if err != nil {
			return nil, err
		}
		// конвертация sql price = []uint8 -> uint
		o.Price = ConvertPriceToUint(price[:])
		orders = append(orders, o)
	}
	return orders, nil
}
 
func (r *OrderDB) GetByIdDB(userId, id int) (models.OrdersRead, error) {
	var order models.OrdersRead
	query := fmt.Sprintf(`SELECT o.OrderID, c.Name + ' ' + c.Surname as 'Client', ` +
							`d.Name + ' ' + d.Surname as 'Driver', ` +
							`m.Name + ' ' + m.Surname as 'Manager', ` +
							`o.CargoWeight, o.Price, o.Departure, o.Destination, o.Date FROM (SELECT * FROM %s od WHERE od.OrderID = %d) o ` +
								`JOIN %s c ON o.ClientID = c.ClientID ` +
								`JOIN %s d ON o.DriverID = d.DriverID ` +
								`LEFT JOIN %s m ON o.ManagerID = m.ManagerID `,
								 	ordersTable, id, clientTable, driverTable, managerTable)
	row := r.db.QueryRow(query)
	var price []uint8
	err := row.Scan(&order.OrderID, &order.Client, &order.Driver, &order.Manager, &order.CargoWeight, &price, &order.Departure, &order.Destination, &order.Date)
	if err != nil {
		return models.OrdersRead{}, err
	}
	order.Price = ConvertPriceToUint(price[:])
	return order, nil
}


func (r *OrderDB) DeleteManagerDB(managerId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s where OrderID = %d", ordersTable, id)
	_, err := r.db.Exec(query)
	return err
}

	
func (r *OrderDB) UpdateManagerDB(managerId, id int, input models.Orders) error {
	query := fmt.Sprintf("UPDATE %s SET ClientID = %d, DriverID = %d, ManagerID = %d, CargoWeight = %d, Price = %d, Departure = '%s', Destination = '%s' WHERE OrderID = %d",
		ordersTable, input.ClientID, input.DriverID,
		managerId, input.CargoWeight, input.Price, input.Departure, input.Destination, id)
	_, err := r.db.Exec(query)
	return err
}

func (r *OrderDB) SearchOrdersByCityManagerDB(managerId int, city string) ([]models.OrdersRead, error){
	var orders []models.OrdersRead
	query := fmt.Sprintf(`SELECT o.OrderID, c.Name + ' ' + c.Surname as 'Client',` +
						`d.Name + ' ' + d.Surname as 'Driver',` +
						`m.Name + ' ' + m.Surname as 'Manager',` +
						`o.CargoWeight, o.Price, o.Departure, o.Destination, o.Date FROM (SELECT * FROM %s od WHERE od.Departure LIKE '%%%s%%' OR od.Destination LIKE '%%%s%%') o ` +
							`JOIN %s c ON o.ClientID = c.ClientID ` +
							`JOIN %s d ON o.DriverID = d.DriverID ` +
							`LEFT JOIN %s m ON o.ManagerID = m.ManagerID`, ordersTable, city, city, clientTable, driverTable, managerTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := models.OrdersRead{}
		var price []uint8
		err := rows.Scan(
			&o.OrderID, &o.Client, &o.Driver, &o.Manager, &o.CargoWeight, &price, &o.Departure, &o.Destination, &o.Date)
		if err != nil {
			return nil, err
		}
		// конвертация sql price = []uint8 -> uint
		o.Price = ConvertPriceToUint(price[:])
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrderDB) GetUserOrderDB(clientId int) ([]models.OrdersRead, error) {
	var orders []models.OrdersRead
	query := fmt.Sprintf(`SELECT o.OrderID, c.Name + ' ' + c.Surname as 'Client', ` +
							`d.Name + ' ' + d.Surname as 'Driver', ` +
							`m.Name + ' ' + m.Surname as 'Manager', ` +
							`o.CargoWeight, o.Price, o.Departure, o.Destination, o.Date FROM (SELECT * FROM %s od WHERE od.ClientID = %d) o ` +
								`JOIN %s c ON o.ClientID = c.ClientID ` +
								`JOIN %s d ON o.DriverID = d.DriverID ` +
								`LEFT JOIN %s m ON o.ManagerID = m.ManagerID `,
								 	ordersTable, clientId, clientTable, driverTable, managerTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := models.OrdersRead{}
		var price []uint8
		err := rows.Scan(
			&o.OrderID, &o.Client, &o.Driver, &o.Manager, &o.CargoWeight, &price, &o.Departure, &o.Destination, &o.Date)
		if err != nil {
			return nil, err
		}
		// конвертация sql price = []uint8 -> uint
		o.Price = ConvertPriceToUint(price[:])
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrderDB) SearchOrdersByCityClientDB(clientId int, city string) ([]models.OrdersRead, error){
	var orders []models.OrdersRead
	query := fmt.Sprintf(`SELECT o.OrderID, c.Name + ' ' + c.Surname as 'Client',` +
						`d.Name + ' ' + d.Surname as 'Driver',` +
						`m.Name + ' ' + m.Surname as 'Manager',` +
						`o.CargoWeight, o.Price, o.Departure, o.Destination, o.Date FROM (SELECT * FROM %s od WHERE od.ClientID = %d AND (od.Departure LIKE '%%%s%%' OR od.Destination LIKE '%%%s%%')) o ` +
							`JOIN %s c ON o.ClientID = c.ClientID ` +
							`JOIN %s d ON o.DriverID = d.DriverID ` +
							`LEFT JOIN %s m ON o.ManagerID = m.ManagerID`, ordersTable, clientId, city, city, clientTable, driverTable, managerTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := models.OrdersRead{}
		var price []uint8
		err := rows.Scan(
			&o.OrderID, &o.Client, &o.Driver, &o.Manager, &o.CargoWeight, &price, &o.Departure, &o.Destination, &o.Date)
		if err != nil {
			return nil, err
		}
		// конвертация sql price = []uint8 -> uint
		o.Price = ConvertPriceToUint(price[:])
		orders = append(orders, o)
	}
	return orders, nil
}

// конвертация sql price = []uint8 -> uint
func ConvertPriceToUint (price []uint8) (uint){
	b := make([]byte, len(price))
	for i, v := range price {
		b[i] = byte(v)
	}
	str := string(b)
	priceINT, _ := strconv.ParseFloat(str, 64)
	return uint(priceINT)
}