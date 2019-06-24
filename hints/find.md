1. 在当前目录及子目录下，查找名称为foo的文件

```sh
find . -name foo
```

2. 指定查找类型是文件夹

```sh
find . -name foo -type d
```

3. 指定查找子文件夹的最大深度为3

```sh
find . -name foo -max-depth 3
```