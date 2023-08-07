package auth

type KetoRelationInput struct {
	Namespace  string      `json:"namespace"`
	Object     string      `json:"object"`
	Relation   string      `json:"relation"`
	SubjectID  *string     `json:"subjectId"`
	SubjectSet *SubjectSet `json:"subjectSet"`
}

type SubjectSet struct {
	Namespace string `json:"namespace"`
	Object    string `json:"object"`
	Relation  string `json:"relation"`
}
