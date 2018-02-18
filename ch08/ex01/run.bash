#!/usr/bin/env bash

TZ=US/Eastern ./clock2 -port 8010 &
TZ=Asia/Tokyo ./clock2 -port 8020 &
TZ=Europa/London ./clock2 -port 8030 &
./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
