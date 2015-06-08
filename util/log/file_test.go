// Go support for leveled logs, analogous to https://code.google.com/p/google-clog/
//
// Copyright 2013 Google Inc. All Rights Reserved.
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

// Author: Bram Gruneir (bram@cockroachlabs.com)

package log

import (
	"testing"
	"time"
)

// Test that logName and parseLogFilename work as advertised.
func TestLogFilenameParsing(t *testing.T) {
	testCases := []struct {
		Level Level
		Time  time.Time
	}{
		{INFO, time.Now()},
		{WARNING, time.Now().AddDate(-10, 0, 0)},
		{ERROR, time.Now().AddDate(0, 0, -1)},
	}

	for i, testCase := range testCases {
		filename, _ := logName(testCase.Level, testCase.Time)
		details, err := parseLogFilename(filename)
		if err != nil {
			t.Fatal(err)
		}
		if details.Level != testCase.Level {
			t.Errorf("%d: Levels do not match, expected:%s - actual:%s", i, testCase.Level, details.Level)
		}
		if details.Time.Format(time.RFC3339) != testCase.Time.Format(time.RFC3339) {
			t.Errorf("%d: Times do not match, expected:%v - actual:%v", i, testCase.Time.Format(time.RFC3339), details.Time.Format(time.RFC3339))
		}
	}
}
