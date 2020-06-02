package lib

import (
	"io"
	"os"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/sirupsen/logrus"
	"github.com/yyewolf/dca-disgord"
)

var ctx atlas.Context

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

// EncodeSession creates DCA audio
func EncodeSession(inFile string, outFile string) {
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 96
	options.Application = "lowdelay"

	encodeSession, err := dca.EncodeFile(inFile, options)
	defer encodeSession.Cleanup()

	output, err := os.Create(outFile) // Include full path including *.dca TODO: assume path based on inFile
	if err != nil {
		logrus.Fatal(err)
	}
	io.Copy(output, encodeSession)

}
