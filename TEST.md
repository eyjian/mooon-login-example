通过网关执行以下请求：

```shell
# 正确的请求
curl -i -X POST -H "Content-Type: application/json"  -d '{"username":"mooon","password":"123456789a"}' 'http://127.0.0.1:6688/v1/login'

# 密码错误请求
curl -i -X POST -H "Content-Type: application/json"  -d '{"username":"mooon","password":"123456789x"}' 'http://127.0.0.1:6688/v1/login'

# 用户不存在请求
curl -i -X POST -H "Content-Type: application/json"  -d '{"username":"moon","password":"123456789a"}' 'http://127.0.0.1:6688/v1/login'

# 无效的请求
curl -i -X POST -H "Content-Type: application/json" 'http://127.0.0.1:6688/v1/login'
curl -i -X POST -H "Content-Type: application/json"  -d '{"user":"mooon","password":"123456789"}' 'http://127.0.0.1:6688/v1/login'

```