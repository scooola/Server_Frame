# Server_Frame

A simple server-side framework.

## Request Example

### Health Check

**Method:** Get, Post

**Request**

```python
import requests

url = "http://127.0.0.1:12138/healch_check"
# Get
res = requests.get(url)
# Post
res = requests.post(url)
```

**Response**

```python
{"status":"ok", "code":0}
```

### Hello World

**Method:** Post

**Request**

```python
import requests
url = "http://127.0.0.1:12138/frontend"
# must have ActionId param
data = {"Name": "scola", "ActionId": 10000}
res = requests.post(url, json=data)
```

**Reqponse**

```python
{"hello":"world scola", "code":0}
```

## ServerStatus

```go
//Server_Frame\share\enum\error_code.go
const (
	Success = 0
	Failed  = 1
)

//e.p.:
//response success:
{"code":0}
//response failed:
{"code": 1}
```