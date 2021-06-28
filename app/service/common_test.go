package service

import (
	"evaluate_backend/app"
	"testing"
)

func Init() {
	app.Init()
}

func TestCreateQrCode(t *testing.T) {
	Init()
	_, _ = CreateQrCodeSrv("")
}
