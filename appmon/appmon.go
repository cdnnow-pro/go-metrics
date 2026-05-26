// SPDX-License-Identifier: MIT

package appmon

import (
	"errors"
	"runtime/debug"

	"github.com/cdnnow-pro/go-metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var ErrBuildInfo = errors.New("failed to read build info for application metrics")

type ApplicationInfo struct {
	Name     string
	Version  string
	VCS      string
	Revision string
	Date     string
}

// Register registers application metrics:
//   - build_info{name, version, vcs, revision, date}
//   - go_dependencies{path, version}
func Register(registerer prometheus.Registerer, appInfo ApplicationInfo) error {
	metrics.NewGaugeFor(registerer, prometheus.GaugeOpts{
		Name: "build_info",
		Help: "Application build info",
		ConstLabels: prometheus.Labels{
			"name":     appInfo.Name,
			"version":  appInfo.Version,
			"vcs":      appInfo.VCS,
			"revision": appInfo.Revision,
			"date":     appInfo.Date,
		},
	}).Set(1)

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ErrBuildInfo
	}

	for _, d := range info.Deps {
		metrics.NewGaugeFor(registerer, prometheus.GaugeOpts{
			Name: "go_dependencies",
			Help: "Application Go modules dependencies with versions",
			ConstLabels: prometheus.Labels{
				"path":    d.Path,
				"version": d.Version,
			},
		}).Set(1)
	}
	return nil
}
