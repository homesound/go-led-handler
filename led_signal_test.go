package homesound

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// FIXME: Currently, all of these tests are hard-coded for the raspberry pi
func commonTest(lh *LedHandler, require *require.Assertions) {
	err := lh.WriteFile(lh.GetLedFile("trigger"), "default-on")
	require.Nil(err, "Failed to set trigger to default-on", err)

	data, err := ioutil.ReadFile(lh.GetLedFile("trigger"))
	require.Nil(err, "Failed to read trigger file", err)

	require.True(strings.Contains(string(data), "[default-on]"), "Did not set trigger to default-on")
}

func TestLedOn(t *testing.T) {
	require := require.New(t)

	lh := NewRpiZeroLedHandler()
	commonTest(lh, require)

	err := lh.WriteFile(lh.GetLedFile("brightness"), "1")
	require.Nil(err, "Failed to turn off LED", err)
	v, err := lh.ReadFile(lh.GetLedFile("brightness"))
	require.Nil(err)
	require.Equal("255", strings.TrimSpace(v))

	err = lh.LedOn()
	require.Nil(err, "Failed to turn on LED", err)

	v, err = lh.ReadFile(lh.GetLedFile("brightness"))
	require.Nil(err)
	require.Equal("0", strings.TrimSpace(v))
}

func TestLedOff(t *testing.T) {
	require := require.New(t)

	lh := NewRpiZeroLedHandler()
	commonTest(lh, require)

	err := lh.WriteFile(lh.GetLedFile("trigger"), "none")
	require.Nil(err)

	err = lh.WriteFile(lh.GetLedFile("brightness"), "0")
	require.Nil(err, "Failed to turn on LED", err)
	v, err := lh.ReadFile(lh.GetLedFile("brightness"))
	require.Nil(err)
	require.Equal("0", strings.TrimSpace(v))

	err = lh.LedOff()
	require.Nil(err, "Failed to turn off LED", err)

	v, err = lh.ReadFile(lh.GetLedFile("brightness"))
	require.Nil(err)
	require.Equal("255", strings.TrimSpace(v))
}

func TestBlinkLed(t *testing.T) {
	require := require.New(t)

	lh := NewRpiZeroLedHandler()
	commonTest(lh, require)

	err := lh.BlinkLed(1*time.Second, 5*time.Second)
	require.Nil(err, "Failed to blink LED", err)
}
