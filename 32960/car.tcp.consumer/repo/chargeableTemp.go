package repo

import (
	"context"

	"car.tcp.consumer/entity/model"
	"gorm.io/gorm/clause"
)

// ChargeableTempUpsert chargeableTemp
func ChargeableTempUpsert(ctx context.Context, chargeableTemp model.ChargeableTemp) error {
	db := getDBWithTable(ctx, &model.ChargeableTemp{})
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "vin"}},
		UpdateAll: true,
	}).Create(&chargeableTemp).Error
	if err != nil {
		return err
	}
	return nil
}

// ChargeableTempLogCreate chargeableTemp log
func ChargeableTempLogCreate(ctx context.Context, chargeableTempLog model.ChargeableTempLog) error {
	err := getCkDBWithTable(ctx, &model.ChargeableTempLog{}).Create(&chargeableTempLog).Error
	if err != nil {
		return err
	}
	return nil
}
