## 路由设计
使用RESTFUL API设计。
- GET: 读取和检索
- POST: 新增和新建
- PUT: 更新动作，用于更新一个完整的资源，要求冥等
- PATCH: 更新动作，用于更新某一个资源的一个组成部分，可以不冥等
- DELETE: 删除操作

### 标签管理
| 功能 | HTTP方法 | 路径 |   
| 新增标签 | POST | /tags |  
| 删除指定标签 | DELETE | /tags/:id |  
| 更新指定标签 | PUT | /tags/:id |  
| 获取标签列表 | GET | /tags |  


### 文章管理
| 功能 | HTTP方法 | 路径 |  
| 新增文件 | POST | /articles |  
| 删除指定文章 | DELETE | /articles/:id |  
| 更新指定文章 | PUT | /articles/:id |  
| 获取指定文章 | GET | /articles/:id |  
| 获取文章列表 | GET | /articles |  

