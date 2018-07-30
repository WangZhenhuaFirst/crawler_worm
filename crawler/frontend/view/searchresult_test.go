package frontend

import (
	"os"
	"testing"

	"github.com/crawler/engine"
	"github.com/crawler/frontend/model"
	common "github.com/crawler/model"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1826404337",
		Type: "zhenai",
		Id:   "1826404337",
		Payload: common.Profile{
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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}

}
