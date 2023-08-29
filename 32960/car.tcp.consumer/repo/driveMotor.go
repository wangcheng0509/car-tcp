package repo

import (
	"context"

	"car.tcp.consumer/entity/model"
	"gorm.io/gorm/clause"
)

// DrivemotorUpsert Drivemotor
func DrivemotorUpsert(ctx context.Context, drivemotors model.Drivemotors) error {
	db := getDBWithTable(ctx, &model.Drivemotors{})
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "vin"}},
		UpdateAll: true,
	}).Create(&drivemotors).Error
	if err != nil {
		return err
	}
	return nil
}

// DrivemotorLogCreate drivemotorLog
func DrivemotorLogCreate(ctx context.Context, drivemotorLog model.DrivemotorsLog) error {
	err := getCkDBWithTable(ctx, &model.DrivemotorsLog{}).Create(&drivemotorLog).Error
	if err != nil {
		return err
	}
	return nil
}
