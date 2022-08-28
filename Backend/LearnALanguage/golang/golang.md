# golang

\========================================

golangを使うときの自分用メモ

## install

- anyenvからサクッとインストール

  ```bash
  anyenv install goenv
  goenv install -l
  goenv install xxx
  ```

- [チュートリアル](https://go-tour-jp.appspot.com/list/)

## memo

- \<\< >>
  - bit演算子
  - 1\<\<2だったら 100 になる(左に2bitシフトする)
- uint
  - 符号付数字
- constants
  - := にて宣言不可
  - 文字(character)、文字列(string)、boolean、数値(numeric)のみ
- Pow
  - 乗算
    - Pow(3, 3) -> 3*3*3(3の3乗)
- Case
  - case だけを実行してそれに続く全ての case は実行されない。
    - break は自動的に提供される。
  - Go の switch の case は定数である必要はない
  - 関係する値は整数である必要はない
- defer
  - defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させるもの
  - defer へ渡した関数の引数は、すぐに評価されるが、その関数自体は呼び出し元の関数がreturnするまで実行されない
  - defer へ渡した関数は LIFO(last-in-first-out) の順番で実行される
- pointer
  - ポインタは値のメモリアドレス
  - 変数 chome の値が 777 の場合
    - `p :~ &chom4`
      - `&` を使う
      - 変数pにchome変数の値のアドレスを格納
    - `fmt.Println(*p)`
      - `*` を使う
      - 変数pが指すメモリアドレスに記録された値(ここではchomeの値)を表示する
    - `*p = *p*2`
      - 変数pが指すメモリアドレスに記録された値(ここではchomeの値)を二倍する
- E表記
  - 巨大な数を表すのに利用する
  - `1e9` であれば 1 × 10 の9乗
    - E表記の基数は通常10
    - int型として扱える
- array vs slices
  - 配列は固定長

  - スライスは可変長

    - スライスは配列への参照のようなもの

      - スライスの要素を変更すると、その元となる配列の対応する要素が変更される

    - lowとhighはどちらも添字が基準

    - low ≤ i \< high の範囲が切り取られる

    - 上限または下限を省略可能

      - 既定値は下限が 0 で上限はスライスの長さ

      - 下記スライス式は等価

        ```txt
        a[0:10]
        a[:10]
        a[0:]
        a[:]
        ```

    - スライスは長さ( length )と容量( capacity )の両方を持っている

    - スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素数

  - range

    - 2つの変数を返す
    - 1つ目の変数はインデックス( index )
    - 2つ目はインデックスの場所の要素のコピー

  - インデックスや値を捨てる

    - " _ "(アンダーバー) へ代入

  - map

    - キーと値とを関連付ける

  - 変数のように関数を渡すことができる

  - Pow

    - べき乗
      ― closure
    - 関数のスコープを超えて変数を共有できる
      - closure内の変数が静的に保持される
    - 関数型プログラミングでよく使われる

  - レシーバ

    - golangにはクラスはない

    - 型(struct)にメソッドを定義出来る

    - この際に、レシーバ引数というものを下記がイメージ

      ```txt
      まずは構造体定義
      type CHOMETARO struct {X,Y int}

      メソッドを定義する際funcのすぐ後に(c CHOMETARO)というレシーバ関数を定義
      func (c CHOMETARO) addNums() int{
        return c.X + c.Y
      }

      あとは構造体の変数を作って、当該変数からメソッドを呼び出すだけ
      func main() {
        c := CHOMETARO{
          3,
          5,
        }
        fmt.Println(c.adNums())
      }
      ```

    - structの型だけではなく、任意の型(type)にもメソッドを宣言できる

      - イメージ

        ```txt
        type ChomeString string

        func (c ChomeString) SayYes() string {
          msg := stirng(c) + " " + Yes!!
          return string
        }

        func main() {
          c := ChomeString("ChageAsu Say??")
          fmt.Println(c.SayYes())
        }
        ```

    - ポインタレシーバのメリット

      - 呼び出し時に、変数、または、ポインタのいずれかを取れる(変数を暗黙的にポイントへ変換)
        - 通常関数にて引数をポイントとした場合はポインタでないとコンパイルエラー

    - 通常関数の場合、同じ名前を利用できずメソッド名が長くなりがち

    - メソッドの場合は、別レシーバ名があるので、同じ名前を付けられる

  - goroutine

    - Goのランタイムに管理される軽量なスレッド
    - `go f(x, y, z)` で呼び出せる
      - fは関数
    - goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要がある

  - scan

    - 単純系

      ```go
      var input string
      fmt.Scan(&変数)
      ```

    - 配列/スライスに格納したい

      ```go
      inputS := make([]string,要素数)
      for i := 0; i < 要素数; i ++ {
        fmt.Scan(&変数[i])
      }
      ```

    - Scan, Scanf, Scanln はos.Stdin から読み取る。

    - Fscan, Fscanf, Fscanln は，指定された io.Reader から読み取る。

    - Sscan, Sscanf, Sscanln は引数文字列から読み取る

    - Scan, Fscan, Sscan は入力の改行をスペースとして扱う。

    - Scanln, Fscanln, Sscanln は改行でスキャンを停止する。

      - アイテムの後に改行または EOF が続く必要がある。

    - Scanf, Fscanf, Sscanf は，フォーマット文字列に従って引数を解析する。

  - シングルクォーテとダブルクォーテーション

    - シングルクォーテはrune型というものになる
    - ダブルクォーテーションで文字列を定義すること

  - 階乗(factorial)のライブラリがない

    - 再帰で作る

  - strings.Join(配列,結合文字列)

    - 配列の文字を結合する

  - 空スライスの作成方法

    - スライス名 := \[\]型名{}
      - `emptyS := []string{}`

  - sort

    - `import sort`
      - `sort.Ints` で数値並べ替え昇順

  - 文字/数字変換

    - "strconv"
      - 文字を数値にする `i,_ = strconv.Atoi(s)`
      - 数値を文字にする `s = strconv.Itoa(i)`

  - mapのValueに構造体や配列を定義する方法

    ```text
    type chome struct{
      grade []map[string]int
      memo  []string
    }

    test := chome{
      grade: map[string]int{"ドラゴン級": 1,"ライオン級": 2},
      memo :{"メモ1","メモ2",}
    }
    ```

## tips

### 文字列中の文字の判定

#### 前提

- Sがstring
- Nはint

#### メモ

- N番目の文字取得したい
- S\[N-1:N\]

```go
if S[N-1:N] == "🎸" {
  fmt.Println("Yes")
} else {
  fmt.Println("No")
}
```

## 文字列分割　-> 配列へ格納

```go
xy := strings.Split(s, ".")
```

## 数値を文字に変換(返り値は1個)

```go
strconv.Itoa(i + 1)
```

## 文字列を数値にう変換(返り値は2個)

```go
x, _ := strconv.Atoi(xy[1])
```

## 配列の並べかえ

```go
sort.Strings(arr)
```

## 重複を見つける(mapを作りforで要素を追加していき、要素が被るかを判定)

```go
m := make(map[string]bool)

for _, ele := range fullname {
  if !m[ele] {
    m[ele] = true
  } else {
    fmt.Println("Yes")
    return
  }
}
```

## 行単位で読み込む

- スペースがあるとsliceに入れる際に一つの要素としてCountされる

- 何かしらの文字列で結合して一文字列にする

  ```golang
  var s = bufio.NewScanner(os.Stdin)

  for i := 0; i < N; i++ {
    s.Scan()
    tmp := []string{}
    tmp = append(tmp, s.Text())
    numStr := strings.Join(tmp, " ")
    numStr = strings.ReplaceAll(numStr, " ", "-")
    L[i] = numStr
    }
  ```

## byte型の使い方

- 文字列及び数値を1文字ずつ処理するのに便利
  - 文字列の場合、通常でいいえばstring\[0:1\]のように範囲指定が必要
  - byte型の場合はbyte\[0\]のように一発で指定可能

## 最小値と最大値

- sliceはsort.IntSlice(slice) -> Min:slice\[0\] Max:slice\[len(slice)-1\]
- 二値はmath.Max/Min
