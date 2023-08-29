package repo

import (
	"context"

	"car.tcp.consumer/entity/model"
	"gorm.io/gorm/clause"
)

// ChargeableVoltageUpsert chargeableVoltage
func ChargeableVoltageUpsert(ctx context.Context, chargeableVoltage model.ChargeableVoltage) error {
	db := getDBWithTable(ctx, &model.ChargeableVoltage{})
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "vin"}},
		UpdateAll: true,
	}).Create(&chargeableVoltage).Error
	if err != nil {
		return err
	}
	return nil
}

// ChargeableVoltageLogCreate chargeableVoltage log
func ChargeableVoltageLogCreate(ctx context.Context, chargeableVoltageLog model.ChargeableVoltageLog) error {
	err := getCkDBWithTable(ctx, &model.ChargeableVoltageLog{}).Create(&chargeableVoltageLog).Error
	if err != nil {
		return err
	}
	return nil
}
