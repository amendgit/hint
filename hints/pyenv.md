1. 将brew安装的python版本交给pyenv进行管理

```sh
# python 2
ln -s $(brew --cellar python@2)/* ~/.pyenv/versions/
# python 3
ln -s $(brew --cellar python)/* ~/.pyenv/versions/
```

* python3的版本，需要进入到对应的文件夹(如：`~/.pyenv/versions/3.7.3/bin`)将python3拷贝一份命名为python。