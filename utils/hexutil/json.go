package hexutil 

import (
	"encoding/json" 
	"encoding/hex" 
	"fmt"
	"math/big" 
	"reflect" 
	"strconv"
)

var (
	bytesT = reflect.TypeOf(Bytes(nil))
	bigT = reflect.TypeOf((*Big)(nil))
	uintT = reflect.TypeOf(Uint(0))
	uint64T = reflect.TypeOf(Uint64(0))
)

// Bytes marhshals/unmarshals as a JSON string with 0x prefix 
// The empty slice  marshals as "0x" 
type Bytes []byte 

func (b Bytes) MarshalText()([]byte, error){
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`) 
	hex.Encode(result[2:], b)
	return result, nil 
}

// Unmarshal JSON implements Json.Unmarshal 
func (b *Bytes) UnmarshalJSON(input []byte) error{
	if isString(input){
		return errNonString(bytesT)
	}
	return wrapTypeError(b.UnmarshalText(input[1:len(input)-1]), bytesT)
}

func (b *Bytes) UnmarshalText(input []byte) error{
	raw, err := checkText(input, true)
	if err != nil{
		return err
	}
	dec := make([]byte, len(raw)/2)
	if _, err:= hex.Decode(dec, raw); err!= nil{
		err = mapError(err)
	} else{
		*b = dec
	}
	return err
}

// String returns the hex encoding of b 
func (b Bytes) String() string{
	return Encode(b)
}

// ImplementsGraphQLType returns true if Bytes implement the specified GraphQL Type 
func (b Bytes) ImplementsGraphQLType(name string) bool { return name == "Bytes" } 

// UnMarshalGraphQL unmarshalls the provided GraphQL query data 
func (b *Bytes) UnMarshalGraphQL(input interface{}) error{
	var err error 
	switch input := input.(type){
	case string: 
		data, err := Decode(input)
		if err != nil{
			return err
		}
		*b = data
	default: 
		err =  fmt.Errorf("unexpected type %T for Bytes", input)
	}
	return err
}

// UnMarshalFixedJson decodes the input as string with 0x prefix. the length 
// of Out determines the required input length. this function is commonly used 
// to implement the UnmarshalJSON for fixed-size types 
func UnMarshalFixedJson(typ reflect.Type, input, out []byte) error{
	if !isString(input){
		return errNonString(typ) 
	}
	return wrapTypeError(UnMarshalFixedText(typ.String(), input[1:len(input)-1], out), typ)
}

// UnMarshalFixedText decodes the input as a string with 0x prefix. the length  of 
// out determines the required input length. This function is commonly used to implement 
// The UnMarshalText method for fixed size types. 
func UnMarshalText(typename string, input, out []byte) error{
	raw, err:= checkText(input, true)
	if err != nil{
		return err
	}
	if len(raw)/2 != len(out){
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(out)*2, typename)
	}

	// Pre verify syntax before modifying 
	for _, b := range raw{
		if decodeNibble(b) == badNibble{
			return ErrSyntax
		}
	}
	hex.Decode(out, raw)
	return nil 
}

// UnMarshalFixedUnPrefixedText decodes the input as a string with Optional 0x prefix 
// Teh length of the out determines the required input length. This function 	
// Is commonly used to implement the UnMarshalText method for fixed sized inpût.
func UnMarshalFixedUnPrefixedText(typename string, input, out []byte) error{
	raw, err := checkText(input, false)
	if err != nil{
		return err 
	}
	if len(raw)/2 != len(out){
		return fmt.Errorf("hex string has length %d, want %d for %s", len(raw), len(out)*2, typename)
	}

	for _, b:= range raw{
		if decodeNibble(b) == badNibble{
			return ErrSyntax
		}
	}
	hex.Decode(out, raw)
	return nil
}

// Big marshals/unmarshals as a JSON with 0x prefix 
// The Zero value marshals as "0x0" 
// 
//  Negative integers are not supported. Attempting to 
// marshall them will return an error. values larger than 256bits 
// are rejected by UnMarshal but will  be marshalled without error 
type Big big.Int 


// TODO: Tomorrow implement Big class methods




// Utilities 
func isString(input []byte) bool{
	return len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"'
}

func bytesHave0xPrefix(input []byte) bool{
	return len(input) >= 2 && input[0] = '0' && (input[1] == 'x' || input[1] == 'X')
}

func checkText(input []byte, wantPrefix bool)([]byte, error){
	if len(input) == 0{
		return nil, nil 
	}
	if bytesHave0xPrefix(input){
		input = input[2:]
	} else if wantPrefix{
		return nil, ErrMissingPrefix
	}
	if len(input)%2 != 0{
		return nil, ErrOddLength
	}

	return input, nil
}

func checkNumberText(input []byte) (raw []byte, err error){
	if len(input) == 0{
		return nil, nil 
	}
	if !bytesHave0xPrefix(input){
		return nil, ErrMissingPrefix
	}
	input = input[2:] 
	if len(input) == 0{
		return nil, ErrEmptyNumber 
	}
	if len(input) > 1 && input[0] == '0'{
		return nil, ErrLeadingZero
	}

	return input, nil 
}

func wrapTypeError(err error, typ reflect.Type) error{
	if _, ok := err.(*decError); ok{
		return &json.UnmarshalTypeError{Value: err.Error(), Type: typ}
	}
	return err
}

func errNonString(typ reflect.Type) error{
	return &json.UnmarshalTypeError{Value: "non-string", Type: typ}
}