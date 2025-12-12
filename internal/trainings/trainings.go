package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/meetjoeblack13/fitness-tracker/internal/personaldata"
	"github.com/meetjoeblack13/fitness-tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("input error")
	}
	t.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("steps parsing error")
	}
	if t.Steps <= 0 {
		return errors.New("steps calculation error")
	}
	if !strings.EqualFold(t.TrainingType, "Ходьба") && !strings.EqualFold(t.TrainingType, "Бег") {
		return errors.New("неизвестный тип тренировки") // Текст этой ошибки пришлось оставить на русском, так как в тестах проверяется наличие именно этого текста
	}
	t.Duration, err = time.ParseDuration(parts[2])
	if err != nil {
		return errors.New("duration parsing error")
	}
	if t.Duration <= 0 {
		return errors.New("duration calculation error")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// считаем общие значения
	distanceKm := spentenergy.Distance(t.Steps, t.Height) // одинаково для ходьбы и бега
	speedKmH := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var (
		calories float64
		err      error
	)
	switch t.TrainingType { // Получение кол-ва сгоревших калорий в зависимости от типа тренировки
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distanceKm, speedKmH, calories)
	return result, nil
}
