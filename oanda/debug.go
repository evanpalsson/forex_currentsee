// Copyright 2014 Tjerk Santegoeds
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

// +build debug

package oanda

import (
	"io"
	"log"
	"os"
)

var dbgLogger *log.Logger

func init() {
	dbgLogger = log.New(os.Stderr, "DEBUG: ", log.LstdFlags|log.Lmicroseconds)
}

func debug(fmt string, args ...interface{}) {
	dbgLogger.Printf(fmt, args...)
}

func trace(rdr io.Reader) io.Reader {
	return io.TeeReader(rdr, os.Stderr)
}
