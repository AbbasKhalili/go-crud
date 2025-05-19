package Dtos

type GenericResponseDto struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

//type GenericResponseDto[T any] struct {
//	Message string `json:"message"`
//	Data    T      `json:"data,omitempty"`
//}
//
//func SuccessResponse[T any](message string, data T) GenericResponseDto[T] {
//	return GenericResponseDto[T]{Message: message, Data: data}
//}
