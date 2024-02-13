package types

type User struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"_firstName" json:"firstName"`
	Lastname  string `bson:"_lastName" json:"lastName"`
}

func types() {

}
