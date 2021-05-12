#!/usr/bin/env bash
[ "${BASH_VERSINFO:-0}" -ge 4 ] || { echo "Upgrade your bash to at least version 4"; exit 1; }


APPS=(
  nodejs
  pythonapp
  javaapp
)

declare -A IMAGES=(
  [nodejs]="imagerepository/nodejs"
  [pythonapp]="imagerepository/pythonapp"
  [javaapp]="anotherimagerepository/javaapp"
)

declare -A REPLICAS=(
  [nodejs]="2"
  [pythonapp]="1"
  [javaapp]="3"
)

declare -A PORTS=(
  [nodejs]="3000"
  [pythonapp]="5000"
  [javaapp]="8080"
)

for APP in ${APPS[@]}
do
  IMAGE=${IMAGES[$APP]}
  REPLICA=${REPLICAS[$APP]}
  PORT=${PORTS[$APP]}
  eval "echo \"$(cat k8s/deployment.yaml)\"" 
done

