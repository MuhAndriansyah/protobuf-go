package main

import (
	"log"
	"protocol-buffer-go/protogen"
	"protocol-buffer-go/utils"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	version, _ := anypb.New(&protogen.Detail{
		Version: "1.0",
	})

	userOne := &protogen.User{
		Id:          1,
		FullName:    "Muhammad Andriansyah",
		PhoneNumber: "081112345678",
		Email:       "x@icloud.com",
		Gender: &protogen.User_Man{
			Man: true,
		},
		Address: &protogen.Address{
			StreetName:  "Merdeka",
			HouseNumber: 123,
			Village:     "Sukamaju",
			District:    "Medan Baru",
			City:        "Medan",
			PostalCode:  20122,
			Province:    "Sumater Utara",
			Country:     "Indonesia",
		},
		Notification: &protogen.User_Notification{
			Id:     1,
			UserId: 1,
			Title:  "Email Notification",
			Content: &protogen.User_Notification_Email{
				Email: &protogen.EmailContent{
					Subject: "Test",
					Body:    "lorem ipsum",
					From:    "no-reply@icloud.com",
					Cc:      []string{"x@icloud.com", "x@gmail.com"},
				},
			},
			STATUS: protogen.User_Notification_SEND,
			Attributes: map[string]string{
				"device": "Iphone 15",
			},
			SendAt:  timestamppb.New(time.Now()),
			Details: version,
		},
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}

	jsonByte, err := protojson.Marshal(userOne)

	if err != nil {
		log.Fatalf(err.Error())
	}

	utils.WriteFormattedFile(jsonByte, "user", "json")

	protoByte, err := proto.Marshal(userOne)

	if err != nil {
		log.Fatalf(err.Error())
	}

	utils.WriteFormattedFile(protoByte, "user", "bin")
}
