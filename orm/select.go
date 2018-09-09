package orm

// Select Test
func Select(number int) (int, error) {
	// Prepare statement for reading data
	stmtOut, err := pool.Prepare("SELECT squareNumber FROM squareNum WHERE number = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(number).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		return 0, err
	}

	return squareNum, nil
}
