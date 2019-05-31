package types

import (
	"github.com/jinzhu/gorm"

	"github.com/suda/go-swagger-example/models"
)

type Tweet struct {
	gorm.Model
	models.Tweet
}
