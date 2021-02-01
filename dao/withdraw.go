package dao

import "gorm.io/gorm"

type Withdraw struct {
	Common
	GroupId   uint
	Recipient string `gorm:"type:char(64)"`
	Amount    uint64
}

type WithdrawDao struct {
	db *gorm.DB
}

func NewWithdrawDao(db *gorm.DB) (*WithdrawDao, error) {
	return &WithdrawDao{
		db: db,
	}, nil
}

func (d *WithdrawDao) Create(model *Withdraw) (*Withdraw, error) {
	if err := d.db.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (d *WithdrawDao) Get(id uint) (*Withdraw, error) {
	var withdraw Withdraw
	if err := d.db.Model(&Withdraw{}).Where("`id` = ?", id).First(&withdraw).Error; err != nil {
		return nil, err
	}

	return &withdraw, nil
}
