package jsonq

// func BenchmarkKeepSmallSimple(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallSimple(false, false)
// 	}
// }

// func BenchmarkKeepSmallSimpleParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallSimple(false, true)
// 	}
// }

// func BenchmarkKeepSmallSimpleAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallSimple(true, true)
// 	}
// }

// func BenchmarkKeepSmallMedium(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallMedium(false, false)
// 	}
// }

// func BenchmarkKeepSmallMediumParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallMedium(false, true)
// 	}
// }

// func BenchmarkKeepSmallMediumAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallMedium(true, true)
// 	}
// }
// func BenchmarkKeepSmallHard(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallHard(false, false)
// 	}
// }

// func BenchmarkKeepSmallHardParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallHard(false, true)
// 	}
// }

// func BenchmarkKeepSmallHardAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepSmallHard(true, true)
// 	}
// }

// func BenchmarkKeepMediumSimple(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumSimple(false, false)
// 	}
// }

// func BenchmarkKeepMediumSimpleParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumSimple(false, true)
// 	}
// }

// func BenchmarkKeepMediumSimpleAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumSimple(true, true)
// 	}
// }

// func BenchmarkKeepMediumMedium(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumMedium(false, false)
// 	}
// }

// func BenchmarkKeepMediumMediumParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumMedium(false, true)
// 	}
// }

// func BenchmarkKeepMediumMediumAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumMedium(true, true)
// 	}
// }
// func BenchmarkKeepMediumHard(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumHard(false, false)
// 	}
// }

// func BenchmarkKeepMediumHardParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumHard(false, true)
// 	}
// }

// func BenchmarkKeepMediumHardAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepMediumHard(true, true)
// 	}
// }

// func BenchmarkKeepLargeSimple(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeSimple(false, false)
// 	}
// }

// func BenchmarkKeepLargeSimpleParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeSimple(false, true)
// 	}
// }

// func BenchmarkKeepLargeSimpleAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeSimple(true, true)
// 	}
// }

// func BenchmarkKeepLargeMedium(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeMedium(false, false)
// 	}
// }

// func BenchmarkKeepLargeMediumParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeMedium(false, true)
// 	}
// }

// func BenchmarkKeepLargeMediumAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeMedium(true, true)
// 	}
// }
// func BenchmarkKeepLargeHard(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeHard(false, false)
// 	}
// }

// func BenchmarkKeepLargeHardParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeHard(false, true)
// 	}
// }

// func BenchmarkKeepLargeHardAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		KeepLargeHard(true, true)
// 	}
// }

// // Check

// func BenchmarkCheckSmallSimple(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallSimple(false, false)
// 	}
// }

// func BenchmarkCheckSmallSimpleParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallSimple(false, true)
// 	}
// }

// func BenchmarkCheckSmallSimpleAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallSimple(true, true)
// 	}
// }

// func BenchmarkCheckSmallMedium(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallMedium(false, false)
// 	}
// }

// func BenchmarkCheckSmallMediumParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallMedium(false, true)
// 	}
// }

// func BenchmarkCheckSmallMediumAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallMedium(true, true)
// 	}
// }
// func BenchmarkCheckSmallHard(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallHard(false, false)
// 	}
// }

// func BenchmarkCheckSmallHardParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallHard(false, true)
// 	}
// }

// func BenchmarkCheckSmallHardAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckSmallHard(true, true)
// 	}
// }

// func BenchmarkCheckMediumSimple(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumSimple(false, false)
// 	}
// }

// func BenchmarkCheckMediumSimpleParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumSimple(false, true)
// 	}
// }

// func BenchmarkCheckMediumSimpleAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumSimple(true, true)
// 	}
// }

// func BenchmarkCheckMediumMedium(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumMedium(false, false)
// 	}
// }

// func BenchmarkCheckMediumMediumParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumMedium(false, true)
// 	}
// }

// func BenchmarkCheckMediumMediumAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumMedium(true, true)
// 	}
// }
// func BenchmarkCheckMediumHard(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumHard(false, false)
// 	}
// }

// func BenchmarkCheckMediumHardParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumHard(false, true)
// 	}
// }

// func BenchmarkCheckMediumHardAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckMediumHard(true, true)
// 	}
// }

// func BenchmarkCheckLargeSimple(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeSimple(false, false)
// 	}
// }

// func BenchmarkCheckLargeSimpleParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeSimple(false, true)
// 	}
// }

// func BenchmarkCheckLargeSimpleAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeSimple(true, true)
// 	}
// }

// func BenchmarkCheckLargeMedium(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeMedium(false, false)
// 	}
// }

// func BenchmarkCheckLargeMediumParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeMedium(false, true)
// 	}
// }

// func BenchmarkCheckLargeMediumAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeMedium(true, true)
// 	}
// }
// func BenchmarkCheckLargeHard(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeHard(false, false)
// 	}
// }

// func BenchmarkCheckLargeHardParseCMD(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeHard(false, true)
// 	}
// }

// func BenchmarkCheckLargeHardAll(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		CheckLargeHard(true, true)
// 	}
// }
