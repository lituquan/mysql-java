### 主从设置
https://www.cnblogs.com/codehome/p/9356496.html

#### 查看主服务
启动配置my.ini：
    port = 3306
    server_id = 1
    log-bin=mysql-bin

连接主服务：mysql -uroot -p -hlocalhost -P3306
设置账户：
mysql> GRANT REPLICATION SLAVE,RELOAD,SUPER ON *.*  TO backup@'localhost'  IDENTIFIED BY '1234';

mysql> SHOW SLAVE STATUS

+------------------+----------+--------------+------------------+-------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+------------------+----------+--------------+------------------+-------------------+
| mysql-bin.000002 |      120 |              |                  |                   |
+------------------+----------+--------------+------------------+-------------------+

#### 查看从服务
启动配置：
    port=3307
    server_id=2
    log-bin=mysql-bin
    relay_log = mysql-relay-bin
    log_slave_updates = 1
    read_only= 1
连接从服务：mysql -uroot -p -hlocalhost -P3307
主拜从：
mysql> CHANGE MASTER TO MASTER_HOST='localhost',

    -> MASTER_USER='backup',

    -> MASTER_PASSWORD='1234',

    -> MASTER_LOG_FILE='mysql-bin.000002',

    -> MASTER_LOG_POS=0;
从服务状态：
mysql> show slave status\G
开始复制：    
mysql> START SLAVE;
注意：由于拷贝同一份mysql,所以data/auto.cnf 里面的server-uuid是相同的，把从服务里的这个文件删掉即可。

### https://blog.csdn.net/sunbocong/article/details/81634296