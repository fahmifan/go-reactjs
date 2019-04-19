#!/bin/bash
go build -o output/simplerest ./cmd/simplerest
cd web && npm run build && cp -r build ../output/static