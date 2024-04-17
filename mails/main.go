package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime"
	"os"
	"time"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/emersion/go-message/charset"
	"github.com/emersion/go-message/mail"
	"github.com/k0kubun/pp"
	"go.uber.org/zap"
)

type MailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

var DefaultFetchOptions = &imap.FetchOptions{
	Flags:         true,
	InternalDate:  true,
	RFC822Size:    true,
	Envelope:      true,
	BodySection:   []*imap.FetchItemBodySection{{}},
	UID:           true,
	BodyStructure: &imap.FetchItemBodyStructure{Extended: true},
}

func GetMockMailConfig() *MailConfig {
	return &MailConfig{
		Host:     "imap.feishu.cn",
		Port:     "993",
		Username: "test@seekthought.com",
		Password: "o8HlzhbehavmGQ0z",
	}
}

type attachmentDetail struct {
	filename  string
	partsData *bytes.Buffer
}

type MemoryMail struct {
	SeqNum       uint32
	Flags        []imap.Flag
	Envelope     *imap.Envelope
	InternalDate time.Time
	RFC822Size   int64
	UID          imap.UID
	// BodySection   map[*imap.FetchItemBodySection][]byte
	Html              string
	PlainText         string
	Attachment        []*attachmentDetail
	BodyStructure     imap.BodyStructure
	BodySectionReader imap.LiteralReader
	BodySectionData   *imap.FetchItemBodySection
	// BinarySection     map[*imap.FetchItemBinarySection][]byte
	BinarySectionSize []imapclient.FetchItemDataBinarySectionSize
	ModSeq            uint64
}

func saveAttachment(mm *MemoryMail) error {
	pp.Println("开始存入附件......")
	for _, item := range mm.Attachment {
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(item.partsData)
		if err != nil {
			return err
		}
		if buf.Len() == 0 {
			pp.Println("No data read from item.partsData")
		}
		outFile, err := os.Create(item.filename)
		if err != nil {
			return err
		}
		defer outFile.Close()
		_, err = io.Copy(outFile, buf)
		if err != nil {
			return err
		}
		pp.Println(item.filename + " 保存成功！！！")
	}

	// reader := io.NopCloser(mm.Attachment.partsData)
	// defer reader.Close()

	return nil

}

func parseBodySection(mm *MemoryMail, item imapclient.FetchItemDataBodySection) error {
	mm.BodySectionReader = item.Literal
	mm.BodySectionData = item.Section
	mr, err := mail.CreateReader(item.Literal)
	if err != nil {
		pp.Println("err1 = ", err)
	}
	for {
		p, err := mr.NextPart()
		if err != nil {
			if errors.Is(err, io.EOF) || err.Error() == "multipart: NextPart: EOF" {
				return nil
			}
			pp.Println("err= ", err)
			return err
		}
		if p == nil {
			return nil
		}
		switch h := p.Header.(type) {
		case *mail.InlineHeader:
			mtype, _, err := h.ContentType()
			if err != nil {
				err = pp.Errorf("can not parse content type: %w", err)
				return err
			}
			pp.Println("mimeType = ", mtype)
			//Content-Disposition Content-Transfer-Encoding:
			switch mtype {
			case "text/html":
				html, err := io.ReadAll(p.Body)
				if err != nil {
					err = fmt.Errorf("can not read html part: %w", err)
					return err
				}
				mm.Html = string(html)
			case "text/plain":
				plain, err := io.ReadAll(p.Body)
				if err != nil {
					err = fmt.Errorf("can not read plain/text part: %w", err)
					return err
				}
				mm.PlainText = string(plain)
			default:
				pp.Println("we do not support read other content type", zap.String("content-type", mtype))
			}

		case *mail.AttachmentHeader:
			// contentType := p.Header.Get("Content-Type")
			// pp.Println("---------", contentType)
			filename, err := h.Filename()
			if err != nil {
				err = pp.Errorf("can not parse content type: %w", err)
				return err
			}
			pp.Println("filename = ", filename)

			// buf := new(bytes.Buffer)
			// _, err = buf.ReadFrom(p.Body)
			// if err != nil {
			// 	return err
			// }
			// outFile, err := os.Create(filename)
			// if err != nil {
			// 	return err
			// }
			// _, err = io.Copy(outFile, buf)
			// if err != nil {
			// 	return err
			// }
			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(p.Body)
			if err != nil {
				pp.Println("ee = ", err)
				return err
			}
			act := &attachmentDetail{
				filename:  filename,
				partsData: buf,
			}
			mm.Attachment = append(mm.Attachment, act)

		}
	}
}

