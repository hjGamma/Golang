[https://github.com/gin-gonic/gin
https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/
https://www.kancloud.cn/jiajunxi/ginweb100/1801414

gin语法
---------------------------------------------------------------------------------------------
1. gin.BasicAuth 是一个 Gin 中间件，用于处理 HTTP 基本认证。它会拦截请求，并验证请求头中的 Authorization 字段。

2. gin.Accounts 是用于定义一个 map[string]string，其中键是用户名，值是密码。它用于存储一组用户名和密码的组合。

3. gin.Default()用于创建一个包含默认中间件（如日志和恢复中间件）的 *gin.Engine 实例。

-----------------------------------------------------------------------------------------------
r := gin.Default()中r包含的方法
1. 路由方法
GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS 等方法： 这些方法用于绑定 HTTP 请求的路由，并指定对应的处理函数。每种方法对应 HTTP 协议的一个请求类型。
r.GET("/path", handler)
r.POST("/path", handler)
r.PUT("/path", handler)
r.DELETE("/path", handler)
r.PATCH("/path", handler)
r.OPTIONS("/path", handler)
r.HEAD("/path", handler)

6. 路由组的中间件管理
r.Group("/api").Use(authMiddleware())
    定义一个路由组，并在该组内使用中间件。

3. 中间件管理
r.Use(gin.Logger())
    gin.Logger()：记录日志
r.Use(gin.Recovery())
    gin.Recovery()：在发生 panic 时恢复，防止服务器崩溃。

4. 静态文件服务
r.Static("/static", "./static")
    用于提供整个目录的静态文件服务。通过设置一个目录，允许访问该目录下的所有文件。
r.StaticFile("/favicon.ico", "./resources/favicon.ico")
    用于提供单个静态文件服务。
r.StaticFS(relativePath, http.FS)
    支持从内存中的文件系统、压缩包文件系统、或者其他定制的文件系统来提供文件

5. 模板渲染
r.LoadHTMLGlob("templates/*")
    加载指定目录下的所有 HTML 模板文件。

6. 错误处理 
r.NoRoute(handler)
    定义一个处理未匹配路由的处理函数。

7. 启动服务器
r.Run(":8080")
    启动服务器，监听指定的端口。
-----------------------------------------------------------------------------------------
Context中包含的方法
1. 路由参数和请求数据的获取
c.Param(key string) string：获取 URL 路径中的路由参数。例如，定义路径 /:name 后可以用 c.Param("name") 获取路径参数。
c.Query(key string) string：获取 URL 查询参数（即 ?key=value 部分）。如果参数不存在，返回空字符串。
c.PostForm(key string) string：获取 POST 表单中的数据。
c.DefaultQuery(key, defaultValue string) string：获取 URL 查询参数，如果不存在则返回指定的默认值。
c.DefaultPostForm(key, defaultValue string) string：获取 POST 表单中的数据，如果不存在则返回指定的默认值。
c.BindJSON(obj interface{}) error：绑定 JSON 请求体的数据到结构体上，用于处理 JSON 数据的请求。
c.ShouldBind(obj interface{}) error：自动根据 Content-Type 来决定解析方式，支持 JSON、XML、表单等格式。
c.GetHeader(key string) string：获取请求头中的指定字段值。

2. 响应数据的设置
c.JSON(statusCode int, obj interface{})：返回 JSON 格式的响应。
    通常使用gin.H来构建 JSON 响应，例如 gin.H{"message": "Hello, Gin!"}。
c.String(statusCode int, format string, values ...interface{})：返回字符串格式的响应。
c.XML(statusCode int, obj interface{})：返回 XML 格式的响应。
c.HTML(statusCode int, name string, obj interface{})：返回 HTML 模板渲染结果的响应（需要先加载模板文件）。
c.Data(statusCode int, contentType string, data []byte)：返回原始字节数据的响应。
c.File(filepath string)：返回文件内容给客户端，常用于文件下载。

3. 请求上下文和控制
c.Next()：执行当前中间件链中的下一个处理程序。
    由于中间件的执行与其他函数执行是分开的，当脱离当前代码段与上下文时，如果不使用c.Next()，后续的处理程序将无法执行。
c.Abort()：终止当前请求的处理，后续中间件将不再执行。
c.AbortWithStatus(statusCode int)：终止处理，并立即返回指定的状态码。
c.Set(key string, value interface{}) 和 c.Get(key string) (interface{}, bool)：设置和获取上下文中的键值对，用于在请求生命周期中共享数据。

4. Cookies 和 Headers
c.Cookie(name string) (string, error)：获取指定名称的 cookie 值。
c.SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)：设置一个 cookie。
c.GetHeader(key string) string：获取请求头中的指定字段值。
c.Header(key, value string)：设置响应头的字段和值。

