modfiles=($(find . -type f -name go.mod))
for f in "${modfiles[@]}"; do
  pushd "$(dirname "$f")"
  go mod tidy
  popd
done
