package repo

import (
	"log"

	"car.tcp.consumer/entity/dbmodel"
)

func FuelcellCreate(v dbmodel.Fuelcell, vLog dbmodel.FuelcellLog) error {
	if err := ckDB.Create(&vLog).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	var temp dbmodel.Fuelcell
	mysqlDB.Model(&temp).Where("vin=?", v.Vin).Find(&temp)
	if len(temp.Vin) > 0 {
		if err := mysqlDB.Model(&temp).Where("vin=?", temp.Vin).Save(&v).Error; err != nil {
			log.Println(err.Error())
			return err
		}
	} else {
		if err := mysqlDB.Create(&v).Error; err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}
