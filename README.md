# File Admin

简易的文件管理，支持上传，下载

预览

![](https://cdn.jsdelivr.net/gh/mouday/img/2024/05/31/ksocfvg.png)

使用示例

```python
# -*- coding: utf-8 -*-
import requests


def main():
    url = 'http://127.0.0.1:8088/api/upload'

    files = {
        "file": open("demo.txt", "rb")
       }

    res = requests.post(url=url, files=files)
    if res.ok:
        print(res.json())


if __name__ == '__main__':
    main()

```

返回结果
```js   
{
  'code': 0,
  'data': {
  'fileUrl': 'http://127.0.0.1:8088/upload/2024/05/31/nDbw8KLM1N2Bv7v9hbsro.txt'
   },
  'msg': 'success'
}
```

本项目采用前后端分离的架构，前端代码使用React，如需获取前端代码，关注公众号：

![](https://mouday.github.io/img/2024/06/06/spcnbk3.png)

回复：`file-admin-web`