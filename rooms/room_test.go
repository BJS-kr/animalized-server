package rooms_test

import (
	"animalized/message"
	"animalized/packet"
	"animalized/rooms"
	"animalized/users"
	"fmt"
	"math"
	"net"
	"testing"
)

func isEquallyDistributed(types map[string]message.Room_CharacterType) bool {
	typemap := make(map[message.Room_CharacterType]int)

	max := 0
	min := math.MaxInt

	for _, t := range types {
		typemap[t]++
	}

	for _, v := range typemap {
		if v < max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max-min <= 1
}

func TestPickCharacterRandomTypes(t *testing.T) {
	rs := rooms.New()
	r1, _ := rs.Create("test1", 2)
	r2, _ := rs.Create("test2", 5)
	r3, _ := rs.Create("test3", rooms.MAX_USERS_LIMIT)

	ps := packet.NewStore()
	conn, _ := net.Pipe()

	for i := 0; i < r1.Max; i++ {
		u, err := users.NewUser(conn, fmt.Sprintf("test%d", i), ps)

		if err != nil {
			t.Fail()
		}

		r1.Join(u)
	}

	for i := 0; i < r2.Max; i++ {
		u, err := users.NewUser(conn, fmt.Sprintf("test%d", i), ps)

		if err != nil {
			t.Fail()
		}

		r2.Join(u)
	}

	for i := 0; i < r3.Max; i++ {
		u, err := users.NewUser(conn, fmt.Sprintf("test%d", i), ps)

		if err != nil {
			t.Fail()
		}

		r3.Join(u)
	}

	for i := 0; i < 100; i++ {
		r1types := r1.PickCharacterRandomTypes()
		r2types := r2.PickCharacterRandomTypes()
		r3types := r3.PickCharacterRandomTypes()

		if !isEquallyDistributed(r1types) || !isEquallyDistributed(r2types) || !isEquallyDistributed(r3types) {
			fmt.Println(r1types, r2types, r3types)
			t.Fail()
		}
	}
}
