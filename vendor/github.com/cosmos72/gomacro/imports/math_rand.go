// this file was generated by gomacro command: import _b "math/rand"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"math/rand"
)

// reflection: allow interpreted code to import "math/rand"
func init() {
	Packages["math/rand"] = Package{
	Binds: map[string]Value{
		"ExpFloat64":	ValueOf(rand.ExpFloat64),
		"Float32":	ValueOf(rand.Float32),
		"Float64":	ValueOf(rand.Float64),
		"Int":	ValueOf(rand.Int),
		"Int31":	ValueOf(rand.Int31),
		"Int31n":	ValueOf(rand.Int31n),
		"Int63":	ValueOf(rand.Int63),
		"Int63n":	ValueOf(rand.Int63n),
		"Intn":	ValueOf(rand.Intn),
		"New":	ValueOf(rand.New),
		"NewSource":	ValueOf(rand.NewSource),
		"NewZipf":	ValueOf(rand.NewZipf),
		"NormFloat64":	ValueOf(rand.NormFloat64),
		"Perm":	ValueOf(rand.Perm),
		"Read":	ValueOf(rand.Read),
		"Seed":	ValueOf(rand.Seed),
		"Uint32":	ValueOf(rand.Uint32),
		"Uint64":	ValueOf(rand.Uint64),
	},Types: map[string]Type{
		"Rand":	TypeOf((*rand.Rand)(nil)).Elem(),
		"Source":	TypeOf((*rand.Source)(nil)).Elem(),
		"Source64":	TypeOf((*rand.Source64)(nil)).Elem(),
		"Zipf":	TypeOf((*rand.Zipf)(nil)).Elem(),
	},Proxies: map[string]Type{
		"Source":	TypeOf((*Source_math_rand)(nil)).Elem(),
		"Source64":	TypeOf((*Source64_math_rand)(nil)).Elem(),
	},
	}
}

// --------------- proxy for math/rand.Source ---------------
type Source_math_rand struct {
	Object	interface{}
	Int63_	func(interface{}) int64
	Seed_	func(_proxy_obj_ interface{}, seed int64) 
}
func (Proxy *Source_math_rand) Int63() int64 {
	return Proxy.Int63_(Proxy.Object)
}
func (Proxy *Source_math_rand) Seed(seed int64)  {
	Proxy.Seed_(Proxy.Object, seed)
}

// --------------- proxy for math/rand.Source64 ---------------
type Source64_math_rand struct {
	Object	interface{}
	Int63_	func(interface{}) int64
	Seed_	func(_proxy_obj_ interface{}, seed int64) 
	Uint64_	func(interface{}) uint64
}
func (Proxy *Source64_math_rand) Int63() int64 {
	return Proxy.Int63_(Proxy.Object)
}
func (Proxy *Source64_math_rand) Seed(seed int64)  {
	Proxy.Seed_(Proxy.Object, seed)
}
func (Proxy *Source64_math_rand) Uint64() uint64 {
	return Proxy.Uint64_(Proxy.Object)
}
