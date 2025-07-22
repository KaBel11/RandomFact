package repository

import (
	"context"
	"errors"

	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/model"
	"github.com/jackc/pgx/v5"
)

type PostgresFactsRepository struct {
	db *pgx.Conn
}

func NewPostgresFactsRepository(conn *pgx.Conn) *PostgresFactsRepository {
	return &PostgresFactsRepository{db: conn}
}

func (r *PostgresFactsRepository) GetAll() ([]model.Fact, error) {
	rows, err := r.db.Query(context.Background(), `SELECT id, text FROM facts`)
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

func (r *PostgresFactsRepository) GetRandom() (*model.Fact, error) {
	row := r.db.QueryRow(context.Background(), `
        SELECT id, text FROM facts
        OFFSET floor(random() * (SELECT COUNT(*) FROM facts))::int LIMIT 1
    `)

	var fact model.Fact
	if err := row.Scan(&fact.ID, &fact.Text); err != nil {
		return nil, err
	}
	return &fact, nil
}

func (r *PostgresFactsRepository) GetByID(id uint64) (*model.Fact, error) {
	row := r.db.QueryRow(context.Background(), `SELECT id, text FROM facts WHERE id=$1`, id)

	var fact model.Fact
	if err := row.Scan(&fact.ID, &fact.Text); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("fact not found")
		}
		return nil, err
	}
	return &fact, nil
}

func (r *PostgresFactsRepository) Create(req dtos.CreateFactRequest) (*model.Fact, error) {
	var fact model.Fact
	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO facts (text) VALUES ($1) RETURNING id, text`,
		req.Text,
	).Scan(&fact.ID, &fact.Text)

	if err != nil {
		return nil, err
	}
	return &fact, nil
}

func (r *PostgresFactsRepository) Update(req dtos.UpdateFactRequest) (*model.Fact, error) {
	var fact model.Fact
	err := r.db.QueryRow(
		context.Background(),
		`UPDATE facts SET text=$1 WHERE id=$2 RETURNING id, text`,
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

func (r *PostgresFactsRepository) Delete(id uint64) error {
	cmdTag, err := r.db.Exec(context.Background(), `DELETE FROM facts WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("fact not found")
	}
	return nil
}
