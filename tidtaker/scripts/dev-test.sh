#!/usr/bin/env bash
set -euo pipefail

go run . superuser upsert test@example.com testpassword123 --dir=./pb_test_data
go run . serve --http=127.0.0.1:8091 --dir=./pb_test_data
