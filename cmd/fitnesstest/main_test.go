package ftracker

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowTrainingInfo(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	actionsNum := int(rnd.Int63n(10000-1000) + 1000)
	durationNum := float64(rnd.Int63n(3)) + rnd.Float64()
	weightNum := float64(rnd.Int63n(140-80) + 80)
	heightNum := float64(rnd.Int63n(220-150) + 150)
	lengthPoolNum := int(rnd.Int63n(50-10) + 10)
	countPoolNum := int(rnd.Int63n(10-1) + 1)

	t.Run("rinning", func(t *testing.T) {
		trainingType := "Бег"
		res := ShowTrainingInfo(actionsNum, trainingType, durationNum, weightNum, heightNum, lengthPoolNum, countPoolNum)

		distance := testDistance(actionsNum)
		speed := testMeanSpeed(actionsNum, durationNum)
		calories := RunningSpentCalories(actionsNum, weightNum, durationNum)
		expected := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, durationNum, distance, speed, calories)

		assert.Equal(t, expected, res, "Результат выполнения функции ShowTrainingInfo не совпадает с ожидаемым")
	})

	t.Run("walking", func(t *testing.T) {
		trainingType := "Ходьба"
		res := ShowTrainingInfo(actionsNum, trainingType, durationNum, weightNum, heightNum, lengthPoolNum, countPoolNum)

		distance := testDistance(actionsNum)
		speed := testMeanSpeed(actionsNum, durationNum)
		calories := WalkingSpentCalories(actionsNum, durationNum, weightNum, heightNum)
		expected := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, durationNum, distance, speed, calories)

		assert.Equal(t, expected, res, "Результат выполнения функции ShowTrainingInfo не совпадает с ожидаемым")
	})

	t.Run("swimming", func(t *testing.T) {
		trainingType := "Плавание"
		res := ShowTrainingInfo(actionsNum, trainingType, durationNum, weightNum, heightNum, lengthPoolNum, countPoolNum)

		distance := testDistance(actionsNum)
		speed := testSwimmingMeanSpeed(lengthPoolNum, countPoolNum, durationNum)
		calories := SwimmingSpentCalories(lengthPoolNum, countPoolNum, durationNum, weightNum)
		expected := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, durationNum, distance, speed, calories)

		assert.Equal(t, expected, res, "Результат выполнения функции ShowTrainingInfo не совпадает с ожидаемым")
	})

	t.Run("unknown", func(t *testing.T) {
		actionsNum := int(rnd.Int63n(10000-1000) + 1000)
		trainingType := randString(3, 15)
		durationNum := float64(rnd.Int63n(3)) + rnd.Float64()
		weightNum := float64(rnd.Int63n(140-80) + 80)
		heightNum := float64(rnd.Int63n(220-150) + 150)
		lengthPoolNum := int(rnd.Int63n(50-10) + 10)
		countPoolNum := int(rnd.Int63n(10-1) + 1)

		res := ShowTrainingInfo(actionsNum, trainingType, durationNum, weightNum, heightNum, lengthPoolNum, countPoolNum)
		assert.Equal(t, "неизвестный тип тренировки", res)
	})
}

func TestWalkingSpentCalories(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	actionsNum := int(rnd.Int63n(10000-1000) + 1000)
	durationNum := float64(rnd.Int63n(3)) + rnd.Float64()
	weightNum := float64(rnd.Int63n(140-80) + 80)
	heightNum := float64(rnd.Int63n(220-150) + 150)

	meanSpeed := testMeanSpeed(actionsNum, durationNum)
	expected := (_walkingCaloriesWeightMultiplier*weightNum + (math.Pow(meanSpeed*_kmhInMsec, 2.0)/(heightNum/_cmInM))*_walkingSpeedHeightMultiplier*weightNum) * durationNum * _minInH

	res := WalkingSpentCalories(actionsNum, durationNum, weightNum, heightNum)
	assert.InDelta(t, expected, res, 0.05, "Значение полученное из функции WalkingSpentCalories не совпадает с ожидаемым")
}

func TestRunningSpentCalories(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	actionsNum := int(rnd.Int63n(10000-1000) + 1000)
	durationNum := float64(rnd.Int63n(3)) + rnd.Float64()
	weightNum := float64(rnd.Int63n(140-80) + 80)

	meanSpeed := testMeanSpeed(actionsNum, durationNum)
	expected := ((_runningCaloriesMeanSpeedMultiplier * meanSpeed * _runningCaloriesMeanSpeedShift) * weightNum / _mInKm * durationNum * _minInH)

	res := RunningSpentCalories(actionsNum, weightNum, durationNum)
	assert.InDelta(t, expected, res, 0.05, "Значение полученное из функции RunningSpentCalories не совпадает с ожидаемым")
}

func TestSwimmingSpentCalories(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	lengthPoolNum := int(rnd.Int63n(50-10) + 10)
	countPoolNum := int(rnd.Int63n(10-1) + 1)
	durationNum := float64(rnd.Int63n(3)) + rnd.Float64()
	weightNum := float64(rnd.Int63n(140-80) + 80)

	meanSpeed := testSwimmingMeanSpeed(lengthPoolNum, countPoolNum, durationNum)
	expected := (meanSpeed + _swimmingCaloriesMeanSpeedShift) * _swimmingCaloriesWeightMultiplier * weightNum * durationNum

	res := SwimmingSpentCalories(lengthPoolNum, countPoolNum, durationNum, weightNum)
	assert.InDelta(t, expected, res, 0.05, "Значение полученное из функции SwimmingSpentCalories не совпадает с ожидаемым")
}

func randString(minLen, maxLen int) string {
	var letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFJHIJKLMNOPQRSTUVWXYZ"

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	slen := rnd.Intn(maxLen-minLen) + minLen

	s := make([]byte, 0, slen)
	i := 0
	for len(s) < slen {
		idx := rnd.Intn(len(letters) - 1)
		char := letters[idx]
		if i == 0 && '0' <= char && char <= '9' {
			continue
		}
		s = append(s, char)
		i++
	}

	return string(s)
}

const (
	_lenStep   = 0.65
	_mInKm     = 1000
	_minInH    = 60
	_kmhInMsec = 0.278
	_cmInM     = 100

	_walkingCaloriesWeightMultiplier = 0.035
	_walkingSpeedHeightMultiplier    = 0.029

	_runningCaloriesMeanSpeedMultiplier = 18
	_runningCaloriesMeanSpeedShift      = 1.79

	_swimmingLenStep                  = 1.38
	_swimmingCaloriesMeanSpeedShift   = 1.1
	_swimmingCaloriesWeightMultiplier = 2
)

func testMeanSpeed(action int, duration float64) float64 {
	if duration <= 0 {
		return 0
	}
	d := testDistance(action)
	return d / duration
}

func testSwimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / _mInKm / duration
}

func testDistance(action int) float64 {
	return float64(action) * _lenStep / _mInKm
}