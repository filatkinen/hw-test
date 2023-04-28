package hw09structvalidator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type (
	TagString Tag
	TagInt    Tag
)

func in(tagval, val string) error {
	slice := strings.Split(tagval, ",")
	for _, v := range slice {
		v = strings.TrimSpace(v)
		if v == val {
			return nil
		}
	}
	return ErrorValidation
}

func (t *TagInt) Call(val int) (bool, error) {
	if len(t.tagname) == 0 {
		return false, ErrorNotImplemented
	}
	fname := strings.ToUpper(string(t.tagname[0])) + t.tagname[1:]
	v := reflect.ValueOf(t).MethodByName(fname)
	if v == (reflect.Value{}) {
		return false, ErrorNotImplemented
	}
	result := v.Call([]reflect.Value{reflect.ValueOf(val)})
	if result[1].IsZero() {
		return result[0].Bool(), nil
	}
	return result[0].Bool(), result[1].Interface().(error)
}

func (t *TagInt) Min(val int) (bool, error) {
	tval, err := strconv.Atoi(t.tagval)
	if err != nil {
		return false, ErrorParsValidateTag
	}
	return val >= tval, nil
}

func (t *TagInt) Max(val int) (bool, error) {
	tval, err := strconv.Atoi(t.tagval)
	if err != nil {
		return false, ErrorParsValidateTag
	}
	return val <= tval, nil
}

func (t *TagInt) In(val int) (bool, error) {
	result := in(t.tagval, strconv.Itoa(val))
	if result != nil {
		return false, result
	}
	return true, nil
}

func (t *TagString) Call(val string) (bool, error) {
	if len(t.tagname) == 0 {
		return false, ErrorNotImplemented
	}
	fname := strings.ToUpper(string(t.tagname[0])) + t.tagname[1:]
	v := reflect.ValueOf(t).MethodByName(fname)
	if v == (reflect.Value{}) {
		return false, ErrorNotImplemented
	}
	result := v.Call([]reflect.Value{reflect.ValueOf(val)})
	if result[1].IsZero() {
		return result[0].Bool(), nil
	}
	return result[0].Bool(), result[1].Interface().(error)
}

func (t *TagString) Len(val string) (bool, error) {
	tval, err := strconv.Atoi(t.tagval)
	if err != nil {
		return false, ErrorParsValidateTag
	}
	return tval == len(val), nil
}

func (t *TagString) In(val string) (bool, error) {
	result := in(t.tagval, val)
	if result != nil {
		return false, result
	}
	return true, nil
}

func (t *TagString) Regexp(val string) (bool, error) {
	ok, err := regexp.MatchString(t.tagval, val)
	if err != nil {
		return false, ErrorParsValidateTag
	}
	return ok, nil
}
