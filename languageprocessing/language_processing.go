package languageprocessing

const DEFAULT_RESPONSE = "I don't know the answer to that one."

type LanguageProcessing interface {
	FindAnswerByQuery(query string) string
}
type LanguageKnowledge struct {
	KnowledgeBase   map[string]string
	KnowledgeCorpus []string
}

func GetAnswer(q string, lp LanguageProcessing) string {
	return lp.FindAnswerByQuery(q)
}
