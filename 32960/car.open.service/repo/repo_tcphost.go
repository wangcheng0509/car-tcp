package repo

import (
	"context"

	"car.open.service/repo/model"
)

// GetTCPAddr 获取tcp addr
func GetTCPAddr(ctx context.Context, vin string) (*model.VehicleConn, error) {
	db := getDBWithTable(ctx, &model.VehicleConn{})
	db = db.Where("vin=?", vin).Where("status=1")
	var vhicleConn model.VehicleConn
	err := db.Find(&vhicleConn).Error
	if err != nil {
		return nil, err
	}

	return &vhicleConn, nil
}
