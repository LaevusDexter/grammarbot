package grammarbot

import (
	"net/http"
)

type GrammarBot struct {
	Language string
	ApiKey string
	BaseURI string
	Version string
	ApiName string

	Client *http.Client
}

type Response struct {
	Software BotInfo `json:"software"`
	Warnings BotWarnings `json:"warnings"`
	Language RequestLanguage `json:"language"`
	Matches  []*Match `json:"matches"`
}

type Match struct {
	Message      string `json:"message"`
	ShortMessage string `json:"shortMessage"`
	Replacements []MatchReplacement `json:"replacements"`
	Offset  int `json:"offset"`
	Length  int `json:"length"`
	Context MatchContext `json:"context"`
	Sentence string `json:"sentence"`
	Type MatchType `json:"type"`
	Rule *MatchRule `json:"rule"`
}

type MatchRule struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	IssueType   string `json:"issueType"`
	Category RuleCategory `json:"category"`
}

type RuleCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MatchType struct {
	TypeName string `json:"typeName"`
}

type MatchContext struct {
	Text   string `json:"text"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type MatchReplacement struct {
	Value string `json:"value"`
}

type BotInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	APIVersion  int    `json:"apiVersion"`
	Premium     bool   `json:"premium"`
	PremiumHint string `json:"premiumHint"`
	Status      string `json:"status"`
}

type BotWarnings struct {
	IncompleteResults bool `json:"incompleteResults"`
}

type RequestLanguage struct {
	Name             string `json:"name"`
	Code             string `json:"code"`
	Detected DetectedLanguage `json:"detectedLanguage"`
}

type DetectedLanguage struct {
	Name string `json:"name"`
	Code string `json:"code"`
}