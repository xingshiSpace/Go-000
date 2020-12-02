package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)


var db *sql.DB

func main() {
	var initDbErr error
	db, initDbErr = sql.Open("mysql", "root:root@/orderdb?charset=utf8")
	if initDbErr != nil {
		fmt.Printf("Error: login db err=%+v\n", initDbErr)
	}
	fmt.Printf("%+v", serviceProduct())
}

//service
func serviceProduct()(string, error)  {
	productname, err := daoQueryRowByProductId(10)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//process ErrNoRows
			return  "", nil
		}
			return  "", err
	}
		return  productname, nil
	}

	//DAO
	func daoQueryRowByProductId(id int) (string, error) {
		var productname string
		// if id doesn't exist, the err would be sql.ErrNoRows
		if err := db.QueryRow("select productname FROM product WHERE id=?", id).Scan(&productname);
			err != nil {
			return "", errors.Wrap(err, "failed to find product date")
		}
		return productname, nil
	}
