package vector

import (
	"fmt"
	"testing"

	"github.com/FoundationDB/fdb-go/fdb"
	"github.com/FoundationDB/fdb-go/fdb/directory"
)

func TestClear(t *testing.T) {

	fdb.MustAPIVersion(200)
	db := fdb.MustOpenDefault()

	subspace, err := directory.CreateOrOpen(db, []string{"tests", "vector"}, []byte{0})
	if err != nil {
		panic(err)
	}

	_, e := db.Transact(func(tr fdb.Transaction) (interface{}, error) {

		vector := Vector{subspace: subspace}

		vector.Set(0, "a", tr)
		vector.Set(1, "b", tr)

		i := vector.size(tr)
		if i != 2 {
			return nil, fmt.Errorf("Expected vector to be size 1, got %d instead", i)
		}

		vector.clear(tr)

		i = vector.size(tr)
		if i != 0 {
			return nil, fmt.Errorf("Expected empty vector to be size 0, got %d instead", i)
		}

		return nil, nil

	})

	if e != nil {
		t.Error(e)
	}

}

func TestSize(t *testing.T) {

	fdb.MustAPIVersion(200)
	db := fdb.MustOpenDefault()

	subspace, err := directory.CreateOrOpen(db, []string{"tests", "vector"}, []byte{0})
	if err != nil {
		panic(err)
	}

	_, e := db.Transact(func(tr fdb.Transaction) (interface{}, error) {

		vector := Vector{subspace: subspace}
		vector.clear(tr)

		i := vector.size(tr)
		if i != 0 {
			return nil, fmt.Errorf("Expected empty vector to be size 0, got %d instead", i)
		}

		vector.Set(0, "a", tr)

		i = vector.size(tr)
		if i != 1 {
			return nil, fmt.Errorf("Expected vector to be size 1, got %d instead", i)
		}

		return nil, nil

	})

	if e != nil {
		t.Error(e)
	}

}

func TestPushPop(t *testing.T) {

	fdb.MustAPIVersion(200)
	db := fdb.MustOpenDefault()

	subspace, err := directory.CreateOrOpen(db, []string{"tests", "vector"}, []byte{0})
	if err != nil {
		panic(err)
	}

	_, e := db.Transact(func(tr fdb.Transaction) (interface{}, error) {

		vector := Vector{subspace: subspace}
		vector.clear(tr)

		vector.Push("a", tr)
		vector.Push("b", tr)

		v, err := vector.Pop(tr)
		if err != nil {
			return nil, fmt.Errorf("Pop returned an error")
		}
		if v != "b" {
			return nil, fmt.Errorf("Expected popped value to be 'b', got %s instead", v)
		}

		v, err = vector.Pop(tr)
		if err != nil {
			return nil, fmt.Errorf("Pop returned an error")
		}
		if v != "a" {
			return nil, fmt.Errorf("Expected popped value to be 'a', got %s instead", v)
		}

		i := vector.size(tr)
		if i != 0 {
			return nil, fmt.Errorf("Expected empty vector to be size 0, got %d instead", i)
		}

		return nil, nil

	})

	if e != nil {
		t.Error(e)
	}

}