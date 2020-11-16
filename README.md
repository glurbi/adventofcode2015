# run a test

`bazel test //day1:day1_test --test_output=all`

or

`go test github.com/glurbi/adventofcode2015/day1 -v`

# run all tests

`bazel test //... --test_output=all`

# update bazel BUILD files

`bazel run //:gazelle`
