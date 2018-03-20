package model

import "github.com/jinzhu/gorm"

// Email .
type Email struct {
	gorm.Model
	From        string        `gorm:"column:from;type:varchar(255);INDEX"`
	FromName    string        `gorm:"column:from_name;type:varchar(255);INDEX"`
	To          string        `gorm:"column:to;type:varchar(255);INDEX"`
	ToName      string        `gorm:"column:to_name;type:varchar(255);INDEX"`
	RCPT        string        `gorm:"column:rcpt;type:text"`
	ReplyTo     string        `gorm:"column:reply_to"`
	Subject     string        `gorm:"column:subject;type:text;INDEX"`
	HTML        string        `gorm:"column:html;type:text"`
	Attachments []*Attachment `gorm:"many2many:email_attachments"`
}

// Attachment .
type Attachment struct {
	gorm.Model
	EmailID  int64  `gorm:"column:email_id;INDEX"`
	Filename string `gorm:"column:filename;type:text;INDEX"`
	Filemime string `gorm:"column:filemime;type:text"`
	Filesize int64  `gorm:"column:filesize"`
	Content  []byte `gorm:"column:content;type:bytea"`
}
