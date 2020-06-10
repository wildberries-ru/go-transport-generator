package imports

import (
	"os/exec"
)

// Imports ...
type Imports struct {
}

func (fw *Imports) GoImports(path string) (err error) {
	var execPath string
	if execPath, err = exec.LookPath("goimports"); err != nil {
		return
	}
	err = exec.Command(execPath, "-w", path).Run()
	return
}

// NewImports ...
func NewImports() *Imports {
	return &Imports{}
}
