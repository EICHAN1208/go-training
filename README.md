
# 練習問題

## 1. スライスの合計を求める

整数スライス `[]int` を受け取り、合計を返す関数を書いてください。

```go
func Sum(nums []int) int
```

### 例

```go
fmt.Println(Sum([]int{1, 2, 3, 4}))
```

出力:

```go
10
```

### 鍛えたいこと

最初は `for range` に慣れるのが大事です。
Goではコレクション処理をまず素直なループで書けることが土台になります。ここで `range` の書き方と、初期値を持つ変数更新に慣れておくと次につながります。


## 2. 最大値を返す

整数スライスを受け取り、最大値を返す関数を書いてください。
ただし、空スライスのときは `0` を返すことにします。

```go
func Max(nums []int) int
```

### 例

```go
fmt.Println(Max([]int{4, 9, 2, 7}))
fmt.Println(Max([]int{}))
```

出力:

```go
9
0
```

### 鍛えたいこと

条件分岐と、空データをどう扱うかです。
Goでは「入力が空のときどうするか」を早めに決めておく癖がかなり大事です。この考え方は後で `error` にもつながります。


## 3. 文字列中の文字数を数える

文字列を受け取り、各文字の出現回数を `map[string]int` で返してください。

```go
func CountLetters(s string) map[string]int
```

### 例

```go
fmt.Println(CountLetters("banana"))
```

イメージ:

```go
map[a:3 b:1 n:2]
```

### 鍛えたいこと

`map` の基本操作です。
Goでは「集計」はかなり頻出なので、`m[key]++` を自然に書けるようになると強いです。あわせて、文字列を1文字ずつ見る感覚も身につきます。


## 4. 偶数だけを取り出す

整数スライスを受け取り、偶数だけを入れた新しいスライスを返してください。

```go
func FilterEven(nums []int) []int
```

### 例

```go
fmt.Println(FilterEven([]int{1, 2, 3, 4, 5, 6}))
```

出力:

```go
[2 4 6]
```

### 鍛えたいこと

`append` とスライス操作です。
Goでは配列変換系の処理を自分で書くことが多いので、`append` を使って結果を育てる書き方に慣れる価値があります。


## 5. struct とポインタレシーバ

次の `Person` に、年齢を1増やす `Birthday()` メソッドを実装してください。

```go
type Person struct {
    Name string
    Age  int
}

func (p *Person) Birthday()
```

### 例

```go
p := Person{Name: "Taro", Age: 20}
p.Birthday()
fmt.Println(p.Age)
```

出力:

```go
21
```

### 鍛えたいこと

値レシーバとポインタレシーバの違いです。
レシーバ自身を更新したいなら、なぜポインタが必要なのかを実感できる問題です。Goのメソッド理解の最重要ポイントの1つです。


## 6. 独自の String() を実装する

`Person` に `String()` メソッドを実装し、`fmt.Println(p)` したときに見やすい文字列が出るようにしてください。

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func (p Person) String() string
```

### 例

```go
p := Person{FirstName: "Taro", LastName: "Yamada", Age: 20}
fmt.Println(p)
```

出力例:

```go
Yamada Taro (20)
```

### 鍛えたいこと

interface を「使う側」ではなく「満たす側」から理解することです。
`fmt` が `String()` を見つけて自動で使う、というGoらしい仕組みに触れられます。ここで「interfaceは明示的に宣言しなくても満たせる」という感覚が育ちます。


## 7. エラーを返す関数を書く

割り算をする関数を作ってください。
ただし、0で割ろうとしたら `error` を返してください。

```go
func Divide(a, b int) (int, error)
```

### 例

```go
v, err := Divide(10, 2)
fmt.Println(v, err)

