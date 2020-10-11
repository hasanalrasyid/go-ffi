package purescript_record

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Record.Builder")

	exports["copyRecord"] = func(rec Any) Any {
		cpy := make(Dict)
		for key, value := range rec.(Dict) {
			cpy[key] = value
		}
		return cpy
	}

	exports["unsafeInsert"] = func(l_ Any) Any {
		return func(a Any) Any {
			return func(rec_ Any) Any {
				l, _ := l_.(string)
				rec, _ := rec_.(Dict)
				rec[l] = a
				return rec
			}
		}
	}
}
