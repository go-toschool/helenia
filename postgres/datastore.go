package postgres

import (
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-toschool/helenia"
	"github.com/jmoiron/sqlx"
)

var (
	errMissingRequiredFields = errors.New("Missing required fields. talk_id, speaker_id(user) and assistant_id(user)")
	errMissingAssistID       = errors.New("Missing assist id. Assist id must not be empty")
	errMissingAssistantID    = err.New("Missing assistant id. Assistant id (user_id) must not be empty")
)

// AssistsStore is the backend db connection. This struct implements the
// helenia.Assists interface that is bassically a crud over assists table
type AssistsStore struct {
	*sqlx.DB
}

// AssistsService implements specific logic over assists table on database
type AssistsService struct {
	*AssistsStore
}

// Add implements the method to register a talk from Assists interfacer
// If the speaker id is empty, will return an error
func (as *AssistsStore) Add(aq *helenia.AssistQuery) (a *helenia.Assist, err error) {
	if aq.SpeakerID == "" {
		return a, errMissingRequiredFields
	}

	query, args, err := sq.Insert("assits").
		Columns("talk_id", "speaker_id", "assistant_id").
		Values(aq.TalkID, aq.SpeakerID, aq.AssistantID).
		Suffix("returning *").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return a, err
	}

	row := as.QueryRowx(query, args...)
	err = row.StructScan(a)

	return a, err
}

// Get implements the search by assist id from assists interface
func (as *AssistsStore) Get(aq *helenia.AssistQuery) (a *helenia.Assist, err error) {
	if aq.AssistID == "" {
		return a, errMissingAssistID
	}

	query, args, err := sq.Select("*").
		From("assits").
		Where("deleted_at is not null").
		ToSql()

	row := as.QueryRowx(query, args...)
	err = row.StructScan(a)

	return a, err
}

// Update implements the update function from Assists interface
func (as *AssistsStore) Update(aq *helenia.AssistQuery) (a *helenia.Assist, err error) {
	query, args, err := sq.Update("assits").
		Set("speaker_id", aq.SpeakerID).
		Set("assistant_id", aq.AssistantID).
		Where("id = ?", aq.AssistID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return a, err
	}

	row := as.QueryRowx(query, args...)
	err = row.StructScan(a)

	return a, err
}

// Delete implements the Delete method of Assist interface
func (as *AssistsStore) Delete(aq *helenia.AssistQuery) (a *helenia.Assist, err error) {
	if aq.AssistID == "" {
		return a, errMissingAssistID
	}

	query, args, err := sq.Update("talks").
		Set("deleted_at", time.Now()).
		Where("id = ?", aq.AssistID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return a, err
	}

	row := as.QueryRowx(query, args...)
	err = row.StructScan(a)

	return a, err
}

func (ass *AssistsService) FindAssistsByAssistantID(aq *helenia.AssistQuery) (aa []*helenia.Assist, err error) {
	if aq.AssistantID == "" {
		return aa, errMissingAssistantID
	}

	query, args, err := sq.Select("*").
		From("assists").
		Where("assistant_id = ?", aq.AssistantID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return aa, err
	}

	row := ass.QueryRowx(query, args...)
	err = row.StructScan(aa)

	return aa, err
}
