## LeTaoTao



## 依赖包
 * 数据库
    * **mysql**: github.com/go-sql-driver/mysql
    * **httprouter**: github.com/julienschmidt/httprouter
    * **redis**: github.com/go-redis/redis

## 组织结构
```text
    ___________
   |__REST API_| ...
    ____________________________________
   |             Router                 |
   |____________________________________|
    ____________________________________
   |             Wrapper                |
   |____________________________________|
    ____________________________________
   |             Service                |
   |____________________________________|
    ____________________________________
   |             Repository             |
   |____________________________________|
    ________    _________   _________
   |  mysql |  |  redis  | | mongodb | ..
   |________|  |_________| |_________|
   
```   
    
## TODO
 * 数据库管理模块
 * 完善REST API
 * JWT 验证
 * 缓存中间件 扩展 
 * 日志模块   
 * GO程序运行在Docker环境中
 * GO程序存储到源仓库中 通过包管理工具apt/yum/brew下载运行
 
## DOING
 * systemd 方式管理GO执行文件
 * 
 
 
## DOC
 * https://godoc.org/github.com/go-redis/redis 
 