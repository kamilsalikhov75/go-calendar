package events

import "testing"

func TestTranslate(t *testing.T) {
	result := PriorityLow.Translate()

	if result != "Низкий" {
		t.Error(
			"For", PriorityLow,
			"expected", "Низкий",
			"got", result,
		)
	}

	result = PriorityMedium.Translate()

	if result != "Средний" {
		t.Error(
			"For", PriorityMedium,
			"expected", "Средний",
			"got", result,
		)
	}

	result = PriorityHigh.Translate()

	if result != "Высокий" {
		t.Error(
			"For", PriorityHigh,
			"expected", "Высокий",
			"got", result,
		)
	}
}

func TestPriorityValidate(t *testing.T) {
	err := PriorityLow.Validate()
	if err != nil {
		t.Error(
			"For", PriorityLow,
			"expected", nil,
			"got", err,
		)
	}

	err = PriorityMedium.Validate()
	if err != nil {
		t.Error(
			"For", PriorityMedium,
			"expected", nil,
			"got", err,
		)
	}

	err = PriorityHigh.Validate()
	if err != nil {
		t.Error(
			"For", PriorityHigh,
			"expected", nil,
			"got", err,
		)
	}

	var priority Priority = "test"
	err = priority.Validate()
	if err == nil {
		t.Error(
			"For", priority,
			"expected", "Неверное значение приоритета",
			"got", nil,
		)
	}
}
