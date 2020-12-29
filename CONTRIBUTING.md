# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change or assign yourself to a card on the [Project's board](https://github.com/TibebeJS/yenepay.sdk.go/projects/1)

We'd love to accept your patches and contributions to this project.

## Pull Request Process

1. Ensure any install or build dependencies are removed before the end of the layer when doing a 
   build.
2. Update the README.md with details of changes to the interface, this includes new environment 
   variables, exposed ports, useful file locations and container parameters.
3. Increase the version numbers in any examples files and the README.md to the new version that this
   Pull Request would represent. The versioning scheme we use is [SemVer](http://semver.org/).
4. You may merge the Pull Request in once you have the sign-off of two other developers, or if you 
   do not have permission to do that, you may request the second reviewer to merge it for you.


## Tasks #

You can find the project's Canban board (Todo, Tasks in Progress, etc..) here:\
[Project Board](https://github.com/TibebeJS/yenepay.sdk.go/projects/1)

## Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Testing ##
Before sending a pull-request, please make sure your code is well-tested and linted locally.
To run the tests, use the following command:
```
$ go test -v ./...
```

Alternatively, you may use [`goconvey`](https://github.com/smartystreets/goconvey) (Recommended)

To install [`goconvey`](https://github.com/smartystreets/goconvey), run:
```
$ go get github.com/smartystreets/goconvey
```

After installation, use the following to execute the tests:
```
$ goconvey
```

There are also GitHub Actions/Workflows on the repo to assist with these, make sure your changes do pass these workflows.
