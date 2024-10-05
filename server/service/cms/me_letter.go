package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/model/cms"
	cmsReq "github.com/pauljohn21/cms-gva/server/model/cms/request"
	"gorm.io/gorm"
)

type MeLetterService struct{}

// CreateMeLetter 创建我的保函记录

func (meLetterService *MeLetterService) CreateMeLetter(meLetter *cms.MeLetter) (err error) {
	err = global.GVA_DB.Create(meLetter).Error
	return err
}

// DeleteMeLetter 删除我的保函记录

func (meLetterService *MeLetterService) DeleteMeLetter(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cms.MeLetter{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&cms.MeLetter{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMeLetterByIds 批量删除我的保函记录

func (meLetterService *MeLetterService) DeleteMeLetterByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cms.MeLetter{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&cms.MeLetter{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMeLetter 更新我的保函记录

func (meLetterService *MeLetterService) UpdateMeLetter(meLetter cms.MeLetter) (err error) {
	err = global.GVA_DB.Model(&cms.MeLetter{}).Where("id = ?", meLetter.ID).Updates(&meLetter).Error
	return err
}

// GetMeLetter 根据ID获取我的保函记录

func (meLetterService *MeLetterService) GetMeLetter(ID string) (meLetter cms.MeLetter, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&meLetter).Error
	return
}

// GetMeLetterInfoList 分页获取我的保函记录

func (meLetterService *MeLetterService) GetMeLetterInfoList(info cmsReq.MeLetterSearch) (list []cms.MeLetter, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cms.MeLetter{}).Order("id desc")
	var meLetters []cms.MeLetter
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&meLetters).Error
	return meLetters, total, err
}
