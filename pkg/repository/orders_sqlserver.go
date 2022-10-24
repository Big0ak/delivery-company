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
	creatOrder := fmt.Sprintf("INSERT INTO %s (ClientID, RouteID, DriverID, ManagerID, CargoWeight, Price) OUTPUT Inserted.OrderID VALUES('%d', '%d', '%d', '%d','%d', '%d')",
		ordersTable, order.ClientID, order.RouteID, order.DriverID, managerId, order.CargoWeight, order.Price)
	row := r.db.QueryRow(creatOrder)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrderDB) GetAllDB(managerId int) ([]models.Orders, error) {
	var orders []models.Orders
	query := fmt.Sprintf("SELECT * FROM %s", ordersTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := models.Orders{}
		var price []uint8
		err := rows.Scan(
			&o.OrderID, &o.ClientID, &o.RouteID, &o.DriverID, &o.ManagerID, &o.CargoWeight, &price, &o.Date)
		if err != nil {
			return nil, err
		}
		// конвертация sql price = []uint8 -> uint
		o.Price = ConvertPriceToUint(price[:])
		orders = append(orders, o)
	}
	return orders, nil
}
 
func (r *OrderDB) GetByIdManagerDB(managerid, id int) (models.Orders, error) {
	var order models.Orders
	query := fmt.Sprintf("SELECT * FROM %s as od where od.OrderID = %d", ordersTable, id)
	row := r.db.QueryRow(query)
	var price []uint8
	err := row.Scan(&order.OrderID, &order.ClientID, &order.RouteID, &order.DriverID, &order.ManagerID, &order.CargoWeight, &price, &order.Date)
	if err != nil {
		return models.Orders{}, err
	}
	order.Price = ConvertPriceToUint(price[:])
	return order, nil
}


func (r *OrderDB) DeleteManagerDB(managerid, id int) error {
	query := fmt.Sprintf("DELETE FROM %s where OrderID = %d", ordersTable, id)
	_, err := r.db.Exec(query)
	return err
}

	
func (r *OrderDB) UpdateManagerDB(managerid, id int, input models.Orders) error {
	query := fmt.Sprintf("UPDATE %s SET ClientID = %d, RouteID = %d, DriverID = %d, ManagerID = %d, CargoWeight = %d, Price = %d WHERE OrderID = %d",
		ordersTable, input.ClientID, input.RouteID, input.DriverID,
		managerid, input.CargoWeight, input.Price, id)
	_, err := r.db.Exec(query)
	return err
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