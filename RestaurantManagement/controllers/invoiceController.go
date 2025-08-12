package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get all invoices
	}
}	

func GetInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get invoice by ID
	}
}
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to create a new invoice
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to update an existing invoice
	}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to delete an invoice
	}
}