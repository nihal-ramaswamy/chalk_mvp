package dto

type Skills struct {
	SkillName     string `json:"skill_name"`
	SkillCategory string `json:"skill_category"`
}

func NewSkills() *Skills {
	return &Skills{}
}
