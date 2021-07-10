# compitition
AtCoderとかの過去問などを説いていった備忘録
基本Goです

# go tips
## set
goにはsetがないのでmapのkeyが重複を許さない事を利用する
```
slice := []int{1,2,3,}
 
set := make(map[int]struct{})
for _, v := range slice {
	set[v] = struct{}{}
}
```

sliceの中身が重複してるのか否かの判定は上記のsliceとsetの長さを比べる
```
if len(slice) != len(set) {}
```

## 無名関数
```
// 個人的にはdfsで使うことが多いので

var dfs func(start int) 
dfs = func(start int) {
  hoge
  for {
    dfs(start+1)
  }
}
```
  
