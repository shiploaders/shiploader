APPS=(
  nodejs-1
  nnodjs-2
  pythonapp
  javaapp
)

declare -A IMAGES=(
  [nodejs-1]="ïmagerepository/nodejs-n"
  [nodejs-2]="imagerepository/nodejs"
  [pythonapp]="imagerepository/pythonapp"
  [javaapp]="anotherimagerepository/javaapp"
)

declare -A REPLICAS=(
  [nodejs-1]="2"
  [nodejs-2]="2"
  [pythonapp]="1"
  [javaapp]="3"
)

declare -A PORTS=(
  [nodejs-1]="3000"
  [nodejs-2]="3000"
  [pythonapp]="5000"
  [javaapp]="8080"
)
