go-fiber 主要定位为一个轻量级、高性能的 Web 框架

由于fiber使用 fasthttp 作为 HTTP 引擎，使得 Fiber 的性能非常出色



学习fiber的大部分人，应该是有gin做基础的吧？如果是零基础学fiber的，可以反馈给我，人多的话

出一期fiber的零基础入门课程

本课程是针对有gin基础的人，但是处于其他原因需要使用fib 做开发



## fiber怎么用



### 运行项目

运行是通过Listen方法

```Go
package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New(fiber.Config{})
  app.Get("/", func(c *fiber.Ctx) error {
    return nil
  })
  app.Post("/", func(c *fiber.Ctx) error {
    return nil
  })
  app.Put("/", func(c *fiber.Ctx) error {
    return nil
  })
  app.Delete("/", func(c *fiber.Ctx) error {
    return nil
  })
  err := app.Listen(":80")
  fmt.Println(err)
}
```



### 响应

响应字符串

```Go
c.SendString("string")
```



响应json

```Go
c.JSON(map[string]any{"code": 0, "msg": "成功"})
type Info struct {
  Name string `json:"name"`
  Age  int    `json:"age"`
}
c.JSON(Info{"fengfeng", 18})
```



响应html

```Go
c.Set("Content-Type", "text/html;charset=utf-8")
c.SendString("<h1>你好</h1>")
```



```Go
c.SendFile("index.html")
```



静态路由

```Go
app.Static("/static", "static")
```



### 路由分组

Route

Group

```Go
app.Route("/system", func(router fiber.Router) {
router.Get("health", func(c *fiber.Ctx) error { return c.SendString("system.health") })
router.Get("info", func(c *fiber.Ctx) error { return c.SendString("system.info") })
router.Get("user/list", func(c *fiber.Ctx) error { return c.SendString("system.user/list") })
})

group := app.Group("/video")
group.Get("info", func(c *fiber.Ctx) error { return c.SendString("video.info") })
group.Get("progress", func(c *fiber.Ctx) error { return c.SendString("video.progress") })

```



### 动态路由

除了动态路由，fiber还支持*通配符路由

```Go
app.Get("/user/:name", func(c *fiber.Ctx) error {
    return c.SendString("动态路由" + c.Params("name"))
  })
app.Get("/user/*", func(c *fiber.Ctx) error {
  return c.SendString("通配符路由")
})
```



### 反向解析

在定义路由的时候，可以给路由设置一个名字

在其他地方，就可以基于这个路由名称反向得到路由的路径

```Go
// 1. 给路由命名
app.Get("/user/:id", func(c *fiber.Ctx) error {
    return c.SendString("用户ID：" + c.Params("id"))
}).Name("user.detail") // 命名为 user.detail

// 2. 反向解析 URL（动态生成 /user/123）
app.Get("/test", func(c *fiber.Ctx) error {
    // 传入路由名 + 参数映射，生成 URL
    url, err := c.GetRouteURL("user.detail", fiber.Map{"id": 123})
    if err != nil {
        return err
    }
    return c.SendString("用户详情页 URL：" + url) // 输出：/user/123
})
```



### 文件上传

SaveFile不会自动创建目录

```Go
package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New(fiber.Config{})
  app.Post("upload", func(c *fiber.Ctx) error {
    fileHeader, err := c.FormFile("file")
    if err != nil {
      return err
    }
    fmt.Println(fileHeader.Filename)
    return c.SaveFile(fileHeader, "uploads/"+fileHeader.Filename)
  })
  app.Post("uploads", func(c *fiber.Ctx) error {
    form, err := c.MultipartForm()
    if err != nil {
      return err
    }
    for s, headers := range form.File {
      for _, header := range headers {
        fmt.Printf("%s %s\n", s, header.Filename)
      }
    }
    return c.Status(201).SendString("成功")
  })
  app.Listen(":80")
}

```



### 参数绑定