v, err = Divide(10, 0)
fmt.Println(v, err)
```

出力例:

```go
5 <nil>
0 cannot divide by zero
```

### 鍛えたいこと

Goのエラーハンドリングです。
例外ではなく、戻り値でエラーを返すというGoの基本姿勢を練習できます。ここを避けて通ると、実務コードが一気に読みにくくなります。


## 8. io.Reader を使って文字数を数える

`string` ではなく、`io.Reader` を受け取り、そこから読んだ英字の出現回数を数える関数を書いてください。

```go
func CountFromReader(r io.Reader) (map[string]int, error)
```

### 条件

* `Read()` を使って読む
* 英字だけ数える
* 読み終わったら結果を返す
* `io.EOF` は正常終了として扱う

### 鍛えたいこと

`io.Reader`、`Read`、`EOF` の理解です。
これはGoのI/Oの入口です。ファイル、HTTPレスポンス、文字列入力などが同じ形で扱える理由が見えてきます。ここまで来ると、かなりGoらしいコードに触れたことになります。


## 9. クロージャでカウンタを作る

呼び出すたびに 1, 2, 3... と増えていく関数を返してください。

```go
func NewCounter() func() int
```

### 例

```go
counter := NewCounter()
fmt.Println(counter())
fmt.Println(counter())
fmt.Println(counter())
```

出力:

```go
1
2
3
```

### 鍛えたいこと

クロージャが外側の変数を保持する感覚です。
Goでは関数も値なので、「関数を返す」「状態を閉じ込める」ができます。関数型っぽい要素ですが、実務でも意外と使います。


## 10. goroutine と channel で並行処理する

1から5までの整数を goroutine 側で `channel` に送り、main側で受け取って表示するプログラムを書いてください。

### 条件

* goroutine を1つ使う
* `chan int` を使う
* 送り終わったら `close(ch)` する
* 受け取る側は `for v := range ch` で読む

### 例のイメージ

```go
1
2
3
4
5
```

### 鍛えたいこと

Goの並行処理の入口です。
`goroutine` は「軽い並行実行」、`channel` は「値の受け渡し」です。この組み合わせはGoの大きな特徴なので、最低1回は自分で書いて動かすと理解が進みます。


## 11. interface を自分で定義する

以下の `Shape` interface を定義し、`Circle` と `Rectangle` にそれぞれ `Area()` を実装してください。
そして、`[]Shape` を受け取り、合計面積を返す関数を書いてください。

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width, Height float64
}

func TotalArea(shapes []Shape) float64
```

### 例

```go
shapes := []Shape{
    Circle{Radius: 3},
    Rectangle{Width: 4, Height: 5},
}
fmt.Println(TotalArea(shapes))
```

出力例（円周率は `math.Pi` を使用）:

```
48.274333882308138
```

### 鍛えたいこと

interface を「定義する側」から実践することです。
6問目では `fmt.Stringer` という既存の interface を満たしましたが、今回は自分で interface を作ります。「異なる型を同じように扱う」というGoのポリモーフィズムの核心を体感できます。


## 12. カスタムエラー型を作る

ファイル操作を想定した関数を書いてください。
エラーが発生したときに、ファイル名と理由を含むカスタムエラー型を返すようにします。

```go
type FileError struct {
    FileName string
    Reason   string
}

func (e *FileError) Error() string

func ReadConfig(filename string) (string, error)
```

### 条件

* `filename` が空文字のとき、`FileError{FileName: "", Reason: "filename is empty"}` を返す
* `filename` が `"config.json"` のとき、`"debug=true"` という文字列を返す
* それ以外のとき、`FileError{FileName: filename, Reason: "file not found"}` を返す

### 例

```go
v, err := ReadConfig("config.json")
fmt.Println(v, err)

_, err = ReadConfig("missing.txt")
var fe *FileError
if errors.As(err, &fe) {
    fmt.Println("ファイル名:", fe.FileName)
    fmt.Println("理由:", fe.Reason)
}
```

出力例:

```
debug=true <nil>
ファイル名: missing.txt
理由: file not found
```

### 鍛えたいこと

`error` interface を自分で実装し、`errors.As` で型を取り出す感覚です。
標準の `errors.New` では「何が起きたか」しか伝えられませんが、カスタムエラーにすると「どのファイルで」「なぜ」起きたかを構造化して持てます。実務では必須のパターンです。


## 13. sync.WaitGroup で複数 goroutine を待つ

整数スライスを受け取り、各要素を別々の goroutine で処理して、結果をスライスで返す関数を書いてください。
各要素に対して「2倍にする」処理を並行で行います。

