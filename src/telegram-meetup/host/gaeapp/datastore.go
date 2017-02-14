package gaeapp

import (
	"log"
	. "telegram-meetup/types"
	"time"

	"strings"

	"golang.org/x/net/context"
	glog "google.golang.org/appengine/log"

	"google.golang.org/appengine/datastore"
)

const DatastoreUserKind = "users"
const DatastoreMeetupKind = "meetups"
const DatastoreMeetupUserLinkKind = "meetups_users"
const DatastorePollUserLinkKind = "polls_users"

type userStore struct {
}

func NewUserStore() *userStore {
	return &userStore{}
}

func property(name string, value interface{}) datastore.Property {
	return datastore.Property{Name: name, Value: value}
}

func patchProperties(pl *datastore.PropertyList, a datastore.Property) *datastore.PropertyList {
	for i, _ := range *pl {
		if (*pl)[i].Name == a.Name {
			(*pl)[i].Value = a.Value
			return pl
		}
	}
	*pl = append(*pl, a)
	return pl
}

/* ------------------------------------- methods -----------------------------------------------*/

func (us *userStore) SetUserProperties(ctx context.Context, userid UserID, m KVMap) error {
	k := datastore.NewKey(ctx, DatastoreUserKind, string(userid), 0, nil)
	pl := datastore.PropertyList{}
	err := datastore.Get(ctx, k, &pl)
	newpl := &pl
	if err == nil || err == datastore.ErrNoSuchEntity {
		for k, v := range m {
			newpl = patchProperties(newpl, property(string(k), string(v)))
		}
	} else {
		return err
	}
	_, err = datastore.Put(ctx, k, newpl)
	return err
}

func (us *userStore) GetUserProperties(ctx context.Context, userid UserID, ks ...Key) (KVMap, error) {
	k := datastore.NewKey(ctx, DatastoreUserKind, string(userid), 0, nil)
	pl := datastore.PropertyList{}
	err := datastore.Get(ctx, k, &pl)
	if err == datastore.ErrNoSuchEntity {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	m := KVMap{}
	for _, k := range ks {
		m[k] = Value("")
	}
	for i, _ := range pl {
		if _, ok := m[Key(pl[i].Name)]; ok {
			m[Key(pl[i].Name)] = Value(pl[i].Value.(string))
		}
	}
	return m, nil
}

func composePairKey(a, b string) string {
	return a + " " + b
}
func decomposePairKey(k string) (a, b string) {
	var ss = strings.SplitN(k, " ", 2)
	if len(ss) > 0 {
		a = ss[0]
	}
	if len(ss) > 1 {
		b = ss[1]
	}
	return
}

func (us *userStore) SetMeetupUserLinkProperties(ctx context.Context, l MeetupUserLink, m KVMap) error {
	keyStr := composePairKey(string(l.UserID), string(l.MeetupID))
	k := datastore.NewKey(ctx, DatastoreMeetupUserLinkKind, keyStr, 0, nil)
	pl := datastore.PropertyList{}
	err := datastore.Get(ctx, k, &pl)
	newpl := &pl
	if err == nil || err == datastore.ErrNoSuchEntity {
		for k, v := range m {
			newpl = patchProperties(newpl, property(string(k), string(v)))
		}
		newpl = patchProperties(newpl, property("Meetup", string(l.MeetupID)))
		newpl = patchProperties(newpl, property("User", string(l.UserID)))

	} else {
		return err
	}
	_, err = datastore.Put(ctx, k, newpl)
	return err

}

func (us *userStore) FindMeetupUserLinksNotMatch(ctx context.Context, mid MeetupID, k Key, v Value) (<-chan MeetupUserLink, <-chan error) {
	outC := make(chan MeetupUserLink)
	errC := make(chan error, 1) // buffered!
	log.Println(k, v, outC, errC)
	baseQuery := datastore.NewQuery(DatastoreMeetupUserLinkKind).Filter("Meetup=", string(mid))
	cb := func(ctx context.Context, uid string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case outC <- MeetupUserLink{mid, UserID(uid)}:
			return nil
		}
	}
	go func() {
		defer close(errC)
		defer close(outC)
		qLT := baseQuery.Filter(string(k)+"<", string(v))
		err := runQueryCB(ctx, qLT, cb)
		if err == nil {
			qGT := baseQuery.Filter(string(k)+">", string(v))
			err = runQueryCB(ctx, qGT, cb)
		}
		if err != nil {
			errC <- err
		}
	}()
	return outC, errC
}

