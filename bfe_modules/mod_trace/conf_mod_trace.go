// Copyright (c) 2019 Baidu, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mod_trace

import (
	"fmt"
)

import (
	"github.com/baidu/go-lib/log"
	"gopkg.in/gcfg.v1"
)

import (
	"github.com/baidu/bfe/bfe_modules/mod_trace/trace/zipkin"
	"github.com/baidu/bfe/bfe_util"
)

const (
	defaultDataPath = "mod_trace/trace_rule.data"
)

type ConfModTrace struct {
	Basic struct {
		DataPath    string // The path of rule data
		ServiceName string // The name of this service
	}

	Log struct {
		OpenDebug bool
	}

	Zipkin zipkin.Config
}

func ConfLoad(filePath string, confRoot string) (*ConfModTrace, error) {
	var err error
	var cfg ConfModTrace

	err = gcfg.ReadFileInto(&cfg, filePath)
	if err != nil {
		return nil, err
	}

	err = cfg.Check(confRoot)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *ConfModTrace) Check(confRoot string) error {
	if len(cfg.Basic.DataPath) == 0 {
		cfg.Basic.DataPath = defaultDataPath
		log.Logger.Warn("ModTrace.DataPath not set, use default value: %s", defaultDataPath)
	}
	cfg.Basic.DataPath = bfe_util.ConfPathProc(cfg.Basic.DataPath, confRoot)

	err := cfg.Zipkin.Check()
	if err != nil {
		return fmt.Errorf("ModTrace.Zipkin %v", err)
	}

	return err
}
