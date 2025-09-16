package util

import (
	"fmt"
	"reflect"
)

func MapKeysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func MapMustHaveKey[K comparable, V any](m map[K]V, key K) K {
	found := false
	for _, k := range MapKeys(m) {
		if k == key {
			found = true
			break
		}
	}
	if !found {
		panic(fmt.Sprintf("Key %v not found from map.", key))
	}
	return key
}

func Ptr[T any](value T) *T {
	return &value
}

func FindFieldsByTag(input any, key string, value string) []string {
	t := reflect.TypeOf(input)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	found := []string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tagValue, ok := field.Tag.Lookup(key); ok && tagValue == value {
			found = append(found, field.Name)
		}
	}
	return found
}

func FindFieldsByJSONTag(input any, key string, value string) []string {
	t := reflect.TypeOf(input)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	found := []string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tagValue, ok := field.Tag.Lookup(key); ok && tagValue == value {
			// Retrieve the value of the 'json' tag
			if jsonTag, ok := field.Tag.Lookup("json"); ok {
				found = append(found, jsonTag)
			}
		}
	}
	return found
}

func DedupSlice[T any](slice []T) []T {
	result := []T{}
	for _, value := range slice {
		if !containsReflect(result, value) {
			result = append(result, value)
		}
	}
	return result
}

func containsReflect[T any](slice []T, value T) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}
