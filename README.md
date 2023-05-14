# common
golang develop common kit




## 如何配置ssh连接
1. 配置本地ssh：
```
git config --global user.name "XXX"
git config --global user.email "xxxx@163.com"
ssh-keygen -t rsa -C "xxxx@163.com"
```
2. 登录github账号，点击“头像”->“settings”-> "ssh&GPG keys" 复制 ～/.ssh中id_rsa.pub中全部的文字，起一个 title name保存。

3. git clone 自己的项目