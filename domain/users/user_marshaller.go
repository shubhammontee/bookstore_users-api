package users

import "encoding/json"

type (
	//PublicUser for publi
	PublicUser struct {
		ID          int64  `json:"id,omitempty"`
		CreatedDate string `json:"created_date,omitempty"`
		Status      string `json:"status,omitempty"`
	}
	//PrivateUser for private
	PrivateUser struct {
		ID          int64  `json:"id,omitempty"`
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Email       string `json:"email,omitempty"`
		CreatedDate string `json:"created_date,omitempty"`
		Status      string `json:"status,omitempty"`
	}
)

//Marshall for array of users
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

//Marshall returns either publuc or private user
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {

		//if the input json tag of user and Public user are not same
		//than we do this way individually construct a json
		//we have just shown for eg our  json tag sre same in user and publicuser
		return PublicUser{
			ID:          user.ID,
			CreatedDate: user.CreatedDate,
			Status:      user.Status,
		}
	}
	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser

}
