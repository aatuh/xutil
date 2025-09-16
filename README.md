# xutil

Small, generic helpers: map utilities, pointer helper, simple
reflection-based tag scans, and slice dedup.

## Install

```go
import "github.com/aatuh/xutil"
```

## Quick start

```go
// Maps
m := map[string]int{"a": 1, "b": 2}
keys := util.MapKeys(m)          // ["a" "b"]
vals := util.MapValues(m)        // [1 2]
ks, vs := util.MapKeysAndValues(m)

// Pointer helper
p := util.Ptr(42)                // *int

// Dedup
u := util.DedupSlice([]int{1,1,2,3,2}) // [1 2 3]

// Tag scans
type T struct {
  ID   int    `db:"id" json:"id"`
  Name string `db:"name" json:"name"`
}
_ = util.FindFieldsByTag(T{}, "db", "id")       // ["ID"]
_ = util.FindFieldsByJSONTag(T{}, "db", "name") // ["name"]
```

## Notes

- Utilities prefer clarity and minimal surprises.
- Reflection helpers operate on concrete types; pass pointers or values.
