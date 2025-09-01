#!/bin/bash

protoc \
--proto_path=. \
--proto_path=../third_party \
--proto_path=../../ \
--go_out=. \
--go_opt=paths=source_relative \
--goose_out=. \
--goose_opt=paths=source_relative \
*/*.proto