绑定json，绑定form

都是用BodyParser，通过请求的ContentType进行区分的

如果是json，就用json的tag

如果是form，就用form的tag

```Go
type RequestData struct {
  Name string `json:"name" form:"name"`
  Age  int    `json:"age" form:"age"`
}
var user RequestData
err := c.BodyParser(&user)
```



绑定query

```Go
type RequestData struct {
  Name string `query:"name"`
  Age  int    `query:"age"`
}
var user RequestData
err := c.QueryParser(&user)
```



绑定uri

要和动态路由结合

```Go
app.Get("/user/:name/:id", func(c *fiber.Ctx) error {
  type RequestData struct {
    Name string `params:"name"`
    ID   int    `params:"id"`
  }
  var user RequestData
  c.ParamsParser(&user)
  return c.JSON(user)
})
```



### 参数校验

和gin用的同一个参数校验库，不过使用逻辑有点不太一样

gin是直接把validate库内置了，但是fiber需要手动操作

```Go
go get github.com/go-playground/validator/v10
```



```Go
package main

import (
  "fmt"
  "github.com/go-playground/validator/v10"
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New(fiber.Config{})
  validate := validator.New()

  app.Get("/validate", func(c *fiber.Ctx) error {
    type RequestData struct {
      Name  string `json:"name" params:"name" validate:"required,min=2,max=10"` // 必须存在，长度2-10
      ID    int    `json:"id" params:"id" validate:"required,min=1"`            //  必须存在，最小值1
      Email string `json:"email" params:"email" validate:"email"`               //  存在就得是合法的邮箱格式
    }
    var user RequestData
    err := c.QueryParser(&user)
    if err != nil {
      fmt.Println("参数绑定失败", err)
      return err
    }
    err = validate.Struct(user)
    if err != nil {
      fmt.Println("参数校验失败", err)
      return err
    }
    return c.JSON(user)
  })
  err := app.Listen(":80")
  fmt.Println(err)
}

```



#### 校验错误显示中文

1. 创建翻译器
2. 给validate 注册翻译器
3. 遇到校验错误使用翻译器

```Go
package main

import (
  "fmt"
  "github.com/go-playground/locales/zh"
  ut "github.com/go-playground/universal-translator"
  "github.com/go-playground/validator/v10"
  zh_translations "github.com/go-playground/validator/v10/translations/zh"
  "github.com/gofiber/fiber/v2"
  "strings"
)

var trans ut.Translator
var validate *validator.Validate

func init() {
  // 创建翻译器
  uni := ut.New(zh.New())
  trans, _ = uni.GetTranslator("zh")

  // 初始化 validator 并注册中文翻译
  validate = validator.New()
  if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
    panic(fmt.Sprintf("注册中文翻译失败：%v", err))
  }
}

func main() {
  app := fiber.New(fiber.Config{})

  app.Get("/validate", func(c *fiber.Ctx) error {
    type RequestData struct {
      Name  string `json:"name" params:"name" validate:"required,min=2,max=10"` // 必须存在，长度2-10
      ID    int    `json:"id" params:"id" validate:"required,min=1"`            //  必须存在，最小值1
      Email string `json:"email" params:"email" validate:"email"`               //  存在就得是合法的邮箱格式
    }
    var user RequestData
    err := c.QueryParser(&user)
    if err != nil {
      fmt.Println("参数绑定失败", err)
      return err
    }
    err = validate.Struct(user)
    if err != nil {
      //fmt.Println("参数校验失败", err)
      errs := err.(validator.ValidationErrors)
      var errMsgs []string
      for _, e := range errs {
        // 使用翻译器
        errMsgs = append(errMsgs, e.Translate(trans))
      }
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "error": strings.Join(errMsgs, ";"),
      })
    }
    return c.JSON(user)
  })
  err := app.Listen(":80")
  fmt.Println(err)
}

```



### 请求头响应头操作

