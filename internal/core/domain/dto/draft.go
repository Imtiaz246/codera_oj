package dto

import (
	"github.com/imtiaz246/codera_oj/internal/core/domain/models/problem"
)

type CreateProblemOption struct {
	Title string `json:"title" validate:"required"`
}

type UpdateProblemOption struct {
	TimeLimit                         float64             `json:"timeLimit"`
	MemoryLimit                       float64             `json:"memoryLimit"`
	Statement                         string              `json:"statement"`
	InputStatement                    string              `json:"inputStatement"`
	OutputStatement                   string              `json:"outputStatement"`
	NoteStatement                     string              `json:"noteStatement"`
	StatementsVisibilityDuringContest bool                `json:"statementsVisibilityDuringContest"`
	CheckerType                       problem.CheckerType `json:"checkerType"`
}

type DatasetOption struct {
	Title      string                `json:"title"`
	Weight     int64                 `json:"weight"`
	IsSample   bool                  `json:"isSample"`
	InputContent  []byte `json:"inputFile"`
	OutputContent []byte `json:"outputFile"`
}

type ShareProblemOption struct {
	ShareWith  string             `json:"shareWith"`
	PermitType problem.PermitType `json:"permitType"`
}

type SolutionOption struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}
