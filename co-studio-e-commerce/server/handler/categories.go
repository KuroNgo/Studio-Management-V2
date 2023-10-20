package handler

import (
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Category struct {
	service  service.ICategory
	category model.Categories
	user     service.IUser
}

func NewCategory(service service.ICategory) *Category {
	return &Category{service: service}
}

func (c *Category) GetAllCategories(ctx *gin.Context) {
	categories, err := c.service.GetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"category": categories}})
}

func (c *Category) GetCategory(category model.Categories) {

}

func (c *Category) CreateCategory(ctx *gin.Context) {
	// Lấy user hiện đang đăng nhập
	currentUser := ctx.MustGet("currentUser")
	var category model.Categories
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// Thực hiện lấy thông tin người dùng
	//userName, err := c.user.FindUserByID(fmt.Sprint(currentUser))
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  "error",
	//		"message": err.Error(),
	//	})
	//}

	category.UpdatedAt = time.Now()
	category.CreatedAt = time.Now()

	// Thực hiện thêm category
	newCategory := model.Categories{
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Enable:       1,
		IsUpdate:     category.IsUpdate,
		UpdatedAt:    category.UpdatedAt,
		WhoUpdate:    fmt.Sprint(currentUser),
		IsDelete:     0,
	}

	categoryResponse, err := c.service.CreateCategory(newCategory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   categoryResponse,
	})

}