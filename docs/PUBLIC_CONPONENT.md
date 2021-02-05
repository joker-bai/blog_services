## 公共组件
### 错误代码
在应用程序运行过程中，我们常常需要与客户端进行交互。交互一般分为两点：一个是返回正确响应下的结果集；另一个是返回错误响应下的错误码和消息体，以便告诉客户端，这一次请求发生了什么事，以及请求失败的原因

#### 公共错误代码
- Success    0:成功
- ServerError 100000:服务内部错误
- InvalidParams 100001:入参错误
- NotFound  100002:找不到
- UnauthorizedAuthNotExist  100003:鉴权失败,找不到对应的AppKey和AppSecret
- UnauthorizedTokenError 100004:鉴权失败,Token错误
- UnauthorizedTokenTimeout  100005:鉴权失败,Token超时
- UnauthorizedTokenGenerate  100006:鉴权失败,Token生成失败
- TooManyRequests 100007:请求过多

### 配置中心
使用viper进行配置管理

### 数据库
使用gorm进行数据库管理操作

### 日志管理
使用lumberjack进行日志管理