func (us *userStore) FindMeetupUserLinksMatch(ctx context.Context, mid MeetupID, k Key, v Value) (<-chan MeetupUserLink, <-chan error) {
	outC := make(chan MeetupUserLink)
	errC := make(chan error, 1) // buffered!
	log.Println(k, v, outC, errC)
	baseQuery := datastore.NewQuery(DatastoreMeetupUserLinkKind).Filter("Meetup=", string(mid))
	cb := func(ctx context.Context, uid string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case outC <- MeetupUserLink{mid, UserID(uid)}:
			return nil
		}
	}
	go func() {
		defer close(errC)
		defer close(outC)
		qEq := baseQuery.Filter(string(k)+"=", string(v))
		err := runQueryCB(ctx, qEq, cb)

		if err != nil {
			errC <- err
		}
	}()
	return outC, errC
}

type queryCallback func(context.Context, string) error

func runQueryCB(ctx context.Context, q *datastore.Query, cb queryCallback) error {
	log.Println(q)
	for t := q.Run(ctx); ; {
		k, err := t.Next(nil)
		if err == datastore.Done {
			break
		}
		if err != nil {
			return err
		}
		u, _ := decomposePairKey(k.StringID())
		err = cb(ctx, u)
		if err != nil {
			return err
		}
	}
	return nil
}

func (us *userStore) GetPollResult(ctx context.Context, pid PollID, k Key) (PollResult, error) {
	q := datastore.NewQuery(DatastorePollUserLinkKind).
		Filter("Poll=", string(pid))

	var res = PollResult{}
	for t := q.Run(ctx); ; {
		pl := datastore.PropertyList{}
		_, err := t.Next(&pl)
		if err == datastore.Done {
			break
		}

		glog.Debugf(ctx, "pl %v", pl)
		for _, v := range pl {
			if v.Name == string(k) {
				vote := v.Value.(string)
				count := res[vote]
				res[vote] = count + 1
			}
		}
	}
	return res, nil
}

func (us *userStore) SetPollUserLinkProperties(ctx context.Context, l PollUserLink, m KVMap) error {
	keyStr := composePairKey(string(l.UserID), string(l.PollID))
	k := datastore.NewKey(ctx, DatastorePollUserLinkKind, keyStr, 0, nil)
	pl := datastore.PropertyList{}
	err := datastore.Get(ctx, k, &pl)
	newpl := &pl
	if err == nil || err == datastore.ErrNoSuchEntity {
		for k, v := range m {
			newpl = patchProperties(newpl, property(string(k), string(v)))
		}
		newpl = patchProperties(newpl, property("Poll", string(l.PollID)))
		newpl = patchProperties(newpl, property("User", string(l.UserID)))

	} else {
		return err
	}
	_, err = datastore.Put(ctx, k, newpl)
	return err

}
func (us *userStore) FindPollUserLinksNotMatch(ctx context.Context, pid PollID, k Key, v Value) (<-chan PollUserLink, <-chan error) {
	outC := make(chan PollUserLink)
	errC := make(chan error, 1) // buffered!
	log.Println(k, v, outC, errC)
	baseQuery := datastore.NewQuery(DatastorePollUserLinkKind).Filter("Poll=", string(pid))
	cb := func(ctx context.Context, uid string) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case outC <- PollUserLink{pid, UserID(uid)}:
			return nil
		}
	}
	go func() {
		defer close(errC)
		defer close(outC)
		qLT := baseQuery.Filter(string(k)+"<", string(v))
		err := runQueryCB(ctx, qLT, cb)
		if err == nil {
			qGT := baseQuery.Filter(string(k)+">", string(v))
			err = runQueryCB(ctx, qGT, cb)
		}
		if err != nil {
			errC <- err
		}
	}()
	return outC, errC

}

type datastoreVolatile struct {
	ctx context.Context
}

type volatileBinaryEntity struct {
	Data    []byte
	Expires time.Time
}

func (v *datastoreVolatile) setBinaryData(key string, value []byte, ttl time.Duration) error {
	k := datastore.NewKey(v.ctx, DatastoreVolatileKind, key, 0, nil)
	var e volatileBinaryEntity
	e.Data = value
	e.Expires = time.Now().Add(ttl)
	var err error
	k, err = datastore.Put(v.ctx, k, &e)
	return err
}
func (v *datastoreVolatile) getBinaryData(key string) ([]byte, error) {
	k := datastore.NewKey(v.ctx, DatastoreVolatileKind, key, 0, nil)
	var e volatileBinaryEntity
	var err error
	err = datastore.Get(v.ctx, k, &e)
	return e.Data, err
}

func (v *datastoreVolatile) deleteExpiredEntries() error {
	q := datastore.NewQuery(DatastoreVolatileKind).Filter("Expires <=", time.Now())
	var err error
	var k *datastore.Key
	for err == nil {
		i := q.Run(v.ctx)
		k, err = i.Next(nil)
		glog.Debugf(v.ctx, "Next: err: %v", err)
		if err == nil {
			err = datastore.Delete(v.ctx, k)
			glog.Debugf(v.ctx, "Delete: err: %v", err)
		}
	}
	if err == datastore.Done {
		return nil
	}
	return err
}
