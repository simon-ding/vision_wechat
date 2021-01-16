package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"path"
)

type Account struct {
	gorm.Model
	UserID   string
	Name     string
	Username string
	Password string
	Cookie   string
}

//var DB *gorm.DB

type DB struct {
	db *gorm.DB
}

var DefaultDB *DB

func InitConnection() {
	dataDir := viper.GetString("server.dataDir")
	dataDir = path.Join(dataDir, "wechat.db")
	db, err := gorm.Open("sqlite3", dataDir)
	if err != nil {
		panic("failed to connect database")
	}
	DefaultDB = &DB{db: db}
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) GetInstagram(openID string) Account {
	var ins Account
	d.db.Where("user_id = ?", openID).Where("name = ?", "instagram").First(&ins)
	if ins.ID == 0 {
		logrus.Errorf("no instagram account for user id %s", openID)
		d.db.Create(&Account{UserID: openID})
	}
	return ins
}

func (d *DB) GetAll500px() []Account {
	var px []Account
	d.db.Where("name = ?", "500px").Find(&px)
	return px
}
func (d *DB) Set500pxAccount(openID string, username, password string) {
	var acc Account
	d.db.Where("name = ?", "500px").Where("user_id = ?", openID).First(&acc)
	if acc.ID == 0 {
		logrus.Info("用户没有登记500px账号，创建一个")
		acc.Name = "500px"
		acc.UserID = openID
	}
	acc.Username = username
	acc.Password = password
	d.db.Save(&acc)
}

func (d *DB) Migrate() {
	d.db.AutoMigrate(&Account{})
}
