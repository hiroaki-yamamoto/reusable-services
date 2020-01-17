#!/bin/sh -e
# -*- coding: utf-8 -*-

go mod download
exec go test -coverprofile="./${PKGNAME//\//-}.cov" ./${PKGNAME}/...
