if [ -f dev.env ]
then
  export $(cat dev.env | sed 's/#.*//g' | xargs)
  go run ./cmd/.
  # go run ./internal/repository/test/.
fi