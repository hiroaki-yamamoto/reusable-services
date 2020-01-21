package random

import "strings"
import "math/rand"

func genTxt(size int, baseTxts ...string) (txt string, err error) {
	var builder strings.Builder
	if len(baseTxts) > 0 {
		for _, base := range baseTxts {
			if _, err = builder.WriteString(base); err != nil {
				return
			}
		}
	} else {
		if _, err = builder.WriteString(
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		); err != nil {
			return
		}
	}
	mapTxt := builder.String()
	var resBuilder strings.Builder
	for i := 0; i < size; i++ {
		resBuilder.WriteByte(mapTxt[rand.Intn(len(mapTxt))])
	}
	txt = resBuilder.String()
	return
}
