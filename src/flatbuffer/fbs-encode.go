package main

import (
	sample "MyGame/Sample"
	"time"

	"os"

	mcap "github.com/foxglove/mcap/go/mcap"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	builder := flatbuffers.NewBuilder(1024)
	weaponOne := builder.CreateString("Sword")
	weaponTwo := builder.CreateString("Axe")

	// Create the first `Weapon` ("Sword").
	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponOne)
	sample.WeaponAddDamage(builder, 3)
	sword := sample.WeaponEnd(builder)

	// // Create the second `Weapon` ("Axe").
	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponTwo)
	sample.WeaponAddDamage(builder, 5)
	axe := sample.WeaponEnd(builder)

	// Serialize a name for our monster, called "Orc".
	name := builder.CreateString("Orc")

	// Create a `vector` representing the inventory of the Orc. Each number
	// could correspond to an item that can be claimed after he is slain.
	// Note: Since we prepend the bytes, this loop iterates in reverse.
	sample.MonsterStartInventoryVector(builder, 10)
	for i := 9; i >= 0; i-- {
		builder.PrependByte(byte(i))
	}
	inv := builder.EndVector(10)

	// Create a FlatBuffer vector and prepend the weapons.
	// Note: Since we prepend the data, prepend them in reverse order.
	sample.MonsterStartWeaponsVector(builder, 2)
	builder.PrependUOffsetT(axe)
	builder.PrependUOffsetT(sword)
	weapons := builder.EndVector(2)

	sample.MonsterStartPathVector(builder, 2)
	sample.CreateVec3(builder, 1.0, 2.0, 3.0)
	sample.CreateVec3(builder, 4.0, 5.0, 6.0)
	path := builder.EndVector(2)

	// Create our monster using `MonsterStart()` and `MonsterEnd()`.
	sample.MonsterStart(builder)
	sample.MonsterAddPos(builder, sample.CreateVec3(builder, 1.0, 2.0, 3.0))
	sample.MonsterAddHp(builder, 300)
	sample.MonsterAddName(builder, name)
	sample.MonsterAddInventory(builder, inv)
	sample.MonsterAddColor(builder, sample.ColorRed)
	sample.MonsterAddWeapons(builder, weapons)
	sample.MonsterAddPath(builder, path)
	orc := sample.MonsterEnd(builder)

	// Call `Finish()` to instruct the builder that this monster is complete.
	builder.Finish(orc)

	// This must be called after `Finish()`.
	buf := builder.FinishedBytes() // Of type `byte[]`.

	// Take the buffer and write to an mcap
	file, _ := os.Create("output.mcap")
	opts := mcap.WriterOptions{
		IncludeCRC: true,
	}
	writer, _ := mcap.NewWriter(file, &opts)
	writer.WriteHeader(&mcap.Header{})

	// Need to include the binary flatbuffer schema in order to decode correctly
	fbs, _ := os.ReadFile("monster.bfbs")

	schema := mcap.Schema{
		ID:       1,
		Name:     "MyGame.Sample.Monster",
		Encoding: "flatbuffer",
		Data:     fbs,
	}
	writer.WriteSchema(&schema)

	channel := mcap.Channel{
		ID:              0,
		SchemaID:        1,
		Topic:           "example",
		MessageEncoding: "flatbuffer",
	}
	writer.WriteChannel(&channel)

	for i := uint32(0); i < 10; i++ {
		msg := mcap.Message{
			ChannelID:   0,
			Sequence:    i,
			PublishTime: uint64(time.Now().UnixNano()),
			LogTime:     uint64(time.Now().UnixNano()),
			Data:        buf,
		}
		writer.WriteMessage(&msg)
		time.Sleep(time.Duration(1e8))
	}
	writer.Close()
}
