#!/bin/bash
if golangci-lint \
	run ./... \
	--timeout=5m \
	--out-format colored-line-number \
	--skip-dirs-use-default; then
	echo "OK"
else
	echo "FAILED"
fi

