package scheduler

import (
	"backend/models"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

func init() {
	db := models.GetAccountDB()
	go cleanDB(db)
}

func cleanDB(db *leveldb.DB) {
	for {
		iter := db.NewIterator(nil, nil)
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()
			valueByte := []byte(value)
			timeStampByte := valueByte[0:4]
			expiryTime := binary.LittleEndian.Uint32(timeStampByte)
			currentTime := uint32(time.Now().Unix())
			if currentTime > expiryTime {
				fmt.Printf("Cleaning cache for user: %s\n", string(key))
				db.Delete(key, nil)
			}
		}
		iter.Release()
		db.Close()
	}

}
