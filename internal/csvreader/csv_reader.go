package csvreader

import (
	"UEBA/config"
	"UEBA/internal/domain"
	"encoding/csv"
	"fmt"
	"go.uber.org/zap"
	"os"
)

type CSVReader struct {
	logger *zap.Logger
	path   string
	format rune
}

func NewCsvReader(cfg *config.Config, logger *zap.Logger) *CSVReader {
	formatRune := []rune(cfg.ScanningOpts.Format)[0]
	return &CSVReader{path: cfg.ScanningOpts.Path, format: formatRune, logger: logger}
}

func (CR *CSVReader) GetRecordByID(id string) (*domain.UEBAS, error) {
	file, err := os.Open(CR.path)
	if err != nil {
		CR.logger.Error("Failed to open file", zap.Error(err))
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = CR.format

	records, err := reader.ReadAll()
	if err != nil {
		CR.logger.Error("Failed to read CSV", zap.Error(err))
		return nil, err
	}

	for _, record := range records {
		if record[1] == id {
			return &domain.UEBAS{
				SequenceNumber:            record[0],
				ID:                        record[1],
				UID:                       record[2],
				Domain:                    record[3],
				CN:                        record[4],
				Department:                record[5],
				Title:                     record[6],
				Who:                       record[7],
				LogonCount:                record[8],
				NumLogons7:                record[9],
				NumShare7:                 record[10],
				NumFile7:                  record[11],
				NumAd7:                    record[12],
				NumN7:                     record[13],
				NumLogons14:               record[14],
				NumShare14:                record[15],
				NumFile14:                 record[16],
				NumAd14:                   record[17],
				NumN14:                    record[18],
				NumLogons30:               record[19],
				NumShare30:                record[20],
				NumFile30:                 record[21],
				NumAd30:                   record[22],
				NumN30:                    record[23],
				NumLogons150:              record[24],
				NumShare150:               record[25],
				NumFile150:                record[26],
				NumAd150:                  record[27],
				NumN150:                   record[28],
				NumLogons365:              record[29],
				NumShare365:               record[30],
				NumFile365:                record[31],
				NumAd365:                  record[32],
				NumN365:                   record[33],
				HasUserPrincipalName:      record[34],
				HasMail:                   record[35],
				HasPhone:                  record[36],
				FlagDisabled:              record[37],
				FlagLockout:               record[38],
				FlagPasswordNotRequired:   record[39],
				FlagPasswordCantChange:    record[40],
				FlagDontExpirePassword:    record[41],
				OwnedFiles:                record[42],
				NumMailboxes:              record[43],
				NumMemberOfGroups:         record[44],
				NumMemberOfIndirectGroups: record[45],
				MemberOfIndirectGroupsIDs: record[46],
				MemberOfGroupsIDs:         record[47],
				IsAdmin:                   record[48],
				IsService:                 record[49],
			}, nil
		}
	}

	err = fmt.Errorf("record with ID %s not found", id)
	CR.logger.Error("Failed to find record", zap.Error(err))
	return nil, err

}
