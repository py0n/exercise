# 自分の各種演習用

* [Project Euler](https://projecteuler.net/)
* [言語100本ノック](http://www.cl.ecei.tohoku.ac.jp/nlp100/) (予定)
* [A Tour of Go](https://go-tour-jp.appspot.com/list)

## A Tour of Go の覺書

[A Tour of Go (日本語版)](http://go-tour-jp.appspot.com/)を讀んだ際の覺書。

### Chapter 3 - Go Playground

Goのプログラムをサーバで受取り、コンパイル、リンク、實行して結果を返す仕組み。

* [The Go Playground](http://play.golang.org/)

以下の制限が有る。

* 標準ライブラリが使用出來る。
  * ネットワークアクセスは出來ない。
  * ファイルアクセスも出來ない。
* 日時は常に *2009-11-10 23:00:00 UTC* で固定される。
* マルチスレッドは使へない。

### chapter 4 - Package

* パッケージ名はインポートパスの最後の要素と同じ。
  * `math/rand` パッケージは `package rand` で始まるファイル群で構成される。
  * `code.google.com/p/go.net/websocket` なら `package websocket` で始まるファイル群で構成される.

### Chapter 6 - Exported names

* 頭文字が大文字の名前は *外部へ公開* される。
  * 小文字の名前は公開されない。

### Chapter 13 - Short variable declarations

* `var` 宣言は凾數内外で使用出來る。
* `:=`  は凾數内でしか使用出來ない。

```go
var x,  y float64 = 1.2,  1.4
i, j := 3, 4
```

### Chapter 15 - Constants

* 定數は `const` キーワードを使用して定義する。
* 定數に出來るのは以下の四つ。
  * character
  * string
  * boolean
  * numeric (數値)

### Chapter 29 - The new function

* `new(T)` はTを *ゼロ初期化* してそのポインタを返す。

```go
vat t *T = new(T)
t := new(T)
```

### Chapter 32 - Making slices

* *slice* は `make` 凾數で生成する。

```go
a := make([]int, 4, 5)
// len(a) = 4
// len(a) = 5
```


### Chapter 33 - Nil slice

* *slice* の初期値は `nil` となる。
  * `nil` の *slice* は *length* と *cpacity* が0。

### Chapter 36 - Exercise: Slices

* 二重 *slice* : `m := maake([][]int, 5, 5)` で *length* と *capacity* が5の `[]int` の*slice* が生成される。

### Chapter 37 - Maps

* *map* は `make` 凾數で生成する。
* `string` と `int` のマップの例。

```go
m := make(map[string]int)
m["a"] = 1
m["b"] = 2
```

### Chapter 40 - Mutating Maps

* `m[key] = value` -- `map` への要素の插入・更新。
* `value = m[key]` -- 要素の抽出。
* `delete(m, key)` -- 要素の削除。
* `value, ok = m[key]` -- キーの存在確認。要素が存在すれば `ok` は `ture` になる。

### Chapter 43 - Function closures

* *Go* の凾數はクロージャ( *closure* )。

### Chapter 45 - Swith

* *Go* の `swith` には `break` がない。
* 次の `case` に通したい場合は `fallthrough` 文を `case` の末尾に記述する。

### Chapter 47 - Switch with no condition

* 條件のない `switch` は 長い `if-then-else` の代はりになる。

```go
t := time.Now()
swith {
case t.Hour() < 12:
    fmt.Println("朝だよ")
case t.Hour() < 17:
    fmt.Println("晝だよ")
default:
    fmt.Println("夜だよ")
}
```

### Chapter 53 - Interface

* インターフェース型( *interface type* )はメソッド群で定義される。

### Chapter 55 - Errors

* 或る型 *T* に `func (t T) Error() string` と云ふメソッドが實裝されてゐれば、其乃型は *interface type* `error` を滿たす。


### Chapter 61 - Exercise: Rot13 Reader

* `io.Reader` は以下のやうなインターフェース型。

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

* `Read` は讀込んだバイト數を返す。
  * `err == nil` の時は正常。
  * `err != nil` の時は何らかのエラー。
    * `err == EOF` の時はファイルの末尾。
* *guncip* みたいなフィルタのサンプル。

```go
package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	reader, _ := gzip.NewReader(os.Stdin)
	io.Copy(os.Stdout, reader)
}
```

### Chapter 64 - Channels

* チャネル( *channel* )は `make` で生成する。

```go
ch := make(chan int)
ch <- v   // vをチャネルchへ送る。
v := <-ch // chから受信して變數vへ割當てる。
```

* デフォルトでは送受信は *ブロック* される。

### Chapter 65 - Buffered Channels

* チャネルはバッファ出來る。
* バッファの長さは `make` の第二引數で指定する。
* バッファの長さは `cap` で取得する。

```go
ch := make(chan int, 100)
//cap(ch) == 100
```

* バッファが滿杯の状態で更に送信を行ふとエラーになる。

```
fatal error: all goroutines are asleep - deadlock!
```

### Chapter 65 - Range and Close

* 受信式に第二變數を與へる事でチャネルが閉じてゐるか確認可能。

```go
v, ok := <- ch
// ok == true  // 閉じてゐない
// ok == false // 閉じてゐる
```

* `range` を使つてチャネルから繰返し受信する事が出來る。
* `range` でインデックスを取得する事は出來ない。

```go
for v := range c {
    // do something
}
```

### Chapter 67 - Select

* `select` は `goroutine` を複數の通信操作で待たせる。
* 複數の `case` が一致した場合、ランダムに一つが擇ばれる。

### Chapter 68 - Default Selection

* `select` でどの `case` にも一致しない場合、 `default` が實行される。
