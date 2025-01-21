package utils

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type SrtEntry struct {
	Start, End time.Duration
	Content    string
}

func (se SrtEntry) String() string {
	return "TODO"
}

type Srt []SrtEntry

func (srt Srt) String() string {
	result := ""
	for i, entry := range srt {
		result += fmt.Sprintf("%d\n", i+1)
		result += fmt.Sprintf("%s --> %s\n", durationToSrtTimeString(entry.Start), durationToSrtTimeString(entry.End))
		result += entry.Content
		result += "\n\n"
	}
	return result
}

var indexRe = regexp.MustCompile(`^\d+r$`)

func ParseSrt(bs []byte) (Srt, error) {
	var result Srt
	for begin, seg := range UntilNextEmptyLine(string(bs)) {
		if segmentContainsOnlyBlankLines(seg) {
			continue // sometimes there's a couple newlines at the end of them, don't fail on that
		}
		if len(seg) < 3 { // last newline could be missing
			return nil, fmt.Errorf("found malformed section beginning on line %d", begin)
		}
		if !indexRe.MatchString(seg[0]) {
			return nil, fmt.Errorf("expected to find index on line %d, found: %s", begin, seg[0])
		}
		beginDur, endDur, err := parseSrtTimes(seg[1])
		if err != nil {
			return nil, errors.Join(fmt.Errorf("line %d", begin+1), err)
		}
		contents := strings.Join(seg[2:], "\n")
		result = append(result, SrtEntry{
			Start:   beginDur,
			End:     endDur,
			Content: contents,
		})
	}
	return result, nil
}

func parseSrtTimes(s string) (begin, end time.Duration, err error) {
	segments := strings.Split(s, "-->")
	if len(segments) != 2 {
		err = fmt.Errorf("malformed time segment, found: %s", s)
		return
	}
	if begin, err = parseSrtTime(segments[0]); err != nil {
		return
	}
	end, err = parseSrtTime(segments[1])
	return
}

func parseSrtTime(s string) (time.Duration, error) {
	var result time.Duration
	segments := strings.Split(s, ",")
	if len(segments) > 2 {
		return result, fmt.Errorf("found invalid time string: %s", s)
	}
	millis := time.Duration(0)
	if len(segments) == 2 { // just in case someone leaves off the milliseconds, don't fail to parse
		if i, err := strconv.Atoi(segments[2]); err != nil {
			return result, err
		} else {
			millis = time.Duration(i)
		}
	}
	result += millis * time.Millisecond
	timeUnits := strings.Split(segments[0], ":")
	if len(segments) > 3 {
		return result, fmt.Errorf("found invalid time string: %s", s)
	}
	slices.Reverse(timeUnits) // reversing it this way makes the bigger times optional e.g. 00:43:27 == 43:27
	for t, unit := range Zip(timeUnits, []time.Duration{time.Second, time.Minute, time.Hour}) {
		if tAsInt, err := strconv.Atoi(t); err != nil {
			return result, err
		} else {
			result += time.Duration(tAsInt) * unit
		}
	}
	return result, nil
}

func durationToSrtTimeString(t time.Duration) string {
	hours := t / time.Hour
	t -= hours * time.Hour
	minutes := t / time.Minute
	t -= minutes * time.Minute
	seconds := t / time.Second
	t -= seconds * time.Second
	millis := t / time.Millisecond
	return formatSrtTimeStr(int(hours), int(minutes), int(seconds), int(millis))
}

func formatSrtTimeStr(hours, minutes, seconds, millis int) string {
	hString := LeftPad(fmt.Sprint(hours), 2, '0')
	mString := LeftPad(fmt.Sprint(minutes), 2, '0')
	sString := LeftPad(fmt.Sprint(seconds), 2, '0')
	msString := LeftPad(fmt.Sprint(millis), 3, '0')
	return fmt.Sprintf("%d:%d:%d,%d", hString, mString, sString, msString)
}
