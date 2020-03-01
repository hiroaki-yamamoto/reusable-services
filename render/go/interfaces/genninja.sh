#!/usr/bin/env sh
# -*- coding: utf-8 -*-

cat > build.ninja << EOF
rule mock
  command = mockgen -package mocks -source \$in > \$out
EOF

for f in *.go; do
cat >> build.ninja << EOF
build mocks/template.go: mock ${f}
EOF
done
