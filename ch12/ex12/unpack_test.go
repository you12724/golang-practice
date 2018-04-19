package params

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestUnpack(t *testing.T) {
	type Data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}

	for _, test := range []struct {
		url  string
		data Data
	}{
		{`http://localhost:12345/search`, Data{nil, 10, false}},
		{`http://localhost:12345/search?l=golang&l=programming`,
			Data{[]string{"golang", "programming"}, 10, false}},
		{`http://localhost:12345/search?l=golang&l=programming&max=100`,
			Data{[]string{"golang", "programming"}, 100, false}},
		{`http://localhost:12345/search?x=true&l=golang&l=programming`,
			Data{[]string{"golang", "programming"}, 10, true}},
	} {
		var data Data
		data.MaxResults = 10 // set default

		req, err := newRequest(test.url)
		if err != nil {
			t.Errorf("Parse: %v\n", err)
			continue
		}

		if err := Unpack(req, &data); err != nil {
			t.Errorf("Unpack: %v\n", err)
			continue
		}

		if !reflect.DeepEqual(data, test.data) {
			t.Errorf("%q => \n%+v but want %+v\n", test.url, data, test.data)
		}
	}
}

func TestUnpackRXP(t *testing.T) {
	type Data struct {
		Mail string `http:"m"`
	}

	for _, test := range []struct {
		input  string
		err    bool
		result Data
	}{
		{`http://localhost:12345/search`, false, Data{}},
		{`http://localhost:12345/search?m=test@gmail.com`, false,
			Data{Mail: "test@gmail.com"}},
		{`http://localhost:12345/search?m=@`, true, Data{}},
		{`http://localhost:12345/search?m=test@`, true, Data{}},
		{`http://localhost:12345/search?m=@gmail.com`, true, Data{}},
	} {
		var data Data

		req, err := newRequest(test.input)
		if err != nil {
			t.Errorf("Parse: %v\n", err)
			continue
		}

		if err := Unpack(req, &data); err != nil {
			if test.err == true {
				continue
			}

			t.Errorf("Unpack: %v\n", err)
			continue
		}

		if !reflect.DeepEqual(data, test.result) {
			t.Errorf("%q => \n%+v but want %+v\n", test.input, data, test.result)
		}
	}
}

func newRequest(rawurl string) (*http.Request, error) {
	var req http.Request
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	req.URL = url
	return &req, nil
}
