package repository

import (
	"context"
	"errors"

	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/model"
	"github.com/jackc/pgx/v5"
)

type FactRepository struct {
	db *pgx.Conn
}

func NewFactRepository(conn *pgx.Conn) *FactRepository {
	return &FactRepository{db: conn}
}

func (r *FactRepository) GetAll(ctx context.Context) ([]model.Fact, error) {
	query := `SELECT id, text FROM facts`

	rows,err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var facts []model.Fact

	for rows.Next() {
		var fact model.Fact
		if err := rows.Scan(&fact.ID, &fact.Text); err != nil {
			return nil, err
		}
		facts = append(facts, fact)
	}

	return facts, nil
}

func (r *FactRepository) GetRandom(ctx context.Context) (*model.Fact, error) {
	query := `
        SELECT id, text FROM facts
        OFFSET floor(random() * (SELECT COUNT(*) FROM facts))::int LIMIT 1
    `
	var fact model.Fact
	err := r.db.QueryRow(ctx, query).Scan(&fact.ID, &fact.Text)
	if err != nil {
		return nil, err
	}
	return &fact, nil
}

func (r *FactRepository) GetByID(ctx context.Context, id uint64) (*model.Fact, error) {
	query := `SELECT id, text FROM facts WHERE id=$1`

	var fact model.Fact
	err := r.db.QueryRow(ctx, query, id).Scan(&fact.ID, &fact.Text)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("fact not found")
		}
		return nil, err
	}
	return &fact, nil
}

func (r *FactRepository) Create(ctx context.Context, req dtos.CreateFactRequest) (*model.Fact, error) {
	query := `INSERT INTO facts (text) VALUES ($1) RETURNING id, text`
	
	var fact model.Fact
	err := r.db.QueryRow(
		ctx,
		query,
		req.Text,
	).Scan(&fact.ID, &fact.Text)

	if err != nil {
		return nil, err
	}
	return &fact, nil
}

func (r *FactRepository) Update(ctx context.Context, req dtos.UpdateFactRequest) (*model.Fact, error) {
	query := `UPDATE facts SET text=$1 WHERE id=$2 RETURNING id, text`
	
	var fact model.Fact
	err := r.db.QueryRow(
		ctx,
		query,
		req.Text, req.ID,
	).Scan(&fact.ID, &fact.Text)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("fact not found")
		}
		return nil, err
	}
	return &fact, nil
}

func (r *FactRepository) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM facts WHERE id=$1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("fact not found")
	}
	return nil
}