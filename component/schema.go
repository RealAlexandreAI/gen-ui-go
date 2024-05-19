package component

import (
	"encoding/json"
	"fmt"
)

type Metadata struct {
	ExamplePath string
}

type Input interface {
	any
	fmt.Stringer
	json.Marshaler
	json.Unmarshaler
	Payload() any
}

type Component interface {
	Metadata() Metadata
	Input() Input
}

type payload = any

type WrappedInput struct {
	payload `json:",inline"`
}

func (wi WrappedInput) String() string {
	switch v := wi.payload.(type) {
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (wi WrappedInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(wi.payload)
}

func (wi WrappedInput) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &wi.payload)
}

func (wi WrappedInput) Payload() any {
	return wi.payload
}

func WrapInput(p any) Input {
	return WrappedInput{payload: p}
}

func UnwrapInput(i Input, ref any) error {
	b, err := i.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(b, ref)
}
