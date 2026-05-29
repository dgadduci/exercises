#!/usr/bin/env bash

cd grammars
antlr -Dlanguage=Go -visitor -o ../parser Test.g4