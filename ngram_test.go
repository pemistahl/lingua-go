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
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeOfLowerOrderNgrams(t *testing.T) {
	n := newNgram("äbcde")
	assert.Equal(
		t,
		[]ngram{
			newNgram("äbcde"),
			newNgram("äbcd"),
			newNgram("äbc"),
			newNgram("äb"),
			newNgram("ä"),
		},
		n.rangeOfLowerOrderNgrams())
}

func TestNgram_MarshalJSON(t *testing.T) {
	serialized, err := json.Marshal(newNgram("äbcde"))
	assert.Equal(t, "\"äbcde\"", string(serialized))
	assert.Equal(t, nil, err)
}

func TestNgram_UnmarshalJSON(t *testing.T) {
	var ngram ngram
	err := json.Unmarshal([]byte("\"äbcde\""), &ngram)
	assert.Equal(t, newNgram("äbcde"), ngram)
	assert.Equal(t, nil, err)
}
