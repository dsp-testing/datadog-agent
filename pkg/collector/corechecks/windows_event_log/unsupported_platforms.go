// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2021-present Datadog, Inc.

//go:build !windows

// Package evtlog is not implemented on non-Windows platforms
package evtlog

import (
	"github.com/DataDog/datadog-agent/pkg/collector/check"
	"github.com/DataDog/datadog-agent/pkg/util/optional"
)

const (
	// CheckName is the name of the check
	CheckName = "win32_event_log"
)

// Factory creates a new check factory
func Factory() optional.Option[func() check.Check] {
	return optional.NewNoneOption[func() check.Check]()
}
