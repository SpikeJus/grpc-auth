package env

// var env Env

// type Env struct {
// 	RSAKeys rsa.PrivateKey
// }

func init() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Print("No .env file found")
	// } else {
	// 	publicKeyNString, exists := os.LookupEnv("RSA_PUBLIC_KEY_N")
	// 	exists(exists)
	// 	publicKeyN := new(big.Int)
	// 	_, err := fmt.Sscan(publicKeyNString, publicKeyN)
	// 	errConversion(err)

	// 	publicKeyEString, exists := os.LookupEnv("RSA_PUBLIC_KEY_E")
	// 	exists(exists)
	// 	publicKeyE, err := strconv.Atoi(publicKeyEString)
	// 	errConversion(err)

	// 	privateKeyD, exists := os.LookupEnv("RSA_D")
	// 	exists(exists)
	// 	d := new(big.Int)
	// 	_, err := fmt.Sscan(privateKeyD, d)
	// 	errConversion(err)

	// 	primes := []*big.Int{}
	// 	privateKeyPrimes, exists := os.LookupEnv("RSA_PRIMES")
	// 	exists(exists)

	// 	env = Env{
	// 		RSAKeys: rsa.PrivateKey{
	// 			PublicKey: rsa.PublicKey{publicKeyN, publicKeyE},
	// 			D: d,
	// 			Primes:
	// 		},
	// 	}
	// }
}

// func exists(exists bool) {
// 	if !exists {
// 		log.Println("No var found")
// 	}
// }

// func errConversion(err error) {
// 	if err != nil {
// 		log.Println("err converting key string")
// 	}
// }

// func GetPublicKey() string {
// 	return env.PublicKey
// }

// func GetPrivateKey() string {
// 	return env.PrivateKey
// }
