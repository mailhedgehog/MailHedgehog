package emailSharing

import (
	"encoding/csv"
	"errors"
	"github.com/google/uuid"
	"github.com/mailhedgehog/logger"
	"io"
	"os"
	"time"
)

var ExpiredAdFormat = "2006-01-02 15:04:05"

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

		expiredAt, err := time.Parse(ExpiredAdFormat, rec[3])

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

func (csvEmailSharing *CsvEmailSharing) Create(emailSharingRecord *EmailSharingRecord) (*EmailSharingRecord, error) {
	f, err := os.OpenFile(csvEmailSharing.csvFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	logger.PanicIfError(err)

	defer f.Close()

	emailSharingRecord.Id = uuid.New().String()

	row := []string{
		emailSharingRecord.Id,
		emailSharingRecord.Room,
		emailSharingRecord.EmailId,
		emailSharingRecord.ExpiredAt.Format(ExpiredAdFormat),
	}

	csvWriter := csv.NewWriter(f)
	err = csvWriter.Write(row)
	if err != nil {
		return nil, err
	}
	csvWriter.Flush()

	return emailSharingRecord, nil
}

func (csvEmailSharing *CsvEmailSharing) DeleteExpired() (bool, error) {

	tmpFile := csvEmailSharing.csvFilePath + ".tmp"

	f, err := os.Open(csvEmailSharing.csvFilePath)
	logger.PanicIfError(err)
	defer f.Close()

	outFile, err := os.Create(tmpFile)
	logger.PanicIfError(err)
	defer outFile.Close()

	csvReader := csv.NewReader(f)
	csvWriter := csv.NewWriter(outFile)

	rowFound := false

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		logger.PanicIfError(err)

		expiredAt, err := time.Parse(ExpiredAdFormat, rec[3])

		if err == nil && expiredAt.After(time.Now().UTC()) {
			_ = csvWriter.Write(rec)
			continue
		}

		rowFound = true
	}

	csvWriter.Flush()
	f.Close()
	outFile.Close()

	defer os.Remove(tmpFile)
	_ = os.Remove(csvEmailSharing.csvFilePath)

	_ = os.Rename(tmpFile, csvEmailSharing.csvFilePath)

	return rowFound, nil
}
