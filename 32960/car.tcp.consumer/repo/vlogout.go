package repo

import (
	"context"

	"car.tcp.consumer/entity/model"
	"gorm.io/gorm/clause"
)

// VlogoutUpsert vlogout
func VlogoutUpsert(ctx context.Context, vlogout model.Vlogout) error {
	db := getDBWithTable(ctx, &model.Vlogout{})
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "vin"}},
		UpdateAll: true,
	}).Create(&vlogout).Error
	if err != nil {
		return err
	}
	return nil
}

// VlogoutLogCreate vlogout log
func VlogoutLogCreate(ctx context.Context, vlogoutLog model.VlogoutLog) error {
	err := getCkDBWithTable(ctx, &model.VlogoutLog{}).Create(&vlogoutLog).Error
	if err != nil {
		return err
	}
	return nil
}
