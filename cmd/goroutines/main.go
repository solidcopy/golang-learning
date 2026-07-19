package main

import (
	"fmt"
	"sync"
	"time"
)

// Goルーチンは同じプロセスで実行されるのでヒープ、ファイルやソケットなどは共有される
// ルーチンを実行するスレッドは複数あるが1つずつではない
// ランタイムがスレッド数を管理していてそれぞれのルーチンをどのスレッドで実行するか割り振る

// 関数リテラルはクロージャーにできるが、実行されるタイミングで変数の状態が変わっている可能性があるので注意

func main() {
	join()
	loop()
	select_stmt()
}

func join() {
	// チャンネルを使って完了を待つ
	done := make(chan bool)
	go output(done, "Hello")
	go output(done, "World")
	<-done
	<-done

	// WaitGroupを使って完了を待つ
	var wg sync.WaitGroup
	wg.Go(func() {
		fmt.Println("By function literal-1")
	})
	wg.Go(func() {
		fmt.Println("By function literal-2")
	})
	wg.Wait()
}

// 単にchanとしてもいいが、<-chan（受信）かchan<-（送信）に限定した方が誤用が防げる
// chanは参照なのでポインタで渡す必要はない
func output(done chan<- bool, s string) {
	fmt.Println(s)
	done <- true
}

func magicNumber(ch chan<- int) {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
}

func loop() {
	// 受信をループ
	ch := make(chan int, 2) // バッファサイズを指定
	go magicNumber(ch)
	// 送信側でcloseされるとループ終了
	for n := range ch {
		fmt.Printf("n: %d\n", n)
	}
}

func select_stmt() {
	ch1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "from ch1"
	}()

	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "from ch2"
	}()

	// 最初に受信した値のみ処理する
	// 受信するまで待つ
	// default:があると待たない
	// どちらも処理可能な場合はランダムに選ぶ
	select {
	case s := <-ch1:
		fmt.Printf("message %v\n", s)
	case s := <-ch2:
		fmt.Printf("output %v\n", s)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	}
}
