package schema

// ActiveUser struct
type ActiveUser struct { //Payload
	Id   string `bson:"_id,omitempty" json:"_id,omitempty" redis:"_id"`
	Jwt  string `bson:"jwt,omitempty" json:"jwt,omitempty" redis:"jwt"`
	Role int    `bson:"role,omitempty" json:"role,omitempty" redis:"role"`
}

func (user ActiveUser) Database() string {
	return ""
}

func (user ActiveUser) Collection() string {
	return ""
}

func (user ActiveUser) Key() string {
	return user.Id
}
