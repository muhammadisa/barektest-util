package er

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Ebl(code codes.Code, message string, systemerror error) error {
	if systemerror != nil {
		return status.Error(code, fmt.Sprintf("{%s} | {%s}", message, systemerror.Error()))
	}
	return status.Error(code, fmt.Sprintf("{%s} | {%s}", message, errors.New("no system error")))
}
