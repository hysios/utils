package response

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/fatih/structs"
	"github.com/hysios/utils/errors"
	"github.com/iancoleman/strcase"
)

var (
	rStatus  = reflect.ValueOf("status")
	rSuccess = reflect.ValueOf("success")
)

func Jsonify(w http.ResponseWriter, val interface{}) {
	var (
		b   []byte
		v   = reflect.ValueOf(val)
		err error
	)
	v = reflect.Indirect(v)

	if v.Kind() == reflect.Map {
		if v.CanSet() {
			v.SetMapIndex(rStatus, rSuccess)
			b, err = json.Marshal(val)
			if err != nil {
				panic(err)
			}
		} else {
			panic("can't set map")
		}
	} else if v.Kind() == reflect.Struct {
		m := structs.Map(val)
		m["status"] = "success"

		b, err = json.Marshal(m)
		if err != nil {
			panic(err)
		}
	} else if val == nil {
		var m = map[string]interface{}{
			"status": "success",
		}
		b, _ = json.Marshal(m)
	} else {
		panic("invalid type")
	}
	w.Write(b)
}

func camelCase(str string) string {
	return strcase.ToLowerCamel(str)
}

func AbortErr(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	var m = map[string]interface{}{
		"status":  "error",
		"errCode": errors.ErrorCode(err),
		"errors":  err.Error(),
	}
	b, _ := json.Marshal(m)
	w.Write(b)
}
