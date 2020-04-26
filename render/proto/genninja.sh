#!/bin/sh
# -*- coding: utf-8 -*-

DIR=$(dirname $(realpath $0))
packagePath="render"

cat > $(dirname $0)/build.ninja << EOF
goFlags = --go_out=plugins=grpc
goOut = $(realpath $(dirname $DIR)/go/rpc --relative-to=$DIR)
pyOut = $(realpath $(dirname $DIR)/py/reusable_services/render/rpc --relative-to=$DIR)

inc = .

rule protoc
  command = protoc \$flags:\$outdir -I \$inc \$in

rule mock
  command = mockgen -package mocks -destination \$out -source \$in

rule genPy
  command = $(realpath -s $VIRTUAL_ENV/bin/python --relative-to=$DIR) -m grpc_tools.protoc -I\$inc --python_out=\$outdir --grpc_python_out=\$outdir \$in

rule getInitPy
  command = cp init.py \$out

rule genRPCInitPy
  command = cp initRPC.py \$out

build \$pyOut/__init__.py: genRPCInitPy
build \$pyOut/../__init__.py: getInitPy
build \$pyOut/../../__init__.py: getInitPy
EOF

cat >> $(dirname $0)/build.ninja << EOF
build \$goOut/mocks/services.go: mock \$goOut/services.pb.go
EOF

for f in $(find $DIR -type f -name '*.proto'); do
fname=$(basename -s '.proto' ${f})
childDir=$(realpath $(dirname ${f}) --relative-to=$DIR)

cat >> $(dirname $0)/build.ninja << EOF
build \$goOut/$childDir/${fname}.pb.go: protoc $(realpath $f --relative-to=$DIR)
  flags = \$goFlags
  outdir = \$goOut/$childDir
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

grep -qe "build \$pyOut/__init__.py: genRPCInitPy" $(dirname $0)/build.ninja
rootOK=${?}
grep -qe "build \$pyOut/$childDir/__init__.py: getInitPy" $(dirname $0)/build.ninja
childOK=${?}
if [ $rootOK -ne 0 -a $childOK -ne 0 ]; then
  echo build "\$pyOut/$childDir/__init__.py: getInitPy" >> $(dirname $0)/build.ninja
fi
done
