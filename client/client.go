package client

import (
	"bytes"
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/wayn3h0/go-uuid" // UUID (RFC 4122)
)

/*
Create builds a new `Client` */
func Create(id string) *Client {
	c := Client{
		FirstName: "Site",
		LastName:  "Visitor",
		Name:      "Site Visitor",
	}

	if id == "" {
		log.Println(WARNING, "No client ID specified")

		uuid, _ := uuid.NewRandom()

		c.Uid = uuid.String()
	} else {
		c.Uid = id
	}

	return &c
}

// func UnmarshalJSON

func (c *Client) GetFullName() string {
	if c != nil {
		return fmt.Sprintf("%s %s", c.GetFirstName(), c.GetLastName())
	}
	return ""
}

func (c *Client) UnmarshalJSON(data []byte) error {
	u := jsonpb.Unmarshaler{}
	buf := bytes.NewBuffer(data)

	if err := u.Unmarshal(buf, &*c); err != nil {
		return err
	}

	return nil
}

func (c Client) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	m := jsonpb.Marshaler{}

	if err := m.Marshal(&buf, &c); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c Client) StoreKey() string {
	return fmt.Sprintf("client.%s", c.Uid)
}
