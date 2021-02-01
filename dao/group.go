package dao

import (
	"fmt"

	"gorm.io/gorm"
)

type Group struct {
	Common
	Address      string `gorm:"type:char(64)"`
	RedeemScript string `gorm:"type:char(256)"`
}

type GroupKeeper struct {
	Common
	GroupId  uint
	KeeperId uint
}

type GroupDao struct {
	db *gorm.DB
}

func NewGroupDao(db *gorm.DB) (*GroupDao, error) {
	return &GroupDao{
		db: db,
	}, nil
}

func (d *GroupDao) Create(model *Group) (*Group, error) {
	if err := d.db.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (d *GroupDao) Get(id uint) (*Group, error) {
	var group Group
	if err := d.db.Model(&Group{}).Where("`id` = ?", id).First(&group).Error; err != nil {
		return nil, err
	}

	return &group, nil
}

type GroupKeeperDao struct {
	db *gorm.DB
}

func NewGroupKeeperDao(db *gorm.DB) (*GroupKeeperDao, error) {
	return &GroupKeeperDao{
		db: db,
	}, nil
}

func (d *GroupKeeperDao) Create(model *GroupKeeper) (*GroupKeeper, error) {
	var count int64
	if err := d.db.Model(&GroupKeeper{}).Where("`group_id` = ? and `keeper_id` = ?",
		model.GroupId, model.KeeperId).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("group keeper %d:%d already exist", model.GroupId, model.KeeperId)
	}
	if err := d.db.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (d *GroupKeeperDao) CheckInGroup(groupId, keeperId uint) (bool, error) {
	var count int64
	if err := d.db.Model(&GroupKeeper{}).Where("`group_id` = ? and `keeper_id` = ?",
		groupId, keeperId).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
