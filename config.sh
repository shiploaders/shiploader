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
