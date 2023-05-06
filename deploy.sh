v=v1.0.6
git tag $v
git push --tags
go install github.com/ymzuiku/option@$v
echo "done."