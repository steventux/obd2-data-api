package helpers

import "os"

func Authenticate(id string) bool {
	return (id == os.Getenv("TORQUE_ID"))
}
