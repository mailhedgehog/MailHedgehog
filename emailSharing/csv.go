package emailSharing

import (
	"encoding/csv"
	"errors"
	"github.com/mailhedgehog/MailHedgehog/logger"
	"io"
	"os"
	"time"
)

type CsvEmailSharing struct {
	csvFilePath string
}

func CreateSharingEmailUsingCSV(csvFilePath string) *CsvEmailSharing {
	csvFile := &CsvEmailSharing{
		csvFilePath: csvFilePath,
	}

	return csvFile
}

func (csvEmailSharing *CsvEmailSharing) Find(id string) (*EmailSharingRecord, error) {
	f, err := os.Open(csvEmailSharing.csvFilePath)
	logger.PanicIfError(err)

	defer f.Close()

	csvReader := csv.NewReader(f)

	var emailSharingRecord *EmailSharingRecord

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		logger.PanicIfError(err)

		if rec[0] != id {
			continue
		}

		expiredAt, err := time.Parse("2006-01-02 15:04:05", rec[3])

		if err == nil && expiredAt.After(time.Now().UTC()) {
			emailSharingRecord = &EmailSharingRecord{
				Id:        rec[0],
				Room:      rec[1],
				EmailId:   rec[2],
				ExpiredAt: expiredAt,
			}

			break
		}
	}

	if emailSharingRecord != nil && emailSharingRecord.Id == id {
		return emailSharingRecord, nil
	}

	return nil, errors.New("row not found")
}
