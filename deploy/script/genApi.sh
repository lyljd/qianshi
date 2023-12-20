# Ensure to use in the root directory of .api

INPUT_PATH=""

while getopts "i:" opt; do
  case $opt in
    i)INPUT_PATH="$OPTARG";;
    \?)exit 1;;
  esac
done

if [ -z "$INPUT_PATH" ]; then
  echo "Usage: sh $0 -i <api file path>"
  exit 1
fi

echo "Ensure to use in the root directory of .api"
goctl api go --api "$INPUT_PATH" --dir ../ --style go_zero
