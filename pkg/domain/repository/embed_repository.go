package repository

type EmbedRepository interface {
	GetEmbeddings(text string) ([]float64, error)
}
