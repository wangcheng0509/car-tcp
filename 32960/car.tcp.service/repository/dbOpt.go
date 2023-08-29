package repository

import (
	"log"
	"time"

	"car.tcp.service/entity/handleEntity"
)

// 记录客户端链接
func CreateConn(reqMsg handleEntity.CacheModel) error {
	var dbModel handleEntity.CacheModel
	ServiceDBA.Mysql.Where("vin=?", reqMsg.Vin).Find(&dbModel)
	if len(dbModel.Vin) > 0 {
		return ServiceDBA.Mysql.Model(&handleEntity.CacheModel{Id: dbModel.Id}).Updates(handleEntity.CacheModel{
			Host:       reqMsg.Host,
			Port:       "8080",
			ConnStr:    reqMsg.ConnStr,
			Status:     1,
			OnlineDate: time.Now(),
		}).Error
	} else {
		reqMsg.OnlineDate = time.Now()
		return ServiceDBA.Mysql.Create(&reqMsg).Error
	}
}

// 客户端离线
func OfflineByVin(vin string) error {
	var dbModel handleEntity.CacheModel
	ServiceDBA.Mysql.Where("vin=?", vin).Find(dbModel)
	if len(dbModel.Vin) > 0 {
		return ServiceDBA.Mysql.Model(&handleEntity.CacheModel{Id: dbModel.Id}).Updates(handleEntity.CacheModel{
			Status:      2,
			OfflineDate: time.Now(),
		}).Error
	}
	return nil
}

// 客户端离线
func OfflineByConnIdStr(host, connStr string) error {
	var dbModel handleEntity.CacheModel
	ServiceDBA.Mysql.Where("host=? and connStr=?", host, connStr).Find(&dbModel)
	if len(dbModel.Vin) > 0 {
		return ServiceDBA.Mysql.Model(&handleEntity.CacheModel{Id: dbModel.Id}).Updates(handleEntity.CacheModel{
			Status:      2,
			OfflineDate: time.Now(),
		}).Error
	}
	return nil
}

func MsgLogAdd(msg handleEntity.Message) {
	if err := ServiceDBA.Mysql.Create(&msg).Error; err != nil {
		log.Println(err.Error())
	}
}
