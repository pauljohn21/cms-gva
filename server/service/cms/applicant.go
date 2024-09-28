package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/model/cms"
    cmsReq "github.com/pauljohn21/cms-gva/server/model/cms/request"
    "gorm.io/gorm"
)

type ApplicantService struct {}

// CreateApplicant 创建申请人记录

func (applicantService *ApplicantService) CreateApplicant(applicant *cms.Applicant) (err error) {
	err = global.GVA_DB.Create(applicant).Error
	return err
}

// DeleteApplicant 删除申请人记录

func (applicantService *ApplicantService)DeleteApplicant(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cms.Applicant{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&cms.Applicant{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteApplicantByIds 批量删除申请人记录

func (applicantService *ApplicantService)DeleteApplicantByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cms.Applicant{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&cms.Applicant{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateApplicant 更新申请人记录

func (applicantService *ApplicantService)UpdateApplicant(applicant cms.Applicant) (err error) {
	err = global.GVA_DB.Model(&cms.Applicant{}).Where("id = ?",applicant.ID).Updates(&applicant).Error
	return err
}

// GetApplicant 根据ID获取申请人记录

func (applicantService *ApplicantService)GetApplicant(ID string) (applicant cms.Applicant, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&applicant).Error
	return
}

// GetApplicantInfoList 分页获取申请人记录

func (applicantService *ApplicantService)GetApplicantInfoList(info cmsReq.ApplicantSearch) (list []cms.Applicant, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cms.Applicant{})
    var applicants []cms.Applicant
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
	
	err = db.Find(&applicants).Error
	return  applicants, total, err
}