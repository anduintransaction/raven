package smtpserver

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/anduintransaction/raven/api/raven/database"
	"github.com/anduintransaction/raven/api/raven/model"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"
	"github.com/jhillyerd/enmime"
	"github.com/palantir/stacktrace"

	gomail "net/mail"
)

var unableToSaveErr = backends.NewResult("554 Error: could not save email")

// Postgres processor
func Postgres() backends.Decorator {
	return func(p backends.Processor) backends.Processor {
		return backends.ProcessWith(
			func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
				if task == backends.TaskSaveMail {
					emails, err := parseSMTPMessage(e)
					if err != nil {
						logrus.Error(err)
						return unableToSaveErr, backends.StorageError
					}
					if len(emails) == 0 {
						logrus.Errorf("no email submitted")
						return unableToSaveErr, backends.StorageError
					}
					err = saveEmails(emails)
					if err != nil {
						logrus.Error(err)
						return unableToSaveErr, backends.StorageError
					}
					return p.Process(e, task)
				}
				return p.Process(e, task)
			},
		)
	}
}

func parseSMTPMessage(e *mail.Envelope) ([]*model.Email, error) {
	env, err := enmime.ReadEnvelope(e.NewReader())
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot read envelope")
	}
	fromAddress, err := getSingleAddress(env, "From")
	if err != nil {
		return nil, err
	}
	rcptAddresses, err := getRcptAddresses(env)
	if err != nil {
		return nil, err
	}
	replyTo, err := getSingleAddress(env, "Reply-To")
	if err != nil {
		return nil, err
	}
	subject := env.GetHeader("Subject")
	htmlContent := env.HTML
	attachments := []*model.Attachment{}
	for _, attachmentData := range env.Attachments {
		attachments = append(attachments, &model.Attachment{
			Filename: attachmentData.FileName,
			Filesize: int64(len(attachmentData.Content)),
			Filemime: attachmentData.ContentType,
			AttachmentData: &model.AttachmentData{
				Content: attachmentData.Content,
			},
		})
	}
	rcpts := displayRCPT(rcptAddresses)
	emails := []*model.Email{}
	for _, rcptAddress := range rcptAddresses {
		email := &model.Email{
			FromEmail: fromAddress.Address,
			FromName:  fromAddress.Name,
			ToEmail:   rcptAddress.Address,
			ToName:    rcptAddress.Name,
			RCPT:      rcpts,
			ReplyTo:   replyTo.String(),
			Subject:   subject,
			EmailContent: &model.EmailContent{
				HTML: htmlContent,
			},
			Attachments: attachments,
		}
		emails = append(emails, email)
	}
	return emails, nil
}

func getSingleAddress(env *enmime.Envelope, key string) (*gomail.Address, error) {
	addresses, err := env.AddressList(key)
	if err != nil && err != gomail.ErrHeaderNotPresent {
		return nil, err
	}
	if len(addresses) == 0 {
		return nil, nil
	}
	return addresses[0], nil
}

func getRcptAddresses(env *enmime.Envelope) ([]*gomail.Address, error) {
	rcptAddressMap := make(map[string]*gomail.Address)
	err := mergeAddresses(env, "To", rcptAddressMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot parse to address")
	}
	err = mergeAddresses(env, "Cc", rcptAddressMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot parse cc address")
	}
	err = mergeAddresses(env, "Bcc", rcptAddressMap)
	if err != nil {
		return nil, stacktrace.Propagate(err, "cannot parse bcc address")
	}
	rcptAddresses := []*gomail.Address{}
	for _, address := range rcptAddressMap {
		rcptAddresses = append(rcptAddresses, address)
	}
	return rcptAddresses, nil
}

func mergeAddresses(env *enmime.Envelope, key string, rcptMap map[string]*gomail.Address) error {
	addresses, err := env.AddressList(key)
	if err != nil && err != gomail.ErrHeaderNotPresent {
		return err
	}
	for _, address := range addresses {
		rcptMap[address.Address] = address
	}
	return nil
}

func displayRCPT(rcpts []*gomail.Address) string {
	addresses := []string{}
	for _, rcpt := range rcpts {
		addresses = append(addresses, rcpt.String())
	}
	return strings.Join(addresses, ", ")
}

func saveEmails(emails []*model.Email) error {
	message := &model.Message{}
	database.Connection.NewRecord(message)
	err := database.Connection.Create(message).Error
	if err != nil {
		return stacktrace.Propagate(err, "cannot create message")
	}
	for _, email := range emails {
		email.MessageID = message.ID
		database.Connection.NewRecord(email)
		err = database.Connection.Create(email).Error
		if err != nil {
			return stacktrace.Propagate(err, "cannot create email")
		}
		logrus.Infof("Email created: %d", email.ID)
	}
	return nil
}
