#!/bin/bash
# @Author: xiaofeng
# @Date:   2021-02-1 09:44:45
# @Last Modified by:   xiaofeng
# @Last Modified time: 2021-02-1 09:44:45

server_path=./
server_name=iponline

case $1 in 
        start)
                nohup $server_path$server_name >/dev/null 2>&1  &
                echo "$server_name服务已启动..."
                sleep 1
        ;;
        stop)
                killall $server_name >/dev/null 2>&1
                echo "$server_name服务已停止..."
                sleep 1
        ;;
        status)
                if [ $(pidof $server_name|wc -w) -eq 0 ] ;then
                        echo "$server_name服务未运行..."
                else
                        echo "$server_name服务正在运行..."
                fi
                sleep 1
        ;;
        restart)
                killall $server_name >/dev/null 2>&1
                sleep 1
                nohup $server_path$server_name >/dev/null 2>&1 &
                echo "$server_name服务已重启..."
                sleep 1
        ;;
        *) 
                echo "$0 {start|stop|status|restart}"
                exit 4
        ;;
esac
