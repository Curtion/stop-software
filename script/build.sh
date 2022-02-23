#!/bin/bash
rm -rf ./bin/*
cp config.json ./bin/
go build
mv stopSoftware.exe ./bin/stopSoftware.exe
go build -ldflags "-H windowsgui"
mv stopSoftware.exe ./bin/stopSoftware_hidden.exe
tar -czf ./bin/stopSoftware.tar.gz ./bin/stopSoftware.exe ./bin/config.json
tar -czf ./bin/stopSoftware_hidden.tar.gz ./bin/stopSoftware_hidden.exe ./bin/config.json
