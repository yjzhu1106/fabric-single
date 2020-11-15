#!/bin/bash

# Exit on first error
set -e

echo "=================== start ==================="
echo start at：$(date +%Y-%m-%d\ %H:%M:%S)
echo "============================================="
./cc.sh install apmc 1.0 go/apmc Synchro
./cc.sh install dacc 1.0 go/dacc Synchro
./cc.sh install pdcc 1.0 go/pdcc Synchro
./cc.sh install vlrc 1.0 go/vlrc Synchro
# ./cc.sh install pdpc 1.0 go/pdpc Synchro
echo "=================== end ==================="
echo end at: $(date +%Y-%m-%d\ %H:%M:%S)
echo "==========================================="
