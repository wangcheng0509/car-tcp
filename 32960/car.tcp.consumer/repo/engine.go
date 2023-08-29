package repo

import (
	"context"

	"car.tcp.consumer/entity/model"
	"gorm.io/gorm/clause"
)

// EngineUpsert engine
func EngineUpsert(ctx context.Context, engine model.Engine) error {
	db := getDBWithTable(ctx, &model.Engine{})
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "vin"}},
		UpdateAll: true,
	}).Create(&engine).Error
	if err != nil {
		return err
	}
	return nil
}

// EngineLogCreate engine log
func EngineLogCreate(ctx context.Context, engineLog model.EngineLog) error {
	err := getCkDBWithTable(ctx, &model.EngineLog{}).Create(&engineLog).Error
	if err != nil {
		return err
	}
	return nil
}
