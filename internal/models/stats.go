package models

// StatsData - репозиторий возрващает структуру данных аналогичную этой.
type StatsData struct {
	UniqueUsers    int
	TotalEvents    int
	MostVisitedUrls []MostVisitedUrlData	// Список самых посещаемых URL.
}

type MostVisitedUrlData struct {
	Url   string
	Count int
}
