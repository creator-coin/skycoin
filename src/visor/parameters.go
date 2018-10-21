package visor

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

const (
	// MaxCoinSupply is the maximum supply of coins
	MaxCoinSupply uint64 = 180000000
	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100
	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal
	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 50
	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5
	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 31536000 // in seconds
	// MaxDropletPrecision represents the decimal precision of droplets
	MaxDropletPrecision uint64 = 3
	//DefaultMaxBlockSize is max block size
	DefaultMaxBlockSize int = 32768 // in bytes
)

var distributionAddresses = [DistributionAddressesTotal]string{
	"fpj635UU5ULbb1yZhbjNgZZgCCDGHdMKtR",
	"VCv63NkX1hcgngD493zpJgAQszBYrhTRNk",
	"2iK9EcYKxsMfVakwKZhuHJdR8vd9wmFFuT",
	"J3kgYkQiLaRn74e25BBpt8s1p3swKZf6UN",
	"z6ck4p7qg2bMUsaT24ybraWM5Dif43TvLD",
	"LjKfuAnBy5yY2sQjoGs8uuEi21CE4sacBk",
	"Vm1JUNVPgjRkC9NQetYuneB41CXRS2BMtz",
	"pHGvxGtmuvMmWft33YQD5X1pM27o1g4oxW",
	"spUFnaNwMvzFg7fUCY2DmQpNnMkrYdrqV9",
	"XQzoRpUz9r8zgcCfx42tFuT5zFUVcJaJb1",
	"W3T4ZvnJj8HYxhpxPynxHT1ZUKJnebEm5U",
	"2SGSdqtgsBwGXtMntnk4cnJuR6kS8kZFfdS",
	"nRqTiVk6VmjPcNkJ88wpYkf9WPYmTaytmZ",
	"7epr1K1EAKk2j3vswCEvpctSaHsj6qyugd",
	"2hU98G6BpCVA7B17rSKPPeQiY7cZWL9KYBH",
	"QPoSYqSbF599UYLgPvfLJYpVmBCrYny4Tc",
	"2cdByGQUpk1733pBbvxvSMD7DjQHu7uy45j",
	"CvpHc6L8KHjSMBztu2iNJNBoQ4BEF1uWUP",
	"GpTubV5Jciudrtsmq6Dyt13swdpZeABBME",
	"umNDGodTw15q1T3k8yTzu5Y1iTDC2Vdbov",
	"2VBxS56TtskwshUYS7SxFg31T9iwuB83s3H",
	"exc5TuQYJ4TkY5UVM8rZjKvsLN3iu1uU8V",
	"KirvENSojZcDpKVVVJqe1bM1RdQGAqGy4b",
	"4YrV3B6nuvgHcmkavBgd1qq6pAe4Fs9VSe",
	"S49Nj9JfFpovkmg6yrvuMqnaepN6Xjf4jr",
	"fD8gfkuwC5KhpfWJRE5zvUNnmxfnp3DSqP",
	"EDBmQXgEkB1veqa32MvQVuec77BsNqCXtC",
	"2koaDKR5Fx2EqaDEDadRtfjQe6Vn4m2gYfV",
	"VWNPrVh2H57pKTW3tXPQknjdnqxLsH4uyX",
	"NXq9CkSLfxi7E6jPQ3uCyjFzbFsajCnRRN",
	"HBCr2G9VHgabUv3QQE454qPiaTKNHz8jp3",
	"2jEQmw9tUtFrkb2HdcWadMp8CE7WTZy3vUU",
	"2kYGsNejnUMypQjK8ghqX7ZGahcoocH5pm4",
	"TMneJzWBvG3Caq27MiKhVW8ENpc9UZcEx1",
	"2SMEbVoQ18s5jda5J4noUj5jHstQb7Eb7To",
	"2NBszTYMQFNsTdFBR9NuhCYQXy7YTPuukgt",
	"29drbU15AdZeE38oqVFPeJwMFvGxYbLCovB",
	"oKHimJdCttpx1u4nc1XuQjnHLJRuEtRrje",
	"2EJGgkKrK74Yw5emuMYn7U7CfQC51AL3edS",
	"28Yt6zmbJZCMmHx24tQtWfB5DT2aR2w43DG",
	"2WZQurTPXxgEjuE7QCR6y8sEv7iRteUSfjx",
	"pF226Hq1jn9zQQifr2Nd5cD4nr5rLDHvj5",
	"2mJqv8iqSvs3bMGrDqfqMjPaRQmX83VTnt8",
	"DrHJHsM6tUfbtuqc7YAaC2C7tWR7C1kWjo",
	"2871MDHLggtNhPWBVwcygHycavUy6ob4K3u",
	"FTzztkA1ZrgyNjzasBshbruur7YUMpfdRE",
	"jowtK4Q6q4qRZTGhzyyb1PCXKMde64pLvR",
	"RTHTx4NXPEW9dsQMVmdvw8yfj3msZauxNo",
	"fzq7yR9WNHVQVaH5q4Xkkq9H8F8GXXKgmh",
	"82n7VwYm7wP8dkTHZi6g4XpjZcveZR8uFf",
	"dbqaTpiurT89D4KqMpmbBBwF8BaurivdYK",
	"2ERx6ujEUZURkPhxgtkdvYbN9TCBwWpTs1u",
	"ioXztpu8rrAh2ULToDusyzRgZwnxtvJA9D",
	"2MWuCP8AT1kznquDrR4wBxEPo8deZPGMwcw",
	"pSMp3vjRbpf6qSG3RJcV1TwBUjvCXFkBUC",
	"MEbsuJ7oi5nN89YqoEwaJq9DgDnwuBxTEj",
	"2CRrTL25bqvctAaC3B8KkuXEzx6h9vSXu3L",
	"QXfFRRsz8Nh2DJM2qCM1dwit3KYtuGigEB",
	"C2QXrZSRXzZ75T6PsvMkyemF8GBGyEGqPK",
	"u5vVDtukyZiidVY1RwbuukPhKGBaBMffdg",
	"2BqGWfpEdP7FrMcUrHLotYpe8uaShzRrBQ3",
	"21bb84v86TvLWGVQcihgDwBkUXWi2gumbmS",
	"2bVeHH64mVLndndcYb3gX5PmAoe1rfDXDFC",
	"Uyi6U3nfb9ABD1A8NBt5qnJFG4aMmZbQin",
	"sszZUVRCYwH1oiQuaFw9iQ3iJr8yfKBnB3",
	"yZeg6v8gZa8c9A56BjyKdySU6YS5nUhCY3",
	"rN8dKuPggD5ANuQFHW8XTGK6fmURNEaNZu",
	"2iJgHD1X5vySZG2JptuT5kRGPxKiCpuynHY",
	"2GmnTEbfic8ogaBHc6z6i44rgN5jB6Vf54s",
	"2hucgXrFs8Wzcy1EUMYyRjs967qCQVbXvwZ",
	"2jVq4asXnZtLxXD4xe8FnnyFjfQqTCJjH8a",
	"2KJ1HvBg37DzvhRMfkoyeiozyCwPTohYbhX",
	"YsAgDNMbBXjpPR2Nr2psfSRoxZKWGCddGv",
	"GZbj4w5uMSzQrdcdamhEZJtrHrMZQnSqwZ",
	"26zvUZdreaofTKzoNzNB7BTFMBZvQEVj5Hi",
	"2m87oG3drwJCC3SYxEwhgyf6fqH5vExdDTD",
	"Pm9iCxUGsGFjj1HDLtrkjbt5d9MNwHKs4t",
	"Q4oEt8UDDKR4sArKsdkRqdiujrfyB1bh9j",
	"CSMx9R5hTG6nez7PCnG8iTwycTxgri2maW",
	"2df7uMvj1Yt8CqbMjtcPFLsUEcpmR72Xost",
	"2K72eTr67QqRTkvzaejtZVJLyeG7io8PaQP",
	"2mTiezRAF2usL2FCPHH1Wk9Uk2LVpMrsdmJ",
	"5VocLQjZxvThny4ghYwNHyRXjeCfXJkPJ8",
	"SSYLBPthVM8YDizHxQz6rVs4VnP5U1DM9B",
	"ra829hwPEhWBzDHgN1edGLYJrHwyEtZ5Eg",
	"Kp6H1yJ1Bz71JgYpXRACFRMKEYqdRBmWKv",
	"PesQzGtt4mgyCFRXDMYzYjDiQxiN4momLy",
	"XsKg55DypGhAeZk16YUsqAtCGDbsUu7u6J",
	"jz77DGua4Eu2wpQtA6xJTAP4ThvQ9pfyyn",
	"218q9KsVaU4umecnFezoz9hk8NQFhtmhSaP",
	"HEjrXtsoACQXjkiWcmtmnf4VSwr9JS2RED",
	"rUy4XiWztu9MzndcUXSto28wF2naMf9W1K",
	"Ub4zjFw3R4zkX5yP3TWzvFpMRwE6Gk2cLz",
	"rKtyDKPwjk7ouVQEomaLGzDVWBkBwhSH2N",
	"oFtDnyXwoevLvkxGpr594sS4ix7Fn1mNqk",
	"26dLZKqz3y729HvByFH79dV67KUSHcrMmZF",
	"pGKTHUoVq7bRh9kyge5JdQn4zwrPw9Ro99",
	"2DynhXHuSDGHpC2iVEd6RKJQ73YW9YZxcWD",
	"2ZJt5BzPPHsNTuM9T4TaCfwR9QLTCNBE5C",
	"TuAe6ZfDXieKd4sdz9tPZn5MUBXryCcg21",
}
