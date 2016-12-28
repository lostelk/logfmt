package logfmt

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/Sirupsen/logrus"
)

var (
	SortKeys   = false
	IgnoreData = false
)

var (
	tags = [...]string{
		logrus.PanicLevel: "[PNC] ",
		logrus.FatalLevel: "[FAT] ",
		logrus.ErrorLevel: "[ERR] ",
		logrus.WarnLevel:  "[WRN] ",
		logrus.InfoLevel:  "[INF] ",
		logrus.DebugLevel: "[DBG] ",
	}
)

var DefaultFormatter *PlainFormatter

type PlainFormatter struct{}

func (*PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	buf := entry.Buffer
	if buf == nil {
		buf = bytes.NewBuffer(nil)
	}

	// time
	var tbuf [20]byte
	year, month, day := entry.Time.Date()
	tbuf[0], tbuf[1] = getdigit(byte(year / 100))
	tbuf[2], tbuf[3] = getdigit(byte(year % 100))
	tbuf[4] = '-'
	tbuf[5], tbuf[6] = getdigit(byte(month))
	tbuf[7] = '-'
	tbuf[8], tbuf[9] = getdigit(byte(day))
	tbuf[10] = '/'
	hour, min, sec := entry.Time.Clock()
	tbuf[11], tbuf[12] = getdigit(byte(hour))
	tbuf[13] = ':'
	tbuf[14], tbuf[15] = getdigit(byte(min))
	tbuf[16] = ':'
	tbuf[17], tbuf[18] = getdigit(byte(sec))
	tbuf[19] = ' '
	buf.Write(tbuf[:])

	// level
	buf.WriteString(tags[entry.Level])

	if !IgnoreData && len(entry.Data) > 0 {
		// fields
		buf.WriteByte('{')
		if !SortKeys {
			more := false
			for k, v := range entry.Data {
				if more {
					buf.WriteByte(',')
					buf.WriteByte(' ')
				} else {
					more = true
				}
				buf.WriteString(k)
				buf.WriteByte(':')
				buf.WriteByte(' ')
				fmt.Fprint(buf, v)
			}
		} else {
			keys := make([]string, len(entry.Data))
			p := 0
			for k := range entry.Data {
				keys[p] = k
				p++
			}
			keys = keys[:p]
			sort.Strings(keys)
			more := false
			for _, k := range keys {
				if more {
					buf.WriteByte(',')
					buf.WriteByte(' ')
				} else {
					more = true
				}
				buf.WriteString(k)
				buf.WriteByte(':')
				buf.WriteByte(' ')
				fmt.Fprint(buf, entry.Data[k])
			}
		}
		buf.WriteByte('}')
		buf.WriteByte(' ')
	}

	// message
	buf.WriteString(entry.Message)

	buf.WriteByte('\n')
	return buf.Bytes(), nil
}

func getdigit(n byte) (byte, byte) {
	return n/10 + '0', n%10 + '0'
}
