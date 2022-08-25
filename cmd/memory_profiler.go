/*
 * Copyright © 2021-today Peter M. Stahl pemistahl@gmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either expressed or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/pemistahl/lingua-go"
	"runtime"
)

func main() {
	lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MB\n", megabytes(m.Alloc))
	fmt.Printf("TotalAlloc = %v MB\n", megabytes(m.TotalAlloc))
	fmt.Printf("Sys = %v MB\n", megabytes(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func megabytes(bytes uint64) uint64 {
	return bytes / 1000000
}
