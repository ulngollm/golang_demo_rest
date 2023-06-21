package main

import (
	repo "app/db"
	"strconv"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	router.GET("/", home)
	router.GET("/transactions", getTransactionsList)
	router.POST("/transactions/", addTransaction)
	router.GET("/transactions/:id", getTransaction)
	router.DELETE("/transactions/:id", deleteTransaction)

	router.Run(":3333")
}

func home(c *gin.Context) {
	c.String(200, "Homepage")
}

func getTransactionsList(c *gin.Context) {
	transactions, err := repo.GetList()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, transactions)
}

func addTransaction(c *gin.Context) {
	var t repo.Transaction
	err := c.BindJSON(&t)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	t, err = repo.Save(t)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, t)
}

func getTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	t, err := repo.GetOne(id)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, t)
}

func deleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repo.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}