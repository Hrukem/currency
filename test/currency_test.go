package test

import (
	"currency/request"
	"reflect"
	"sort"
	"testing"
)

func TestParseRequest(t *testing.T) {
	ss := []string{
		"USDCAD:1.282720",
		"USDCHF:0.917500",
		"USDEUR:0.855060",
		"USDGBP:0.734072",
		"USDRUB:74.229920",
	}

	m1 := map[string]interface{}{
		"USDCAD": 1.28272,
		"USDCHF":0.9175,
		"USDEUR":0.85506,
		"USDGBP":0.734072,
		"USDRUB":74.22992,
	}

	m := map[string]interface{}{
		"quotes": m1,
	}

	res := request.ParseRequest(m)

	sort.Strings(res)
	sort.Strings(ss)

	if !(reflect.DeepEqual(res, ss)){
		t.Error("Expected ", ss, "got ", res)
	}
}
