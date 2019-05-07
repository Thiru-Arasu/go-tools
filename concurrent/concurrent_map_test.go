package concurrent

import (
	"testing"
)

func Test_AddMap(t *testing.T) {
	m := NewConcurrentMap()
	m.Add("test1", "test1")
	if m.Get("test1", "").(string) != "test1" {
		t.Errorf("Returned unknown value: %v", m.Get("test1", ""))
	}
}

func Test_GetMap(t *testing.T) {
	m := NewConcurrentMap()
	m.Add("test1", "test1")
	if m.Get("test2", "default").(string) != "default" {
		t.Errorf("Returned unknown value: %v", m.Get("test2", ""))
	}
}

func Test_AppendMap(t *testing.T) {
	m := NewConcurrentMap()
	m.Add("test", []string{"test1"})
	m.Append(func(cm MapType) MapType {
		test1 := cm["test"].([]string)
		test1 = append(test1, "test2")
		cm["test"] = test1
		return cm
	})
	result := m.Get("test", "").([]string)
	if result[1] != "test2" {
		t.Errorf("Returned unknown value: %v", m.Get("test", ""))
	}
}

func Test_ForEach(t *testing.T) {
	m := NewConcurrentMap()
	m.Add("test1", "test1")
	m.Add("test2", "test2")
	m.ForEach(func(key string, value interface{}) {
		if key != "test1" && key != "test2" {
			t.Errorf("Returned unknown key: %s", key)
		}

		if value != "test1" && value != "test2" {
			t.Errorf("Returned unknown value: %v", value)
		}
	})
}
