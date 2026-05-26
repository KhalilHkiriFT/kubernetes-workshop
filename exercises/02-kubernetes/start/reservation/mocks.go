package main

import "reservation/model"

var MockPedelecs []model.Pedelec = []model.Pedelec{
	{
		ID:            "1",
		Charge:        100,
		IsAvailable:   true,
		WorkingRadius: 100,
	},
	{
		ID:            "2",
		Charge:        100,
		IsAvailable:   true,
		WorkingRadius: 100,
	},
	{
		ID:            "3",
		Charge:        100,
		IsAvailable:   true,
		WorkingRadius: 100,
	},
}

var MockReservations []model.Reservation
