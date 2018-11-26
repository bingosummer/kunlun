// Code generated by "esc -pkg builtinmanifests -prefix  -ignore  -include  -o resources.go manifests"; DO NOT EDIT.

package builtinmanifests

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return []os.FileInfo(fis[0:limit]), nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/manifests/large_php.yml": {
		name:    "large_php.yml",
		local:   "manifests/large_php.yml",
		size:    3832,
		modtime: 1542944515,
		compressed: `
H4sIAAAAAAAC/+xW34vcNhB+918xXF6yZXfv9gih+K0tLQ0ECrkLfSjFyPKsrZwsuZqR9zZ/fZGstb0/
7pIrfWghBMJqZjQz/ub75rRarTKSDbYih/5mvckyNH1hRIs5PHijvckckvVOYlE767vke/36gnmxyLSV
gpU1IeLwe7HIMiUE5SA+e4fwCrhRBFIYKDHZrAOyLYLlBh10WvDWujbL+nbIT3kGsIKh+CffdqV9zAAA
WmSRx18AQye8705ipPWGc9jEAz34HO5YmEq4qvhxQ9E6XOrbk+b6NnTWt0RLkN45NKz3YI3eB5ciIN91
1jFWMYulonN2qzQeWhJVq8yI2XRaLFKAVsY/FtKaraq9G7BLLgCiZjoAdL7UShYPuKe5OQBzyE3UrKew
VIXYOlGPPal2dgCw2y26HD6W3rC/Q9ejy44qUhP8PwljjZJCT90FJDdv1zdvVu/v70Zzj44iA7RgJE52
S0Wl6GEq2wojaqyiNc1snMr7D1M6KWSjTJ3DBxTV704xTi6HgrGw3cC4X5xt34VvSwGVYBHTz9BafXXd
Zyuf1f657Xg/c8fspD5jUZc5bG5uRt+rROMw8x7NAJX3qppRr8RouY6EXoLaxuMSrBkcsLNeV4GgtUEn
GCsQFLPGMgZ5Z9302SsgXxrkxEMyyKvN2FCKLgild4r3RzonagpDdXZCQdXlQCxYyehxVuOs2nx9rMMU
AhrZTMA7LFcUmUYvFeg3nX3T2f9TZ9qKqiiFFkaiSyHHtqdiSyEf0FSFqCqHREVnrU4JLrq+Uts7LAcR
jgofdHwASBF4wgrYQo3D9wM3CBV22u5bNAwkneqYvqD8y96YZXhDSFtNM++FO5Jd5+wnlFyk50aILTrB
TRD6U75R9vMEteLCYa8ovU8u2ceLJ912TZdlvUE+eoj004QPA6BOSAxMXMd/15u3w2qIxDjbkScUccLU
88u3b0ZXLRh3Yj86N1l2RJKjvi7R6mitRssl7px1+AzBGhSam7CJy/Pt3zB3gyubzYGttDqHe9nBKxg0
LTT0QvvAuz/uZbeEX5nT//TndNc6zuH7ubwd/uWROHHh6voq2ZTDKkj5UC4+05ADjUPS8KCLydfwW3hs
7hThEhSHMGMZhNZ2h9X6dFLUPPc5J43e3g568voJZILnpcB8rLol/KD1hMrWWcNhOmfwHMb2pOPlu2Q+
7uJsyNnlVXNEy/OFc3bpImJxJKtQbAaZsuFCDreznV8ph3L4Q/HOlNabanQJKZEoD/jZ3RdmmLZJwK5I
krz67mqqgsTKxCfEUcwM5ZRhRNLhVj0+neVy3Jx6/wQwouYCXpv/DF5JJP8SXlkvnBJlQmNE4vBamz0i
w+nvAAAA//8nTRVE+A4AAA==
`,
	},

	"/manifests/maximum_php.yml": {
		name:    "maximum_php.yml",
		local:   "manifests/maximum_php.yml",
		size:    3832,
		modtime: 1542944515,
		compressed: `
H4sIAAAAAAAC/+xW34vcNhB+918xXF6yZXfv9gih+K0tLQ0ECrkLfSjFyPKsrZwsuZqR9zZ/fZGstb0/
7pIrfWghBMJqZjQz/ub75rRarTKSDbYih/5mvckyNH1hRIs5PHijvckckvVOYlE767vke/36gnmxyLSV
gpU1IeLwe7HIMiUE5SA+e4fwCrhRBFIYKDHZrAOyLYLlBh10WvDWujbL+nbIT3kGsIKh+CffdqV9zAAA
WmSRx18AQye8705ipPWGc9jEAz34HO5YmEq4qvhxQ9E6XOrbk+b6NnTWt0RLkN45NKz3YI3eB5ciIN91
1jFWMYulonN2qzQeWhJVq8yI2XRaLFKAVsY/FtKaraq9G7BLLgCiZjoAdL7UShYPuKe5OQBzyE3UrKew
VIXYOlGPPal2dgCw2y26HD6W3rC/Q9ejy44qUhP8PwljjZJCT90FJDdv1zdvVu/v70Zzj44iA7RgJE52
S0Wl6GEq2wojaqyiNc1snMr7D1M6KWSjTJ3DBxTV704xTi6HgrGw3cC4X5xt34VvSwGVYBHTz9BafXXd
Zyuf1f657Xg/c8fspD5jUZc5bG5uRt+rROMw8x7NAJX3qppRr8RouY6EXoLaxuMSrBkcsLNeV4GgtUEn
GCsQFLPGMgZ5Z9302SsgXxrkxEMyyKvN2FCKLgild4r3RzonagpDdXZCQdXlQCxYyehxVuOs2nx9rMMU
AhrZTMA7LFcUmUYvFeg3nX3T2f9TZ9qKqiiFFkaiSyHHtqdiSyEf0FSFqCqHREVnrU4JLrq+Uts7LAcR
jgofdHwASBF4wgrYQo3D9wM3CBV22u5bNAwkneqYvqD8y96YZXhDSFtNM++FO5Jd5+wnlFyk50aILTrB
TRD6U75R9vMEteLCYa8ovU8u2ceLJ912TZdlvUE+eoj004QPA6BOSAxMXMd/15u3w2qIxDjbkScUccLU
88u3b0ZXLRh3Yj86N1l2RJKjvi7R6mitRssl7px1+AzBGhSam7CJy/Pt3zB3gyubzYGttDqHe9nBKxg0
LTT0QvvAuz/uZbeEX5nT//TndNc6zuH7ubwd/uWROHHh6voq2ZTDKkj5UC4+05ADjUPS8KCLydfwW3hs
7hThEhSHMGMZhNZ2h9X6dFLUPPc5J43e3g568voJZILnpcB8rLol/KD1hMrWWcNhOmfwHMb2pOPlu2Q+
7uJsyNnlVXNEy/OFc3bpImJxJKtQbAaZsuFCDreznV8ph3L4Q/HOlNabanQJKZEoD/jZ3RdmmLZJwK5I
krz67mqqgsTKxCfEUcwM5ZRhRNLhVj0+neVy3Jx6/wQwouYCXpv/DF5JJP8SXlkvnBJlQmNE4vBamz0i
w+nvAAAA//8nTRVE+A4AAA==
`,
	},

	"/manifests/medium_php.yml": {
		name:    "medium_php.yml",
		local:   "manifests/medium_php.yml",
		size:    3832,
		modtime: 1542944515,
		compressed: `
H4sIAAAAAAAC/+xW34vcNhB+918xXF6yZXfv9gih+K0tLQ0ECrkLfSjFyPKsrZwsuZqR9zZ/fZGstb0/
7pIrfWghBMJqZjQz/ub75rRarTKSDbYih/5mvckyNH1hRIs5PHijvckckvVOYlE767vke/36gnmxyLSV
gpU1IeLwe7HIMiUE5SA+e4fwCrhRBFIYKDHZrAOyLYLlBh10WvDWujbL+nbIT3kGsIKh+CffdqV9zAAA
WmSRx18AQye8705ipPWGc9jEAz34HO5YmEq4qvhxQ9E6XOrbk+b6NnTWt0RLkN45NKz3YI3eB5ciIN91
1jFWMYulonN2qzQeWhJVq8yI2XRaLFKAVsY/FtKaraq9G7BLLgCiZjoAdL7UShYPuKe5OQBzyE3UrKew
VIXYOlGPPal2dgCw2y26HD6W3rC/Q9ejy44qUhP8PwljjZJCT90FJDdv1zdvVu/v70Zzj44iA7RgJE52
S0Wl6GEq2wojaqyiNc1snMr7D1M6KWSjTJ3DBxTV704xTi6HgrGw3cC4X5xt34VvSwGVYBHTz9BafXXd
Zyuf1f657Xg/c8fspD5jUZc5bG5uRt+rROMw8x7NAJX3qppRr8RouY6EXoLaxuMSrBkcsLNeV4GgtUEn
GCsQFLPGMgZ5Z9302SsgXxrkxEMyyKvN2FCKLgild4r3RzonagpDdXZCQdXlQCxYyehxVuOs2nx9rMMU
AhrZTMA7LFcUmUYvFeg3nX3T2f9TZ9qKqiiFFkaiSyHHtqdiSyEf0FSFqCqHREVnrU4JLrq+Uts7LAcR
jgofdHwASBF4wgrYQo3D9wM3CBV22u5bNAwkneqYvqD8y96YZXhDSFtNM++FO5Jd5+wnlFyk50aILTrB
TRD6U75R9vMEteLCYa8ovU8u2ceLJ912TZdlvUE+eoj004QPA6BOSAxMXMd/15u3w2qIxDjbkScUccLU
88u3b0ZXLRh3Yj86N1l2RJKjvi7R6mitRssl7px1+AzBGhSam7CJy/Pt3zB3gyubzYGttDqHe9nBKxg0
LTT0QvvAuz/uZbeEX5nT//TndNc6zuH7ubwd/uWROHHh6voq2ZTDKkj5UC4+05ADjUPS8KCLydfwW3hs
7hThEhSHMGMZhNZ2h9X6dFLUPPc5J43e3g568voJZILnpcB8rLol/KD1hMrWWcNhOmfwHMb2pOPlu2Q+
7uJsyNnlVXNEy/OFc3bpImJxJKtQbAaZsuFCDreznV8ph3L4Q/HOlNabanQJKZEoD/jZ3RdmmLZJwK5I
krz67mqqgsTKxCfEUcwM5ZRhRNLhVj0+neVy3Jx6/wQwouYCXpv/DF5JJP8SXlkvnBJlQmNE4vBamz0i
w+nvAAAA//8nTRVE+A4AAA==
`,
	},

	"/manifests/small_kubernetes.yml": {
		name:    "small_kubernetes.yml",
		local:   "manifests/small_kubernetes.yml",
		size:    3924,
		modtime: 1543198791,
		compressed: `
H4sIAAAAAAAC/+xWW6vbuBN/96cYTl56/iRpEkopfvuzF7ZQWOhp2YdlMbI8sdUjS17NyGn66RfJiu1c
Ti/LPuxCOXCI5jeaGf/molmtVhnJBluRQ79Zb7MMTV8Y0WIOj95obzKHZL2TWNTO+i5hz57dEN/fZ9pK
wcqaoHH6fX+fZUoIykF88g5hAdwoAikMlJhk1gHZFsFygw46LXhvXZtlfTvYpzwDWMHg/INvu9J+zAAA
WmSRx18AQyR87C50pPWGc9jGAz36HB5YmEq4qvjxYVv0uwgM9/r2Ir6+DcH1LdESpHcODesjWKOPAVIE
5LvOOsYqWrFUdM7ulcZTVKJqlRlpm07390lBK+M/FtKavaq9G+hLEABRMx0AOl9qJYtHPNJcHLg52SZq
1pNa8kJsnajHmFQ7OwDY/R5dDu9Lb9g/oOvRZWceqQn4D8JYo6TQU3SBzO3L9ebF6s27h1Hco6NYBFow
Eie5paJS9Di5bYURNVZRmtI2JubN28mcFLJRps7hLYrqN6cYJ8ihYCxsNxTdz862r8O3JYVKsIjmZ2yt
vtrvZz1f+f6p7fg4g6N1Up+wqMsctpvNiC1SJYec92gGqrxX1az0SoyS57Gml6D28bgEawYADtbrKhRo
bdAJxgoERavRjUE+WDd99grIlwY51SEZ5NV2DChpF4TSO8XHs1YnagpDdXZRgqrLgViwkhFxVuPM23yC
rEMWAhtZNodKdAYZ6csturtq0e+t9r3V/puttgBtRVWUQgsj0SWlc9nT2qWQj2iqQlSVQ6Kis1YnEzeh
r2zxqRnHTh/6+cSSIvCEFbCFGgcSgBuECjttjy0aBpJOdUxfmAA30YtZMGVtuh/9DAuHtBWOegC9cCP1
i6GNnP2Akou0nwT9ohPchJnwFDZOiHMTteLCYa8orTS35LOrF1F3TZdlvUE+2176qSBOuaJOSAyFu45/
z7cvh0kS6+hqql5UlBOmnl/evRihWjAexHEEt1l2Vk/R9BT2dRUuLqdxkt0qtmRsHumTNRnUGhSamzDC
S7xxt2HuBnCEYl7YSqtzeCc7WMAwEISGXmgf6vX3d7Jbwi/M6T/9Mb9tHefwajMTLcDhnx6JU4XcPb9L
MuWwCrPg5DKuesihBYLhsBRGB2v4NeysB0W4BMVBzVgGobU9YLW++q7wnH72s64C3u2SyHn9JFMB+3ai
3lfdEv6v9ZylvbOGQ9Zu0HVK6GegbxlNp7vzUiiu0p/dnl1nTXVjgl3dGumbkxfztArepmfYKRsu5LCb
vSSVciiH5+e1Ka031QgJKZEoD0Taw8zMeVLTMz4MnkBgkTr37n93kxckViYuJmc6rzaXFkYyHe7Vx6et
3NabV+PfIYyoucHX9l/D1273T/KV9cIpUSY2RiZOO+BsNw2nvwIAAP//Ps/0l1QPAAA=
`,
	},

	"/manifests/small_php.yml": {
		name:    "small_php.yml",
		local:   "manifests/small_php.yml",
		size:    3839,
		modtime: 1543198791,
		compressed: `
H4sIAAAAAAAC/+xWTYvkOA++51eI7svUS1V1VzMML7kt+8EODCxM97CHZQmOo0o87dhZS051za9f7LiS
1Ef3fLCHXRgamrIkS/KjR4pWq1VGssFW5NDfrjdZhqYvjGgxh0dvtDeZQ7LeSSxqZ32XdK9eXRAvFpm2
UrCyJlgcfi8WWaaEoBzEJ+8QroEbRSCFgRKTzDog2yJYbtBBpwVvrWuzrG8H/5RnACsYgn/0bVfapwwA
oEUWefwFMGTC++7ERlpvOIdNPNCjz+GehamEq4qf7jdFfxcVw72+Pcmvb0NyfUu0BOmdQ8N6D9bofVAp
AvJdZx1jFb1YKjpnt0rjIStRtcqMsE2nxSIZaGX8UyGt2arauwG+pAIgaqYDQOdLrWTxiHuaiwM2B99E
zXoyS1GIrRP1mJNqZwcAu92iy+FD6Q37e3Q9uuwoIjVB/6Mw1igp9JRdAHPzZn37evXu4X4U9+gokkAL
RuIkt1RUih6nsK0wosYqSlPZxsK8ez+5k0I2ytQ5vEdR/e4U46RyKBgL2w2k+8XZ9m14WzKoBIvofobW
6ovjvhj5LPbPbcf7mTp6J/UJi7rMYXN7O+quE5NDzXs0A1Teq2pGvRKj5CZyeglqG49LsGZQwM56XQWC
1gadYKxAUPQawxjknXXTs1dAvjTIiYdkkFebMaFkXRBK7xTvj1qdqCkM1dkJBVWXA7FgJaPGWY2zaPMJ
sg5VCGhksx7eYbmiyDT6lh793mvfe+2/2WvaiqoohRZGoksmx7LnbEshH9FUhagqh0RFZ61ODi6qvrC/
d1gOjTh2+dDLB4AUgSesgC3UOLwfuEGosNN236JhIOlUx/SZ7r+sjV6GVULaaqp5L9xR23XOfkTJRdo6
gm3RCW5Coz+nG9t+7qBWXDjsFaU15ZJ8vHiSbdd0WdYb5KN9pJ8qfCgAdUJiYOI6/t1s3gyjIRLjbE6e
UMQJU88v370eVbVg3In9qNxk2RFJjvK6RKuj0Roll7hzluELBGtQaG7CJC7PvwANczeoslkd2Eqrc3iQ
HVzD0NNCQy+0D7z740F2S/iVOf2nP6e71nEO/5+3t8O/PBInLlzdXCWZcliFVj6Ei6sacqBxcBqWuuh8
Db+FnXOnCJegOJgZyyC0tjus1qeVoual55wkejd8tZzXzyATNF8LzIeqW8IPWk+obJ01HKpzBs+hbM8q
vn6WzMtdnBU5uzxqjmh5PnDOLl1ELJZkFYLNIFM2XMjhbjbzK+VQDh+Kt6a03lSjSkiJRHnAz+4+U8M0
TQJ2RWrJq/9dTVGQWJm4QhzZzFBOHkYkHW7V0/NeLtvNqfctgBE1F/Da/GvwSk3yD+GV9cIpUSY0RiQO
29psiwynvwMAAP//GjIHz/8OAAA=
`,
	},

	"/manifests": {
		name:  "manifests",
		local: `manifests`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"manifests": {
		_escData["/manifests/large_php.yml"],
		_escData["/manifests/maximum_php.yml"],
		_escData["/manifests/medium_php.yml"],
		_escData["/manifests/small_kubernetes.yml"],
		_escData["/manifests/small_php.yml"],
	},
}
