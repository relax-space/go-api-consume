package models

import (
	"context"
	"time"

	"github.com/go-xorm/xorm"

	"go-api-consumer/factory"
)

type Fruit struct {
	Id        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Price     int64     `json:"price"`
	StoreCode string    `json:"storeCode"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

func (d *Fruit) Create(ctx context.Context) (affectedRow int64, err error) {
	affectedRow, err = factory.DB(ctx).Insert(d)
	return
}

func (d *Fruit) Create2(db *xorm.Engine) (affectedRow int64, err error) {
	affectedRow, err = db.Insert(d)
	return
}
