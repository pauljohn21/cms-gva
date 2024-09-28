package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/model/cms"
    cmsReq "github.com/pauljohn21/cms-gva/server/model/cms/request"
    "gorm.io/gorm"
)

type CourtService struct {}

// CreateCourt 创建法院记录

func (courtService *CourtService) CreateCourt(court *cms.Court) (err error) {
	err = global.GVA_DB.Create(court).Error
	return err
}

// DeleteCourt 删除法院记录

func (courtService *CourtService)DeleteCourt(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cms.Court{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&cms.Court{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCourtByIds 批量删除法院记录

func (courtService *CourtService)DeleteCourtByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cms.Court{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&cms.Court{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCourt 更新法院记录

func (courtService *CourtService)UpdateCourt(court cms.Court) (err error) {
	err = global.GVA_DB.Model(&cms.Court{}).Where("id = ?",court.ID).Updates(&court).Error
	return err
}

// GetCourt 根据ID获取法院记录

func (courtService *CourtService)GetCourt(ID string) (court cms.Court, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&court).Error
	return
}

// GetCourtInfoList 分页获取法院记录

func (courtService *CourtService)GetCourtInfoList(info cmsReq.CourtSearch) (list []cms.Court, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cms.Court{})
    var courts []cms.Court
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&courts).Error
	return  courts, total, err
}