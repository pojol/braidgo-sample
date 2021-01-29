## BraidSample
`braid@v1.1.13`
> 使用braid微服务构建的框架，可适用于网站，应用，游戏等业务的后端。

[![image.png](https://i.postimg.cc/QtrF7jsR/image.png)](https://postimg.cc/JyP7VVsq)

> 开启并发机器人
```shell
./bot -num 1000
```

> 开启含有生命周期的机器人
> `每间隔1秒生成100个机器人，生命周期是5秒。`
```shell
./bot -num 100 -lifetime 5 -increase true
```

> **注1** 这里只是使用braid来进行构建一个基础的框架，用户可以再此基础上进行任意的自定义调整。如果业务量级小可以合并 login,base,mail 等功能节点，如果业务量大也可以从此基础上进行更多的拆分。

> **注2** 这里为了方便所以将工程都聚合到了一个目录中，正式开发环境中建议分开管理。

> **注3** 如果有一些例如游戏业务，需要高频同步，也可以支持tcp/udp服务。