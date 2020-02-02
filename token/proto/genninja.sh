#!/bin/sh
# -*- coding: utf-8 -*-

DIR=$(dirname $(realpath $0))
packagePath="github.com/hiroaki-yamamoto/token/rpc"

cat > $(dirname $0)/build.ninja << EOF
backFlags = --go_out=plugins=grpc
backOut = $(
  realpath $(dirname $DIR)/rpc --relative-to=$DIR
)
inc = .

rule protoc
  command = protoc \$flags:\$outdir -I \$inc \$in

rule mv
  command = mv \$in \$out
EOF

for f in $(find $DIR -type f -name '*.proto'); do
fname=$(basename -s '.proto' ${f})
childDir=$(realpath $(dirname ${f}) --relative-to=$DIR)

cat >> $(dirname $0)/build.ninja << EOF

build tmp/${packagePath}/$childDir/${fname}.pb.go: protoc $(
  realpath $f --relative-to=$DIR
)
  flags = \$backFlags
  outdir = tmp

build \$backOut/$childDir/${fname}.pb.go: mv tmp/${packagePath}/$childDir/${fname}.pb.go

$frontSc
EOF
done
