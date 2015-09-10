package rccsv

/*
	CSV reader for RingCentral Call Log CSV files

	SYNOPSIS

 	csv := rccsv.NewCallLogRecordsCsvReader()
	err := csv.ReadFile("/path/to/my.csv")
	if err != nil {
		fmt.Printf("ERR %v", err)
	}
	stats := csv.CallLogRecordsCsv.GetStatsForVoiceRecordings()
*/

import (
	"errors"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/grokify/gotilla/encoding/csvutil"
	"github.com/grokify/gotilla/time/timeutil"
	"github.com/ttacon/libphonenumber"
)

type CallLogStats struct {
	Days           float64
	NumCalls       int64
	TotalSeconds   int64
	TotalMinutes   float32
	TotalHours     float32
	AverageMinutes float32
	DailyCalls     float32
	DailyMinutes   float32
	DailyHours     float32
	Dt14Start      int64
	Dt14End        int64
}

func NewCallLogStats(days int64) CallLogStats {
	s := CallLogStats{}
	s.Days = -1
	s.NumCalls = 0
	s.TotalSeconds = 0
	s.TotalMinutes = 0.0
	s.TotalHours = 0.0
	s.AverageMinutes = 0.0
	s.DailyCalls = 0.0
	s.DailyMinutes = 0.0
	s.DailyHours = 0.0
	s.Dt14Start = -1
	s.Dt14End = -1
	return s
}

func (s *CallLogStats) Inflate() {
	if s.Dt14Start > 0 && s.Dt14End > 0 && s.Dt14Start < s.Dt14End {
		dtStart, err1 := timeutil.TimeForDt14(s.Dt14Start)
		if err1 == nil {
			dtEnd, err2 := timeutil.TimeForDt14(s.Dt14End)
			if err2 == nil {
				dur := dtEnd.Sub(dtStart)
				durHrs := dur.Hours()
				days := durHrs / 24
				s.Days = days
			}
		}
	}
	s.TotalMinutes = float32(s.TotalSeconds / 60)
	s.TotalHours = float32(s.TotalSeconds / 60 / 60)
	s.AverageMinutes = s.TotalMinutes / float32(s.NumCalls)
	s.DailyCalls = float32(s.NumCalls) / float32(s.Days)
	s.DailyMinutes = s.DailyCalls * s.AverageMinutes
	s.DailyHours = s.DailyMinutes / 60
}

type CallLogRecordsCsv struct {
	CallLogRecords []CallLogRecordCsv
}

func (rs *CallLogRecordsCsv) GetStatsForVoiceRecordings() CallLogStats {
	stats := NewCallLogStats(int64(30))
	stats.NumCalls = 0
	for _, rec := range rs.CallLogRecords {
		rec.Inflate()
		if rec.Action != "Phone Call" && rec.Action != "VoIP Call" && rec.Action != "FindMe" {
			//continue
		}
		if rec.ActionResult != "Accepted" && rec.ActionResult != "Call connected" {
			continue
		}
		stats.NumCalls += 1
		stats.TotalSeconds += rec.DurationSeconds

		if len(rec.TimeRfc3339) > 0 {
			dt, err := time.Parse(time.RFC3339, rec.TimeRfc3339)
			if err == nil {
				dt14 := timeutil.Dt14ForTime(dt)
				//fmt.Println(dt14)
				if stats.Dt14Start < 0 || dt14 < stats.Dt14Start {
					stats.Dt14Start = dt14
				}
				if stats.Dt14End < 0 || dt14 > stats.Dt14Start {
					stats.Dt14End = dt14
				}
				//panic("B")
			}
		}
	}
	stats.Inflate()
	return stats
}

type CallLogRecordCsv struct {
	Type                  string
	PhoneNumber           string
	PhoneNumberE164       string
	PhoneNumberIso31661a2 string
	Name                  string
	Date                  string
	Time                  string
	TimeRfc3339           string
	Action                string
	ActionResult          string
	ResultDescription     string
	Duration              string
	DurationSeconds       int64
}

func (r *CallLogRecordCsv) SetCountry() {}

func (r *CallLogRecordCsv) Inflate() {
	r.PhoneNumberIso31661a2 = ""
	if 1 == 0 && len(r.PhoneNumber) > 0 {
		country := "US"
		num, err := libphonenumber.Parse(r.PhoneNumber, country)
		if err == nil {
			r.PhoneNumberE164 = libphonenumber.Format(num, libphonenumber.E164)
			r.PhoneNumberIso31661a2 = country
		}
	}
	rx1 := regexp.MustCompile(`^([0-9]{1,2}):([0-9]{1,2}):([0-9]{1,2})$`)
	rs1 := rx1.FindStringSubmatch(r.Duration)
	if len(rs1) > 0 {
		hr, err := strconv.ParseInt(rs1[1], 10, 64)
		if err != nil {
			hr = 0
		}
		mn, err := strconv.ParseInt(rs1[2], 10, 64)
		if err != nil {
			mn = 0
		}
		sc, err := strconv.ParseInt(rs1[3], 10, 64)
		if err != nil {
			sc = 0
		}
		sc = sc + (mn * 60) + (hr * 60 * 60)
		r.DurationSeconds = sc
	}
	if len(r.Date) > 0 && len(r.Time) > 0 {
		dateTimeRaw := r.Date + " " + r.Time
		layout := "Mon 01/02/2006 3:04 PM"
		myTime, err := time.Parse(layout, dateTimeRaw)
		if err == nil {
			r.TimeRfc3339 = myTime.Format(time.RFC3339)
		}
	}
}

func (r *CallLogRecordCsv) LoadRow(cols []string, vals []string) error {
	lc := len(cols)
	lv := len(vals)
	if lc != lv {
		return errors.New("RC CSV COL ROW LENGTH MISMATCH")
	}
	rx1 := regexp.MustCompile(`\s+`)
	for i := 0; i < lc; i++ {
		col := cols[i]
		val := vals[i]
		key := rx1.ReplaceAllString(col, "")
		reflect.ValueOf(r).Elem().FieldByName(key).Set(reflect.ValueOf(val))
	}
	return nil
}

type CallLogRecordsCsvReader struct {
	CallLogRecordsCsv CallLogRecordsCsv
}

func NewCallLogRecordsCsvReader() CallLogRecordsCsvReader {
	rd := CallLogRecordsCsvReader{}
	return rd
}

func (rd *CallLogRecordsCsvReader) ReadFile(path string) error {
	reader, file, err := csvutil.NewReader(path, ',')
	if err != nil {
		return err
	}
	i := 0
	cols := []string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if i == 0 {
			cols = record
			i += 1
			continue
		}
		obj := CallLogRecordCsv{}
		_ = obj.LoadRow(cols, record)
		obj.Inflate()
		rd.CallLogRecordsCsv.CallLogRecords = append(rd.CallLogRecordsCsv.CallLogRecords, obj)
		i += 1
	}
	file.Close()
	return nil
}
