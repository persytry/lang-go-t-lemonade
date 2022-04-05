package server

import (
	"github.com/atotto/clipboard"
	"github.com/lemonade-command/lemonade/lemon"
)

type Clipboard struct{}

var dummyClipboard string

func (_ *Clipboard) Copy(text string, _ *struct{}) error {
	<-connCh
	// Logger instance needs to be passed here somehow?
    err := clipboard.WriteAll(lemon.ConvertLineEnding(text, LineEndingOpt))
	if err == nil {
		return nil
	}
    dummyClipboard = text
    return nil;
}

func (_ *Clipboard) Paste(_ struct{}, resp *string) error {
	<-connCh
	t, err := clipboard.ReadAll()
    if err == nil{
        *resp = t
        return err
    }
    *resp = dummyClipboard
    return nil;
}
