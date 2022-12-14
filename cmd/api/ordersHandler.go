package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type order struct {
	Id          uuid.UUID `json:"id"`
	Items       []item    `json:"items"`
	CreatedDate time.Time `json:"createdDate"`
}

type item struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

var orders = []order{
	{
		Id: uuid.New(),
		Items: []item{
			{
				Id:    uuid.New(),
				Name:  "Fit Aid - Citrus Medley",
				Price: 2.50,
			},
			{
				Id:    uuid.New(),
				Name:  "Fit Aid - Citrus Medley",
				Price: 2.50,
			},
		},
		CreatedDate: time.Now(),
	},
	{
		Id: uuid.New(),
		Items: []item{
			{
				Id:    uuid.New(),
				Name:  "NOCCO - Lemon Del Sol",
				Price: 1.99,
			},
			{
				Id:    uuid.New(),
				Name:  "Bolt24 - Antioxidant",
				Price: 1.50,
			},
		},
		CreatedDate: time.Now(),
	},
}

func (app *application) getOrdersHandler(c *gin.Context) {
	time.Sleep(time.Second * 2)

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (app *application) saveOrderHandler(c *gin.Context) {
	time.Sleep(time.Second * 2)

	var data []item
	var err error
	if err = c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newOrder := order{
		Id:          uuid.New(),
		CreatedDate: time.Now(),
		Items:       data,
	}
	orders = append(orders, newOrder)
	c.JSON(http.StatusCreated, newOrder)
}
