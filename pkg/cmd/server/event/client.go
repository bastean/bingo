package event

type event struct {
	PutAuthorization    string
	DeleteAuthorization string
}

var Client = event{PutAuthorization: "bingo:put-authorization", DeleteAuthorization: "bingo:delete-authorization"}
