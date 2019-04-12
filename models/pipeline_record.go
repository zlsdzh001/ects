package models

import (
	"time"
)

type PipelineRecords struct {
	Id         string    `xorm:"not null pk comment('ID') CHAR(36)"`
	PipelineId string    `xorm:"not null comment('流水线ID') index CHAR(36)"`
	TeamId     string    `xorm:"not null comment('团队ID') index CHAR(36)"`
	NodeId     string    `xorm:"not null comment('节点ID') index CHAR(36)"`
	WorkerName string    `xorm:"not null comment('节点名称') VARCHAR(255)"`
	Spec       string    `xorm:"comment('定时器') CHAR(64)"`
	Status     int       `xorm:"not null default 1 comment('状态') TINYINT(1)"`
	CreatedAt  time.Time `xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt  time.Time `xorm:"not null updated comment('更新于') DATETIME"`
}

// 定义模型的数据表名称
func (records *PipelineRecords) TableName() string {
	return "pipeline_records"
}