func parseItemData(fetched imapclient.FetchItemData, memmail *MemoryMail) error {
	switch item := fetched.(type) {
	case imapclient.FetchItemDataBodySection:
		return parseBodySection(memmail, item)
	case imapclient.FetchItemDataFlags:
		memmail.Flags = item.Flags
	case imapclient.FetchItemDataEnvelope:
		memmail.Envelope = item.Envelope
	case imapclient.FetchItemDataInternalDate:
		memmail.InternalDate = item.Time
	case imapclient.FetchItemDataRFC822Size:
		memmail.RFC822Size = item.Size
	case imapclient.FetchItemDataUID:
		memmail.UID = item.UID
	case imapclient.FetchItemDataBodyStructure:
		memmail.BodyStructure = item.BodyStructure
	case imapclient.FetchItemDataBinarySectionSize:
		memmail.BinarySectionSize = append(memmail.BinarySectionSize, item)
	case imapclient.FetchItemDataModSeq:
		memmail.ModSeq = item.ModSeq
	default:
		panic(fmt.Errorf("unsupported fetch item data %T", item))
	}
	return nil
}

func SaveImap(msg *imapclient.FetchMessageData) error {

	if msg == nil {
		return errors.New("nil message can not be save")
	}
	mm := &MemoryMail{SeqNum: msg.SeqNum}
	for {
		item := msg.Next()
		if item == nil {
			pp.Println("all the message part is accessed")
			break
		}
		err := parseItemData(item, mm)
		if err != nil {
			pp.Printf("mail part can not be parsed: %w", err)
			continue
		}

	}
	pp.Println(mm.Envelope.Subject)

	// parentID, err := saveInlinePart(mm)
	// if err != nil {
	// 	err = fmt.Errorf("can not save inline part of mail: %w", err)
	// 	return err
	// }
	if len(mm.Attachment) != 0 {
		return saveAttachment(mm)
	}
	return nil
}

func handleMessage(msg *imapclient.FetchMessageData) error {
	if msg == nil {
		return nil
	}
	pp.Println("read one mail for saving...")
	err := SaveImap(msg)
	if err != nil {
		pp.Println("save mail body error")
		return err
	}
	return nil
}

func FetchOneBox(c *imapclient.Client, box string) {
	sc, err := c.Select(box, nil).Wait()
	if err != nil {
		pp.Println("选择邮箱出错！", err)
	}
	pp.Println("邮件封数为：", sc.NumMessages)
	uidset := imap.UIDSetNum()
	uidset.AddRange(1, 0)
	// uidset.AddNum(1, 2, 3)

	fc := c.Fetch(uidset, DefaultFetchOptions)
	// fmb, err := fc.Collect()
	// if err != nil {
	// 	pp.Println("Fetch错误", err)
	// }
	// for _, data := range fmb {
	// 	pp.Println(data.Envelope.Subject)

	// }
	for {
		msgData := fc.Next()
		if msgData == nil {
			pp.Println("fetch from one mailbox done")
			break
		}
		err := handleMessage(msgData)
		if err != nil {
			pp.Println("handle Message error", zap.Error(err))
		}
	}

}

func main() {
	mailconfig := GetMockMailConfig()
	options := &imapclient.Options{

		// DebugWriter:           os.Stdout,
		UnilateralDataHandler: &imapclient.UnilateralDataHandler{},
		WordDecoder:           &mime.WordDecoder{CharsetReader: charset.Reader},
	}
	c, err := imapclient.DialTLS(mailconfig.Host+":"+mailconfig.Port, options)

	if err != nil {
		pp.Println("连接错误！", err)
		panic(err)

	}
	if c.Login(mailconfig.Username, mailconfig.Password).Wait(); err != nil {
		pp.Println("登陆错误！", err)
		panic(err)
	}
	// boxes := []string{"INBOX", "测试1", "测试2"}
	boxes := []string{"INBOX"}

	for _, box := range boxes {
		pp.Println("开始获取：", box)
		FetchOneBox(c, box)
	}
	defer c.Close()

}
