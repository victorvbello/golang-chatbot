package languageprocessing

import (
	"sort"
	"strings"
)

type pairknowledge struct {
	Key   string
	Value int
}

type pairknowledgeList []pairknowledge

func (p pairknowledgeList) Len() int           { return len(p) }
func (p pairknowledgeList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p pairknowledgeList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type SiplePorcessor struct {
	LanguageKnowledge
}

func (lp *SiplePorcessor) FindAnswerByQuery(query string) string {

	var result string
	var index int

	knowledgePoints := map[string]int{}
	lowerQuery := strings.ToLower(query)
	sliceQuery := strings.Split(lowerQuery, " ")

	for _, v := range lp.KnowledgeCorpus {
		for _, vQ := range sliceQuery {
			if strings.Contains(v, vQ) {
				knowledgePoints[v]++
			}
		}
	}

	kl := make(pairknowledgeList, len(knowledgePoints))

	for k, v := range knowledgePoints {
		kl[index] = pairknowledge{k, v}
		index++
	}
	sort.Sort(sort.Reverse(kl))

	if len(kl) == 0 {
		return DEFAULT_RESPONSE
	}

	mK := kl[0].Key

	result = lp.KnowledgeBase[mK]

	return result
}
