#!/bin/bash

env GOOS=linux GOARCH=amd64 go build -o dist/MailHedgehog-amd64-linux .
env GOOS=darwin GOARCH=amd64 go build -o dist/MailHedgehog-amd64-darwin .
env GOOS=windows GOARCH=amd64 go build -o dist/MailHedgehog-amd64-windows.exe .

rm -f MailHedgehog.zip
zip -r MailHedgehog.zip dist/
