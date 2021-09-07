#!/bin/bash
set -x
curl http://localhost:8080/time
echo
curl http://localhost:8080/time?tz=UTC
echo
curl http://localhost:8080/time?tz=Local
echo
curl http://localhost:8080/time?tz=Asia/Tokyo
echo
