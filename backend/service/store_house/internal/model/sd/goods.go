package sd

import (
	"context"
	"net/http"

	"github.com/tinkler/mqttadmin/pkg/db"
	"github.com/tinkler/mqttadmin/pkg/status"
)

type Goodsview struct {
	GoodsID       int     `gorm:"column:GOODSID"`
	Code          string  `gorm:"column:CODE"`
	Name          string  `gorm:"column:NAME"`
	GoodsTypeName string  `gorm:"column:GOODSTYPENAME"`
	Unit          string  `gorm:"column:UNIT"`
	OnHand        float64 `gorm:"column:ONHAND"`
}

func (m *Goodsview) TableName() string {
	return "goodsview"
}

type GoodsviewCRUD struct{}

func (m *GoodsviewCRUD) SearchForGoodsViewByName(ctx context.Context, name string) ([]*Goodsview, error) {
	if name == "" {
		return nil, status.NewCn(http.StatusOK, "name is required", "名称不能为空")
	}
	var data []*Goodsview
	se := db.DB().Where(" name like ?", "%"+name+"%").Find(&data)
	if se.Error != nil {
		return nil, status.StatusInternalServer(se.Error)
	}
	return data, nil
}
