package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	mcap "github.com/foxglove/mcap/go/mcap"
)

func main() {
	file, _ := os.Create("output.mcap")
	opts := mcap.WriterOptions{
		IncludeCRC: true,
	}
	writer, _ := mcap.NewWriter(file, &opts)
	channel := mcap.Channel{
		ID:       1,
		SchemaID: 1,
		Topic:    "Example",
	}
	header := mcap.Header{}
	writer.WriteHeader(&header)
	writer.WriteChannel(&channel)

	schema := mcap.Schema{
		ID:       1,
		Name:     "Test Schema",
		Encoding: mcap.schema,
	}

	// Example writing a json message
	testMsgContent := map[string]int{"hello": 10, "world": 20}
	payload, _ := json.Marshal(testMsgContent)
	fmt.Println(payload)
	msg := mcap.Message{
		ChannelID:   1,
		Sequence:    1,
		PublishTime: uint64(time.Now().Unix()),
		LogTime:     uint64(time.Now().Unix()),
		Data:        payload,
	}
	writer.WriteMessage(&msg)
	writer.WriteFooter(&mcap.Footer{})
	writer.Close()

	// file_in, _ := os.Open("output.mcap")
	// // Test reading
	// reader, _ := mcap.NewReader(file_in)
	// msgs, _ := reader.Messages()

	// fmt.Println()
}
