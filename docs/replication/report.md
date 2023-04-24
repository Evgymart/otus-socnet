# Репликация

## Настройка

Создадим в докере 3 контейнера mysql_master, mysql_slave_first и mysql_slave_second.
Узнаем IP узнаем командой
```
 docker inspect \
  -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <container name>
```

Для начала создадим на мастере пользователя для репликации
```
CREATE USER 'replicator_one'@'172.18.0.4' IDENTIFIED WITH mysql_native_password BY 'password';
CREATE USER 'replicator_two'@'172.18.0.5' IDENTIFIED WITH mysql_native_password BY 'password';
```

И дадим ему нужные права

GRANT REPLICATION SLAVE ON *.* TO 'replicator_one'@'172.18.0.4';
GRANT REPLICATION SLAVE ON *.* TO 'replicator_two'@'172.18.0.5';
FLUSH PRIVILEGES;

Получим статус мастера
```
SHOW MASTER STATUS;
```
![image](./img/master_status.png)

Из этой таблицы нам нужны данные:
```
Binlog: 1.000003
Position: 157
```
Выгрузим данные на слейвы

```
mysqldump -h 127.0.0.1 -P 3306 -u root -p database > db.sql
mysql -h 127.0.0.1 -P 4406 -u root -p database < ./db.sql
mysql -h 127.0.0.1 -P 5506 -u root -p database < ./db.sql
rm db.sql
```

Дальше, запустим репликацию на слейвах:
```
CHANGE REPLICATION SOURCE TO
SOURCE_HOST='172.18.0.3',
SOURCE_USER='replicator_one',
SOURCE_PASSWORD='password',
SOURCE_LOG_FILE='1.000003',
SOURCE_LOG_POS=157;

START REPLICA;
```
  
Запросим статус и убеждаемся что ошибок нет:

```
SHOW REPLICA STATUS\G;
*************************** 1. row ***************************
             Replica_IO_State: Waiting for source to send event
                  Source_Host: 172.18.0.3
                  Source_User: replicator_one
                  Source_Port: 3306
                Connect_Retry: 60
              Source_Log_File: 1.000003
          Read_Source_Log_Pos: 157
               Relay_Log_File: 0977c6396938-relay-bin.000002
                Relay_Log_Pos: 318
        Relay_Source_Log_File: 1.000003
           Replica_IO_Running: Yes
          Replica_SQL_Running: Yes
              Replicate_Do_DB:
          Replicate_Ignore_DB:
           Replicate_Do_Table:
       Replicate_Ignore_Table:
      Replicate_Wild_Do_Table:
  Replicate_Wild_Ignore_Table:
                   Last_Errno: 0
                   Last_Error:
                 Skip_Counter: 0
          Exec_Source_Log_Pos: 157
              Relay_Log_Space: 535
              Until_Condition: None
               Until_Log_File:
                Until_Log_Pos: 0
           Source_SSL_Allowed: No
           Source_SSL_CA_File:
           Source_SSL_CA_Path:
              Source_SSL_Cert:
            Source_SSL_Cipher:
               Source_SSL_Key:
        Seconds_Behind_Source: 0
Source_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error:
               Last_SQL_Errno: 0
               Last_SQL_Error:
  Replicate_Ignore_Server_Ids:
             Source_Server_Id: 1
                  Source_UUID: bfb40824-e25d-11ed-bba5-0242ac120003
             Source_Info_File: mysql.slave_master_info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
    Replica_SQL_Running_State: Replica has read all relay log; waiting for more updates
           Source_Retry_Count: 86400
                  Source_Bind:
      Last_IO_Error_Timestamp:
     Last_SQL_Error_Timestamp:
               Source_SSL_Crl:
           Source_SSL_Crlpath:
           Retrieved_Gtid_Set:
            Executed_Gtid_Set:
                Auto_Position: 0
         Replicate_Rewrite_DB:
                 Channel_Name:
           Source_TLS_Version:
       Source_public_key_path:
        Get_Source_public_key: 0
            Network_Namespace:
```