package main

import (
	"fmt"
	"mime"
	"sync"
	"testing"

	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/emersion/go-message/charset"
	"github.com/k0kubun/pp"
)

func nilOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func getTestMailClient() *imapclient.Client {
	var (
		err                error
		loadTestClientOnce sync.Once
		testIMAPCli        *imapclient.Client
	)
	config := GetMockMailConfig()
	loadTestClientOnce.Do(func() {
		options := &imapclient.Options{

			// DebugWriter:           os.Stdout,
			UnilateralDataHandler: &imapclient.UnilateralDataHandler{},
			WordDecoder:           &mime.WordDecoder{CharsetReader: charset.Reader},
		}
		testIMAPCli, err = imapclient.DialTLS(config.Host+":"+config.Port, options)
		if err != nil {
			err = fmt.Errorf("can not get test mail client: %w", err)
		}
		nilOrPanic(err)
		if err = testIMAPCli.Login(config.Username, config.Password).Wait(); err != nil {
			nilOrPanic(err)
		}

	})
	return testIMAPCli
}

func TestFetchOneBox(t *testing.T) {
	c := getTestMailClient()
	FetchOneBox(c, "INBOX")
	boxes := []string{"INBOX"}
	for _, box := range boxes {
		pp.Println("开始获取：", box)
		FetchOneBox(c, box)
	}
}
