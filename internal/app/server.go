package app

func StartApp() {
	application := App{}
	application.Initialize()
	application.Run("8000")
}