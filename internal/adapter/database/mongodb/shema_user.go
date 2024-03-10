package mongodb

type UserSchema struct {
	FName string `bson:"fname"`
	LName string `bson:"lname"`
}
