// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/typeutil"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package typeutil

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/typeutil"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/typeutil"] = imports.Package{
		Binds: map[string]r.Value{
			"Identical":           r.ValueOf(Identical),
			"IdenticalIgnoreTags": r.ValueOf(IdenticalIgnoreTags),
			"MakeHasher":          r.ValueOf(MakeHasher),
		},
		Types: map[string]r.Type{
			"Hasher": r.TypeOf((*Hasher)(nil)).Elem(),
			"Map":    r.TypeOf((*Map)(nil)).Elem(),
		},
		Proxies: map[string]r.Type{}}
}
