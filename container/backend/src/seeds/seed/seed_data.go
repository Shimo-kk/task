package seed

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type ModelInterface interface{}

func SeedData(url string) error {
	return nil
}

func seed(db *gorm.DB, records []ModelInterface) error {
	for _, record := range records {
		result := db.First(record)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			// その他のエラーが発生した場合
			return result.Error
		}

		// エラーがない場合、もしくはレコードが見つからない場合は挿入を試みる
		if result.Error == gorm.ErrRecordNotFound {
			err := db.Create(record).Error
			if err != nil {
				return err
			}
			fmt.Printf("Inserted: %+v\n", record)
			continue
		}

		existingRecord := reflect.ValueOf(record).Elem()

		// 同じIDのものが存在し、かつ変更されている場合は更新
		newRecord := reflect.ValueOf(record).Elem()
		if existingRecord.FieldByName("Id").Interface() == newRecord.FieldByName("Id").Interface() && !reflect.DeepEqual(existingRecord.Interface(), newRecord.Interface()) {
			err := db.Model(record).Updates(record).Error
			if err != nil {
				return err
			}
			fmt.Printf("Updated: %+v\n", record)
		} else {
			// 同じIDのものが存在していて、かつ変更されていない場合は何もしない
			fmt.Printf("Skipped: %+v\n", record)
		}
	}

	return nil
}
