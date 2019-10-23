#!/bin/sh
#mpg123 --list playlist.m3u

mkfifo /tmp/mpg123in
#mkfifo /tmp/mpg123out

mpg123 --remote --fifo /tmp/mpg123in > /dev/null 2>&1 &
sleep 1
#ps ax | grep mpg123
#> /tmp/mpg123out &
echo "VOLUME 3" > /tmp/mpg123in
echo "LOADLIST 1 playlist.m3u" > /tmp/mpg123in
