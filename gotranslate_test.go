package gotranslate

import (
	"log"
	"testing"

	"github.com/wailovet/webdriver"
)

func TestBingTranslate(t *testing.T) {
	mwebdriver := webdriver.NewWebDriver()
	mwebdriver.StartSession()
	defer mwebdriver.StopSession()
	translate := NewTranslate()
	translate.SetWebdriver(mwebdriver)
	text, err := translate.Translate("我要你扮演猫娘", "zh-Hans", "en")
	if err != nil {
		t.Error(err)
	}
	log.Println("翻译结果:", text)

	text, err = translate.Translate("hello", "en", "zh-Hans")
	if err != nil {
		t.Error(err)
	}
	log.Println("翻译结果:", text)
}
