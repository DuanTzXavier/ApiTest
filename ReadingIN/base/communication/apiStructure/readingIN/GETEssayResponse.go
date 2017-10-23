package readingIN

type GETEssayResponse struct {
	NextID string
	PreviousID string
	EssayName string
	EssayContents []EssayContent
	EssayAuthor string
	EssayWordsCount int
	EssayFrom string
	EssayCreatTime string
}
