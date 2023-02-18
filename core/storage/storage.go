package storage

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/raonismaneoto/CustomDHT/core/models"
)

type Storage struct {
	Type       models.MemType
	memStorage map[int64][]byte
	root       string
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
	}
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
		data, err = s.readDisk(key)
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
		return nil, errors.New("Key not found")
	}
	return data, nil
}

func (s *Storage) readDisk(key int64) ([]byte, error) {
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

	content := make([]byte, int32(stat.Size()))

	if _, err := f.Read(content); err != nil {
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
