package initialize

import (
	"gorm.io/gorm"
	"github.com/pauljohn21/cms-gva/server/model/cms"
)

func bizModel(db *gorm.DB) error {
	return db.AutoMigrate(cms.Applicant{}, cms.Court{}, cms.MeLetter{}, cms.Contacts{})
}
