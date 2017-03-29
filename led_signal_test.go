package homesound

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func commonTest(require *require.Assertions) {
	err := WriteFile(GetLedFile("trigger"), "default-on")
	require.Nil(err, "Failed to set trigger to default-on", err)

	data, err := ioutil.ReadFile(GetLedFile("trigger"))
	require.Nil(err, "Failed to read trigger file", err)

	require.True(strings.Contains(string(data), "[default-on]"), "Did not set trigger to default-on")
}

func TestLedOn(t *testing.T) {
	require := require.New(t)

	commonTest(require)

	err := WriteFile(GetLedFile("brightness"), "1")
	require.Nil(err, "Failed to turn off LED", err)
	v, err := ReadFile(GetLedFile("brightness"))
	require.Nil(err)
	require.Equal("255", strings.TrimSpace(v))

	err = LedOn()
	require.Nil(err, "Failed to turn on LED", err)

	v, err = ReadFile(GetLedFile("brightness"))
	require.Nil(err)
	require.Equal("0", strings.TrimSpace(v))
}

func TestLedOff(t *testing.T) {
	require := require.New(t)

	commonTest(require)

	err := WriteFile(GetLedFile("trigger"), "none")
	require.Nil(err)

	err = WriteFile(GetLedFile("brightness"), "0")
	require.Nil(err, "Failed to turn on LED", err)
	v, err := ReadFile(GetLedFile("brightness"))
	require.Nil(err)
	require.Equal("0", strings.TrimSpace(v))

	err = LedOff()
	require.Nil(err, "Failed to turn off LED", err)

	v, err = ReadFile(GetLedFile("brightness"))
	require.Nil(err)
	require.Equal("255", strings.TrimSpace(v))
}
