package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//////////////////////////////////////////
//////////////////////////////////////////
/////// Not in use!
//////////////////////////////////////////
//////////////////////////////////////////


/// Get ...
func Get(r *http.Request, key string) string {
	v, ok := r.URL.Query()[key]
	if !ok {
		return ""
	}
	return v[0]
}

/// GetSlice ...
func GetSlice(r *http.Request, key string) (v []string) {
	v, ok := r.URL.Query()[key]
	if !ok {
		return v
	}
	return v
}

// AjaxGetData
func ajaxGetData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Println("ajax/get-data foldered")

	fmt.Printf("%v\n", p)

	get := r.URL.Query()
	fmt.Printf("%v\n", get)
	fmt.Println(get["pair_id"])
	fmt.Println(Get(r, "pair_id"))

	fmt.Println(Get(r, "asfs"))

}
