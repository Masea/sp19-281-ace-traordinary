/*
	picscmdapi in Go (Version 1)
*/

package main


type Picture struct {
	PictureId    string   `json:"pictureId"`
	UserId       string   `json:"userId"`
	Title        string   `json:"title"`
	Price 	     float32  `json:"price"`
	Description  string   `json:"description"`
	TumbnailUrl  string   `json:"tumbnailUrl"`
	OrigUrl      string   `json:"origUrl"`
}	
type Payload struct {
	RequestId     string   `json:"requestId" bson:"requestId"`
	PictureId     string   `json:"pictureId" bson:"pictureId"`
	UserId        string   `json:"userId" bson:"userId"`
	Title         string   `json:"title" bson:"title"`
	Price 	      float32  `json:"price" bson:"price"`
	Description   string   `json:"description" bson:"description"`
	RequestStatus string   `json:"requestStatus" bson:"requestStatus"`
	IpAddress	  string   `json:"ipaddress" bson:"ipaddress"`
}

type UploadRequest struct {
	RequestId    string  `json:"requestId" bson:"requestId"`
	PictureId    string  `json:"pictureId" bson:"pictureId"`
	UserId       string  `json:"userId" bson:"userId"`
	Title        string  `json:"title" bson:"title"`
	Price 	     float32 `json:"price" bson:"price"`
	Description  string  `json:"description" bson:"description"`
	File         []byte  `json:"file" bson:"file"`
}

type DeleteRequest struct {
	RequestId  string  	`json:"requestId" bson:"requestId"`
	PictureId   string 	`json:"pictureId" bson:"pictureId"`
	UserId   string  	`json:"userId" bson:"userId"`
}

var logs map[string]Payload
