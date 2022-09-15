package hexutil 

import (
	"bytes"
	"encoding/hex"
	"encoding/json" 
	"errors"
	"math/big"
	"testing"
)

func checkError(t *testing.T, input string, got, want error) bool{
	if got == nil{
		if want != nil{
			t.Errorf("input %s: got no error, want %q", input, want)
			return false 
		}
		return true
	}
	if want == nil{
		t.Errorf("input %s: unexpected error %q", input, got)
	} else if got.Error() != want.Error(){
		t.Errorf("input %s: got error %q, want %q", input, got, want)
	}
	return false
}

func referenceBig(s string) *big.Int{
	b, ok := new(big.Int).SetString(s, 16)
	if !ok{
		panic("Invalid")
	}
	return b
}

func referenceBytes(s string) []byte{
	b, err := hex.DecodeString(s)
	if err != nil{
		panic(err)
	}
	return b
}

var errJSONEOF = errors.New("unexpected end of JSON input") 

var unmarshalBytesTests = []unmarshalTest{
	// Invalid encoding 
	{input: "", wantErr: errJSONEOF}, 
	{input: "null", wantErr: errNonString(bytesT)}, 
	{input: "10", wantErr: errNonString(bytesT)}, 
	
}