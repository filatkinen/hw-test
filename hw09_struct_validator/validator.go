package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type Tag struct {
	tagname string
	tagval  string
}

var (
	ErrorNotStructure    = errors.New("error variable is not structure")
	ErrorValidation      = errors.New("error validation")
	ErrorParsValidateTag = errors.New("error parsing or validation tag")
	ErrorNotImplemented  = errors.New("error tag has not been implemented yet")
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	sb := strings.Builder{}
	for _, val := range v {
		sb.WriteString(fmt.Sprintf("error in field:%s: %s\n", val.Field, val.Err))
	}
	return sb.String()
}

func Validate(v interface{}) error {
	valerr := ValidationErrors{}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return ErrorNotStructure
	}
	tval := val.Type()
	for i := 0; i < tval.NumField(); i++ {
		ft := tval.Field(i)
		fv := val.Field(i)
		e := validateElem(ft, fv)
		for _, er := range e {
			valerr = append(valerr, ValidationError{
				Field: ft.Name,
				Err:   er,
			})
		}
	}
	if len(valerr) == 0 {
		return nil
	}
	return valerr
}

func validateElem(ft reflect.StructField, fv reflect.Value) (errslice []error) { //nolint
	if ft.Tag == "" || fv.IsZero() {
		return nil
	}
	if !fv.CanInterface() {
		return nil
	}
	tag := ft.Tag.Get("validate")
	if tag == "" {
		return nil
	}
	switch fv.Kind() { //nolint
	case reflect.Int, reflect.Uint8, reflect.Int32, reflect.Int64:
		if e := validateValue(fv.Int(), tag); e != nil {
			errslice = append(errslice, e)
		}
	case reflect.String:
		if e := validateValue(fv.Interface(), tag); e != nil {
			errslice = append(errslice, e)
		}
	case reflect.Slice:
		switch t := fv.Interface().(type) {
		case []string:
			for _, value := range t {
				if e := validateValue(value, tag); e != nil {
					errslice = append(errslice, e)
				}
			}
		case []int:
			for _, value := range t {
				if e := validateValue(int64(value), tag); e != nil {
					errslice = append(errslice, e)
				}
			}
		case []byte:
			for _, value := range t {
				if e := validateValue(int64(value), tag); e != nil {
					errslice = append(errslice, e)
				}
			}
		}
	}
	return errslice
}

func wraperrors(err1, err2 error) error {
	switch {
	case err1 == nil && err2 == nil:
		return nil
	case err1 == nil:
		return err2
	case err2 == nil:
		return err1
	default:
		return fmt.Errorf("%w;%w;", err1, err2)
	}
}

func validateValue(val any, tag string) (resultError error) {
	result := true
	slicetag, err := parseTag(tag)
	resultError = wraperrors(resultError, err)
	if len(slicetag) == 0 {
		resultError = wraperrors(resultError, ErrorParsValidateTag)
		return resultError
	}

	r := true //nolint
	for _, tagelem := range slicetag {
		var e error
		switch v := val.(type) {
		case int64:
			intelem := TagInt(tagelem)
			r, e = intelem.Call(int(v))
		case string:
			strtelem := TagString(tagelem)
			r, e = strtelem.Call(v)
		default:
			continue
		}
		result = result && r
		resultError = wraperrors(resultError, e)
	}
	if !result {
		resultError = wraperrors(resultError, ErrorValidation)
	}
	return resultError
}

func parseTag(tag string) ([]Tag, error) {
	slicetag := []Tag{}
	tags := strings.Split(tag, "|")
	for _, v := range tags {
		val := strings.SplitN(v, ":", 2)
		if len(val) != 2 {
			return []Tag{}, ErrorParsValidateTag
		}
		slicetag = append(slicetag, Tag{
			tagname: val[0],
			tagval:  val[1],
		})
	}

	return slicetag, nil
}
