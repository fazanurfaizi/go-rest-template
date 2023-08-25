package models

type RolePermission struct {
	RoleID       uint `gorm:"foreignKey:role_id;references:auth.roles.id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PermissionID uint `gorm:"foreignKey:permission_id;references:auth.permissions.id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (rp *RolePermission) TableName() string {
	return "auth.role_permissions"
}
