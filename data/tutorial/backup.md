
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