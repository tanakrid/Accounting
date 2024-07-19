package route

import (
	"fmt"
	"log"
	"net/http"
	"github.com/labstack/echo"
	"github.com/tanakrid/accounting/transaction"
)

func postTransaction(c echo.Context) error {
	record := &transaction.Record{}
	if err := c.Bind(record); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := transaction.Add(*record)
	return c.JSON(http.StatusCreated, id)
}

func getTransaction(c echo.Context) error {
	return c.JSON(http.StatusOK, transaction.Show())
}

func getTransactionById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "Can't find id path param")
	}
	record := transaction.ShowById(id)
	return c.JSON(http.StatusOK, record)
}

func putTransaction(c echo.Context) error {
	// if _, has := transaction.Show()[id]; !has {
	// 	return c.JSON(http.StatusNotFound, "Can't edit, Not found this id:"+id)
	// }

	record := &transaction.Record{}
	fmt.Printf("record's value: %v\n", record) // test value
	if err := c.Bind(record); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Printf("record: %#v\n", record)
	transaction.Edit(*record)
	return c.JSON(http.StatusOK, nil)
}

func deleteTransaction(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "Can't find id path param")
	}
	// if _, has := transaction.Show()[id]; !has {
	// 	return c.JSON(http.StatusNotFound, "Can't delete, Not found this id:"+id)
	// }
	transaction.Del(id)
	return c.JSON(http.StatusOK, nil)
}

func InitRoute() {
	fmt.Println("init route and start service")
	port := "4444"
	e := echo.New()

	e.GET("/transaction", getTransaction)
	e.GET("/transaction/:id", getTransactionById)
	e.POST("/transaction", postTransaction)
	e.DELETE("/transaction/:id", deleteTransaction)
	e.PUT("/transaction", putTransaction)

	log.Println("starting... port:", port)
	log.Fatal(e.Start(":"+port))
}