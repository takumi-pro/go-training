# 合成型

https://zenn.dev/spiegel/articles/20210315-array-and-slice#%E3%82%B9%E3%83%A9%E3%82%A4%E3%82%B9%E3%81%AF%E5%8F%82%E7%85%A7%E3%81%A7%E3%81%82%E3%82%8A%E5%80%A4%E3%81%A7%E3%81%82%E3%82%8B

上記の記事でスライスが「参照のよう振る舞う」意味が理解できた。

## map
- ゼロ値は`nil`
- マップ同士の比較はできない
- `make`関数を使用してサイズを指定して宣言できる

## 構造体
- 構造体の比較はフィールドに依存する
  - 構造体のフィールドが全て比較可能であれば構造体も比較可能」



## 参考
- [MAPは連想配列ではなく連想配列への「参照」である - text.Bldanders.info](https://text.baldanders.info/golang/map-as-a-associative-array/)
- [配列とスライス - zenn](https://zenn.dev/spiegel/articles/20210315-array-and-slice#fn-573f-5)
