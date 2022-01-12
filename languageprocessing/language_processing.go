package languageprocessing

import (
	"sort"
	"strings"

	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

type pairknowledge struct {
	Key   string
	Value int
}

type pairknowledgeList []pairknowledge

func (p pairknowledgeList) Len() int           { return len(p) }
func (p pairknowledgeList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p pairknowledgeList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type LanguageProcessing struct {
	KnowledgeBase   map[string]string
	KnowledgeCorpus []string
}

func (lp *LanguageProcessing) Simple(query string) string {

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

	mK := kl[0].Key

	result = lp.KnowledgeBase[mK]

	return result
}

func (lp *LanguageProcessing) Complex(query string) string {

	var result string

	vectoriser := nlp.NewCountVectoriser(" ")
	transformer := nlp.NewTfidfTransformer()

	reducer := nlp.NewTruncatedSVD(4)

	matrix, _ := vectoriser.FitTransform(lp.KnowledgeCorpus...)
	matrix, _ = transformer.FitTransform(matrix)
	lsi, _ := reducer.FitTransform(matrix)

	matrix, _ = vectoriser.Transform(query)
	matrix, _ = transformer.Transform(matrix)
	queryVector, _ := reducer.Transform(matrix)

	highestSimilarity := -1.0
	var matched int
	_, docs := lsi.Dims()
	for i := 0; i < docs; i++ {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), lsi.(mat.ColViewer).ColView(i))
		if similarity > highestSimilarity {
			matched = i
			highestSimilarity = similarity
		}
	}

	if highestSimilarity == -1 {
		result = "I don't know the answer to that one."
	} else {
		result = lp.KnowledgeBase[lp.KnowledgeCorpus[matched]]
	}

	return result
}
