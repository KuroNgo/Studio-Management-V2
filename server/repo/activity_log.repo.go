package repo

import (
	"co-studio-e-commerce/model"
)

func (r *Repo) GetActivityLog(log *model.ActivityLog) error {
	// GetActivityLog là hàm lấy thông tin log của user
	if err := r.db.Where(log).First(log).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) ViewHistoryUserActitvity(log *[]model.ActivityLog) error {
	// GetAllActivityLog là hàm lấy thông tin tất cả log của user
	if err := r.db.Find(log).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) CreateActivityLog(log *model.ActivityLog) error {
	// CreateActivityLog là hàm tạo mới log của user
	if err := r.db.Create(log).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateActivityLog(log *model.ActivityLog) error {
	// UpdateActivityLog là hàm cập nhật thông tin log của user
	if err := r.db.Save(log).Error; err != nil {
		return err
	}
	return nil
}

// DeleteActivityLog lập lịch xóa ( admin không có quyền xóa log của user )
// Dữ liệu activity log sẽ bị xóa khi user bị xóa
// Dữ liệu activity log sẽ bị xóa khi > 30 days không đăng nhập
func (r *Repo) DeleteActivityLog(log *model.ActivityLog) error {
	// DeleteActivityLog là hàm xóa thông tin log của user
	if err := r.db.Delete(log).Error; err != nil {
		return err
	}
	return nil
}
