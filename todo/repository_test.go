package todo

import (
	"os"
	"strconv"
	"testing"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/essentier/spickspan"
)

func tSaveNew(t *testing.T) {
	//t.Parallel()
	originalVal := os.Getenv("SPICKSPAN_MODE")
	defer os.Setenv("SPICKSPAN_MODE", originalVal)
	os.Setenv("SPICKSPAN_MODE", "local")

	repo, dbSession := getRepo(t)
	defer dbSession.Close()

	todo := Todo{Name: "todo1", Completed: false}
	savedTodo, err := repo.saveNew(todo)
	handleFatalError(t, "Failed to find todo by ID.", err)

	retrievedTodo, err := repo.findById(savedTodo.Id)
	handleFatalError(t, "Failed to find todo by ID.", err)
	if retrievedTodo.Name != "todo1" {
		t.Errorf("Expect todo's name to be todo1 but got %v", retrievedTodo.Name)
	}
}

func getRepo(t *testing.T) (repository, *mgo.Session) {
	provider, err := spickspan.GetDefaultServiceProvider()
	handleFatalError(t, "Failed to get spickspan provider.", err)

	mgoService, err := spickspan.GetMongoDBService(provider, "todo-db")
	handleFatalError(t, "Failed to get MongoDB service.", err)

	mgoUrl := mgoService.IP + ":" + strconv.Itoa(mgoService.Port)
	dbSession, err := createDBSession(mgoUrl)
	handleFatalError(t, "Failed to create DB session.", err)

	reqSession := dbSession.Clone()
	db := reqSession.DB("tododb")
	repo := getRepository(db)
	return repo, reqSession
}

func createDBSession(url string) (*mgo.Session, error) {
	dialInfo, err := mgo.ParseURL(url)
	if err != nil {
		return nil, err
	}

	dialInfo.FailFast = false
	dialInfo.Timeout = 100 * time.Second
	return mgo.DialWithInfo(dialInfo)
}

func handleFatalError(t *testing.T, message string, err error) {
	if err != nil {
		t.Fatalf("%v Error is %v", message, err)
	}
}