5. 文件上传
c.FormFile(name string) (*multipart.FileHeader, error)：获取上传的文件。
c.SaveUploadedFile(file *multipart.FileHeader, dst string) error：将上传的文件(通过 c.FormFile方法获得)保存到服务器指定路径dst。

6. 错误处理
c.Error(err error)：记录请求过程中的错误。
c.Errors：包含所有已记录的错误，可以在请求结束后进行统一处理或日志记录。

7. 重定向
c.Redirect(statusCode int, url string)：重定向到指定的 URL。
](https://github.com/gin-gonic/gin
https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/
https://www.kancloud.cn/jiajunxi/ginweb100/1801414

web

1. 进入main.go,初始化路由，以及端口号
2. 根据浏览器输入的URL地址，在router路由器中找到对应的路由函数方法
3. 根据路由中URL后指定的函数，在Controller中找到对应的方法函数
4. Controller中调用models关于数据库方面的方法函数
5. 渲染html页面，执行js,css等效果

gin语法
---------------------------------------------------------------------------------------------
1. gin.BasicAuth 是一个 Gin 中间件，用于处理 HTTP 基本认证。它会拦截请求，并验证请求头中的 Authorization 字段。

2. gin.Accounts 是用于定义一个 map[string]string，其中键是用户名，值是密码。它用于存储一组用户名和密码的组合。

3. gin.Default()用于创建一个包含默认中间件（如日志和恢复中间件）的 *gin.Engine 实例。

-----------------------------------------------------------------------------------------------
# r := gin.Default()中r包含的方法
1. 路由方法
GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS 等方法： 这些方法用于绑定 HTTP 请求的路由，并指定对应的处理函数。每种方法对应 HTTP 协议的一个请求类型。
r.GET("/path", handler)
r.POST("/path", handler)
r.PUT("/path", handler)
r.DELETE("/path", handler)
r.PATCH("/path", handler)
r.OPTIONS("/path", handler)
r.HEAD("/path", handler)

2. 路由组的中间件管理
r.Group("/api").Use(authMiddleware())
    定义一个路由组，并在该组内使用中间件。

3. 中间件管理
r.Use(gin.Logger())
    gin.Logger()：记录日志
r.Use(gin.Recovery())
    gin.Recovery()：在发生 panic 时恢复，防止服务器崩溃。

4. 静态文件服务
r.Static("/static", "./static")
    用于提供整个目录的静态文件服务。通过设置一个目录，允许访问该目录下的所有文件。
r.StaticFile("/favicon.ico", "./resources/favicon.ico")
    用于提供单个静态文件服务。
r.StaticFS(relativePath, http.FS)
    支持从内存中的文件系统、压缩包文件系统、或者其他定制的文件系统来提供文件

5. 模板渲染
r.LoadHTMLGlob("templates/*")
    加载指定目录下的所有 HTML 模板文件。

6. 错误处理 
r.NoRoute(handler)
    定义一个处理未匹配路由的处理函数。

7. 启动服务器
r.Run(":8080")
    启动服务器，监听指定的端口。
-----------------------------------------------------------------------------------------
# Context中包含的方法
1. 路由参数和请求数据的获取
c.Param(key string) string：获取 URL 路径中的路由参数。例如，定义路径 /:name 后可以用 c.Param("name") 获取路径参数。
c.Query(key string) string：获取 URL 查询参数（即 ?key=value 部分）。如果参数不存在，返回空字符串。
c.PostForm(key string) string：获取 POST 表单中的数据。默认解析的是x-www-form-urlencoded或from-data格式的参数
c.DefaultQuery(key, defaultValue string) string：获取 URL 查询参数，如果不存在则返回指定的默认值。
c.DefaultPostForm(key, defaultValue string) string：获取 POST 表单中的数据，如果不存在则返回指定的默认值。
c.BindJSON(obj interface{}) error：绑定 JSON 请求体的数据到结构体上，用于处理 JSON 数据的请求。
c.ShouldBind(obj interface{}) error：自动根据 Content-Type 来决定解析方式，支持 JSON、XML、表单等格式。
c.GetHeader(key string) string：获取请求头中的指定字段值。

2. 响应数据的设置
c.JSON(statusCode int"http.StatusOK", obj interface{})：返回 JSON 格式的响应。
    通常使用gin.H来构建 JSON 响应，例如 gin.H{"message": "Hello, Gin!"}。
c.String(statusCode int, format string, values ...interface{})：返回字符串格式的响应。
c.XML(statusCode int, obj interface{})：返回 XML 格式的响应。
c.HTML(statusCode int, name string, obj interface{})：返回 HTML 模板渲染结果的响应（需要先加载模板文件）。
c.Data(statusCode int, contentType string, data []byte)：返回原始字节数据的响应。
c.File(filepath string)：返回文件内容给客户端，常用于文件下载。

3. 请求上下文和控制
c.Next()：执行当前中间件链中的下一个处理程序。
    由于中间件的执行与其他函数执行是分开的，当脱离当前代码段与上下文时，如果不使用c.Next()，后续的处理程序将无法执行。
c.Abort()：终止当前请求的处理，后续中间件将不再执行。
c.AbortWithStatus(statusCode int)：终止处理，并立即返回指定的状态码。
c.Set(key string, value interface{}) 和 c.Get(key string) (interface{}, bool)：设置和获取上下文中的键值对，用于在请求生命周期中共享数据。

4. Cookies 和 Headers
c.Cookie(name string) (string, error)：获取指定名称的 cookie 值。
c.SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)：设置一个 cookie。
c.GetHeader(key string) string：获取请求头中的指定字段值。
c.Header(key, value string)：设置响应头的字段和值。

