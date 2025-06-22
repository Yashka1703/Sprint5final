package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataToSlice := strings.Split(datastring, ",")
	if len(dataToSlice) != 2 {
		return fmt.Errorf("wrong incoming format")
	}

	ds.Steps, err = strconv.Atoi(dataToSlice[0])
	if err != nil {
		return fmt.Errorf("error in convertation steps: %w", err)
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("steps can't be less or equal zero")
	}

	ds.Duration, err = time.ParseDuration(dataToSlice[1])
	if err != nil {
		return fmt.Errorf("error in convertation duration: %w", err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("duration can't be less or equal zero")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", fmt.Errorf("steps can't be less or equal zero")
	}
	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration can't be less or equal zero")
	}
	if ds.Height <= 0 {
		return "", fmt.Errorf("height can't be less or equal zero")
	}
	if ds.Weight <= 0 {
		return "", fmt.Errorf("weight can't be less or equal zero")
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", nil
	}

	result := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)
	return result, nil
}
