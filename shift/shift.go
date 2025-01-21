package shift

import (
	"errors"
	"os"
	"strconv"
	"subs/utils"
	"time"
)

const Command = "shift"
const Help = `shift - shift subtitles by an amount (parsed as a golang duration e.g. 1h2m45s185ms)
	--start-after {duration}: optional, time to start after
	-o/--output {output filename}: file name for output, defaults to the input filename, modifying it in place
	example: subs shift [--start-after 1h2m38s488ms] [-]2m14s {filename}`

func Do(args ...string) error {
	if len(args) < 2 {
		return errors.New("missing filename or shift amount")
	}
	fileName := args[len(args)-1]
	shiftAmount, err := time.ParseDuration(args[len(args)-2])
	if err != nil {
		return err
	}
	flags := utils.ParseFlags(args)
	output := utils.Coalesce(flags["o"], flags["output"], fileName)
	startAfterStr := utils.Coalesce(flags["start-after"], "0")
	startAfterInt, err := strconv.Atoi(startAfterStr)
	if err != nil {
		return err
	}
	startAfter := time.Duration(int64(startAfterInt))
	bs, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	srt, err := utils.ParseSrt(bs)
	if err != nil {
		return err
	}
	for i := range srt {
		if srt[i].Start >= startAfter {
			srt[i].Start += shiftAmount
			srt[i].End += shiftAmount
		}
	}
	return os.WriteFile(output, []byte(srt.String()), 0777)
}
