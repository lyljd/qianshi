package errorxs

import "google.golang.org/grpc/status"

func Is(statusErr, err error) bool {
	if se, ok := status.FromError(statusErr); ok {
		return se.Message() == err.Error()
	}
	return statusErr == err
}
