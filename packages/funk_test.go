package packages

import (
	"testing"

	"github.com/mattn/anko/core"
	"github.com/mattn/anko/vm"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		t = ["a","b","c","d"]
		return funk.Contains(t, "a")
	`)
	assert.NoError(t, err)
	assert.True(t, res.(bool))
	res, err = env.Execute(`
		var funk = import("funk")
		t = ["a","b","c","d"]
		return funk.Contains(t, "e")
	`)
	assert.NoError(t, err)
	assert.False(t, res.(bool))
}

func TestIndexOf(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		t = ["a","b","c","d"]
		return funk.IndexOf(t, "a")
	`)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)
	res, err = env.Execute(`
		var funk = import("funk")
		t = ["a","b","c","d"]
		return funk.IndexOf(t, "e")
	`)
	assert.NoError(t, err)
	assert.Equal(t, -1, res)
}

func TestLastIndexOf(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		t = ["a","b","c","d"]
		return funk.LastIndexOf(t, "a")
	`)
	assert.NoError(t, err)
	assert.Equal(t, 0, res)
	res, err = env.Execute(`
		var funk = import("funk")
		t = ["a","b","c","d"]
		return funk.LastIndexOf(t, "e")
	`)
	assert.NoError(t, err)
	assert.Equal(t, -1, res)
}

// func TestMap(t *testing.T) {
// 	env := vm.NewEnv()
// 	core.Import(env)
// 	DefineImport(env)
// 	res, err := env.Execute(`
// 		var funk = import("funk")
// 		t = [1,2,3]
// 		temp = funk.Map(t, func(n){
// 			return n *2
// 		})
// 		return temp
// 	`)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, res)
// 	res, err = env.Execute(`
// 			var funk = import("funk")
// 			t = ["a","b","c","d"]
// 			return funk.LastIndexOf(t, "e")
// 		`)
// 	assert.NoError(t, err)
// 	assert.Equal(t, -1, res)
// }

func TestGet(t *testing.T) {
	type testStruct struct {
		Val string
	}
	env := vm.NewEnv()
	core.Import(env)
	env.Define("test", &testStruct{Val: "this is a test"})
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		temp = funk.Get(test, "Val")
		return temp
	`)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestKeys(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		m = {"foo": "bar", "bar": "baz"}
		temp = funk.Keys(m)
		return temp
	`)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestValues(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		m = {"foo": "bar", "bar": "baz"}
		temp = funk.Values(m)
		return temp
	`)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestForEach(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		t = [1,2,3,4]
		t2 = []
		func f(n){
			t2 += n
		}
		funk.ForEach(t,f)
		return t2
	`)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUniq(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		t = ["a","b","c","d","a","b","c"]
		return funk.Uniq(t)
	`)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestSum(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		t = [2,2,2,2,2]
		return funk.Sum(t)
	`)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestReverse(t *testing.T) {
	env := vm.NewEnv()
	core.Import(env)
	DefineImport(env)
	res, err := env.Execute(`
		var funk = import("funk")
		t = ["a","b","c","d","a","b","c"]
		return funk.Reverse(t)
	`)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
