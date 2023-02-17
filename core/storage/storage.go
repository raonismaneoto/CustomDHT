package storage

import (
	"errors"
	"log"
	"os"
)

type MemType int32

const (
	Mem MemType = iota
	Disk
)

type Storage struct {
	Type       MemType
	memStorage map[int64][]byte
	root       string
}

func New(t MemType) Storage {
	s := Storage{}
	s.memStorage = make(map[int64][]byte)
	s.Type = t
	if s.Type == Disk {
		os.Mkdir("./data", 0644)
		s.root = "./data"
	}
	return s
}

func (s *Storage) Save(data struct {
	key  int64
	data []byte
}) error {
	var err error
	if s.Type == Mem {
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
	if s.Type == Mem {
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

func (s *Storage) saveMem(data struct {
	key  int64
	data []byte
}) error {
	s.memStorage[data.key] = data.data
	return nil
}

func (s *Storage) saveDisk(data struct {
	key  int64
	data []byte
}) error {
	f, err := os.OpenFile(s.root+"/"+string(data.key), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("unable to open/create %v", data.key)
		return err
	}

	defer f.Close()

	if _, err := f.Write(data.data); err != nil {
		log.Println("unable to write to %v", data.key)
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
	f, err := os.Open(s.root + "/" + string(key))
	if err != nil {
		log.Println("unable to open file %v", string(key))
		log.Println(err)
		return nil, err
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
