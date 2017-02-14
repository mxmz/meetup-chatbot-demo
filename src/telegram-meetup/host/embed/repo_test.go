package embed

import (
	"log"
	"math/rand"
	"os"
	. "telegram-meetup/types"
	types "telegram-meetup/types"
	"testing"
	"time"

	"database/sql"

	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

type envMock map[string]string

func (e *envMock) Get(k string) string {
	return (*e)[k]
}

func must(f func() error) {
	err := f()
	if err != nil {
		panic(err)
	}
}

func init() {

}

func dbFixture() (*sql.DB, string) {
	var dbFile, err = ioutil.TempFile("", "embedrepotest")
	if err != nil {
		panic(err)
	}
	log.Println(dbFile)
	db, err := sql.Open("sqlite3", dbFile.Name())
	//db, err := sql.Open("ql", dbFile.Name())
	if err != nil {
		panic(err)
	}
	must(func() error { return DBCreate(db) })
	return db, dbFile.Name()
}

func TestNewRepo(t *testing.T) {
	RegisterTestingT(t)
	env := &envMock{}
	var db *sql.DB
	_, err := NewRepo(env, db)
	Expect(err.Error()).To(BeEquivalentTo("undefined db"))

	db, dbPath := dbFixture()
	defer os.Remove(dbPath)
	repo, err := NewRepo(env, db)
	Expect(err).To(BeNil())
	Expect(repo).NotTo(BeNil())
}

func TestSetUserProperties_ok(t *testing.T) {
	RegisterTestingT(t)
	env := &envMock{}
	db, dbPath := dbFixture()
	defer os.Remove(dbPath)
	_ = dbPath
	repo, err := NewRepo(env, db)
	Expect(err).To(BeNil())
	Expect(repo).NotTo(BeNil())
	err = repo.SetUserProperties(context.TODO(), types.UserID("utente1"), types.KVMap{
		"a1": types.Value("v1"),
		"a2": types.Value("v2"),
	})
	Expect(err).To(BeNil())
	{
		rows, err := db.Query("SELECT key, value FROM node WHERE kind='user' AND id=?", "utente1")
		Expect(err).To(BeNil())
		var res = map[string]string{}
		for rows.Next() {
			var value string
			var name string
			err = rows.Scan(&name, &value)
			Expect(err).To(BeNil())
			res[name] = value
		}
		Expect(res).To(BeEquivalentTo(map[string]string{
			"a1": "v1",
			"a2": "v2"}))
	}
	err = repo.SetUserProperties(context.TODO(), types.UserID("utente1"), types.KVMap{
		"a1": types.Value("v1_2"),
		"a2": types.Value("v2_2"),
	})
	Expect(err).To(BeNil())
	{
		rows, err := db.Query("SELECT key, value FROM node WHERE kind='user' AND id=?", "utente1")
		Expect(err).To(BeNil())
		var res = map[string]string{}
		for rows.Next() {
			var value string
			var name string
			err = rows.Scan(&name, &value)
			Expect(err).To(BeNil())
			res[name] = value
		}
		Expect(res).To(BeEquivalentTo(map[string]string{
			"a1": "v1_2",
			"a2": "v2_2"}))
	}
}

func TestGetUserProperties_ok(t *testing.T) {
	RegisterTestingT(t)
	env := &envMock{}
	db, dbPath := dbFixture()
	defer os.Remove(dbPath)
	_ = dbPath
	repo, err := NewRepo(env, db)
	Expect(err).To(BeNil())
	Expect(repo).NotTo(BeNil())
	Expect(err).To(BeNil())
	{
		_, err = db.Exec("INSERT INTO node (kind,id,key,value) values('user',?,?,?)", "utente1", "a1", "newv1")
		Expect(err).To(BeNil())
		_, err = db.Exec("INSERT INTO node (kind,id,key,value) values('user',?,?,?)", "utente1", "a2", "newv2")
		Expect(err).To(BeNil())
	}
	props, err := repo.GetUserProperties(context.TODO(), types.UserID("utente1"), types.Key("a1"), types.Key("a2"))
	Expect(props).To(BeEquivalentTo(types.KVMap{
		"a1": "newv1",
		"a2": "newv2"}))

}

func TestSetMeetupUserLinkProperties_ok(t *testing.T) {
	RegisterTestingT(t)
	env := &envMock{}
	db, dbPath := dbFixture()
	defer func() {
		//log.Println(dbPath)
		os.Remove(dbPath)
	}()

	_ = dbPath
	repo, err := NewRepo(env, db)
	Expect(err).To(BeNil())
	Expect(repo).NotTo(BeNil())
	err = repo.SetMeetupUserLinkProperties(context.TODO(),
		types.MeetupUserLink{types.MeetupID("meetup1"), types.UserID("utente1")},
		types.KVMap{
			"A1": types.Value("V1"),
			"A2": types.Value("V2"),
		})
	Expect(err).To(BeNil())
	{
		rows, err := db.Query("SELECT key, value FROM edge WHERE kind1='meetup' AND kind2='user' AND id1=? AND id2=?", "meetup1", "utente1")
		Expect(err).To(BeNil())
		var res = map[string]string{}
		for rows.Next() {
			var value string
			var name string
			err = rows.Scan(&name, &value)
			Expect(err).To(BeNil())
			res[name] = value
		}
		Expect(res).To(BeEquivalentTo(map[string]string{
			"A1": "V1",
			"A2": "V2"}))
	}
	err = repo.SetMeetupUserLinkProperties(context.TODO(),
		types.MeetupUserLink{types.MeetupID("meetup1"), types.UserID("utente1")}, types.KVMap{
			"A1": types.Value("V1_2"),
			"A2": types.Value("V2_2"),
		})
	Expect(err).To(BeNil())
	{
		rows, err := db.Query("SELECT key, value FROM edge WHERE kind1='meetup' AND kind2='user' AND id1=? AND id2=?", "meetup1", "utente1")
		Expect(err).To(BeNil())
		var res = map[string]string{}
		for rows.Next() {
			var value string
			var name string
			err = rows.Scan(&name, &value)
			Expect(err).To(BeNil())
			res[name] = value
		}
		Expect(res).To(BeEquivalentTo(map[string]string{
			"A1": "V1_2",
			"A2": "V2_2"}))
	}

}
func randStr(c int) string {
	var s = ""
	for n := 0; n < c; n++ {
		var ch = rune(rand.Intn(26) + 65)
		s = s + string(ch)
	}
	return s
}

func TestFindMeetupUserLinksNotMatch_ok(t *testing.T) {
	rand.Seed(time.Now().Unix())
	RegisterTestingT(t)
	env := &envMock{}
	db, dbPath := dbFixture()
	defer func() {
		log.Println(dbPath)
		os.Remove(dbPath)
	}()
	tag := randStr(10)
	_ = dbPath
	repo, err := NewRepo(env, db)
	Expect(err).To(BeNil())
	Expect(repo).NotTo(BeNil())
	{
		_, err = db.Exec("INSERT INTO edge (kind1,id1,kind2,id2, key,value) values('meetup',?,'user',?,?,?)", "meetup1", "utente1", "tag", randStr(10))
		Expect(err).To(BeNil())
		_, err = db.Exec("INSERT INTO edge (kind1,id1,kind2,id2, key,value) values('meetup',?,'user',?,?,?)", "meetup1", "utente2", "tag", randStr(10))
		Expect(err).To(BeNil())
		_, err = db.Exec("INSERT INTO edge (kind1,id1,kind2,id2, key,value) values('meetup',?,'user',?,?,?)", "meetup1", "utente3", "tag", tag)
		Expect(err).To(BeNil())
		_, err = db.Exec("INSERT INTO edge (kind1,id1,kind2,id2, key,value) values('meetup',?,'user',?,?,?)", "meetup1", "utente4", "tag", tag)
		Expect(err).To(BeNil())
		_, err = db.Exec("INSERT INTO edge (kind1,id1,kind2,id2, key,value) values('meetup',?,'user',?,?,?)", "meetup1", "utente5", "tag", randStr(10))
		Expect(err).To(BeNil())
		_, err = db.Exec("INSERT INTO edge (kind1,id1,kind2,id2, key,value) values('meetup',?,'user',?,?,?)", "meetup2", "utente5", "tag", randStr(10))
		Expect(err).To(BeNil())
	}

	outC, errC := repo.FindMeetupUserLinksNotMatch(context.TODO(), types.MeetupID("meetup1"), types.Key("tag"), types.Value(tag))
	var found = map[string]struct{}{}
	for l := range outC {
		found[string(l.MeetupID)+"_"+string(l.UserID)] = struct{}{}
	}

	err = <-errC
	Expect(err).To(BeNil())
	Expect(found).To(BeEquivalentTo(map[string]struct{}{
		"meetup1_utente1": struct{}{},
		"meetup1_utente2": struct{}{},
		"meetup1_utente5": struct{}{},
	}))

}
