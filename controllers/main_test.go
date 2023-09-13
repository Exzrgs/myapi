package controllers_test

import (
	"testing"

	"github.com/Exzrgs/myapi/controllers"
	"github.com/Exzrgs/myapi/controllers/testdata"

	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMoc()

	// 変数名は短いほうがいい
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
