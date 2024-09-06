package dto

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponseError(status int, message string) Response[string] {
	return Response[string]{
		Code:    status,
		Message: message,
		Data:    "",
	}
}

func CreateResponseErrorData(status int, message string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Code:    status,
		Message: message,
		Data:    data,
	}
}

func CreateResponseSuccess[T any](status int, data T) Response[T] {
	return Response[T]{
		Code:    status,
		Message: "sukses",
		Data:    data,
	}
}
