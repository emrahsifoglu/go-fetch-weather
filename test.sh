#!/usr/bin/env sh

set -e

echo "=== Test GET ==="
curl -X GET -s localhost:9090/alerts?area=NY | rg -q "alerts" 

echo "Success"
