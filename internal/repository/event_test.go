package repository_test

import (
	"log"
	"sync"
	"testing"
	"time"
)

type Batch struct {
	sync.Mutex

	BatchData []int
	CollName  string
}

type BatchAdvanced struct {
	sync.Mutex

	BatchData []ToSaveData
	CollName  string
}

type ToSaveData struct {
	ID       int
	CollName string
}

func Test_Batch(t *testing.T) {
	log.Println("start testing")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	batch := &Batch{}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			batch.Lock()
			commit(batch.BatchData)
			batch.BatchData = nil
			batch.Unlock()
		}
	}()

	for _, v := range data {
		batch.Lock()
		batch.BatchData = append(batch.BatchData, v)
		if len(batch.BatchData) == 2 {
			commit(batch.BatchData)
			batch.BatchData = nil
		}
		batch.Unlock()
		time.Sleep(1 * time.Second)
	}

	log.Println("end testing")
}

func Test_BatchAdvanced(t *testing.T) {
	log.Println("start testing")
	toSaveData := make([]ToSaveData, 0, 10)

	for i := 0; i < 27; i++ {
		toSaveData = append(toSaveData, ToSaveData{
			ID:       i,
			CollName: getCollName(i),
		})
	}

	batchColA := &BatchAdvanced{}
	batchColB := &BatchAdvanced{}

	batchMaps := map[string]*BatchAdvanced{
		"A": batchColA,
		"B": batchColB,
	}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			for _, v := range batchMaps {
				v.Lock()
				if len(v.BatchData) > 0 {
					commitAdvanced(v.BatchData)
					v.BatchData = nil
				}
				v.Unlock()
			}
		}
	}()

	for _, v := range toSaveData {
		batch, ok := batchMaps[v.CollName]
		if !ok {
			log.Println("error: no batch found, should be auto creating here")
			batchMaps[v.CollName] = &BatchAdvanced{}
			continue
		}

		batch.Lock()
		batch.BatchData = append(batch.BatchData, v)
		if len(batch.BatchData) == 5 {
			commitAdvanced(batch.BatchData)
			batch.BatchData = nil
		}
		batch.Unlock()
		// time.Sleep(1 * time.Second)
	}

	log.Println("waiting to end testing")
	time.Sleep(5 * time.Second)
	log.Println("end testing")
}

func commit(batch []int) {
	log.Printf("committing batch: %+v\n", batch)
}

func commitAdvanced(batch []ToSaveData) {
	log.Printf("committing batch of coll: %+v\n", batch)
}

func getCollName(ID int) string {
	if ID%2 == 0 {
		return "A"
	}
	return "C"
}
