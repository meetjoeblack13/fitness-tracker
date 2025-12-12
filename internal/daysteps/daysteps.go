package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/meetjoeblack13/fitness-tracker/internal/personaldata"
	"github.com/meetjoeblack13/fitness-tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("input error")
	}
	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("steps parsing error")
	}
	if ds.Steps <= 0 {
		return errors.New("steps calculation error")
	}
	ds.Duration, err = time.ParseDuration(parts[2])
	if err != nil {
		return errors.New("duration parsing error")
	}
	if ds.Duration <= 0 {
		return errors.New("duration calculation error")
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	var err error
	if ds.Steps <= 0 {
		return "", err
	}
	distanceKilometers := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distanceKilometers, calories), err
}
