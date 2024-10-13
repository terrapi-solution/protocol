PROTO_NAMES=(
    "activity"
    "deployment"
)

for name in "${PROTO_NAMES[@]}"; do
  protoc ${name}/${name}.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.
  if [ $? -ne 0 ]; then
      echo "error processing ${name}.proto"
      exit $?
  fi
done