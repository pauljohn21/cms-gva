package cms

import (
	"github.com/pauljohn21/cms-gva/server/global"
	"github.com/pauljohn21/cms-gva/server/model/cms"
    cmsReq "github.com/pauljohn21/cms-gva/server/model/cms/request"
    "gorm.io/gorm"
)

type ContactsService struct {}

// CreateContacts 创建联系人记录

func (contactsService *ContactsService) CreateContacts(contacts *cms.Contacts) (err error) {
	err = global.GVA_DB.Create(contacts).Error
	return err
}

// DeleteContacts 删除联系人记录

func (contactsService *ContactsService)DeleteContacts(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cms.Contacts{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&cms.Contacts{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteContactsByIds 批量删除联系人记录

func (contactsService *ContactsService)DeleteContactsByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cms.Contacts{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&cms.Contacts{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateContacts 更新联系人记录

func (contactsService *ContactsService)UpdateContacts(contacts cms.Contacts) (err error) {
	err = global.GVA_DB.Model(&cms.Contacts{}).Where("id = ?",contacts.ID).Updates(&contacts).Error
	return err
}

// GetContacts 根据ID获取联系人记录

func (contactsService *ContactsService)GetContacts(ID string) (contacts cms.Contacts, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&contacts).Error
	return
}

// GetContactsInfoList 分页获取联系人记录

func (contactsService *ContactsService)GetContactsInfoList(info cmsReq.ContactsSearch) (list []cms.Contacts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cms.Contacts{})
    var contactss []cms.Contacts
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
	
	err = db.Find(&contactss).Error
	return  contactss, total, err
}