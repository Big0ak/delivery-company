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

func NewOrderDB (db *sql.DB) *OrderDB {
	return &OrderDB{db: db}
}

func (r *OrderDB) CreateDB (managerId int, order models.Orders) (int, error) {
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

	for rows.Next(){
		o := models.Orders{}
		var x []uint8
		err := rows.Scan(&o.OrderID, &o.ClientID, &o.RouteID, &o.DriverID, &o.ManagerID, &o.CargoWeight, &x, &o.Date)
		
		// b := make([]byte, len(x))
    	// for i, v := range x {
        // 	b[i] = byte(v)
    	// }
    	// str := string(b)
		// xx, err := strconv.ParseUint(str, 10, 32)
		// o.Price = uint(xx)
		
		if err != nil{
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}