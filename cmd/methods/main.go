package main

import (
	"fmt"
)

type Reservation struct {
	Date string
}

// 基本型に別名を付けることでメソッドを定義できる
type MyInteger int

func main() {
	reservation := Reservation{Date: "2026-07-12"}
	fmt.Printf("reservation.CheckDate(): %v\n", reservation.CheckDate())
	// ポインタ渡しのメソッドを呼ぶ時、x.method -> (&x).methodと解釈する
	fmt.Printf("reservation.CheckDatePointer(): %v\n", reservation.CheckDatePointer())
	var nilReservation *Reservation
	fmt.Printf("nilReservation.CheckDatePointer(): %v\n", nilReservation.CheckDatePointer())

	n := MyInteger(10)
	fmt.Printf("n.Add(15): %d\n", n.Add(15))
}

// funcの次にあるのはレシーバー
// メソッドを定義できるのは同じパッケージの型に対してのみ
func (reservation Reservation) CheckDate() string {
	return fmt.Sprintf("Reservation for %v", reservation.Date)
}

// レシーバーをポインタにすると構造体をコピーせずにメソッドが呼ばれる
// 値を変更する、構造体のコピーを避ける、Mutexなどコピーすべきでないメンバーを含んでいるなどの場合はポインタにする
// 同じ構造体のメソッドに値渡しとポインタ渡しが混在すると扱いにくいので、
// 粒度の小さい値などは値渡し、モデルのような構造体ならポインタ渡しにする
func (reservation *Reservation) CheckDatePointer() string {
	if reservation != nil {
		return fmt.Sprintf("Reservation for %v", reservation.Date)
	} else {
		// レシーバーがnilでもメソッドは実行される
		// ただし、インタフェースの参照に何も設定されていなければ型からメソッドが特定できないのでエラーになる
		return "No reservation"
	}
}

func (n MyInteger) Add(x int) MyInteger {
	return MyInteger(int(n) + x)
}
