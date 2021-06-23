package utils

import (
	"testing"

	"github.com/tj/assert"
)

func TestParsePkg(t *testing.T) {

	// is.Equal("foo/bar:latest", reference.Name())
	// is.Equal("foo/bar", reference.ShortName())
	// is.Equal("latest", reference.Tag())
	// is.Equal("docker.io", reference.Registry())
	// is.Equal("docker.io/foo/bar", reference.Repository())
	// is.Equal("docker.io/foo/bar:latest", reference.Remote())
	test(t, "foo/bar", "", "foo/bar")
	test(t, "docker.io", "", "docker.io")
	test(t, "docker.io/foo/bar", "docker.io", "foo/bar")
	test(t, "docker.io/foo.bar", "docker.io", "foo.bar")

}

func test(t *testing.T, name string, domain, path string) {
	pkg, err := ParsePkg(name)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
	assert.Equal(t, pkg, &PkgName{Fullname: name, Domain: domain, Path: path})
}
