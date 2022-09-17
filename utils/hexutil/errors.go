package hexutil 

import (
	"fmt"
)


const uintBits = 32 << (uint64(^uint(0)) >> 63)

// Errors 
var (
	ErrEmptyString = &decError{"empty hex string"}
	ErrSyntax = &decError{"invalid hex string"}
	ErrMissingPrefix = &decError{"hex string withouth a 0x prefix"} 
	ErrOddLength = &decError{"hex string of odd length"} 
	ErrEmptyNumber = &decError{"hex string \"0x\""} 
	ErrLeadingZero = &decError{"hex number without leading zero digits"} 
	ErrUint64Range = &decError{"hex number > 64 bits"} 
	ErrUintRange= &decError{fmt.Sprintf("hex number > %d bits", uintBits)}
	ErrBig256Range = &decError{"hex number > 256 bits"}
)


type decError struct{ msg string}  

func (err decError) Error() string { return err.msg }