```Go
// 1. 获取请求头
app.Get("/header", func(c *fiber.Ctx) error {
    token := c.Get("Authorization") // 获取 Authorization 头
    userAgent := c.Get("User-Agent")// 获取 User-Agent
    return c.JSON(fiber.Map{
        "token":      token,
        "user_agent": userAgent,
    })
})

// 2. 设置响应头
app.Get("/set-header", func(c *fiber.Ctx) error {
    c.Set("Content-Type", "application/json;charset=utf-8") // 显式设置响应头
    c.Set("Cache-Control", "max-age=3600") // 设置缓存时间（1小时）
    return c.JSON(fiber.Map{"msg": "已设置响应头"})
})
```



### 中间件

在中间件里面可以很方便的获取请求体和响应体的数据



#### 视图中间件

```Go
package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
  fmt.Println("请求前")
  fmt.Println("请求体数据：", string(c.Body()))
  // return c.JSON(map[string]any{"code": -1}) // 如果需要拦截
  c.Next()
  fmt.Println("响应后")
  fmt.Println("响应体数据：", string(c.Response().Body()))
  return nil
}

func main() {
  app := fiber.New(fiber.Config{})

  app.Post("/user", AuthMiddleware, func(c *fiber.Ctx) error {
    fmt.Println("视图函数")
    return c.JSON(map[string]any{"data": "用户数据"})
  })
  err := app.Listen(":80")
  fmt.Println(err)
}

```



#### 分组中间件

可以通过group的后续参数挂载中间件

也可以基于group的对象的use方法挂载中间件，两者是等价的

```Go
app.Group("/", LoggerMiddleware).Use(LoggerMiddleware)
```





#### 全局中间件

直接基于app的use方法挂载中间件即可

这个中间件会对所有路由生效（包括静态路由），并且优先级最高

```Go
func LoggerMiddleware(c *fiber.Ctx) error {
  // 请求前：记录开始时间
  start := time.Now()

  // 放行请求（执行后续中间件和路由处理函数）
  err := c.Next()

  // 响应后：计算耗时并打印日志
  duration := time.Since(start)
  fmt.Printf("[%s] %s %s - %v\n",
    c.Method(), // 请求方法（GET/POST 等）
    c.Path(),   // 请求路径
    c.IP(),     // 客户端 IP
    duration,   // 处理耗时
  )
  return err
}

app.Use(LoggerMiddleware)
```



### 中间件之间的参数传递

举个例子：每次在视图里面做参数绑定和校验 都挺麻烦的，是否可以在中间件里面做参数绑定和参数校验，然后在视图函数中获取校验之后的数据

```Go
package main

import (
  "github.com/go-playground/validator/v10"
  "github.com/gofiber/fiber/v2"
)

var validate1 *validator.Validate

func init() {
  validate1 = validator.New()
}
func bindMiddleware[T any](c *fiber.Ctx) error {
  var cr T
  err := c.QueryParser(&cr)
  if err != nil {
    return err
  }
  err = validate1.Struct(cr)
  if err != nil {
    return err
  }
  c.Locals("value", cr)
  c.Next()
  return nil
}

func main() {
  app := fiber.New(fiber.Config{})

  type Info struct {
    Name string `query:"name"`
  }

  app.Get("/", bindMiddleware[Info], func(c *fiber.Ctx) error {
    info := c.Locals("value").(Info)
    return c.JSON(info)
  })

  app.Listen(":80")
}

```





### 官方支持中间件

