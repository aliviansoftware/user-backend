package mongo_test

import (
	"log"
	"testing"
	"user-backend/pkg"
	"user-backend/pkg/mongo"
)

const (
	mongoUrl           = "localhost:27017"
	dbName             = "test_db"
	userCollectionName = "user"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	//Arrange
	mongoConfig := pkg.MongoConfig{
		Ip:     "127.0.0.1:27017",
		DbName: "user-backend"}
	session, err := mongo.NewSession(&mongoConfig)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer func() {
		session.DropDatabase(mongoConfig.DbName)
		session.Close()
	}()

	userService := mongo.NewUserService(session.Copy(), &mongoConfig)

	testUsername := "integration_test_user"
	testPassword := "integration_test_password"
	user := pkg.User{
		Username: testUsername,
		Password: testPassword}

	//Act
	err = userService.CreateUser(&user)

	//Assert
	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	_, resultUser := userService.GetUserByUsername(testUsername)

	if resultUser.Username != user.Username {
		t.Errorf("Incorrect Username. Expected `%s`, Got: `%s`", testUsername, resultUser.Username)
	}
}
