package handler

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil"
)

// UnmarshalXML unmarshals a response body for the XML protocol.
func UnmarshalXML(v interface{}, r *http.Response, operation string) error {

	defer r.Body.Close()

	if v == nil {
		return nil
	}

	decoder := xml.NewDecoder(r.Body)
	err := xmlutil.UnmarshalXML(v, decoder, "")

	return sendError(err)
}

// UnmarshalLBUXML unmarshals a response body for the XML protocol.
func UnmarshalLBUXML(v interface{}, r *http.Response, operation string) error {

	defer r.Body.Close()

	if v == nil {
		return nil
	}

	operationName := operation[7:strings.Index(operation, "&")]

	decoder := xml.NewDecoder(r.Body)
	err := xmlutil.UnmarshalXML(v, decoder, operationName+"Result")

	return sendError(err)
}

//UnmarshalDLHandler ...
func UnmarshalDLHandler(v interface{}, r *http.Response, operation string) error {
	defer r.Body.Close()

	j := struct {
		RequestID string `json:"RequestId" type:"string"`
	}{
		r.Header.Get("X-Amz-Requestid"),
	}

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}

	o := commonStuctre(j, v)

	err = json.Unmarshal([]byte(o), v)
	if err != nil {
		return err
	}

	return err

}

//UnmarshalICUHandler ...
func UnmarshalICUHandler(v interface{}, r *http.Response, operation string) error {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}

	return err
}

func debugResponse(r *http.Response) {

	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func sendError(err error) error {
	if err != nil {
		return errors.New("SerializationError" + "failed decoding query response" + fmt.Sprint(err))
	}

	return nil
}
