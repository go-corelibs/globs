#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += GLOBS
GLOBS_MK_SUMMARY := go-corelibs/globs
GLOBS_MK_VERSION := v1.0.0

include CoreLibs.mk
