package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	dataToSlice := strings.Split(datastring, ",")
	if len(dataToSlice) != 3 {
		return fmt.Errorf("wrong incoming format")
	}

	t.Steps, err = strconv.Atoi(dataToSlice[0])
	if err != nil {
		return fmt.Errorf("error in convertation steps: %w", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("steps can't be less or equal zero")
	}

	t.TrainingType = dataToSlice[1]

	t.Duration, err = time.ParseDuration(dataToSlice[2])
	if err != nil {
		return fmt.Errorf("error in convertation duration: %w", err)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("duration can't be less or equal zero")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error in calculating calories: %w", err)
		}
		result := fmt.Sprintf("Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
			t.TrainingType,
			t.Duration.Hours(),
			distance,
			meanSpeed,
			calories)
		return result, err

	case "Ходьба":
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error in calculating calories: %w", err)
		}
		result := fmt.Sprintf("Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
			t.TrainingType,
			t.Duration.Hours(),
			distance,
			meanSpeed,
			calories)
		return result, err

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
