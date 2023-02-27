package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/HasnathJami/go-echo-redis/database"
	"github.com/HasnathJami/go-echo-redis/models"
	"github.com/HasnathJami/go-echo-redis/utils"
	"github.com/labstack/echo"
)

var (
	ctx = context.Background()
	password string
)	

func GenerateStrongPassword(context echo.Context) error {
	 rdb := database.CreateClient(0)
	 password = context.Param("password")
	 err := rdb.Set(ctx, "password", password, 24*3600*time.Second).Err()
	 utils.CheckSimpleError(err)
	 err = rdb.Set(ctx, "generated_password", utils.GetPasswordByCaesarCipher(password), 24*3600*time.Second).Err()
	 utils.CheckSimpleError(err)
	 defer rdb.Close()
	 
	 return context.JSON(http.StatusOK, "Password Set Successfully")
}

func GetStrongPassword(context echo.Context) error {
	 rdb := database.CreateClient(1)
	 password, err := rdb.Get(ctx, "password").Result()
	 utils.CheckSimpleError(err)

	 generatedPassword, err := rdb.Get(ctx, "generated_password").Result()
	 utils.CheckSimpleError(err)
	 defer rdb.Close()

     return context.JSON(http.StatusOK, &models.Response{
		Password : password,
		GeneratedPassword : generatedPassword,
		Message : "Password Fetched Successfully",	
    })
}