package gen

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"golang.org/x/xerrors"

	"github.com/ogen-go/ogen/internal/ir"
)

type TemplateConfig struct {
	Package    string
	Operations []*ir.Operation
	Types      map[string]*ir.Type
	Interfaces map[string]*ir.Type
}

// FileSystem represents a directory of generated package.
type FileSystem interface {
	WriteFile(baseName string, source []byte) error
}

type writer struct {
	fs    FileSystem
	t     *template.Template
	buf   *bytes.Buffer
	wrote map[string]bool
}

// Generate executes template to file using config.
func (w *writer) Generate(templateName, fileName string, cfg TemplateConfig) error {
	if w.wrote[fileName] {
		return fmt.Errorf("name collision (already wrote %s)", fileName)
	}

	w.buf.Reset()
	if err := w.t.ExecuteTemplate(w.buf, templateName, cfg); err != nil {
		return fmt.Errorf("failed to execute template %s for %s: %w", templateName, fileName, err)
	}
	if err := w.fs.WriteFile(fileName, w.buf.Bytes()); err != nil {
		_ = os.WriteFile(fileName+".dump", w.buf.Bytes(), 0600)
		return fmt.Errorf("failed to write file %s: %w", fileName, err)
	}
	w.wrote[fileName] = true

	return nil
}

// WriteSource writes generated definitions to fs.
func (g *Generator) WriteSource(fs FileSystem, pkgName string) error {
	w := &writer{
		fs:    fs,
		t:     vendoredTemplates(),
		buf:   new(bytes.Buffer),
		wrote: map[string]bool{},
	}

	cfg := TemplateConfig{
		Package:    pkgName,
		Operations: g.operations,
		Types:      g.types,
		Interfaces: g.interfaces,
	}
	for _, name := range []string{
		"schemas",
		"schemas_json",
		"interfaces",
		"params",
		"param_dec",
		"handlers",
		"router",
		"req_enc",
		"req_dec",
		"res_enc",
		"res_dec",
		"validators",
		"server",
		"client",
		"cfg",
	} {
		fileName := fmt.Sprintf("oas_%s_gen.go", name)
		if err := w.Generate(name, fileName, cfg); err != nil {
			return xerrors.Errorf("%s: %w", name, err)
		}
	}

	return nil
}
