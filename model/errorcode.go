package model

type errorCode struct {
	SUCCESS 		int
	ERROR 			int
}

var ErrorCode = errorCode{
	SUCCESS:		200,
	ERROR: 			100,
}