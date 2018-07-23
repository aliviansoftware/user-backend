package mongo

import (
	"user-backend/pkg"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
}

func NewUserService(session *mgo.Session, config *pkg.MongoConfig) *UserService {
	collection := session.DB(config.DbName).C("user")
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection}
}

func (p *UserService) CreateUser(u *pkg.User) error {
	user, err := newUserModel(u)
	if err != nil {
		return err
	}
	return p.collection.Insert(&user)
}

func (p *UserService) GetUserByUsername(username string) (error, pkg.User) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return err, pkg.User{
		Id:       model.Id.Hex(),
		Username: model.Username,
		Password: "-"}
}

func (p *UserService) Login(c pkg.Credentials) (error, pkg.User) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": c.Username}).One(&model)

	err = model.comparePassword(c.Password)
	if err != nil {
		return err, pkg.User{}
	}

	return err, pkg.User{
		Id:       model.Id.Hex(),
		Username: model.Username,
		Password: "-"}
}