|中间件名称|包路径|核心用途|
|-|-|-|
|**CORS 跨域**  |`github.com/gofiber/cors/v2`|处理跨域请求，支持配置允许的源、方法、请求头，解决前端跨域问题。|
|**限流**|`github.com/gofiber/limiter/v2`|限制接口请求频率（如每秒 / 每分钟最大请求数），防止接口被刷，保护服务稳定。  |
|**JWT 认证**|`github.com/gofiber/jwt/v3`|基于 JWT（JSON Web Token）的身份认证，用于用户登录态校验、接口权限控制。|
|**WebSocket**|`github.com/gofiber/websocket/v2`|快速实现 WebSocket 服务，自动处理 HTTP 到 WebSocket 的协议升级，简化双向通信。|
|**静态资源**|内置（无需额外安装）|通过 `app.Static()` 提供静态文件服务，支持压缩、目录浏览、缓存等高级配置（核心功能，虽未单独拆包，但属于官方核心能力）。|
|**请求 ID**|`github.com/gofiber/requestid/v2`|为每个请求生成唯一 ID（如 UUID），并写入响应头（`X-Request-ID`），方便日志追踪、链路排查。|
|**日志**|`github.com/gofiber/fiber/v2/middleware/logger`|内置日志中间件，记录请求方法、路径、耗时、状态码等信息，支持自定义日志格式（如 JSON 格式）。|
|**恢复（Panic 捕获）**|`github.com/gofiber/fiber/v2/middleware/recover`|内置异常捕获中间件，捕获请求处理中的 `panic`，防止服务崩溃，返回友好的 500 错误响应。|
|**压缩**|`github.com/gofiber/compression/v2`|自动压缩响应内容（支持 gzip、brotli 等算法），减少传输带宽，提升接口响应速度（尤其适合文本类响应，如 JSON、HTML）。|
|**ETag**|`github.com/gofiber/etag/v2`|自动生成响应的 ETag（资源标识），支持客户端缓存（相同资源第二次请求返回 304 Not Modified），减少重复传输。|




#### cors中间件

```Go
package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  app := fiber.New()
  // 注册 CORS 中间件
  app.Use(cors.New(cors.Config{
    AllowOrigins:     "https://example.com,http://localhost:3000", // 允许的前端域名
    AllowMethods:     "GET,POST,PUT,DELETE,PATCH",                 // 允许的请求方法
    AllowHeaders:     "Content-Type,Authorization",                // 允许的请求头
    ExposeHeaders:    "Content-Length",                            // 允许前端读取的响应头
    AllowCredentials: true,                                        // 允许携带 Cookie
  }))

  app.Get("/api/data", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"data": "跨域请求成功"})
  })
  app.Listen(":80")
}

```



#### jwt中间件

1. 通过颁发jwt的接口，生成jwt
2. 后续请求需要认证的接口，把jwt带上，在统一的中间件中进行认证

```Go
package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/jwt/v3"
  "github.com/golang-jwt/jwt/v4"
  "time"
)

func main() {
  app := fiber.New()

  // 公开接口（无需认证）：登录生成 Token
  app.Post("/login", func(c *fiber.Ctx) error {
    // 模拟登录成功，生成 JWT Token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "username": "test",
      "exp":      jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
    })
    tokenString, _ := token.SignedString([]byte("your-secret-key")) // 密钥（生产环境需保密）
    return c.JSON(fiber.Map{"token": tokenString})
  })

  // 受保护接口：注册 JWT 中间件
  app.Use(jwtware.New(jwtware.Config{
    SigningKey: []byte("your-secret-key"),
  }))

  // 需认证的接口（只有携带有效 Token 才能访问）
  app.Get("/api/user", func(c *fiber.Ctx) error {
    // 从中间件中获取解析后的 Token 信息
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    return c.JSON(fiber.Map{"username": claims["username"]})
  })

  app.Listen(":80")
}

```



#### 限流中间件

```Go
package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/limiter"
  "time"
)

func GetIp(c *fiber.Ctx) string {
  return c.IP()
}

func main() {
  app := fiber.New()

  // 注册限流中间件
  app.Use(limiter.New(limiter.Config{
    KeyGenerator: GetIp,           // 按 IP 限流（也可按用户 ID 等自定义）
    Expiration:   5 * time.Second, // 限流窗口时间
    Max:          1,               // 窗口内最大请求数
    // 超过限流时的响应
    LimitReached: func(c *fiber.Ctx) error {
      return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
        "code": 429,
        "msg":  "请求过于频繁，请稍后再试",
      })
    },
  }))

  app.Get("/api/limit", func(c *fiber.Ctx) error {
    return c.SendString("限流接口请求成功")
  })
  app.Listen(":80")
}

```



