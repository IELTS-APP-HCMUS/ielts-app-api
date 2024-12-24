package services

import (
	"context"
	"ielts-app-api/internal/models"

	"gorm.io/gorm"
)

func (s *Service) GetVocabById(ctx context.Context, userId string) ([]*models.Vocab, error) {
	vocabs, err := s.vocabRepo.List(ctx, models.QueryParams{}, func(tx *gorm.DB) {
		tx.Where("user_id", userId)
	})
	if err != nil {
		return nil, err
	}
	return vocabs, nil
}

func (s *Service) CreateVocab(ctx context.Context, userId string, body models.VocabRequest) (*models.Vocab, error) {
	vocab := &models.Vocab{
		Value:           body.Value,
		WordClass:       body.WordClass,
		Meaning:         body.Meaning,
		IPA:             body.IPA,
		Example:         body.Example,
		Explanation:     body.Explanation,
		IsLearnedStatus: body.IsLearnedStatus,
		UserId:          userId,
	}

	vocab, err := s.vocabRepo.Create(ctx, vocab)
	if err != nil {
		return nil, err
	}

	return vocab, nil
}

// func (s *Service) VocabSuggest(ctx context.Context, userID string, query models.VocabSuggestQuery) (vocab *models.Vocab, err error) {
// 	filters := []repositories.Clause{}
// 	filters = append(filters, func(tx *gorm.DB) {
// 		tx.Where("user_id = ? AND category = ? AND period_type = ?", userID, common.AIUsageCategoryVocabTranslate, common.PeriodTypeWeek)
// 	})

// 	usage, err := s.aiWritingRepo.GetDetailByConditions(ctx, filters...)
// 	if err != nil {
// 		return nil, common.ErrorWrapper("AIUsageCountRepo.GetDetailByConditions", err)
// 	}
// 	if usage != nil && *usage.Used >= *usage.Total {
// 		return nil, errors.New("quota exceeded: you have used all your AI vocab quota")
// 	}

// 	question, err := s.questionRepo.GetByID(ctx, query.QuestionID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	clauses := []repositories.Clause{}
// 	clauses = append(clauses, func(tx *gorm.DB) {
// 		tx.Where("value = ?", query.Input)
// 		tx.Where("question_id = ?", query.QuestionID)

// 	})
// 	vocabCheck, err := s.vocabRepo.List(ctx, models.QueryParams{}, clauses...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(vocabCheck) == 0 {
// 		vocabCheck = nil
// 	}

// 	if vocabCheck != nil {
// 		return processVocabTransaction(ctx, s, userID, vocabCheck[0].ID, usage, vocabCheck[0])
// 	}

// 	vocabPrompt := &models.VocabPrompt{
// 		Input:    query.Input,
// 		Question: question.Title,
// 		Type:     common.SuggestPhrase,
// 	}

// 	promptResponse, err := s.getPromptVocabData(ctx, vocabPrompt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	synonyms := []string{}
// 	if promptResponse.Synonyms.Synonym1 != "" {
// 		synonyms = append(synonyms, promptResponse.Synonyms.Synonym1)
// 	}
// 	if promptResponse.Synonyms.Synonym2 != "" {
// 		synonyms = append(synonyms, promptResponse.Synonyms.Synonym2)
// 	}

// 	examples := []models.Example{}

// 	for _, example := range promptResponse.Examples {
// 		examples = append(examples, models.Example{
// 			ExampleText:     example.ExampleText,
// 			TranslationText: example.TranslationText,
// 		})
// 	}
// 	jsonExamples, err := json.Marshal(examples)
// 	if err != nil {
// 		return nil, err
// 	}
// 	vocab = &models.Vocab{
// 		Value:       query.Input,
// 		Synonyms:    synonyms,
// 		Examples:    jsonExamples,
// 		Translation: promptResponse.Translation,
// 		QuestionID:  query.QuestionID,
// 		Source:      common.SourceAIFromVieToEng,
// 		QuizID:      query.QuizID,
// 	}

// 	newVocab, err := s.vocabRepo.Create(ctx, vocab)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return processVocabTransaction(ctx, s, userID, newVocab.ID, usage, vocab)
// }

// func processVocabTransaction(ctx context.Context, s *Service, userID string, vocabValue int, usage *models.AIUsageCount, vocabCheck *models.Vocab) (*models.Vocab, error) {
// 	updateColumns := map[string]interface{}{
// 		"used": *usage.Used + 1,
// 	}
// 	_, err := s.aiWritingRepo.UpdateColumns(ctx, usage.ID, updateColumns)
// 	if err != nil {
// 		return nil, err
// 	}

// 	vocabSuggest := map[string]string{
// 		"vocab_id": strconv.Itoa(vocabValue),
// 	}
// 	valueMeta, err := json.Marshal(vocabSuggest)
// 	if err != nil {
// 		return nil, err
// 	}

// 	newTransaction := &models.AIUsageCountTransaction{
// 		UserID:   userID,
// 		Category: common.AIUsageCategoryVocabTranslate,
// 		Meta:     valueMeta,
// 	}
// 	_, err = s.aiWritingRepo.Transaction.Create(ctx, newTransaction)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return vocabCheck, nil
// }
