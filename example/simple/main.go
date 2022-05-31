package main

import (
	"github.com/WilliamNHarvey/go-qrcode/writer/standard"

	"github.com/WilliamNHarvey/go-qrcode/v2"
)

func main() {
	qrc, err := qrcode.NewWith("github.com/WilliamNHarvey/go-qrcode",
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionQuart),
	)
	if err != nil {
		panic(err)
	}

	w, err := standard.New("./simple.png", standard.WithQRWidth(40))
	if err != nil {
		panic(err)
	}

	if err = qrc.Save(w); err != nil {
		panic(err)
	}
}
