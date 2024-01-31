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
上传二进制文件（palworld\_chan\_amd64\_linux）到服务器并给到权限\
web ui使用办法还在写中

```bash
chmod 755 palworld_chan_amd64_linux
```

**启动API服务**

```bash
palworld-chan api
#此时会自动一个端口为3000的web服务器,浏览器打开就能看到了,使用手机访问最佳
```



**\[命令行]本地备份**

```bash
./palworld_chan_amd64_linux backup --source /home/docker-palworld-dedicated-server/game/Pal/Saved --dest /root/backup/
./palworld_chan_amd64_linux backup --source 需要备份的服务器目录 --dest 保存备份文件目录
#备份滚动功能 默认保存200个
./palworld_chan_amd64_linux backup --source 需要备份的服务器目录 --dest 保存备份文件目录 --backupCount 保留备份数
#建议配合定时任务一起使用 暂时仅支持linux
```

**\[命令行]使用七牛备份**

```bash
./palworld_chan_amd64_linux 7backup --source /home/docker-palworld-dedicated-server/game/Pal/Saved --dest /root/backup/  --ak accessToken --sk secretKey --bucket bucket存储空间
./palworld_chan_amd64_linux 7backup --source 需要备份的服务器目录 --dest 保存备份文件目录 --ak accessToken --sk secretKey --bucket bucket存储空间
#建议配合定时任务一起使用 暂时仅支持linux
```

**使用linux crontab自动备份**\
进入到服务器新建文件夹输入backup.sh的内容\
使用crontab -e\
进行计划任务编辑 输入

```bash
0 * * * * /脚本路径/backup.sh
```

\:wq!\
保存并退出

**Todo：**

*   英文，多语言支持
*   等待官方服务端更新以便加新功能

**已知问题**

*   1.由于帕鲁服务端BUG，rcon无法发送中文且rcon发送的文本中无法保留空格，已自动替换为下划线。（仅只能英文下划线，日文测了也不行，估计是编码问题）

**程序实现以及引用开源项目**

*   程序基于go语言开发，使用了fiber作为web部分框架，数据持久层使用的nutsdb，前端部分使用vant+vue3+pinia实现，项目脚手架使用的（<https://github.com/xiangshu233/vue3-vant4-mobile>），cron定时任务使用的（github.com/robfig/cron/v3），go部分实现代码已全部开源。

**软件截图**\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/index.png)\
首页\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/tool.png)\
工具\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/setting.png?raw=true)\
设置\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/backup.png?raw=true)\
备份页\
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/server_setting.png?raw=true)\
服务器设定
