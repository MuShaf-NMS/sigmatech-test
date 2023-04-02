package service

import "github.com/MuShaf-NMS/sigmatech-test/entity"

// For now I just use what in the example for the limit
var cl [][]*entity.LimitPinjaman = [][]*entity.LimitPinjaman{
	{
		&entity.LimitPinjaman{
			Tenor: 1,
			Limit: 100000,
		},
		&entity.LimitPinjaman{
			Tenor: 2,
			Limit: 200000,
		},
		&entity.LimitPinjaman{
			Tenor: 3,
			Limit: 500000,
		},
		&entity.LimitPinjaman{
			Tenor: 4,
			Limit: 700000,
		},
	},
	{
		&entity.LimitPinjaman{
			Tenor: 1,
			Limit: 1000000,
		},
		&entity.LimitPinjaman{
			Tenor: 2,
			Limit: 1200000,
		},
		&entity.LimitPinjaman{
			Tenor: 3,
			Limit: 1500000,
		},
		&entity.LimitPinjaman{
			Tenor: 4,
			Limit: 2000000,
		},
	},
}
