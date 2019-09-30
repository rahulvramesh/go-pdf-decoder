package decoder

import (
	"encoding/base64"
	"fmt"
	"os"
)

func Decode(pdfBase64, filename string) error {

	if pdfBase64 == "" {
		return fmt.Errorf("no pdf base64 provided")
	}

	if filename == "" {
		return fmt.Errorf("output pdf not provided")
	}

	dec, err := base64.StdEncoding.DecodeString(pdfBase64)

	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		return err
	}

	if _, err := f.Write(dec); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}
