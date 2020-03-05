package services

import "os"

func FlagOrEnv(fl *string, envNave string) (value string, ok bool) {
	if fl == nil {
		return *fl, true
	}

	return os.LookupEnv(envNave)
}
//postgres://tagspmotvklkfi:9cb1a3d6f70ad82baecafe26750b184e30e1dfeed0ec884b1f1aee6b119f4f3d@localhost:5432/dcs5aet6f8io8d
