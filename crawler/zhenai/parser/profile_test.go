package parser

import (
	"io/ioutil"
	"testing"

	"github.com/crawler/engine"
	"github.com/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://album.zhenai.com/u/1826404337", "vivian")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1826404337",
		Type: "zhenai",
		Id:   "1826404337",
		Payload: model.Profile{
			Age:        26,
			Height:     168,
			Weight:     53,
			Income:     "3000元以下",
			Gender:     "女",
			Name:       "vivian",
			Xinzuo:     "魔羯座",
			Occupation: "咨询/顾问",
			Marriage:   "未婚",
			House:      "租房",
			Hokou:      "山东济南",
			Education:  "硕士",
			Car:        "未购车",
		},
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
