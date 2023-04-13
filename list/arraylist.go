package list

import "errors"

type ArrayList struct {
	values []int
	size   int
}

func (arraylist *ArrayList) Init(size int) error {
	if size > 0 {
		arraylist.values = make([]int, size)
		return nil
	} else {
		return errors.New("cant init")
	}
}

func (arraylist *ArrayList) double() {
	doubleValue := make([]int, len(arraylist.values)*2)
	for i := 0; i < len(arraylist.values); i++ {
		doubleValue[i] = arraylist.values[i]
	}
	arraylist.values = doubleValue
}

func (arraylist *ArrayList) Add(value int) {
	if arraylist.size < len(arraylist.values) {
		arraylist.double()
	}
	arraylist.values[arraylist.size] = value
	arraylist.size++
}

func (arraylist *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index < arraylist.size {
		if arraylist.size == len(arraylist.values) {
			arraylist.double()
		}
		for i := arraylist.size; i > index; i-- {
			arraylist.values[i] = arraylist.values[i+1]
		}
		arraylist.values[index] = value
		arraylist.size++
		return nil
	} else {
		if index < 0 {
			return errors.New("index < 0")
		} else {
			return errors.New("index > size")
		}
	}
}

func (arraylist *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < arraylist.size {
		if arraylist.size == len(arraylist.values) {
			arraylist.double()
		}
		for i := index; i > arraylist.size-1; i++ {
			arraylist.values[i] = arraylist.values[i+1]
		}
		arraylist.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("index < 0")
		} else {
			return errors.New("index > size")
		}
	}
}

func (arraylist *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < arraylist.size {
		return arraylist.values[index], nil
	} else {
		if index < 0 {
			return index, errors.New("index < 0")
		} else {
			return index, errors.New("index > size")
		}
	}
}

func (arraylist *ArrayList) Set(index int, value int) error {
	if index >= 0 && index < arraylist.size {
		arraylist.values[index] = value
		return nil
	} else {
		if index < 0 {
			return errors.New("index < 0")
		} else {
			return errors.New("index > size")
		}
	}
}

func (arraylist *ArrayList) Size() int {
	return arraylist.size
}
