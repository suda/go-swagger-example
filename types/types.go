package types

import (
	"github.com/jinzhu/gorm"

	"github.com/suda/go-swagger-example/models"
)

// Tweet struct composed of GORM and the Swagger model
type Tweet struct {
	gorm.Model
	models.Tweet
}
