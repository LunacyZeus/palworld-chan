**帕鲁酱（pal-chan）**\
游戏幻兽帕鲁（Palworld）服务端面板(windows/linux) 服务端存档备份/上传云存储/RCON服务器管理

本工具基于golang+vue3+vant+wails开发 支持多个平台(windows/linux)  
接入前端中...  

**目前功能**\
文档等等有空写 画的饼基本完成了反正  
1、命令行设置服务器目录 并自动备份 打包并保存为时间格式的zip文件\
2、自动备份到七牛存储空间 防止炸服找不到存档

**使用方法**\
上传二进制文件（palworld\_chan\_amd64\_linux）到服务器并给到权限

```bash
chmod 755 palworld_chan_amd64_linux
```

**本地备份**

```bash
./palworld_chan_amd64_linux backup --source /home/docker-palworld-dedicated-server/game/Pal/Saved --dest /root/backup/
./palworld_chan_amd64_linux backup --source 需要备份的服务器目录 --dest 保存备份文件目录
#备份滚动功能 默认保存200个
./palworld_chan_amd64_linux backup --source 需要备份的服务器目录 --dest 保存备份文件目录 --backupCount 保留备份数
#建议配合定时任务一起使用 暂时仅支持linux
```

**使用七牛备份**

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

*   上传到七牛备份 √
*   RCON操作服务器
*   服务器内存过大自动重启 并发送公告
*   WEB GUI支持（选择性开启）
*   存档管理
*   定时功能
*   接口账号密码认证


**软件截图**  
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/index.png)  
首页  
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/tool.png)  
工具  
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/setting.png?raw=true)  
设置  
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/backup.png?raw=true)  
备份页  
![image](https://github.com/LunacyZeus/palworld-chan/blob/main/screenshots/server_setting.png?raw=true)  
服务器设定  