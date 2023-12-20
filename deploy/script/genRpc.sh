# Ensure to use in the root directory of .proto

INPUT_PATH=""

while getopts "i:" opt; do
  case $opt in
    i)INPUT_PATH="$OPTARG";;
    \?)exit 1;;
  esac
done

if [ -z "$INPUT_PATH" ]; then
  echo "Usage: sh $0 -i <proto file path>"
  exit 1
fi

echo "Ensure to use in the root directory of .proto"
goctl rpc protoc "$INPUT_PATH" --go_out=. --go-grpc_out=. --zrpc_out=.. --style go_zero
