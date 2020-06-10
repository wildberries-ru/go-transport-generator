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

// mod ...
type Mod struct {
}

func (m *Mod) PkgModPath(pkgName string) string {
	modPath, _ := m.goModPath(".")
	module, require := m.parseMod(modPath)
	if strings.HasPrefix(pkgName, module) {
		return "." + pkgName[len(module):]
	}
	// todo mod
	pkgTokens := strings.Split(pkgName, "/")
	for i := 0; i < len(pkgTokens); i++ {
		pathTry := strings.Join(pkgTokens[:len(pkgTokens)-i], "/")
		for modPkg, modPath := range require {
			if pathTry == modPkg {
				return path.Join(modPath, strings.Join(pkgTokens[len(pkgTokens)-i:], "/"))
			}
		}
	}
	return ""
}

func (m *Mod) parseMod(modPath string) (module string, require map[string]string) {
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
func (m *Mod) goModPath(root string) (string, error) {
	var stdout []byte
	var err error
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
				return "", err
			}
			root = r
			continue
		}
		return "", err
	}
	goModPath := string(bytes.TrimSpace(stdout))
	return goModPath, nil
}

// NewMod ...
func NewMod() *Mod {
	return &Mod{}
}
