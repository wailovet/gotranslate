package gotranslate

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/wailovet/webdriver"
)

type Translate struct {
	wd *webdriver.WebDriver
}

func NewTranslate() *Translate {
	return &Translate{}
}
func jsonEncode(v interface{}) string {
	ret, _ := json.Marshal(v)
	return string(ret)
}

func (t *Translate) SetWebdriver(wd *webdriver.WebDriver) {
	t.wd = wd
}

func (t *Translate) Translate(text string, from string, to string) (string, error) {
	// url := fmt.Sprintf("https://cn.bing.com/translator?ref=TThis&from=%s&to=%s&text=%s", url.QueryEscape(from), url.QueryEscape(to), url.QueryEscape(text))
	// t.wd.SetUrl(url)

	href, _ := t.wd.ExecuteScript(`return location.href;`)
	if !strings.HasPrefix(href.String(), "https://cn.bing.com/translator") {
		url := fmt.Sprintf("https://cn.bing.com/translator?ref=TThis")
		t.wd.SetUrl(url)
	}

	jsSrc := fmt.Sprintf(` 
		var tta_srcsl = document.querySelector('#tta_srcsl');
		tta_srcsl.value = %s; 

		var tta_tgtsl = document.querySelector('#tta_tgtsl');
		tta_tgtsl.value = %s; 

		var input = document.querySelector('#tta_input_ta');
		input.value = '';
		input.click()
		await sleep(100);
		input.value = %s;
		input.click()

		var value = ""
		for (let i = 0; i < 5; i++) {
			value = document.querySelector('#tta_output_ta').value;
			value = value.trim()
			console.log("value:",value);
			if (value=="...") {
				await sleep(1000);
				continue;
			}
			if (value&&value.length>0) {
				return {
					data:value,
					error:""
				};
			}
			await sleep(1000);
		} 
		return {
			data:value,
			error:"翻译失败"
		};
	`, jsonEncode(from), jsonEncode(to), jsonEncode(text))

	// log.Println("jsSrc:", jsSrc)
	value, err := t.wd.ExecuteAwaitScript(jsSrc)
	if err != nil {
		return "", err
	}
	if value.Get("error").String() != "" {
		return "", fmt.Errorf(value.Get("error").String())
	}
	return value.Get("data").String(), nil
}
