/* MIT License

Copyright (c) 2020 Lluis Sanchez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package process

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsLicenseHeader(t *testing.T) {
	assert.True(t, containsLicenseHeader(fakeFileWithTargetLicenseHeader))
	assert.True(t, containsLicenseHeader(fakeFileWithDifferentLicenseHeader))
	assert.False(t, containsLicenseHeader(fakeFileWithoutLicense))
}

func TestExtractHeader(t *testing.T) {
	expected := strings.TrimSpace(fakeTargetLicenseHeader)
	input := fakeFileWithTargetLicenseHeader
	output := extractHeader(input)
	assert.True(t, output == expected)

	expected = "/* copyright */"
	input = "/* copyright */\nlorem ipsum dolor sit amet"
	output = extractHeader(input)
	assert.True(t, output == expected)
}

func TestRemoveHeader(t *testing.T) {
	expected := fakeFileWithoutLicense

	input := fakeFileWithTargetLicenseHeader
	output := removeHeader(input)
	assert.True(t, output == expected)

	input = fakeFileWithDifferentLicenseHeader
	output = removeHeader(input)
	assert.True(t, output == expected)
}

func TestInsertHeader(t *testing.T) {
	expected := fakeFileWithTargetLicenseHeader
	input := fakeFileWithoutLicense
	header := fakeTargetLicenseHeader
	output := insertHeader(input, header)
	assert.True(t, output == expected)
}

func TestReplaceHeader(t *testing.T) {
	expected := fakeFileWithTargetLicenseHeader
	input := fakeFileWithDifferentLicenseHeader
	header := fakeTargetLicenseHeader
	output := replaceHeader(input, header)
	assert.True(t, output == expected)
}
