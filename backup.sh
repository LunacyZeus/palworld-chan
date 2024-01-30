#!/bin/bash

# 设置备份源和目标路径
SOURCE_PATH="/home/docker-palworld-dedicated-server/game/Pal/Saved"
DEST_PATH="/root/backup/"

# 执行备份命令
/root/palworld_chan_amd64_linux backup --source "$SOURCE_PATH" --dest "$DEST_PATH"

