package embed

import (
	"errors"
	"log"
	types "telegram-meetup/types"

	"database/sql"

	"golang.org/x/net/context"
)

type repo struct {
	db *sql.DB
}

func NewRepo(env types.Env, db *sql.DB) (*repo, error) {
	if db == nil {
		return nil, errors.New("undefined db")
	}
	return &repo{db}, nil
}

func (d *repo) SetUserProperties(ctx context.Context, u types.UserID, m types.KVMap) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	for k, v := range m {
		var val string
		err := tx.QueryRow("SELECT value FROM node WHERE kind='user' AND id=? and key=?", string(u), string(k)).Scan(&val)
		if err == sql.ErrNoRows {
			_, err = tx.Exec("INSERT INTO node (kind,id,key,value) values('user',?,?,?)", string(u), string(k), string(v))
		} else if err != nil {
			return err
		} else {
			_, err = tx.Exec("UPDATE node SET value=? WHERE kind='user' AND id=? AND key=?", string(v), string(u), string(k))
			if err != nil {
				return err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *repo) GetUserProperties(ctx context.Context, u types.UserID, ks ...types.Key) (types.KVMap, error) {
	tx := d.db
	var res = types.KVMap{}
	for _, k := range ks {
		var val string
		err := tx.QueryRow("SELECT value FROM node WHERE kind='user' AND id=? and key=?", string(u), string(k)).Scan(&val)
		if err == sql.ErrNoRows {

		} else if err != nil {
			return nil, err
		} else {
			res[k] = types.Value(val)
		}
	}
	return res, nil
}

func (d *repo) SetMeetupUserLinkProperties(ctx context.Context, mu types.MeetupUserLink, m types.KVMap) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	id1 := string(mu.MeetupID)
	id2 := string(mu.UserID)
	for k, v := range m {
		var val string
		err := tx.QueryRow("SELECT value FROM edge WHERE kind1='meetup' AND kind2='user' AND id1=? AND id2=? AND key=?", id1, id2, string(k)).Scan(&val)
		if err == sql.ErrNoRows {
			_, err = tx.Exec("INSERT INTO edge (kind1,id1,kind2,id2,key,value) values('meetup',?,'user',?,?,?)", id1, id2, string(k), string(v))
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			_, err = tx.Exec("UPDATE edge SET value=? WHERE kind1='meetup' AND kind2='user' AND id1=? AND id2=? AND key=?", string(v), id1, id2, string(k))
			if err != nil {
				return err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

func (d *repo) FindMeetupUserLinksNotMatch(ctx context.Context, mid types.MeetupID, k types.Key, v types.Value) (<-chan types.MeetupUserLink, <-chan error) {
	outC := make(chan types.MeetupUserLink)
	errC := make(chan error, 1) // buffered!

	go func() {
		tx := d.db
		id1 := string(mid)
		rows, err := tx.Query("SELECT id2 FROM edge WHERE kind1='meetup' AND id1=? AND key=? AND value <> ?",
			id1, string(k), string(v),
		)
		defer func() {
			close(outC)
			errC <- err
			close(errC)
			if rows != nil {
				rows.Close()
			}
		}()

		if err != nil {
			return
		}

		for rows.Next() {
			var id2 string
			err = rows.Scan(&id2)
			log.Println(id1, id2)
			select {
			case outC <- types.MeetupUserLink{mid, types.UserID(id2)}:
				{
					log.Println(id1, id2)
				}
			case <-ctx.Done():
				{
					err = ctx.Err()
					return
				}
			}
		}

	}()
	return outC, errC
}

func (d *repo) FindMeetupUserLinksMatch(ctx context.Context, mid types.MeetupID, k types.Key, v types.Value) (<-chan types.MeetupUserLink, <-chan error) {
	outC := make(chan types.MeetupUserLink)
	errC := make(chan error, 1) // buffered!

	go func() {
		tx := d.db
		id1 := string(mid)
		rows, err := tx.Query("SELECT id2 FROM edge WHERE kind1='meetup' AND id1=? AND key=? AND value = ?",
			id1, string(k), string(v),
		)
		defer func() {
			close(outC)
			errC <- err
			close(errC)
			if rows != nil {
				rows.Close()
			}
		}()

		if err != nil {
			return
		}

		for rows.Next() {
			var id2 string
			err = rows.Scan(&id2)
			log.Println(id1, id2)
			select {
			case outC <- types.MeetupUserLink{mid, types.UserID(id2)}:
				{
					log.Println(id1, id2)
				}
			case <-ctx.Done():
				{
					err = ctx.Err()
					return
				}
			}
		}

	}()
	return outC, errC
}

func (d *repo) GetPollResult(context.Context, types.PollID, types.Key) (types.PollResult, error) {
	panic("not implemented")
}

func (d *repo) SetPollUserLinkProperties(context.Context, types.PollUserLink, types.KVMap) error {
	panic("not implemented")
}

func (d *repo) FindPollUserLinksNotMatch(context.Context, types.PollID, types.Key, types.Value) (<-chan types.PollUserLink, <-chan error) {
	panic("not implemented")
}

func DBCreate(db *sql.DB) error {
	var ddl = []string{
		"create table node ( kind varchar, id varchar, key varchar, value varchar, primary key (kind,id,key) ) ",
		"create table edge ( kind1 varchar, id1 varchar, kind2 varchar, id2 varchar,  key varchar, value varchar, primary key (kind1,id1,kind2, id2,key) ) ",
		"create index idx_edge_id1_key on edge(kind1,id1,key)",
		"create index idx_edge_id2_key on edge(kind2,id2,key)",
	}
	var err error
	for _, s := range ddl {
		_, err = db.Exec(s)
		if err != nil {
			return err
		}
	}
	return err
}
