package sd

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func initTest() {
	os.Chdir("../../../")
	godotenv.Load(".env")
}

func TestGoodsviewCRUDSearchForGoodsViewByName(t *testing.T) {
	initTest()
	ctx := context.Background()

	m := new(GoodsviewCRUD)
	data, err := m.SearchForGoodsViewByName(ctx, "罗城")
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fail()
	}
	for _, d := range data {
		t.Logf("%+v", d)
	}
}