### websocket支持

在 Fiber 中实现 WebSocket 路由需要借助官方的 `websocket` 中间件，该中间件会升级 HTTP 连接为 WebSocket 连接并处理握手逻辑

```Go
go get github.com/gofiber/websocket/v2
```



```Go
package main

import (
  "log"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/websocket/v2"
)

func main() {
  app := fiber.New()

  // 普通 HTTP 路由（用于测试）
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("请通过 WebSocket 连接 ws://localhost:80/ws")
  })

  // WebSocket 路由：使用 websocket.New 中间件包装处理函数
  app.Get("/ws", websocket.New(func(c *websocket.Conn) {
    // 连接成功后打印客户端信息
    log.Printf("新的 WebSocket 连接：%s", c.RemoteAddr())
    defer log.Printf("WebSocket 连接关闭：%s", c.RemoteAddr())

    // 循环读取客户端消息
    for {
      // 读取消息类型（text/binary）和内容
      msgType, msg, err := c.ReadMessage()
      if err != nil {
        log.Printf("读取消息失败：%v", err)
        break // 退出循环，关闭连接
      }

      // 打印收到的消息
      log.Printf("收到来自 %s 的消息：%s", c.RemoteAddr(), string(msg))

      // 向客户端回声（原样返回消息）
      if err := c.WriteMessage(msgType, msg); err != nil {
        log.Printf("发送消息失败：%v", err)
        break
      }
    }
  }))

  // 启动服务
  log.Fatal(app.Listen(":80"))
}



```



## fiber怎么配置

### 静态文件配置项

```Go
app.Static("/static", "./public", fiber.Static{
    Compress:  true,  // 启用静态文件压缩（节省带宽，适合 JS/CSS/HTML）
    Browse:    true,  // 启用目录浏览（访问 /static 时显示文件列表，类似 Nginx）
    Index:     "home.html", // 自定义索引文件（默认是 index.html）
    MaxAge:    86400, // 客户端缓存时间（1天，单位秒）
})
```



### fiber项目配置

```Go
app := fiber.New(fiber.Config{
    Prefork:       true,         // 启用多核（利用 CPU 多核心，提升性能）
    CaseSensitive: true,         // 路由大小写敏感（如 /User 和 /user 视为不同路由）
    StrictRouting: true,         // 严格路由（如 /user/ 和 /user 视为不同路由）
    AppName:       "MyFiberApp", // 应用名称 可以在视图的 c.App().Config().AppName获取到
    ErrorHandler: func(c *fiber.Ctx, err error) error { // 全局错误处理
      return c.Status(500).JSON(fiber.Map{
        "code": 500,
        "msg":  "服务器内部错误：" + err.Error(),
      })
    },
})
```



## 和gin的区别？

|功能点|Gin|Fiber|
|-|-|-|
|性能引擎|基于 `net/http`|基于 `fasthttp`（性能更高）|
|全局配置|配置项少（需手动扩展）|`fiber.Config` 内置多核、错误处理等（更灵活）|
|路由反向解析|需第三方库（如 `gin-contrib/cors`）|内置 `Name()` + `GetRouteURL()`（更方便）|
|静态文件配置|仅基础功能（无压缩 / 目录浏览）|支持压缩、目录浏览、自定义索引（更强大）|
|文件上传|需手动处理 `multipart.Form`|内置 `FormFile()` + `SaveFile()`（更简洁）|
|WebSocket|需第三方库（如 `gorilla/websocket`）|官方 `websocket` 中间件（开箱即用）|


