#!/usr/bin/env bash
# bash

echo "go test: github.com/d3ta-go/ddd-mod-email/modules/email/infrastructure/repository... "
echo "-------------------------------------------------------------------------------"
echo ""

go test -timeout 120s  github.com/d3ta-go/ddd-mod-email/modules/email/infrastructure/repository -v -cover

echo ""
echo "-------------------------------------------------------------------------------"
echo "go test: DONE "
echo ""