5. 文件上传
c.FormFile(name string) (*multipart.FileHeader, error)：获取上传的文件。
c.SaveUploadedFile(file *multipart.FileHeader, dst string) error：将上传的文件(通过 c.FormFile方法获得)保存到服务器指定路径dst。

6. 错误处理
c.Error(err error)：记录请求过程中的错误。
c.Errors：包含所有已记录的错误，可以在请求结束后进行统一处理或日志记录。

7. 重定向
c.Redirect(statusCode int, url string)：重定向到指定的 URL。
------------------------------------
# Session
1. 初始化Session
    1.1 初始化Session存储引擎
    store := cookie.NewStore([]byte("secret")) //secret为加密密钥
    1.2 初始化Session存储
    sessionStore := session.New(store, sessionConfig)
    1.3 注册Session中间件
    r.Use(sessions.Sessions("mysession", store)) //mysession为Session的名称，store为之前定义的存储器
2. 使用Session
    session := sessions.Default(c) //获取当前请求的Session对象
    2.1 获取Session
    session.Get("secret") //获取Session中的secret字段
    2.2 设置Session
    session.Set("secret", username) 在controller中设置Session，保持登录状态
    2.3 删除Session
    
    2.4 清除所有Session
    2.5 保存Session
    session.Save()
    2.6 刷新Session
    2.7 销毁Session
    2.8 检查Session是否存在
    2.9 获取Session的ID

)
