package structs

import (
	"fmt"
	"github.com/imtiaz246/codera_oj/models"
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type CreateProblemOption struct {
	Title string `json:"title" validate:"required"`
}

type UpdateProblemOption struct {
	TimeLimit                         float64            `json:"timeLimit"`
	MemoryLimit                       float64            `json:"memoryLimit"`
	Statement                         string             `json:"statement"`
	InputStatement                    string             `json:"inputStatement"`
	OutputStatement                   string             `json:"outputStatement"`
	NoteStatement                     string             `json:"noteStatement"`
	StatementsVisibilityDuringContest bool               `json:"statementsVisibilityDuringContest"`
	CheckerType                       models.CheckerType `json:"checkerType"`
}

func (p *UpdateProblemOption) UpdateProblemModelFormat(user *models.User, ID uint) *models.Problem {
	return &models.Problem{
		Model:                             gorm.Model{ID: ID},
		Author:                            user,
		TimeLimit:                         p.TimeLimit,
		MemoryLimit:                       p.MemoryLimit,
		Statement:                         p.Statement,
		InputStatement:                    p.InputStatement,
		OutputStatement:                   p.OutputStatement,
		NoteStatement:                     p.NoteStatement,
		StatementsVisibilityDuringContest: p.StatementsVisibilityDuringContest,
		CheckerType:                       p.CheckerType,
	}
}

type AddTagOption struct {
	TagName string `json:"tagName"`
}

type ShareProblemOption struct {
	ShareWith  string            `json:"shareWith"`
	PermitType models.PermitType `json:"permitType"`
}

type DatasetOption struct {
	Title    string `json:"title"`
	Weight   int64  `json:"weight"`
	IsSample bool   `json:"isSample"`
}

func (d *DatasetOption) NewModelDatasetFormat(inputFile, outputFile *multipart.FileHeader, pid, userid uint) (*models.Dataset, error) {
	input, err := inputFile.Open()
	defer input.Close()
	if err != nil {
		return nil, err
	}
	output, err := outputFile.Open()
	defer output.Close()
	if err != nil {
		return nil, err
	}

	if d.IsSample && inputFile.Size > 2000 {
		return nil, fmt.Errorf("input file is too large for being considered as sample")
	}

	inputContent := make([]byte, inputFile.Size)
	outputContent := make([]byte, outputFile.Size)
	if _, err = input.Read(inputContent); err != nil {
		return nil, err
	}
	if _, err = output.Read(outputContent); err != nil {
		return nil, err
	}

	return &models.Dataset{
		Title:     d.Title,
		Weight:    d.Weight,
		IsSample:  d.IsSample,
		Input:     inputContent,
		Output:    outputContent,
		UserID:    userid,
		ProblemID: pid,
	}, nil
}

type SolutionOption struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

type SolutionResponse struct {
	ID                  uint                 `json:"ID"`
	Code                string               `json:"code"`
	Language            string               `json:"language"`
	LastExecuted        *time.Time           `json:"lastExecuted,omitempty"`
	TimeTaken           *float64             `json:"timeTaken,omitempty"`
	MemoryTaken         *float64             `json:"memoryTaken,omitempty"`
	ProblemUserRelation *ProblemUserRelation `json:"problemUserRelation"`
}

type ProblemUserRelation struct {
	Username  string            `json:"username"`
	Role      models.PermitType `json:"role"`
	ProblemID uint              `json:"problemID"`
}

func NewProblemUserRelation(problemShare *models.ProblemShare) *ProblemUserRelation {
	return &ProblemUserRelation{
		Username:  problemShare.SharedWith.Username,
		Role:      problemShare.PermissionType,
		ProblemID: problemShare.ProblemID,
	}
}

func NewSolutionResponse(solution *models.ProblemSolution, problemShare *models.ProblemShare) *SolutionResponse {
	return &SolutionResponse{
		ID:                  solution.ID,
		Code:                solution.Code,
		Language:            solution.Language.Name,
		LastExecuted:        solution.LastExecuted,
		TimeTaken:           solution.TimeTaken,
		MemoryTaken:         solution.MemoryTaken,
		ProblemUserRelation: NewProblemUserRelation(problemShare),
	}
}

type DiscussionOption struct {
	Message string `json:"message"`
}
