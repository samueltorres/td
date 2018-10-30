package boards

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

// Board reprents a todo board
type Board struct {
	Tasks []*Task `json:"tasks"`
}

// Task represents a task from the user
type Task struct {
	ID          string `json:"id"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

type fileBoardStorage struct {
	path   string
	boards map[string]*Board
}

// BoardStorage represents the generic implementation of a todo board storage
type BoardStorage interface {
	// Boards
	GetBoards() ([]string, error)
	AddBoard(board string) error
	//RemoveBoard(board string) error

	// Tasks
	AddTask(board string, description string) error
	SetTaskStatus(board string, id string, status bool) error
	RemoveTask(board string, id string) error
	RemoveTasks(board string) error
	GetTasks(board string) ([]*Task, error)
}

type boardsFileContent struct {
	Boards map[string]*Board `json:"boards"`
}

// NewBoardFileStorage creates a new board file storage
func NewBoardFileStorage(bc BoardConfig) (BoardStorage, error) {
	sc := &boardsFileContent{}

	bfc, err := ioutil.ReadFile(bc.BoardsFile)
	if err != nil {
		// file does not exist we need to create it
		sc, err = initBoardsFile(bc.BoardsFile)

		if err != nil {
			log.Fatalf("Could not create file %s", bc.BoardsFile)
		}
	} else {
		err = json.Unmarshal(bfc, &sc)
		if err != nil {
			// error parsing the file
			log.Fatalf("Error loading todo board")
		}
	}

	s := &fileBoardStorage{
		path:   bc.BoardsFile,
		boards: sc.Boards,
	}

	return s, nil
}

func initBoardsFile(path string) (*boardsFileContent, error) {
	sc := &boardsFileContent{
		Boards: map[string]*Board{
			"default": &Board{
				Tasks: []*Task{},
			},
		},
	}
	data, err := json.Marshal(sc)
	if err != nil {
		return sc, err
	}

	err = ioutil.WriteFile(path, data, 0644)

	return sc, err
}

func (bs *fileBoardStorage) GetBoards() ([]string, error) {
	boardNames := make([]string, len(bs.boards))
	i := 0

	for k := range bs.boards {
		boardNames[i] = k
		i++
	}

	sort.Strings(boardNames)
	return boardNames, nil
}

func (bs *fileBoardStorage) AddBoard(name string) error {
	if bs.boards[name] != nil {
		return errors.New("Board already exists")
	}

	bs.boards[name] = &Board{
		Tasks: []*Task{},
	}

	return bs.Sync()
}

func (bs *fileBoardStorage) AddTask(board string, description string) error {
	b := bs.boards[board]
	if b == nil {
		return nil
	}

	b.Tasks = append(b.Tasks, &Task{ID: GenerateID(2), Description: description, Status: false})

	return bs.Sync()
}

func (bs *fileBoardStorage) SetTaskStatus(board string, id string, status bool) error {
	b := bs.boards[board]
	if b == nil {
		return nil //return the error, we need to define some errors
	}

	var t *Task

	for _, tt := range b.Tasks {
		if tt.ID == id {
			t = tt
			break
		}
	}

	if t == nil {
		return nil //return the error
	}

	t.Status = status

	return bs.Sync()
}

func (bs *fileBoardStorage) RemoveTask(board string, id string) error {
	b := bs.boards[board]
	if b == nil {
		return fmt.Errorf("board does not exist")
	}

	idx := -1
	for i := 0; i < len(b.Tasks); i++ {
		if b.Tasks[i].ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		return fmt.Errorf("task does not exist")
	}

	b.Tasks = append(b.Tasks[:idx], b.Tasks[idx+1:]...)

	return bs.Sync()
}

func (bs *fileBoardStorage) RemoveTasks(board string) error {
	b := bs.boards[board]
	if b == nil {
		return fmt.Errorf("board does not exist")
	}

	b.Tasks = []*Task{}

	return bs.Sync()
}

func (bs *fileBoardStorage) GetTasks(board string) ([]*Task, error) {
	b := bs.boards[board]
	if b == nil {
		return nil, fmt.Errorf("board does not exist")
	}

	return b.Tasks, nil
}

func (bs *fileBoardStorage) Sync() error {
	sc := boardsFileContent{
		Boards: bs.boards,
	}

	data, err := json.Marshal(sc)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(bs.path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
