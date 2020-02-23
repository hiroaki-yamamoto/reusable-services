#!/bin/sh
# -*- coding: utf-8 -*-

DIR=$(dirname $(realpath $0))
packagePath="github.com/hiroaki-yamamoto/render"

cat > $(dirname $0)/build.ninja << EOF
goFlags = --go_out=plugins=grpc
goOut = $(realpath $(dirname $DIR)/go/rpc --relative-to=$DIR)
pyOut = $(realpath $(dirname $DIR)/py/rpc --relative-to=$DIR)

inc = .

rule protoc
  command = protoc \$flags:\$outdir -I \$inc \$in

rule genPy
  command = $VIRTUAL_ENV/bin/python -m grpc_tools.protoc -I\$inc --python_out=\$outdir --grpc_python_out=\$outdir \$in

rule mv
  command = mv \$in \$out
EOF

for f in $(find $DIR -type f -name '*.proto'); do
fname=$(basename -s '.proto' ${f})
childDir=$(realpath $(dirname ${f}) --relative-to=$DIR)

cat >> $(dirname $0)/build.ninja << EOF
build tmp/${packagePath}/go/rpc/$childDir/${fname}.pb.go: protoc $(
  realpath $f --relative-to=$DIR
)
  flags = \$goFlags
  outdir = tmp

build \$goOut/$childDir/${fname}.pb.go: mv tmp/${packagePath}/go/rpc/$childDir/${fname}.pb.go
EOF

pyOutFiles=(
  "\$pyOut/$childDir/${fname}_pb2.py"
  "\$pyOut/$childDir/${fname}_pb2_grpc.py"
)
cat >> $(dirname $0)/build.ninja << EOF
build ${pyOutFiles[@]}: genPy $(realpath $f --relative-to=$DIR)
  flags = \$pyFlags
  outdir = \$pyOut
EOF
done
