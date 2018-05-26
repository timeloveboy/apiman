# apiman

apiman是一个静态网站服务器，但是提供了html复合封装，可用最大限度的减少html代码中的重复部分

## 启动参数

+ -port=8080

+ -root=/webfolder

## html用法

+ nav.html
```
<p>i am nav</p>
 ```


+ index.html可用将nav部分封装进
 ```
...
<!--{{LoadTemplate(nav.html)}}-->
...
 ```

