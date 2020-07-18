package mod

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

type Mod interface {
	PkgModPath(pkgName string) (mod string)
}

// mod ...
type mod struct {
}

// PkgModPath ...
func (m *mod) PkgModPath(pkgName string) (mod string) {
	modPath, _ := m.goModPath(".")
	module, require := m.parseMod(modPath)
	if strings.HasPrefix(pkgName, module) {
		mod = "." + pkgName[len(module):]
		return
	}
	// todo mod
	pkgTokens := strings.Split(pkgName, "/")
	for i := 0; i < len(pkgTokens); i++ {
		pathTry := strings.Join(pkgTokens[:len(pkgTokens)-i], "/")
		for modPkg, modPath := range require {
			if pathTry == modPkg {
				mod = path.Join(modPath, strings.Join(pkgTokens[len(pkgTokens)-i:], "/"))
				return
			}
		}
	}
	return
}

func (m *mod) parseMod(modPath string) (module string, require map[string]string) {
	var (
		err error
		b   []byte
	)
	if b, err = ioutil.ReadFile(modPath); err != nil {
		return
	}
	mod, err := modfile.Parse(modPath, b, nil)
	if err != nil {
		return
	}
	module = mod.Module.Mod.Path
	goPath := os.Getenv("GOPATH")
	require = make(map[string]string)
	for _, r := range mod.Require {
		require[r.Syntax.Token[0]] = fmt.Sprintf("%s/pkg/mod/%s@%s", goPath, r.Syntax.Token[0], r.Mod.Version)
	}
	return
}

// empty if no go.mod, GO111MODULE=off or go without go modules support
func (m *mod) goModPath(root string) (goModPath string, err error) {
	var stdout []byte
	for {
		cmd := exec.Command("go", "env", "GOMOD")
		cmd.Dir = root
		stdout, err = cmd.Output()
		if err == nil {
			break
		}
		if _, ok := err.(*os.PathError); ok {
			// try to find go.mod on level higher
			r := filepath.Join(root, "..")
			if r == root { // when we in root directory stop trying
				return
			}
			root = r
			continue
		}
		return
	}
	goModPath = string(bytes.TrimSpace(stdout))
	return
}

// NewMod ...
func NewMod() Mod {
	return &mod{}
}
