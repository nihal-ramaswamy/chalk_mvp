package db

import (
	"database/sql"

	"github.com/nihal-ramaswamy/chalk_mvp/internal/dto"
)

func GetAllValues(db *sql.DB) ([]dto.Skills, error) {
	return selectFromSkills(db)
}

func GetValuesForCategory(db *sql.DB, category string) ([]dto.Skills, error) {
	return selectFromSkillsWhereSkillCAtegoryIs(db, category)
}

func selectFromSkills(db *sql.DB) ([]dto.Skills, error) {
	var skills []dto.Skills

	query := `SELECT * FROM "SKILLS"`

	rows, err := db.Query(query)
	if nil != err {
		return skills, err
	}
	defer rows.Close()

	for rows.Next() {
		var skill dto.Skills

		if err := rows.Scan(&skill.SkillName, &skill.SkillCategory); nil != err {
			return skills, err
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func selectFromSkillsWhereSkillCAtegoryIs(db *sql.DB, category string) ([]dto.Skills, error) {
	var skills []dto.Skills

	query := `SELECT * FROM "SKILLS" WHERE SKILL_CATEGORY = $1`

	rows, err := db.Query(query, category)
	if nil != err {
		return skills, err
	}
	defer rows.Close()

	for rows.Next() {
		var skill dto.Skills

		if err := rows.Scan(&skill.SkillName, &skill.SkillCategory); nil != err {
			return skills, err
		}
		skills = append(skills, skill)
	}

	return skills, nil
}
