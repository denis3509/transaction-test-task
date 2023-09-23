package transaction

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transaction/internal/entity"
	"transaction/pkg/db"
)

func CreateRoutes(router gin.IRouter, db *db.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	api := NewApi(service)
	router.GET("/user-balance", api.GetUserBalance)
	router.GET("/create-invoice", api.CreateInvoice)
	router.GET("/create-withdraw", api.CreateWithdraw)
}

type api struct {
	service Service
}

func NewApi(service Service) *api {
	return &api{service}
}

func (a *api) GetUserBalance(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			gin.H{"message": "internal error"})
		return
	}
	userBalance, err := a.service.GetUserBalance(userID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	c.IndentedJSON(http.StatusOK, userBalance)

}

type CreateInvoice struct {
	CardNumber int     `json:"cardId"`
	Currency   string  `json:"currency"`
	Amount     float64 `json:"amount"`
}

func (a *api) CreateInvoice(c *gin.Context) {
	var newInvoice CreateInvoice

	if err := c.BindJSON(&newInvoice); err != nil {
		return
	}

	if _, err := a.service.CreateInvoice(

		newInvoice.CardNumber,
		entity.Currency(newInvoice.Currency),
		newInvoice.Amount,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "invoice created"})
}

func (a *api) CreateWithdraw(c *gin.Context) {
	var newInvoice CreateInvoice

	if err := c.BindJSON(&newInvoice); err != nil {
		return
	}

	if _, err := a.service.CreateWithdraw(
		newInvoice.CardNumber,
		entity.Currency(newInvoice.Currency),
		newInvoice.Amount,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "internal error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "invoice created"})
}
