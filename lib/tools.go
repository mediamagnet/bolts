package lib

import (
	"strconv"

	"github.com/andersfylling/disgord"
)

// SnowflakeToUInt64 returns a uint64 version of a snowflake.
func SnowflakeToUInt64(snowflake disgord.Snowflake) uint64 {
	did, _ := strconv.Atoi(snowflake.String())

	return uint64(did)
}

// StrToSnowflake returns a Snowflake from a string.
func StrToSnowflake(str string) disgord.Snowflake {
	did, _ := strconv.Atoi(str)

	return UInt64ToSnowflake(uint64(did))
}

// UInt64ToSnowflake converts a uint64 to a snowflake.
func UInt64ToSnowflake(i uint64) disgord.Snowflake {
	return disgord.NewSnowflake(i)
}
