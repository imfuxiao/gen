## 代码生成器

因工作需要, 将工作中的对Mongo操作的增删改查做成统一的代码生成器, 所以目前版本只支持mongo.

原理是将固定部分的代码作为成template文件, 编写模型对象yaml文件, 程序读取yaml文件, 自动生成model/repository/service/controller模块代码

## 编译

代码使用了[packr.v2](https://github.com/gobuffalo/packr)对模板文件进行了打包, 所以需要参考 packr.v2 操作

```bash
packr2 # 初始化模板文件
packr2 build # 对程序进行打包
```

> 每次修改完模板文件, 都需要执行 `packr2` 重新打包模板代码文件

### 代码

部分代码参考了以下工程

* https://github.com/ezbuy/redis-orm
