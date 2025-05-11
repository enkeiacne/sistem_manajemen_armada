package commonsHelperPagination

import (
	"gorm.io/gorm"
	"math"
)

type PaginationResult struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

func Paginate(db *gorm.DB, page, limit int, model interface{}, result interface{}) (*PaginationResult, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var total int64
	if err := db.Model(model).Count(&total).Error; err != nil {
		return nil, err
	}

	if err := db.Model(model).Limit(limit).Offset(offset).Find(result).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &PaginationResult{
		Page:       page,
		Limit:      limit,
		TotalRows:  total,
		TotalPages: totalPages,
		Data:       result,
	}, nil
}
