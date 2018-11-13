package mongodb

import mgo "gopkg.in/mgo.v2"

//CreateSession creates a session in mongodb.
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	println("connected")
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
