package records

import (
	"errors"
	"strconv"
)

const (
	messageErr        string = "Записи с таким номером нет в списке"
	messageErrRecords string = "Неверно введен номер записи"
)

func CorrectEntry(recordNumber, entry string) ([]string, error) {

	index, err := strconv.Atoi(recordNumber)
	if err != nil {
		return nil, errors.New(messageErrRecords)
	}

	list := GetRecords()
	if len(list) < index || index <= 0 {
		return nil, errors.New(messageErr)
	}

	list[index-1] = entry

	return list, nil
}
