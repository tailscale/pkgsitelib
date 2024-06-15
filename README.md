# pkgsitelib

pkgsitelib is a fork of [golang.org/x/pkgsite] that is intended to be used as a library for other projects.

pkgsite is built very explicitly to host the [pkg.go.dev] website, and as such,
it doesn't make it super easy to [run locally] (though that is gradually improving).
It is also not easy to customize or pull into another project as a library.
For example, almost the entire codebase is within an "internal" package,
which makes it impossible to import into another project.

pkgsitelib only does a few things, most notably of which is to eliminate the top "internal" package.
There are still several "internal" subpackages, but those have been left alone for now.
It also includes a couple of minor changes to make it easier to use or customize.
Ideally, this project will become obsolete and go away as pkgsite becomes more modular.

This is used internally at Tailscale to host documentation for our private Go packages.

[golang.org/x/pkgsite]: (https://pkg.go.dev/golang.org/x/pkgsite)
[pkg.go.dev]: (https://pkg.go.dev)
[run locally]: https://github.com/golang/go/issues/40371
