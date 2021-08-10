
使用方法

```
git clone https://github.com/Issei0804-ie/parallel.git
cd parallel
go run main.go 

# n=5, weight=6

5 6

# {{Value:5,Weight:1},{Value:2,Weight:3},{Value:5,Weight:1},{Value:2,Weight:3},{Value:9,Weight:4}}
5 1
2 3
5 1
2 3
9 4

#output
w = 6 v = 19 b = 21
```

逐次処理と並列処理の実行時間比較

```
cd parallel
go test -v
```
