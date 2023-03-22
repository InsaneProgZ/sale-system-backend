package handler

var ErrorHandler = func(err any) {
	if err != nil {
		panic(err)
	}
}
