package axisapi

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParesPattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRouter(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRouter("GET", "/hello/axisapi")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "axisapi" {
		t.Fatal("name should be equal to 'axisapi'")
	}
	fmt.Printf("matched path: %s,params['name']:%s\n", n.pattern, ps["name"])

	n, ps = r.getRouter("GET", "/assets/a/b/c")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/assets/*filepath" {
		t.Fatal("should match /assets/*filepath")
	}
	if ps["filepath"] != "a/b/c" {
		t.Fatal("name should be equal to 'a/b/c'")
	}
	fmt.Printf("matched path: %s,params['filepath']:%s\n", n.pattern, ps["filepath"])
}
