package mongo

import (
	"user-backend/pkg"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
	hash       pkg.Hash
}

func NewUserService(session *Session, dbName string, collectionName string, hash pkg.Hash) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection, hash}
}

func (p *UserService) Create(u *pkg.User) error {
	user := newUserModel(u)
	hashedPassword, err := p.hash.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return p.collection.Insert(&user)
}

func (p *UserService) GetByUsername(username string) (*pkg.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootUser(), err
}
