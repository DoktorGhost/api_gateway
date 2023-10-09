package types

type NewsShortDetailed struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewsFullDetailed struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	PubTime int64  `json:"pubtime"`
	Link    string `json:"link"`
}

type Comment struct {
	ID              int    `json:"id"`
	NewsID          int    `json:"news_id"`
	Text            string `json:"text"`
	ParentCommentID int    `json:"parent_id"`
}

type AggregatorResponse struct {
	News   NewsShortDetailed
	Errors []error
}

var VerificationResult struct {
	UniqueID string `json:"uniqueID"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}
