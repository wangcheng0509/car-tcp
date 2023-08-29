package repo

import (
	"context"

	"car.tcp.consumer/entity/dbmodel"
)

// GetTCPAddr 获取tcp addr
func GetTCPAddr(ctx context.Context, vin string) (*dbmodel.VehicleConn, error) {
	db := getDBWithTable(ctx, &dbmodel.VehicleConn{})
	db = db.Where("vin=?", vin).Where("status=1")
	var vhicleConn dbmodel.VehicleConn
	err := db.Find(&vhicleConn).Error
	if err != nil {
		return nil, err
	}

	return &vhicleConn, nil
}
