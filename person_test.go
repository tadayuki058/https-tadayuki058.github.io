package person

import (
	"testing"
	"time"
)

// 全体のcovarage: 85.2%
// go test -run ''  でテストファイルをテスト
//go test -v .でログを表示させてテスト

// TestNewPerons テストを作成。
func TestNewPerons(t *testing.T) {

	// 佐藤太郎(男性,1996/10/11生まれ) を生成
	birthday := time.Date(1996, 10, 11, 0, 0, 0, 0, time.Local)
	sato, _ := NewPerson("佐藤", "太郎", "男", birthday)
	yamada, _ := NewPerson("", "", "女", birthday)

	// NewPerson関数でPerson構造体の中にある返り値が返ってきているのかどうかをテスト
	t.Log("setup") // <setup code>
	t.Run("NewPerson/佐藤/太郎/男/1996-10-11", func(t *testing.T) {
		if sato.FamilyName != "佐藤" {
			t.Fatal("error")
		} else if sato.FirstName != "太郎" {
			t.Fatal("error")
		} else if sato.Gender != "男" {
			t.Fatal("error")
		} else if sato.Birthday != time.Date(1996, 10, 11, 0, 0, 0, 0, time.Local) {
			t.Fatal("error")
		}
	})
	t.Run("NewPerson/ / /女/1996-10-11", func(t *testing.T) {
		if yamada.FamilyName != "" { //何も記載がされていない状態
			t.Fatal("error")
		} else if yamada.FirstName != "" {
			t.Fatal("error")
		} else if yamada.Gender != "女" {
			t.Fatal("error")
		} else if yamada.Birthday != time.Date(1996, 10, 11, 0, 0, 0, 0, time.Local) {
			t.Fatal("error")
		}
	})
	t.Log("tear-down") // <tear-down code>
}

func TestOld(t *testing.T) {

	// 佐藤太郎とこの人が17歳の日にちと18歳の日にちを定義
	sato, _ := NewPerson("佐藤", "太郎", "男", time.Date(1996, 10, 11, 0, 0, 0, 0, time.Local))
	SpecifiedDate1 := time.Date(2014, 10, 10, 0, 0, 0, 0, time.Local)
	SpecifiedDate2 := time.Date(2014, 10, 11, 0, 0, 0, 0, time.Local)

	// セットアップコード
	t.Log("setup") // <setup code>
	t.Run("sato-Old-2014/10/10", func(t *testing.T) {
		if sato.Old(SpecifiedDate1) != 17 {
			t.Fatal("error")
		}
	})
	t.Run("sato-Old-2014/10/11", func(t *testing.T) {
		if sato.Old(SpecifiedDate2) != 18 {
			t.Fatal("error")
		}
	})
	// ティアダウンコード
	t.Log("tear-down") // <tear-down code>

}

// TestMarriage テストを作成
func TestMarriage(t *testing.T) {

	// 佐藤、山田、鈴木と特定の日(SpecifiedDate)を宣言
	sato, _ := NewPerson("佐藤", "太郎", "男", time.Date(1996, 10, 11, 0, 0, 0, 0, time.Local))
	yamada, _ := NewPerson("山田", "花子", "女", time.Date(1994, 7, 30, 0, 0, 0, 0, time.Local))
	suzuki, _ := NewPerson("鈴木", "智子", "女", time.Date(1999, 10, 10, 0, 0, 0, 0, time.Local))
	SpecifiedDate := time.Date(2014, 10, 11, 0, 0, 0, 0, time.Local)

	t.Log("setup") // <setup code>
	// 男18歳以上、女16歳以上
	t.Run("Marriage-yamada-sato-2014/10/10", func(t *testing.T) {
		if Marriage(yamada, sato, SpecifiedDate) != true {
			t.Fatal("error")
		}
	})
	// 男18歳以上、女16歳未満
	t.Run("Marriage-sato-suzuki-2014/10/10", func(t *testing.T) {
		if Marriage(sato, suzuki, SpecifiedDate) != false {
			t.Fatal("error")
		}
	})
	// A.Gender == B.Gender
	t.Run("Marriage-yamada-suzuki-2014/10/10", func(t *testing.T) {
		if Marriage(yamada, suzuki, SpecifiedDate) != false {
			t.Fatal("error")
		}
	})
	t.Log("tear-down") // <tear-down code>

}
