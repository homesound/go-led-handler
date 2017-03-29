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

func GetLedFile(filename string) string {
	return filepath.Join("/sys/class/leds/led0", filename)
}

func ReadFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	return string(data), err
}

func WriteFile(path string, value string) error {
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
func LedOn() error {
	return WriteFile(GetLedFile("brightness"), "0")
}

func LedOff() error {
	return WriteFile(GetLedFile("brightness"), "1")
}

func BlinkLed(frequency, durationMillis int) {
	period := float32(1) / float32(frequency)

	startTime := time.Now()
	for {
		timeNow := time.Now()
		if timeNow.Sub(startTime) > time.Duration(durationMillis)*time.Millisecond {
			break
		} else {
			LedOn()
			time.Sleep(time.Duration(period) * time.Nanosecond)
			LedOff()
		}
	}
}
