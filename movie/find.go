package movie

func Find(imdbbId string) string {
	switch imdbbId {
	case "A00001":
		return "Avatar: Endgame"
	case "A00002":
		return "Spiderman"

	default:
		return "Noting to find"
	}
}
