package version

import "runtime"

type Info[T any] struct {
	AppName   string `json:"app_name,omitempty"`
	Version   string `json:"version,omitempty"`
	Commit    string `json:"commit,omitempty"`
	Branch    string `json:"branch,omitempty"`
	BuildTime string `json:"build_time,omitempty"`
	GoVersion string `json:"go_version,omitempty"`
	Extra     T      `json:"extra,omitempty"`
}

func New[T any](extra T) *Info[T] {
	return &Info[T]{
		AppName:   AppName,
		Version:   Version,
		Commit:    Commit,
		Branch:    Branch,
		BuildTime: BuildTime,
		GoVersion: runtime.Version(),
		Extra:     extra,
	}
}
