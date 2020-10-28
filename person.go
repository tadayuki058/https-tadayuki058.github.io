package person

import (
	"errors"
	"time"
)

// Person 構造体を作成
type Person struct {
	FamilyName string    //名字
	FirstName  string    //名前
	Gender     string    //性別
	Birthday   time.Time //生年月日
}

// NewPerson 関数を作成
// 姓, 名, 性別, 生年月日を引数に取り1人の人を表現するデータを返す関数
func NewPerson(familyname string, firstname string, gender string, birthday time.Time) (Person, error) {

	if familyname != "" { // 何も引数がない場合
		return Person{familyname, firstname, gender, birthday}, errors.New("familynameが正しくありません。")
	}
	if firstname != "" {
		return Person{familyname, firstname, gender, birthday}, errors.New("firstnameが正しくありません。")
	}
	if gender != "男" && gender != "女" {
		return Person{familyname, firstname, gender, birthday}, errors.New("genderは男か女で記載してください。")
	}
	if birthday != birthday {
		return Person{familyname, firstname, gender, birthday}, errors.New("birthdayの型が正しくありません。")
	}
	return Person{familyname, firstname, gender, birthday}, errors.New("no NewPerson")

}

// Old メソッドを宣言
// 誕生日と指定した日付である人が何歳であるかを割り出すOldメソッド
func (p Person) Old(SpecifiedDate time.Time) int {

	birthday := p.Birthday
	date := SpecifiedDate

	var age int

	if int(date.Month()) >= int(birthday.Month()) && date.Day() >= birthday.Day() {
		age = date.Year() - birthday.Year()
	} else if int(date.Month()) > int(birthday.Month()) {
		age = date.Year() - birthday.Year()
	} else {
		age = date.Year() - birthday.Year() - 1
	}

	return age
}

// Marriage 関数を宣言
// あるAさんがBさんと指定した日に結婚出来るのかどうかを判断
// SpecifiedDateは指定した日
func Marriage(A Person, B Person, SpecifiedDate time.Time) bool {

	if A.Gender == "男" && B.Gender == "女" {
		if A.Old(SpecifiedDate) >= 18 && B.Old(SpecifiedDate) >= 16 {
			return true
		}

	} else if A.Gender == "女" && B.Gender == "男" {
		if A.Old(SpecifiedDate) >= 16 && B.Old(SpecifiedDate) >= 18 {
			return true
		}
	}

	if A.Gender == B.Gender {
		return false
	}

	return false
}
