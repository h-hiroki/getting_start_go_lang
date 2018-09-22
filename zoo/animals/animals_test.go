package animals

import "testing"

// TODO 引数の意味調べる。
// Tは、テスト関数に渡される型で、テストの状態を管理し、テストのログの書式化をサポートします。
// Logは実行中に蓄積され、完了時に標準エラー出力に吐かれます。
// 参考：http://golang.jp/pkg/testing
func TestElephantFeed(t *testing.T) {
	expect := "Grass"        // 期待する結果
	actual := ElephantFeed() // 実行するfunc

	if expect != actual { // funcの実行結果と期待する結果に差異がある場合の処理
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestMonkeyFeed(t *testing.T) {
	expect := "Banana"
	actual := MonkeyFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestRabbitFeed(t *testing.T) {
	expect := "Carrot"
	actual := RabbitFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
