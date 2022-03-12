# axisapi

>axisapi:是一个防Gin的Go Web框架，设计思想参照开源Web 框架Gin，体量较小只有400多行代码

axisapi 主要实现了一下功能

- [x] 采用前缀树算法实现动态路由
- [x] 请求上下文处理模式
- [x] 分组路由控制功能
- [x] 默认定义了日志处理和错误处理中间件，用户也可以设置服务自身业务需求的中间件
- [x] 可响应json，data，string，html格式的数据
- [x] 支持POST
- [x] 支持GET



不足之处

- [ ] 处理前端除x-www-form-urlencoded格式以外的数据
- [ ] 模板功能
- [ ] 支持HEAD
- [ ] 支持PUT
- [ ] 支持DELETE
- [ ] 支持CONNECT
- [ ] 支持OPTIONS
- [ ] 支持TRACE
- [ ] 等等。。。（思考不全面，欢迎补充）
