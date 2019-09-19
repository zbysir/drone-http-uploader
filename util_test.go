package main

import "testing"

func TestInjectEnv(t *testing.T) {
	src := injectEnv(`?secret={{SECRET}}?secret=${eee}`, map[string]string{"SECRET": "111"})
	t.Log(src)
}
