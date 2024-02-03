**帕鲁酱（pal-chan）**\
通过web可视化管理游戏(幻兽帕鲁 / PalWorld)专用服务器，支持(Linux/Windows/MacOS)，自动本地存档备份/七牛云存储备份/坚果云备份(后续)，基于RCON实现的玩家列表/踢出/封禁/游戏内广播。 &#x20;

**基本功能**

*   优雅重启（内存占用达到阈值后自动发送公告并关服等待重启）
*   自动备份存档（本地自动备份存档/七牛云端备份）
*   存档备份管理（无需ssh，可直接在线下载对应存档）
*   在线可视化编辑服务器配置文档(.ini)
*   Webhook通知推送(企业微信机器人，未来会接入其他的钉钉或者pushplus)
*   服务器/游戏服务端监控（支持监控CPU/内存/启动时间/硬盘/网络IO）
*   内存监控（自定义阈值触发广播和推送）
*   进程守护（推荐用脚本实现，虽然写了，暂时没测好）
*   轮询获取在线玩家（定时任务获取当前服务器在线人数）
*   提供GUI（wails编译的版本以及WEB UI版本，同时支持仅命令行终端运行）
*   Rcon指令（获取服务器信息、实时玩家列表、踢出/封禁玩家、服务器广播、平滑关闭服务器并广播）
*   RCON 命令官方文档地址（<https://tech.palworldgame.com/server-commands）>

**注意!!!**

有问题可以提交issues给我,不怎么会用github这边的action,还在摸索中,今晚应该可以把自动发releases肝出来 &#x20;



**使用方法  **\
下载并上传二进制文件（palworld\_chan\_amd64\_linux）到服务器并给到权限\

```bash
chmod 755 palworld_chan_amd64_linux
```

**[部署教程]启动API服务**

```bash
palworld-chan api --port=:3001 --authUser=pal --authPassWd=1234
#此时会自动一个端口为3001的web服务器,账号为pal 密码为1234,浏览器打开就能看到了,使用手机访问最佳
#windows启动也是一样的你需要在命令行输入palworld_chan.exe --port=:3001 --authUser=pal --authPassWd=1234
```

**Todo：**

*   英文，多语言支持
*   等待官方服务端更新以便加新功能

**已知问题**

*   1.由于帕鲁服务端BUG，rcon无法发送中文且rcon发送的文本中无法保留空格，已自动替换为下划线。（仅只能英文下划线，日文测了也不行，估计是编码问题）

**程序实现以及引用开源项目**

*   程序基于go语言开发，使用了fiber作为web部分框架，数据持久层使用的nutsdb，前端部分使用vant+vue3+pinia实现，项目脚手架使用的（<https://github.com/xiangshu233/vue3-vant4-mobile>），cron定时任务使用的（github.com/robfig/cron/v3），go部分实现代码已全部开源。
*   配置生成代码部分借鉴于https://github.com/Bluefissure/pal-conf/
* 
**软件截图**\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/1.png?raw=true)\
1\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/2.png?raw=true)\
2\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/3.png?raw=true)\
3\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/4.png?raw=true)\
4\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/5.png?raw=true)\
5\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/6.png?raw=true)\
6\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/7.png?raw=true)\
7\
