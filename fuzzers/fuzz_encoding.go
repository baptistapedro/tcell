package myfuzz

import (
	"github.com/gdamore/tcell"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func Fuzz(data []byte) int {
	tcell.RegisterEncoding("GBK", simplifiedchinese.GBK);
	enc := tcell.GetEncoding("GBK");
	enc.NewDecoder().Bytes(data)
	return 0
}
