package gotranslate

import (
	"log"
	"testing"
)

func TestBingTranslate(t *testing.T) {
	translate := NewTranslate()
	defer translate.Close()

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
