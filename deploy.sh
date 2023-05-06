go test ./... -count=1
v=v1.1.0
git tag $v
git push --tags
go install github.com/ymzuiku/option@$v
echo "done."