```go
func DoubleAll(nums []int) []int
```

### 条件

* 各要素を別々の goroutine で処理する
* `sync.WaitGroup` を使って全 goroutine の完了を待つ
* 結果のスライスは元のスライスと同じ順序で返す（インデックスを保持する）

### 例

```go
result := DoubleAll([]int{1, 2, 3, 4, 5})
fmt.Println(result)
```

出力:

```
[2 4 6 8 10]
```

### 鍛えたいこと

`sync.WaitGroup` による goroutine の同期です。
「goroutine を起動したが、全部終わるまで待ちたい」という場面は非常によく出てきます。また、複数 goroutine から同じスライスに書き込む場合のインデックス管理も重要なポイントです。


## 14. select でタイムアウトを実装する

値を返すのに時間がかかる処理を goroutine で実行し、指定した時間内に完了しなければタイムアウトエラーを返す関数を書いてください。

```go
func FetchWithTimeout(delay time.Duration, timeout time.Duration) (string, error)
```

### 条件

* goroutine 内で `time.Sleep(delay)` の後に `"result"` という文字列を channel に送る
* `select` を使って、結果が来たときとタイムアウトしたときを分岐する
* タイムアウトには `time.After` を使う
* タイムアウト時は `errors.New("timeout")` を返す

### 例

```go
v, err := FetchWithTimeout(50*time.Millisecond, 200*time.Millisecond)
fmt.Println(v, err)

v, err = FetchWithTimeout(300*time.Millisecond, 100*time.Millisecond)
fmt.Println(v, err)
```

出力:

```
result <nil>
 timeout
```

### 鍛えたいこと

`select` 文と `time.After` によるタイムアウトパターンです。
`select` は複数の channel 操作を同時に待ち、最初に準備できたものを実行します。このタイムアウトパターンは HTTP クライアントや外部 API 呼び出しなど、実務でほぼ必ず登場します。


# 解く順番の意図

この14問は、ただバラバラに並べたのではありません。
最初の1から4で **slice、map、append、条件分岐** を固めて、5と6で **struct、メソッド、interface** に進みます。次に7と8で **error と I/O** に触れ、9と10で **クロージャと並行処理** に入る流れです。

11以降はその応用です。11で **interface を自分で設計する**、12で **カスタムエラー型と errors.As**、13で **WaitGroup による goroutine の同期**、14で **select によるタイムアウト** を学びます。

つまり、**Goの基本文法 → Goらしい設計 → Goらしい標準ライブラリ → Goらしい並行処理 → 実務で使うパターン** という順番です。この順でやると、知識が点ではなく線になります。

# ディレクトリ構成
```
go-practice/
├── go.mod
├── README.md
├── problems/
│   ├── 01_sum/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── sum.go
│   │   └── sum_test.go
│   ├── 02_max/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── max.go
│   │   └── max_test.go
│   ├── 03_count_letters/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── count_letters.go
│   │   └── count_letters_test.go
│   ├── 04_filter_even/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── filter_even.go
│   │   └── filter_even_test.go
│   ├── 05_person_birthday/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── person.go
│   │   └── person_test.go
│   ├── 06_person_string/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── person.go
│   │   └── person_test.go
│   ├── 07_divide/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── divide.go
│   │   └── divide_test.go
│   ├── 08_count_from_reader/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── count_from_reader.go
│   │   └── count_from_reader_test.go
│   ├── 09_counter/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── counter.go
│   │   └── counter_test.go
│   ├── 10_goroutine_channel/
│   │   ├── README.md
│   │   ├── main.go
│   │   └── main_test.go
│   ├── 11_shape_interface/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── shape.go
│   │   └── shape_test.go
│   ├── 12_custom_error/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── config.go
│   │   └── config_test.go
│   ├── 13_wait_group/
│   │   ├── README.md
│   │   ├── main.go
│   │   ├── double.go
│   │   └── double_test.go
│   └── 14_select_timeout/
│       ├── README.md
│       ├── main.go
│       ├── fetch.go
│       └── fetch_test.go
└── shared/
    └── .gitkeep
```
