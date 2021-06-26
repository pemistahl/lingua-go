/*
 * Copyright © 2021 Peter M. Stahl pemistahl@gmail.com
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

package lingua

import (
	"archive/zip"
	"bytes"
	"embed"
	"fmt"
	"io"
	"strings"
)

//go:embed language-models
var languageModels embed.FS

func loadJson(language Language, ngramLength int) []byte {
	ngramName := getNgramNameByLength(ngramLength)
	isoCode := strings.ToLower(language.IsoCode639_1().String())
	zipFilePath := fmt.Sprintf("language-models/%s/%ss.json.zip", isoCode, ngramName)
	zipFileBytes, _ := languageModels.ReadFile(zipFilePath)
	zipFile, _ := zip.NewReader(bytes.NewReader(zipFileBytes), int64(len(zipFileBytes)))
	jsonFile, _ := zipFile.File[0].Open()
	defer jsonFile.Close()
	jsonFileContent, _ := io.ReadAll(jsonFile)
	return jsonFileContent
}
