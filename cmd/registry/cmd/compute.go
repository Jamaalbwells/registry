// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

var computeFilter string

func init() {
	rootCmd.AddCommand(computeCmd)
	computeCmd.PersistentFlags().StringVar(&computeFilter, "filter", "", "filter compute arguments")
}

var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Compute properties of resources in the API Registry",
}