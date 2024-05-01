package mappers

import (
	"SEP/internal/models/dataModels"
	"SEP/internal/utils"
)

type RecordMapper struct {
}

func (rm *RecordMapper) AddRecord(record *dataModels.Record) error {
	result := utils.DB.Create(record)
	return result.Error
}

func (rm *RecordMapper) DeleteRecord(record *dataModels.Record) error {
	result := utils.DB.Delete(record)
	return result.Error
}

func (rm *RecordMapper) UpdateRecord(record *dataModels.Record) error {
	result := utils.DB.Save(record)
	return result.Error
}

func (rm *RecordMapper) GetAllRecords() ([]*dataModels.Record, error) {
	var records []*dataModels.Record
	result := utils.DB.Find(&records)
	return records, result.Error
}

func (rm *RecordMapper) GetRecordsByUserId(userId int) ([]*dataModels.Record, error) {
	var records []*dataModels.Record
	result := utils.DB.Find(&records, "user_id=?", userId)
	return records, result.Error
}

func (rm *RecordMapper) GetRecordById(recordId int) (*dataModels.Record, error) {
	var record *dataModels.Record
	result := utils.DB.First(&record, "id=?", recordId)
	return record, result.Error
}

func (rm *RecordMapper) GetRecordsByType(recordType string) ([]*dataModels.Record, error) {
	var records []*dataModels.Record
	result := utils.DB.Find(&records, "type=?", recordType)
	return records, result.Error
}

func (rm *RecordMapper) GetRecordsByPatientName(patientName string) ([]*dataModels.Record, error) {
	var records []*dataModels.Record
	result := utils.DB.Find(&records, "patient_name=?", patientName)
	return records, result.Error
}
