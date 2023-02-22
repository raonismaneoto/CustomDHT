package storage

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/raonismaneoto/CustomDHT/commons/helpers"
	"github.com/raonismaneoto/CustomDHT/core/models"
)

type Storage struct {
	Type       models.MemType
	memStorage map[int64][]byte
	root       string
	chunkLimit int64
}

type Entry struct {
	Key  int64
	Data []byte
}

func New(t models.MemType) Storage {
	if t == models.NotSupported {
		panic("invalid MemType for the storage")
	}

	s := Storage{}
	s.memStorage = make(map[int64][]byte)
	s.Type = t
	if s.Type == models.Disk {
		os.Mkdir("./data", 0644)
		s.root = "./data"
	} else {
		os.Mkdir("./flushed-data", 0644)
		s.root = "./flushed-data"
		helpers.PeriodicInvocation(s.FlushMem, 3600)
	}
	s.chunkLimit = 10000
	return s
}

func (s *Storage) Save(data Entry) error {
	var err error
	if s.Type == models.Mem {
		err = s.saveMem(data)
	} else {
		err = s.saveDisk(data)
	}

	if err != nil {
		log.Println("error while saving data: %v", err)
		return err
	}

	return nil
}

func (s *Storage) Read(key int64) ([]byte, error) {
	var data []byte
	var err error
	if s.Type == models.Mem {
		data, err = s.readMem(key)
	} else {
		data, err = s.readDisk(key, -1, 0)
	}

	if err != nil {
		log.Println("error while reading data: %v", err)
		return nil, err
	}

	return data, nil
}

func (s *Storage) Delete(key int64) error {
	var err error
	if s.Type == models.Mem {
		err = s.deleteMem(key)
	} else {
		err = s.deleteDisk(key)
	}

	if err != nil {
		log.Println("error while deleting data: %v", err)
		return err
	}

	return nil
}

func (s *Storage) FlushMem() {
	for key, value := range s.memStorage {
		entry := Entry{Key: key, Data: value}
		err := s.saveDisk(entry)
		if err != nil {
			log.Println("error when flushing to disk")
			log.Println(err.Error())
			continue
		}
		delete(s.memStorage, key)
	}
}

func (s *Storage) ReadAsync(key int64, cbuffer chan []byte, ebuffer chan error) {
	if s.Type == models.Mem {
		ebuffer <- errors.New("async read is only available for Disk memory type")
	} else {
		s.readDiskAsync(key, cbuffer, ebuffer)
	}
}

func (s *Storage) saveMem(data Entry) error {
	s.memStorage[data.Key] = data.Data
	return nil
}

func (s *Storage) saveDisk(data Entry) error {
	f, err := os.OpenFile(s.root+"/"+fmt.Sprint(data.Key), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("unable to open/create %v", data.Key)
		return err
	}

	defer f.Close()

	if _, err := f.Write(data.Data); err != nil {
		log.Println("unable to write to %v", data.Key)
	}

	return err
}

func (s *Storage) readMem(key int64) ([]byte, error) {
	data, ok := s.memStorage[key]
	if !ok {
		data, err := s.readDisk(key, -1, 0)
		if err != nil {
			return nil, errors.New("Key not found")
		}
		return data, nil
	}
	return data, nil
}

func (s *Storage) readDisk(key, limit, offset int64) ([]byte, error) {
	f, err := os.Open(s.root + "/" + fmt.Sprint(key))
	if err != nil {
		log.Println("unable to open file %v", fmt.Sprint(key))
		log.Println(err)
		return nil, errors.New("Key not found")
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Println("unable to get file stat")
		return nil, err
	}

	var contentSize int64
	if limit == -1 {
		contentSize = stat.Size()
	} else {
		if (int64(stat.Size()) - offset) < limit {
			contentSize = (int64(stat.Size()) - offset)
		} else {
			contentSize = limit
		}
	}

	content := make([]byte, contentSize)

	if _, err := f.ReadAt(content, int64(offset)); err != nil {
		log.Println("unable to read file: ", err)
		return nil, err
	}

	return content, nil
}

func (s *Storage) deleteMem(key int64) error {
	delete(s.memStorage, key)
	return nil
}

func (s *Storage) deleteDisk(key int64) error {
	return os.Remove(s.root + "/" + fmt.Sprint(key))
}

func (s *Storage) readDiskAsync(key int64, cbuffer chan []byte, ebuffer chan error) {
	f, err := os.Open(s.root + "/" + fmt.Sprint(key))
	if err != nil {
		log.Println("unable to open file %v", fmt.Sprint(key))
		log.Println(err)
		ebuffer <- errors.New("Key not found")
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Println("unable to get file stat")
		ebuffer <- err
	}

	limit := s.chunkLimit
	offset := int64(0)

	chuncks := int32(math.Ceil(float64(stat.Size()) / float64(s.chunkLimit)))
	for i := int32(0); i < chuncks; i++ {
		var contentSize int64

		if (int64(stat.Size()) - offset) < limit {
			contentSize = (int64(stat.Size()) - offset)
		} else {
			contentSize = limit
		}

		content := make([]byte, contentSize)

		if _, err := f.ReadAt(content, int64(offset)); err != nil {
			log.Println("unable to read file: ", err)
			ebuffer <- err
		}

		cbuffer <- content

		offset += contentSize
	}

	close(cbuffer)
	close(ebuffer)
}
