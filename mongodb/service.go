package mongodb

import (
	pb "github.com/mbmanthey/foodtruck/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName     = "foodtruck"
	collection = "trucks"
)

type TruckRepository struct {
	session *mgo.Session
}

//Repository is an interface of the functions for managing food trucks.
type Repository interface {
	Create(*pb.Truck) error
	GetAll() ([]*pb.Truck, error)
	Get(id string) (*pb.Truck, error)
	Close()
}

//Find returns a truck by TruckID
func (repo *TruckRepository) Find(id string) (*pb.Truck, error) {
	var truck *pb.Truck
	err := repo.collection().Find(bson.M{
		"id": id,
	}).One(&truck)
	if err != nil {
		return nil, err
	}
	return truck, nil
}

//Create adds a truck to the mongodb.
func (repo *TruckRepository) Create(truck *pb.Truck) error {
	return repo.collection().Insert(truck)
}

//GetAll will retrieve all trucks from mongodb.
func (repo *TruckRepository) GetAll() ([]*pb.Truck, error) {
	var trucks []*pb.Truck
	err := repo.collection().Find(nil).All(&trucks)
	return trucks, err
}

//Get will retrieve a single truck using an ID.
func (repo *TruckRepository) Get(id string) (*pb.Truck, error) {
	var truck *pb.Truck
	err := repo.collection().FindId(id).One(&truck)
	return truck, err
}

func (repo *TruckRepository) DeleteAll() error {
	_, error := repo.collection().RemoveAll(nil)
	return error
}

func (repo *TruckRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(collection)
}

//Close will close the mongodb session.
func (repo *TruckRepository) Close() {
	repo.session.Close()
}
