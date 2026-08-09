// Copyright (C) 2026 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ci_tests

import "testing"

func TestProductRelativePath(t *testing.T) {
	testCases := []struct {
		name          string
		productOut    string
		installedFile string
		wantPath      string
		wantBelongs   bool
	}{
		{
			name:          "relative product file",
			productOut:    "out/target/product/nucleus_x86_64",
			installedFile: "out/target/product/nucleus_x86_64/system/app/Test.apk",
			wantPath:      "system/app/Test.apk",
			wantBelongs:   true,
		},
		{
			name:          "absolute product file",
			productOut:    "/build/out/target/product/nucleus_x86_64",
			installedFile: "/build/out/target/product/nucleus_x86_64/system/app/Test.apk",
			wantPath:      "system/app/Test.apk",
			wantBelongs:   true,
		},
		{
			name:          "absolute host file",
			productOut:    "/build/out/target/product/nucleus_x86_64",
			installedFile: "/build/out/host/linux-x86/bin/tool",
			wantBelongs:   false,
		},
		{
			name:          "product prefix collision",
			productOut:    "/build/out/target/product/nucleus_x86_64",
			installedFile: "/build/out/target/product/nucleus_x86_64_other/system/app/Test.apk",
			wantBelongs:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			gotPath, gotBelongs := productRelativePath(testCase.productOut, testCase.installedFile)
			if gotPath != testCase.wantPath || gotBelongs != testCase.wantBelongs {
				t.Fatalf("productRelativePath(%q, %q) = (%q, %t), want (%q, %t)", testCase.productOut, testCase.installedFile, gotPath, gotBelongs, testCase.wantPath, testCase.wantBelongs)
			}
		})
	}
}
