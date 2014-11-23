package contract_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bmizerany/assert"

	. "github.com/dockpit/pit/contract"
)

func TestFactoryLoading(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	//gven a new factory
	f := NewFactory()

	//when we load json example
	data, err := f.Load(filepath.Join(wd, "auth.json"))
	if err != nil {
		t.Fatal(err)
	}

	//assert contract data
	assert.Equal(t, "auth", data.Name)

	//assert resource
	assert.Equal(t, "/users", data.Resources[0].Pattern)

	//assert cases
	assert.Equal(t, "some users", data.Resources[0].Cases[0].Given["mongodb"].Name)
	assert.Equal(t, "some messages", data.Resources[0].Cases[0].Given["nsq"].Name)
	assert.Equal(t, "GET", data.Resources[0].Cases[0].When.Method)
	assert.Equal(t, "/users", data.Resources[0].Cases[0].When.Path)
	assert.Equal(t, 200, data.Resources[0].Cases[0].Then.StatusCode)
	assert.Equal(t, `[{"id": "32"}]`, data.Resources[0].Cases[0].Then.Body)

	assert.Equal(t, "single user", data.Resources[0].Cases[0].While[0].CaseName)
	assert.Equal(t, "github.com/dockpit/ex-store-customers", data.Resources[0].Cases[0].While[0].ID)
}

func TestFactoryDraft(t *testing.T) {
	var c C
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	//gven a new factory
	f := NewFactory()

	//when we load json example
	c, err = f.Draft(filepath.Join(wd, "auth.json"))
	if err != nil {
		t.Fatal(err)
	}

	//assert contract data
	assert.Equal(t, "auth", c.Name())

	//assert dependencies
	deps, err := c.Dependencies()
	if err != nil {
		t.Fatal(err)
	}

	//should have 1 dependency with empty list of cases
	assert.Equal(t, 1, len(deps))
	assert.Equal(t, []string{}, deps["github.com/dockpit/ex-store-customers"])

	//assert states
	states, err := c.States()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(states))
	assert.Equal(t, []string{"some users"}, states["mongodb"])

	//assert resource
	resources, err := c.Resources()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "/users", resources[0].Pattern())

	//assert actions
	actions, err := resources[0].Actions()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "GET", actions[0].Method())
	assert.Equal(t, "POST", actions[1].Method())

}