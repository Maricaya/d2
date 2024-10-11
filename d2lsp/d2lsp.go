// d2lsp contains functions useful for IDE clients
package d2lsp

import (
	"fmt"
	"strings"

	"oss.terrastruct.com/d2/d2ir"
	"oss.terrastruct.com/d2/d2parser"
	"oss.terrastruct.com/d2/lib/memfs"
)

func GetFieldRefs(path, index string, fs map[string]string, key string) (refs []d2ir.Reference, _ error) {
	if _, ok := fs[index]; !ok {
		return nil, fmt.Errorf(`"%s" not found`, index)
	}
	r := strings.NewReader(fs[index])
	ast, err := d2parser.Parse(path, r, nil)
	if err != nil {
		return nil, err
	}

	mfs, err := memfs.New(fs)
	if err != nil {
		return nil, err
	}

	mk, err := d2parser.ParseMapKey(key)
	if err != nil {
		return nil, err
	}
	if mk.Key == nil {
		return nil, fmt.Errorf(`"%s" is invalid`, key)
	}

	ir, _, err := d2ir.Compile(ast, &d2ir.CompileOptions{
		FS: mfs,
	})
	if err != nil {
		return nil, err
	}

	var f *d2ir.Field
	curr := ir
	for _, p := range mk.Key.Path {
		f = curr.GetField(p.Unbox().ScalarString())
		if f == nil {
			return nil, nil
		}
		curr = f.Map()
	}
	for _, ref := range f.References {
		refs = append(refs, ref)
	}
	return refs, nil
}
