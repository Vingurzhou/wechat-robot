package model

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

func GenerateSnowflakeID() string {
	node, err := snowflake.NewNode(1)

	if err != nil {
		log.Println(err)
	}

	return node.Generate().String()
}
