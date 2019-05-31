package linkedlist

import (
	"testing"
)

func TestCreateEmpty(t *testing.T) {
	actual := CreateEmpty()
	expected := &Node{
		next:  nil,
		value: "",
	}

	// find out how to test it properly
	if actual.next != expected.next && actual.value != expected.value {
		t.Errorf("Create doesn't work as expected, expected: %v, got %v", expected, actual)
	}
}

func TestInsert(t *testing.T) {
	list := CreateEmpty()
	list.Insert("1")

	actual := list.value
	expected := "1"

	if expected != actual {
		t.Errorf("Insert at head doesn't work as expected, expected: %v, got %v", expected, actual)
	}

	list.Insert("2")
	actual = list.next.value
	expected = "2"

	if list.value != "1" {
		t.Errorf("Insert at last doesn't work as expected, expected: %v, got %v", expected, actual)
	}

	if expected != actual {
		t.Errorf("Insert at last doesn't work as expected, expected: %v, got %v", expected, actual)
	}

	list.Insert("3")
	actual = list.next.next.value
	expected = "3"

	if list.value != "1" {
		t.Errorf("Insert at last doesn't work as expected, expected: %v, got %v", expected, actual)
	}

	if list.next.value != "2" {
		t.Errorf("Insert at last doesn't work as expected, expected: %v, got %v", expected, actual)
	}

	if expected != actual {
		t.Errorf("Insert at last doesn't work as expected, expected: %v, got %v", expected, actual)
	}
}

