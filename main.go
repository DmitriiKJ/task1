package main

import (
	"fmt"
	"sync"
)

type Resource struct {
	Data string
}

type AllResource struct {
	All map[int]Resource
	mut sync.Mutex
}

type resManage interface {
	Register(id int, data string) error
	Unregister(id int) error
	Update(id int, data string) error
}

func (ar *AllResource) Register(id int, data string) error {
	ar.mut.Lock()
	defer ar.mut.Unlock()

	_, err := ar.All[id]

	if !err {
		ar.All[id] = Resource{Data: data}
		return nil
	}

	return fmt.Errorf("Resource already exists! Id: %d Data: %s", id, data)
}

func (ar *AllResource) Unregister(id int) error {
	ar.mut.Lock()
	defer ar.mut.Unlock()

	_, err := ar.All[id]

	if err {
		delete(ar.All, id)
		return nil
	}

	return fmt.Errorf("Resource is not found! Id: %d", id)
}

func (ar *AllResource) Update(id int, data string) error {
	ar.mut.Lock()
	defer ar.mut.Unlock()

	_, err := ar.All[id]

	if err {
		ar.All[id] = Resource{Data: data}
		return nil
	}

	return fmt.Errorf("Resource is not found! Id: %d Data: %s", id, data)

}

func main() {

	AllRes := AllResource{All: make(map[int]Resource)}

	var wg sync.WaitGroup

	// 1
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := AllRes.Register(1, "First")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Register(2, "Second")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Unregister(1)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Update(1, "New First")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Update(2, "New Second")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// 2
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := AllRes.Register(1, "First")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Register(3, "Final")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Unregister(1)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Update(3, "Some")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// 3
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := AllRes.Register(5, "Five")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Register(6, "Six")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Unregister(1)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Update(5, "New Five")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Update(2, "Two")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// 4
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := AllRes.Register(2, "Second")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Unregister(2)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Update(2, "Error")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// 5
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := AllRes.Register(11, "One One")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Unregister(2)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Update(1, "New One One")
		if err != nil {
			fmt.Println(err.Error())
		}

		err = AllRes.Update(2, "New Second")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	wg.Wait()
}
