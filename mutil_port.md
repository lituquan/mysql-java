### 开多个端口
    新建：
        my.cnf
        my.ini

### 设置port
    port=3307

    netstat -ano |findstr "3307"

    taskkill -pid 10368 -t -f