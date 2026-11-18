package utils

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func ValidateFindErrorMessage(err error) error {
	if err != nil {
		//data tidak ditemukan untuk gorm
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		//data tidak ditemukan untuk mongo
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}

		//data tidak ditemukan untuk redis
		if errors.Is(err, redis.Nil) {
			return nil
		}
		return err
	}

	return err
}
