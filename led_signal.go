package homesound

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gurupras/go-easyfiles"
)

type LedHandler struct {
	basePath string
	onValue  string
	offValue string
}

func NewRpiZeroLedHandler() *LedHandler {
	return &LedHandler{"/sys/class/leds/led0", "0", "1"}
}

func (lh *LedHandler) GetLedFile(filename string) string {
	return filepath.Join(lh.basePath, filename)
}

func (lh *LedHandler) ReadFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	return string(data), err
}

func (lh *LedHandler) WriteFile(path string, value string) error {
	f, err := easyfiles.Open(path, os.O_WRONLY, easyfiles.GZ_FALSE)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to open file '%v': %v", path, err))
		return err
	}
	defer f.Close()

	writer, err := f.Writer(0)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to get writer to file '%v': %v", path, err))
		return err
	}
	defer writer.Close()
	defer writer.Flush()

	if _, err = writer.Write([]byte(value)); err != nil {
		err = errors.New(fmt.Sprintf("Failed to write '%v' to file '%v': %v", value, path, err))
		return err
	}
	return nil
}
func (lh *LedHandler) LedOn() error {
	return lh.WriteFile(lh.GetLedFile("brightness"), lh.onValue)
}

func (lh *LedHandler) LedOff() error {
	return lh.WriteFile(lh.GetLedFile("brightness"), lh.offValue)
}

func (lh *LedHandler) BlinkLed(period time.Duration, duration time.Duration) error {
	startTime := time.Now()
	for {
		timeNow := time.Now()
		if timeNow.Sub(startTime) > duration {
			break
		} else {
			if err := lh.LedOn(); err != nil {
				return err
			}
			time.Sleep(period)
			if err := lh.LedOff(); err != nil {
				return err
			}
			time.Sleep(period)
		}
	}
	return nil
}
