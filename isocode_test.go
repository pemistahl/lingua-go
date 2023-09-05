/*
 * Copyright © 2021-present Peter M. Stahl pemistahl@gmail.com
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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIsoCode639_1FromValue(t *testing.T) {
	assert.Equal(t, AF, GetIsoCode639_1FromValue("AF"))
	assert.Equal(t, AF, GetIsoCode639_1FromValue("af"))
	assert.Equal(t, ZU, GetIsoCode639_1FromValue("ZU"))
	assert.Equal(t, ZU, GetIsoCode639_1FromValue("zu"))
	assert.Equal(t, UnknownIsoCode639_1, GetIsoCode639_1FromValue("xy"))
	assert.Equal(t, UnknownIsoCode639_1, GetIsoCode639_1FromValue(""))
}

func TestGetIsoCode639_3FromValue(t *testing.T) {
	assert.Equal(t, AFR, GetIsoCode639_3FromValue("AFR"))
	assert.Equal(t, AFR, GetIsoCode639_3FromValue("afr"))
	assert.Equal(t, ZUL, GetIsoCode639_3FromValue("ZUL"))
	assert.Equal(t, ZUL, GetIsoCode639_3FromValue("zul"))
	assert.Equal(t, UnknownIsoCode639_3, GetIsoCode639_3FromValue("xyz"))
	assert.Equal(t, UnknownIsoCode639_3, GetIsoCode639_3FromValue(""))
}
