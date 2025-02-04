package local

import (
	"time"
)

type Service struct {
	Id              uint             `json:"id" gorm:"primarykey"`
	IconId          uint             `json:"icon_id"`
	Name            string           `json:"name"`
	Desc            string           `json:"desc"`
	Address         string           `json:"address"`
	Port            int              `json:"port"`
	ServiceTypeId   uint             `json:"service_type_id"`
	VersionServices []VersionService `gorm:"foreignkey:ServiceId"`
}

type ServiceType struct {
	Id                 uint                `json:"id" gorm:"primarykey"`
	Name               string              `json:"name"`
	Icon               string              `json:"icon"`
	Services           []Service           `gorm:"foreignkey:ServiceTypeId"`
	ScriptServiceTypes []ScriptServiceType `gorm:"foreignkey:ServiceTypeId"`
}

type ScriptServiceType struct {
	Id            uint   `json:"id" gorm:"primarykey"`
	ScriptId      uint   `json:"script_id"`
	ServiceTypeId uint   `json:"service_type_id"`
	Name          string `json:"name"`
}

type Script struct {
	Id                 uint                `json:"id" gorm:"primarykey"`
	IconId             uint                `json:"icon_id"`
	Name               string              `json:"name"`
	DescShort          string              `json:"desc_short"`
	DescLong           string              `json:"desc_long"`
	Created            time.Time           `json:"created"`
	ScriptServiceTypes []ScriptServiceType `gorm:"foreignkey:ScriptId"`
	Versions           []Version           `gorm:"foreignkey:ScriptId"`
	ScriptTag          []ScriptTag         `gorm:"foreignkey:ScriptId"`
}

type Version struct {
	Id              uint             `json:"id" gorm:"primarykey"`
	ScriptId        uint             `json:"script_id"`
	Created         time.Time        `json:"service_type_id"`
	Major           int64            `json:"major"`
	Minor           int64            `json:"minor"`
	Patch           int              `json:"patch"`
	Commit          int              `json:"commit"`
	Uuid            string           `json:"uuid"`
	Change          string           `json:"change"`
	VersionCodes    []VersionCode    `gorm:"foreignkey:VersionId"`
	VersionServices []VersionService `gorm:"foreignkey:VersionId"`
}

type Code struct {
	Id           uint          `json:"id" gorm:"primarykey"`
	Type         string        `json:"type"`
	Value        string        `json:"value"`
	Uuid         string        `json:"uuid"`
	VersionCodes []VersionCode `gorm:"foreignkey:CodeId"`
}

type VersionCode struct {
	Id        uint `json:"id" gorm:"primarykey"`
	VersionId uint `json:"version_id"`
	CodeId    uint `json:"code_id"`
}

type VersionService struct {
	Id         uint   `json:"id" gorm:"primarykey"`
	VersionId  uint   `json:"version_id"`
	ServiceId  uint   `json:"service_id"`
	VersionMin uint   `json:"version_min"`
	VersionMax uint   `json:"version_max"`
	Usage      string `json:"usage"`
	Details    string `json:"details"`
}

type TagType struct {
	Id   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Tag  []Tag  `gorm:"foreignkey:TagTypeId"`
}

type Tag struct {
	Id        uint        `json:"id" gorm:"primarykey"`
	TagTypeId uint        `json:"tag_type_id"`
	Name      string      `json:"name"`
	ScriptTag []ScriptTag `gorm:"foreignkey:TagId"`
}

type ScriptTag struct {
	Id       uint `json:"id" gorm:"primarykey"`
	ScriptId uint `json:"script_id"`
	TagId    uint `json:"tag_id"`
}

type Icon struct {
	Id        uint    `json:"id" gorm:"primarykey"`
	Code      string  `json:"code"`
	Directory string  `json:"directory"`
	Filename  string  `json:"filename"`
	Service   Service `gorm:"foreignkey:IconId"`
	Script    Script  `gorm:"foreignkey:IconId"`
}
