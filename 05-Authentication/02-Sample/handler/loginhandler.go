package handler

import (
	"Authentication/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	var loginObj models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginObj); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})

		badRequest(ctx, http.StatusBadRequest, "Invalid request", errors)
	}

	var claims = &models.JwtClaims{}
	claims.CompanyId = "CompanyId"
	claims.Username = loginObj.UserName
	claims.Roles = []int{1, 2, 3}
	claims.Audience = ctx.Request.Header.Get("Referer")

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(10) * time.Minute)
}
