package storage

import (
	"encoding/json"
	"fmt"
	"tapi/desktop/model"
	"testing"
	"time"
)

type FakeFileStore[F any] struct {
	Cache F
}

func (f *FakeFileStore[F]) getJson() error {
	bytes, err := json.Marshal(f.Cache)
	fmt.Println(string(bytes))
	return err
}
func TestFileStore(t *testing.T) {
	p1, err := NewFileStore("p1.json", []model.Project{})
	if err != nil {
		t.Fatal(err)
	}
	defer p1.Close()
	p1.Update(func(f []model.Project) []model.Project {
		f = append(f, model.Project{
			Id:    1,
			HubId: 2,
			Name:  "123",
		})
		return f
	})
	p2, err2 := NewFileStore("p2.json", model.Project{})
	if err2 != nil {
		t.Fatal(err2)
	}
	defer p2.Close()
	time.Sleep(20 * time.Second)
}
