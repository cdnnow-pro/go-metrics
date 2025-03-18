// SPDX-License-Identifier: MIT

package appmon

import (
	"errors"
	"runtime/debug"

	"github.com/cdnnow-pro/metrics-go"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	buildInfo = metrics.NewGaugeVec(
		"app_build_info",
		"Application build info",
		[]string{"version", "vcs", "revision", "go"},
	)
	dependencies = metrics.NewGaugeVec(
		"go_dependencies",
		"Application Go modules dependencies with versions",
		[]string{"path", "version"},
	)
)

// Register registers application metrics:
//   - app_build_info{version, vcs, revision}
//   - go_dependencies{path, version}
func Register(appVersion, vcs, revision string) error {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return errors.New("failed to read build info for application metrics")
	}

	buildInfo.With(prometheus.Labels{
		"version":  appVersion,
		"vcs":      vcs,
		"revision": revision,
		"go":       info.GoVersion,
	}).Set(1)

	for _, d := range info.Deps {
		dependencies.With(prometheus.Labels{
			"path":    d.Path,
			"version": d.Version,
		}).Set(1)
	}
	return nil
}
