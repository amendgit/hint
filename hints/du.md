1. 获取文件夹自身的大小
```sh
hu -sh dir/
```

2. 获取文件夹下所有子文件夹的大小
```sh
hu -h dir/
```

3. 指定子文件夹的深度
```sh
hu -hd1 dir/
```

4. 按照大小进行排序
```sh
hu -hd1 dir/ | sort -h
```