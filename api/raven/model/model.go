package model

import "github.com/jinzhu/gorm"

// Message .
type Message struct {
	gorm.Model
}

// Email .
type Email struct {
	gorm.Model
	MessageID      uint `gorm:"column:message_id;INDEX"`
	Message        *Message
	FromEmail      string `gorm:"column:from_email;type:varchar(255);INDEX"`
	FromName       string `gorm:"column:from_name;type:varchar(255);INDEX"`
	ToEmail        string `gorm:"column:to_email;type:varchar(255);INDEX"`
	ToName         string `gorm:"column:to_name;type:varchar(255);INDEX"`
	RCPT           string `gorm:"column:rcpt;type:text"`
	ReplyTo        string `gorm:"column:reply_to"`
	Subject        string `gorm:"column:subject;type:text;INDEX"`
	EmailContent   *EmailContent
	EmailContentID uint
	Attachments    []*Attachment `gorm:"many2many:email_attachments"`
}

// EmailContent .
type EmailContent struct {
	gorm.Model
	HTML string `gorm:"column:html;type:text"`
}

// Attachment .
type Attachment struct {
	gorm.Model
	Filename         string `gorm:"column:filename;type:text;INDEX"`
	Filemime         string `gorm:"column:filemime;type:text"`
	Filesize         int64  `gorm:"column:filesize"`
	AttachmentData   *AttachmentData
	AttachmentDataID uint
}

// AttachmentData .
type AttachmentData struct {
	gorm.Model
	Content []byte `gorm:"column:content;type:bytea"`
}
