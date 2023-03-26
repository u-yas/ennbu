#!/bin/bash -l

set -xv

cmds="$INPUT_COMMANDS"

IFS=$'\n' read -d '' -r -a commands <<<"$cmds"

commands_length=${#commands[@]}

for ((i = 0; i < commands_length; i++)); do
  command=${commands[$i]}
  command=$(echo "$command" | sed -e 's/^[[:space:]]*//')
  eval 'ennbu $command'
done
