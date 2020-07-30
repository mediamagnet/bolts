package lib

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/sirupsen/logrus"
	"github.com/yyewolf/dca-disgord"
)

var ctx atlas.Context

// NoArtistURL default avatar
var NoArtistURL = "https://discordapp.com/assets/322c936a8c8be1b803cd94861bdfa868.png"

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

// JoinString joins a string slice with a char, and removes the end char.
func JoinString(strs []string, char string) string {
	return strings.TrimRight(strings.Join(strs, char), char)
}

// JoinStringMap joins a string map with a char, and removes the end char.
func JoinStringMap(strs map[int]string, char string) string {
	// make sure map is sorted in order cause Go likes random orders
	// for some stupid fucking reason
	var keys []int

	for k := range strs {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var vals []string

	for _, k := range keys {
		vals = append(vals, strs[k])
	}

	return JoinString(vals, char)
}

// Ucwords capitalizes the first letter in each word. (Mirror's PHP's ucwords function)
func Ucwords(str string) string {
	return strings.Title(str)
}

// GenAvatarURL generates a URL used to get a user avatar.
func GenAvatarURL(user *disgord.User) string {
	if user.Avatar == "" {
		return NoArtistURL
	}

	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.webp", user.ID.String(), user.Avatar)
}

// AddEmbedFooter adds embed
func AddEmbedFooter(msg *disgord.Message) (f *disgord.EmbedFooter, t disgord.Time) {
	f = &disgord.EmbedFooter{
		IconURL: GenAvatarURL(msg.Author),
		Text:    fmt.Sprintf("Command invoked by: %s#%s", msg.Author.Username, msg.Author.Discriminator),
	}

	t = disgord.Time{
		Time: time.Now(),
	}

	return
}
