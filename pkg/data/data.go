package nextDate

import (
	"errors"
	"fmt"
	"github/Rarik88/go_final_project/pkg/model"
	"log"
	"strconv"
	"strings"
	"time"
)

const Format_yyyymmdd = "20060102"

func NextDate(nd model.NextDate) (string, error) {
	now, err := time.Parse("20060102", nd.Now)

	if err != nil {
		return "", errors.New("неверный формат даты")
	}

	// Парсим исходное время
	startDate, err := time.Parse("20060102", nd.Date)
	log.Println(startDate)
	if err != nil {
		return "", errors.New("неверный формат даты")
	}

	// Разбиваем правило повторения на части
	ruleParts := strings.Fields(nd.Repeat)

	if len(ruleParts) == 0 && nd.Repeat != "y" {
		return "", errors.New("пустое правило повторения")
	}

	switch ruleParts[0] {
	case "d":
		now, err = timeNow(nd)
		if err != nil {
			return "", err
		}

		stringRepeatIntervalDays := strings.TrimPrefix(nd.Repeat, "d ")
		repeatIntervalDays, err := strconv.Atoi(stringRepeatIntervalDays)
		if err != nil {
			return "некоректное повторение", fmt.Errorf("некоректное повторение. не смог распарсить - %s", nd.Repeat)
		}
		if repeatIntervalDays < 1 || repeatIntervalDays > 400 {
			return "некоректное повторение", fmt.Errorf("некоректное повторение - %s. допускается больше 1 и меньше 400", nd.Repeat)
		}
		searchDate := nd.Date

		for searchDate <= now.Format(Format_yyyymmdd) || searchDate <= nd.Date {
			d, err := time.Parse(Format_yyyymmdd, searchDate)
			if err != nil {
				return "неверный формат даты", fmt.Errorf("неверный формат даты - %s", nd.Date)
			}
			searchDate = d.AddDate(0, 0, repeatIntervalDays).Format(Format_yyyymmdd)
		}
		return searchDate, nil
	case "y":
		now, err := timeNow(nd)
		if err != nil {
			return "", err
		}

		formattedNow := now.Format(Format_yyyymmdd)
		searchDate := nd.Date

		for searchDate <= formattedNow || searchDate <= nd.Date {
			d, err := time.Parse(Format_yyyymmdd, searchDate)
			if err != nil {
				return "неверный формат даты", fmt.Errorf("неверный формат даты - %s", nd.Date)
			}
			searchDate = d.AddDate(1, 0, 0).Format(Format_yyyymmdd)
		}
		return searchDate, nil
	default:
		return "некоректное повторение", fmt.Errorf("некоректное повторение - %s", nd.Repeat)
	}

	return "", nil
}

func timeNow(nd model.NextDate) (time.Time, error) {

	var now time.Time
	if nd.Now == "" {
		now = time.Now()
	}
	now, err := time.Parse(Format_yyyymmdd, nd.Now)
	if err != nil {
		return time.Time{}, fmt.Errorf("неверный формат даты - %s", nd.Now)
	}
	return now, nil
}
