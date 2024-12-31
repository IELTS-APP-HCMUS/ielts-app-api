package services

import (
	"context"
	"errors"
	"fmt"
	"ielts-app-api/internal/models"
	"strings"

	"gorm.io/gorm"
)

func (s *Service) GetVocabById(ctx context.Context, userId string) ([]*models.UserVocabBank, error) {
	vocabs, err := s.vocabBankRepo.List(ctx, models.QueryParams{}, func(tx *gorm.DB) {
		tx.Where("user_id", userId)
	})
	if err != nil {
		return nil, err
	}
	return vocabs, nil
}

func (s *Service) CreateVocab(ctx context.Context, userId string, body models.VocabRequest) (*models.UserVocabBank, error) {
	vocabs, err := s.vocabBankRepo.List(ctx, models.QueryParams{}, func(tx *gorm.DB) {
		tx.Where("user_id", userId).Where("value", body.Word)
	})
	vocab := &models.UserVocabBank{
		Key:             fmt.Sprintf("%s_%d", body.Word, len(vocabs)+1),
		Value:           body.Word,
		WordClass:       body.WordType,
		Meaning:         body.Meaning,
		IPA:             body.IPA,
		Example:         body.Example,
		Explanation:     body.Note,
		IsLearnedStatus: body.Status == "Đã học",
		UserId:          userId,
	}

	vocab, err = s.vocabBankRepo.Create(ctx, vocab)
	if err != nil {
		return nil, err
	}

	return vocab, nil
}

func (s *Service) UpdateVocab(ctx context.Context, userId string, vocabValue models.VocabQuery, body models.VocabRequest) (*models.UserVocabBank, error) {
	vocab, err := s.vocabBankRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("user_id = ?", userId).Where("key = ?", vocabValue.Key)
	})
	if err != nil {
		return nil, err
	}

	// Just update the 'IsLearnedStatus' field
	var updateColumns map[string]interface{}

	if body.Status != "" {
		updateColumns = map[string]interface{}{
			"IsLearnedStatus": body.Status == "Đã học",
		}
	}
	vocab, err = s.vocabBankRepo.UpdateColumns(ctx, vocab.ID, updateColumns)
	if err != nil {
		return nil, err
	}
	return vocab, nil
}

func (s *Service) DeleteVocab(ctx context.Context, userId string, vocabValue models.VocabQuery) error {
	err := s.vocabBankRepo.Delete(ctx, func(tx *gorm.DB) {
		tx.Where("user_id = ?", userId).Where("key = ?", vocabValue.Key)
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetReadingVocab(ctx context.Context, request models.LookUpVocabRequest) (*models.Vocab, error) {
	vocabId := fmt.Sprintf("%d_%d_%d", request.QuizId, request.SentenceIndex, request.WordIndex)

	// Tìm vocab trong câu cụ thể trước
	vocab, err := s.vocabRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("vocab_id = ?", vocabId)
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Nếu không tìm thấy trong câu, thử tìm kiếm theo câu
			vocab, err = s.LookUpVocabLinear(ctx, request.QuizId, request.SentenceIndex, request.Word)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// Nếu vẫn không tìm thấy, tìm trên toàn bài
					vocab, err = s.LookUpVocabGlobal(ctx, request.QuizId, request.Word)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	} else {
		// Nếu tìm thấy nhưng giá trị không khớp, tìm lại trên toàn bài
		if vocab.Value != request.Word {
			vocab, err = s.LookUpVocabGlobal(ctx, request.QuizId, request.Word)
			if err != nil {
				return nil, err
			}
		}
	}

	vocab.Explanation = strings.ReplaceAll(vocab.Explanation, "\"", "")
	return vocab, nil
}

func (s *Service) LookUpVocabLinear(ctx context.Context, quizId int, sentenceIndex int, word string) (*models.Vocab, error) {
	vocabIdPattern := fmt.Sprintf("%d_%d_%%", quizId, sentenceIndex)

	vocab, err := s.vocabRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("vocab_id LIKE ? AND value ILIKE ?", vocabIdPattern, word)
	})
	if err != nil {
		return nil, err
	}

	return vocab, nil
}

func (s *Service) LookUpVocabGlobal(ctx context.Context, quizId int, word string) (*models.Vocab, error) {
	vocabIdPattern := fmt.Sprintf("%d_%%", quizId)

	vocab, err := s.vocabRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("vocab_id LIKE ? AND value ILIKE ?", vocabIdPattern, word)
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Nếu không tìm thấy từ, lấy đại một từ nào đó trong bài
			vocab, err = s.vocabRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
				tx.Where("vocab_id LIKE ?", vocabIdPattern).Limit(1)
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return vocab, nil
}
