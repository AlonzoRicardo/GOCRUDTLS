#!/bin/bash

for f in ./scripts/*; do source $f; grep -E '^[[:space:]]*([[:alnum:]_]+[[:space:]]*\(\)|function[[:space:]]+[[:alnum:]_]+)' $f; done