func TestLength(t *testing.T) {
	list := CreateEmpty()

	actual := list.Length()
	expected := 0

	if expected != actual {
		t.Errorf("Length doesn't work for empty list, expected: %v, got %v", expected, actual)
	}

	list.Insert("1")
	list.Insert("2")
	list.Insert("3")

	actual = list.Length()
	expected = 3

	if expected != actual {
		t.Errorf("Length doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}
}

func TestRemove(t *testing.T) {
	list := CreateEmpty()
	list.Remove()

	actual := list.Length()
	expected := 0

	if expected != actual {
		t.Errorf("Remove doesn't work for empty lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("1")
	list.Insert("2")
	list.Insert("3")
	list.Remove()

	actual = list.Length()
	expected = 2

	if expected != actual {
		t.Errorf("Remove doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}
}

func TestIsEmpty(t *testing.T) {
	list := CreateEmpty()

	actual := list.IsEmpty()
	expected := true

	if expected != actual {
		t.Errorf("IsEmpty doesn't work for empty lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("1")

	actual = list.IsEmpty()
	expected = false

	if expected != actual {
		t.Errorf("IsEmpty doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}
}

func TestContains(t *testing.T) {
	list := CreateEmpty()

	actual := list.Contains("0")
	expected := false

	if expected != actual {
		t.Errorf("Contains doesn't work for empty lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("1")

	actual = list.Contains("1")
	expected = true

	if expected != actual {
		t.Errorf("Contains doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("2")

	actual = list.Contains("1")
	expected = true

	if expected != actual {
		t.Errorf("Contains doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	actual = list.Contains("2")
	expected = true

	if expected != actual {
		t.Errorf("Contains doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	actual = list.Contains("3")
	expected = false

	if expected != actual {
		t.Errorf("Contains doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("3")
	actual = list.Contains("2")
	expected = true

	if expected != actual {
		t.Errorf("Contains doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}
}

func TestGet(t *testing.T) {
	list := CreateEmpty()

	actual, err := list.Get(0)

	if actual != "" && err != nil {
		t.Errorf("Get doesn't work for empty lists")
	}

	list.Insert("1")
	list.Insert("2")
	list.Insert("3")

	actual, _ = list.Get(0)
	expected := "1"

	if expected != actual {
		t.Errorf("Get doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	actual, _ = list.Get(1)
	expected = "2"

	if expected != actual {
		t.Errorf("Get doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	actual, _ = list.Get(2)
	expected = "3"

	if expected != actual {
		t.Errorf("Get doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	_, err = list.Get(4)

	if err == nil {
		t.Errorf("Get doesn't work for filled lists, expected: %v, got %v", "error", err)
	}
}

func TestGetFirst(t *testing.T) {
	list := CreateEmpty()

	actual := list.GetFirst()

	if actual != "" {
		t.Errorf("GetFirst doesn't work for empty lists")
	}

	list.Insert("1")

	actual = list.GetFirst()
	expected := "1"

	if expected != actual {
		t.Errorf("GetFirst doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("2")
	actual = list.GetFirst()
	expected = "1"

	if expected != actual {
		t.Errorf("GetFirst doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}
}

func TestGetLast(t *testing.T) {
	list := CreateEmpty()

	actual := list.GetLast()

	if actual != "" {
		t.Errorf("GetLast doesn't work for empty lists")
	}

	list.Insert("1")

	actual = list.GetLast()
	expected := "1"

	if expected != actual {
		t.Errorf("GetLast doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("2")
	actual = list.GetLast()
	expected = "2"

	if expected != actual {
		t.Errorf("GetLast doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}
}

func TestIndexOf(t *testing.T) {
	list := CreateEmpty()

	actual := list.IndexOf("1")
	expected := -1

	if actual != expected {
		t.Errorf("IndexOf doesn't work for empty lists")
	}

	list.Insert("1")

	actual = list.IndexOf("1")
	expected = 0

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("2")
	actual = list.IndexOf("1")
	expected = 0

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	actual = list.IndexOf("2")
	expected = 1

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	actual = list.IndexOf("4")
	expected = -1

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("3")
	list.Insert("4")
	list.Insert("5")
	actual = list.IndexOf("2")
	expected = 1

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}
}

func TestToArray(t *testing.T) {
	list := CreateEmpty()

	actual := list.ToArray()
	expected := 0

	if len(actual) != expected {
		t.Errorf("ToArray doesn't work as expected, expected: %v, got %v", expected, actual)
	}

	list.Insert("1")
	actual = list.ToArray()

	if len(actual) != 1 && actual[0] != "1" {
		t.Errorf("ToArray at last doesn't work as expected, expected: %v, got %v", expected, actual)
	}

	list.Insert("2")
	actual = list.ToArray()

	if len(actual) != 2 && actual[1] != "2" && actual[0] != "1" {
		t.Errorf("ToArray at last doesn't work as expected, expected: %v, got %v", expected, actual)
	}
}

func TestFromArray(t *testing.T) {
	array := []string{"1", "2", "3", "4", "5"}

	list := FromArray(array)
	actual := list.Length()
	expected := 5

	if actual != expected {
		t.Errorf("FromArray doesn't work as expected, expected: %v, got %v", expected, actual)
	}

	list = FromArray([]string{})
	actual = list.Length()
	expected = 0

	if actual != expected {
		t.Errorf("FromArray doesn't work as expected, expected: %v, got %v", expected, actual)
	}
}

func TestLastIndexOf(t *testing.T) {
	list := CreateEmpty()

	actual := list.LastIndexOf("1")
	expected := -1

	if actual != expected {
		t.Errorf("LastIndexOf doesn't work for empty lists")
	}

	list.Insert("1")

	actual = list.LastIndexOf("1")
	expected = 0

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("1")

	actual = list.LastIndexOf("1")
	expected = 1

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("2")

	actual = list.LastIndexOf("1")
	expected = 1

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}

	list.Insert("2")

	actual = list.LastIndexOf("2")
	expected = 3

	if expected != actual {
		t.Errorf("IndexOf doesn't work for filled lists, expected: %v, got %v", expected, actual)
	}
}

func TestTruncate(t *testing.T) {
	list := CreateEmpty()

	list.Truncate()

	actual := list.Length()
	expected := 0

	if actual != expected {
		t.Errorf("Truncate doesn't work for empty lists")
	}

	list.Insert("1")
	list.Insert("2")
	list.Insert("3")
	list.Truncate()

	actual = list.Length()
	expected = 0

	if actual != expected {
		t.Errorf("Truncate doesn't work for filled lists")
	}
}

func TestInsertAt(t *testing.T) {
	list := CreateEmpty()

	err := list.InsertAt(200, "98")

	if err == nil || list.value != "" {
		t.Errorf("InsertAt doesn't work for empty lists")
	}

	err = list.InsertAt(0, "1")

	if err != nil || list.value != "1" || list.next != nil {
		t.Errorf("InsertAt doesn't work for empty lists")
	}

	err = list.InsertAt(0, "2")

	if err != nil || list.value != "2" || list.next.value != "1" || list.next.next != nil {
		t.Errorf("InsertAt doesn't work for filled lists")
	}

	err = list.InsertAt(1, "3")

	if err != nil ||
		list.value != "2" ||
		list.next.value != "3" ||
		list.next.next.value != "1" ||
		list.next.next.next != nil {
		t.Errorf("InsertAt doesn't work for filled lists")
	}

	err = list.InsertAt(3, "4")

	if err != nil ||
		list.value != "2" ||
		list.next.value != "3" ||
		list.next.next.value != "1" ||
		list.next.next.next.value != "4" ||
		list.next.next.next.next != nil {
		t.Errorf("InsertAt doesn't work for filled lists")
	}

	err = list.InsertAt(0, "5")

	if err != nil ||
		list.value != "5" ||
		list.next.value != "2" ||
		list.next.next.value != "3" ||
		list.next.next.next.value != "1" ||
		list.next.next.next.next.value != "4" ||
		list.next.next.next.next.next != nil {
		t.Errorf("InsertAt doesn't work for filled lists at index 0")
	}

	err = list.InsertAt(1, "6")

	if err != nil ||
		list.value != "5" ||
		list.next.value != "6" ||
		list.next.next.value != "2" ||
		list.next.next.next.value != "3" ||
		list.next.next.next.next.value != "1" ||
		list.next.next.next.next.next.value != "4" ||
		list.next.next.next.next.next.next != nil {
		t.Errorf("InsertAt doesn't work for filled lists at index 0")
	}
}

func TestRemoveAt(t *testing.T) {
	list := CreateEmpty()

	err := list.RemoveAt(0)

	if err == nil {
		t.Errorf("RemoveAt doesn't work for empty lists")
	}

	err = list.RemoveAt(1)

	if err == nil {
		t.Errorf("RemoveAt doesn't work for empty lists")
	}

	list.Insert("1")
	list.Insert("2")
	list.Insert("3")

	err = list.RemoveAt(0)

	if err != nil || list.value != "2" || list.next.value != "3" || list.next.next != nil {
		t.Errorf("RemoveAt doesn't work for filled lists at index 0")
	}

	list = CreateEmpty()
	list.Insert("1")
	list.Insert("2")
	list.Insert("3")

	err = list.RemoveAt(1)

	if err != nil || list.value != "1" || list.next.value != "3" || list.next.next != nil {
		t.Errorf("RemoveAt doesn't work for filled lists at index 1")
	}

	list = CreateEmpty()
	list.Insert("1")
	list.Insert("2")
	list.Insert("3")

	err = list.RemoveAt(2)

	if err != nil || list.value != "1" || list.next.value != "2" || list.next.next != nil {
		t.Errorf("RemoveAt doesn't work for filled lists at index 2")
	}

	err = list.RemoveAt(50)

	if err == nil || list.value != "1" || list.next.value != "2" || list.next.next != nil {
		t.Errorf("RemoveAt doesn't work for filled lists with bigger index")
	}
}
