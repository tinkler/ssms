package commodity

import (
	"context"

	"github.com/tinkler/mqttadmin/pkg/db"
	"github.com/tinkler/mqttadmin/pkg/status"
	"gorm.io/gorm"
)

// 商品
type Commodity struct {
	Id          string
	Name        string
	Price       float64
	Stock       int
	Unit        string
	Description string
	DeleteAt    gorm.DeletedAt
}

// Save adds a new commodity to the database when the commodity id is empty
// or updates an existing commodity when the commodity id is not empty
func (m *Commodity) Save(ctx context.Context) error {
	if m.Id != "" {
		if se := db.DB().Model(m).Select("name", "price", "stock", "unit", "description").Updates(m); se.Error != nil {
			return se.Error
		}
	} else {
		if se := db.DB().Model(m).Create(m); se.Error != nil {
			return se.Error
		}
	}
	return nil
}

// Delete deletes a commodity from the database
func (m *Commodity) Delete(ctx context.Context) error {
	if m.Id == "" {
		return status.Ok("no commodity id")
	}
	if se := db.DB().Model(m).Delete(m); se.Error != nil {
		return se.Error
	}
	return nil
}
