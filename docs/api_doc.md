### 测试用例接口   
>   1.文档主要提供本项目骨架已经集成的Api接口使用说明。            
>   2.相关测试全部基于`postman`工具进行。     

### 默认已经集成的路由  

####  门户网站类  
>GET    http://127.0.0.1:20191    
>GET    /api/v1/home/news?newsType=portal&page=1&limit=50 

#### 后台管理类
>GET    /http://127.0.0.1:20201                         
>GET   /Admin/ws         
>POST   /Admin/users/register     
>POST   /Admin/users/login        
>POST   /Admin/users/refreshtoken        
>GET    /Admin/users/index        
>POST   /Admin/users/create       
>POST   /Admin/users/edit         
>POST   /Admin/users/delete       
>POST   /Admin/upload/file           

#### pprof 路由
>调试模式自动开启，以pprof开头的路由   
> http://127.0.0.1:20191/debug/pprof/  
> http://127.0.0.1:20201/debug/pprof/  

### 门户网站类
>   1.ip、端口使用本项目默认配置，即：`http://127.0.0.1:20191`，门户类接口通用  
####    1.首页新闻   
>    <font color=#FF4500>*get*，/api/v1/home/news?newsType=portal&page=1&limit=50 
> 返回示例：
```json
{
    "code": 200,
    "data": {
        "content": "门户新闻内容001",
        "limit": 20,
        "newstype": "potal",
        "page": 1,
        "title": "门户首页公司新闻标题001",
        "user_ip": "127.0.0.1"
    },
    "msg": "Success"
}
```  




### 后台应用类
>   1.ip、端口使用本项目默认配置，即：`http://127.0.0.1:20201`，后端管理类系统通用。  

####    1.用户注册   
>    <font color=#FF4500>*post*，/Admin/users/register   

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
name|form-data|string|必填|admin123
pass|form-data|string|必填|123456abc
phone|form-data|string|必填|1360177xxxx
> 返回示例：
```json
{
    "code": 200,
    "data": "",
    "msg": "Success"
}
```  

####    2.用户登录     
>    <font color=#FF4500>*post*，/Admin/users/login   

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
name|form-data|string|必填|admin123
pass|form-data|string|必填|123456abc
phone|form-data|string|非必填|1360177xxxx
> 返回示例：
```json
{
    "code": 200,
    "data": {
        "name": "admin123",
        "phone": "",
        "real_name": "",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE1ODgyMzkzNzcsIm5iZiI6MTU4ODIzNTc2N30.jOGFKEitESsaT965RNXZMEgG6cVxOCU_pFCacfUU1iU",
        "updated_at": "2020-04-30 16:36:17",
        "userid": 14
    },
    "msg": "Success"
}
```  

####    3.根据关键词查询用户表      
>    <font color=#FF4500>*get*，/Admin/users/index   ，注意该接口需要token鉴权，请在 `header` 头添加 `Authorization` 字段值  
>   CURD相关的其他接口格式与本接口基本一致，例如：/Admin/users/create、/Admin/users/edit、/Admin/users/delete，只不过表单参数不一致。    

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE1ODgyMzkzNzcsIm5iZiI6MTU4ODIzNTc2N30.jOGFKEitESsaT965RNXZMEgG6cVxOCU_pFCacfUU1iU
name|form-data|string|必填|a
page|form-data|int|必填|1
limits|form-data|int|必填|20

> 返回示例：
```json
{
    "code": 200,
    "data": {
        "name": "admin123",
        "phone": "",
        "real_name": "",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE1ODgyMzkzNzcsIm5iZiI6MTU4ODIzNTc2N30.jOGFKEitESsaT965RNXZMEgG6cVxOCU_pFCacfUU1iU",
        "updated_at": "2020-04-30 16:36:17",
        "userid": 14
    },
    "msg": "Success"
}
```  

####    5.token刷新 ，请将旧token防止在header头参数直接提交更新         
>    <font color=#FF4500>*post*，/Admin/users/refreshtoken    

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE1ODgyMzkzNzcsIm5iZiI6MTU4ODIzNTc2N30.jOGFKEitESsaT965RNXZMEgG6cVxOCU_pFCacfUU1iU  

> 返回示例：
```json
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE1ODgyNDQxMDYsIm5iZiI6MTU4ODIzNTc2N30.Yaxah-3GdK6lFOxRaDGAiZ1vBh66uN8vL1mnxcVnlOQ"
    },
    "msg": "Success"
}
```  

####    6.文件上传        
>    <font color=#FF4500>*post*，/Admin/upload/file    

参数字段|参数属性|类型|选项|默认值
---|---|---|---|---
Authorization|Headers|string|必填|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjE0LCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE1ODgyMzkzNzcsIm5iZiI6MTU4ODIzNTc2N30.jOGFKEitESsaT965RNXZMEgG6cVxOCU_pFCacfUU1iU
files|form-data|string|必填|(注意表单键名为files，如果需要修改成别的键名，参见：App\Global\Variable\Variable.go ，UploadFileField=files)
> 返回示例：
```json
{
    "code": 200,
    "data": {
        "path": "E:\\GO\\TestProject\\goskeleton/Storage/app/uploaded/9c7e67955cfc32b3399695c3ae814bef.png"
    },
    "msg": "Success"
}
```  