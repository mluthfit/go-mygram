package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	BaseModel
	Name           string `json:"name" gorm:"not null" valid:"required" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" valid:"required" binding:"required"`
	UserID         uint   `json:"user_id" gorm:"not null" valid:"required"`
	User           User   `valid:"-" binding:"-"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(sm); err != nil {
		return err
	}

	return nil
}

func (sm *SocialMedia) Create(db *gorm.DB) error {
	return db.Debug().Create(sm).Error
}

func (sm *SocialMedia) GetAllWithUser(db *gorm.DB) ([]map[string]any, error) {
	var socialMedias []SocialMedia

	if err := db.Debug().Preload("User.Photos").
		Find(&socialMedias).Error; err != nil {
		return nil, err
	}

	return sm.mappingGetAll(socialMedias), nil
}

func (sm *SocialMedia) Update(db *gorm.DB, newSocialMedia SocialMedia) error {
	if err := db.Debug().Model(sm).
		Updates(newSocialMedia).Error; err != nil {
		return err
	}

	return db.Debug().First(sm).Error
}

func (sm *SocialMedia) Delete(db *gorm.DB) error {
	return db.Debug().Delete(sm).Error
}

func (sm *SocialMedia) mappingGetAll(socialMedias []SocialMedia) (results []map[string]any) {

	for _, socialMedia := range socialMedias {
		var profileImageUrl string
		if len(socialMedia.User.Photos) > 0 {
			profileImageUrl = socialMedia.User.Photos[0].PhotoUrl
		}

		var data = map[string]any{
			"id":               socialMedia.ID,
			"name":             socialMedia.Name,
			"social_media_url": socialMedia.SocialMediaUrl,
			"UserId":           socialMedia.UserID,
			"created_at":       socialMedia.CreatedAt,
			"updated_at":       socialMedia.UpdatedAt,
			"User": map[string]any{
				"id":                socialMedia.User.ID,
				"email":             socialMedia.User.Email,
				"profile_image_url": profileImageUrl,
			},
		}

		results = append(results, data)
	}

	return
}
