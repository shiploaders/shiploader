#!/usr/bin/env bash
[ "${BASH_VERSINFO:-0}" -ge 4 ] || { echo "Upgrade your bash to at least version 4"; exit 1; }

source config.sh

for APP in ${APPS[@]}
do
  IMAGE=${IMAGES[$APP]}
  REPLICA=${REPLICAS[$APP]}
  PORT=${PORTS[$APP]}
  eval "echo \"$(cat k8s/deployment.yaml)\"" 
done

