### GO beginner
* conf:负责防止各种配置文件
* config：负责加载各种配置文件
* handler:响应处理
* model：数据层
* pkg：常量，或者一些错误号的定义
* router：路由以及一些中间件的定义
* service：具体处理业务逻辑的地方
* main.go：主入口
* vendor：包管理的目录
* util：工具类
* db.sql：据库表定义
### request----->middlewares----->router------>service----->model----->handler------>response


# reference https://blog.csdn.net/Survivalist_Boy/article/details/88804284