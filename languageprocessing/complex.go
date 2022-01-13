package languageprocessing

import (
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

type ComplexPorcessor struct {
	LanguageKnowledge
}

func (lp *ComplexPorcessor) FindAnswerByQuery(query string) string {

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
		return DEFAULT_RESPONSE
	}

	result = lp.KnowledgeBase[lp.KnowledgeCorpus[matched]]

	return result
}
