# goctl

魔改版
主要做了以下修改
- 1.logic生成 err类型修改为 github.com/go-kratos/kratos/v2/errors 需要返回 code reason message
- 2.response 统一修改为 code message data 类型的响应
- 3.swagger生成可以使用 https://github.com/sliveryou/goctl-swagger
