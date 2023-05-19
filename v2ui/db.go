package v2ui

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var v2db *gorm.DB

func initDB(dbPath string) error {
	c := &gorm.Config{
		Logger: logger.Discard,
	}
	var err error
	dsn := "yas:Yas2566*7425@tcp(79.137.203.149:3306)/x_ui"
	v2db, err = gorm.Open(mysql.Open(dsn), c)
	//v2db, err = gorm.Open(sqlite.Open(dbPath), c)
	if err != nil {
		return err
	}

	return nil
}

func getV2Inbounds() ([]*V2Inbound, error) {
	inbounds := make([]*V2Inbound, 0)
	err := v2db.Model(V2Inbound{}).Find(&inbounds).Error
	return inbounds, err
}
