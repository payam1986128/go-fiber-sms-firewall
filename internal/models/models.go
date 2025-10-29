package models

import "time"

// SMS represents an outgoing (or attempted) SMS message
type SMS struct {
	ID        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"` // e.g., pending, allowed, blocked, sent
}

// Rule represents a regulatory rule that can match an SMS and decide allow/deny
type Rule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Simple example: a regex applied to Body, or list of blocked recipients, senders, etc.
	BodyRegex string    `json:"body_regex"`
	Senders   []string  `json:"senders"`
	Receivers []string  `json:"receivers"`
	Action    string    `json:"action"` // "allow" or "deny"
	CreatedAt time.Time `json:"created_at"`
}

// EvaluationResult used for evaluate endpoint
type EvaluationResult struct {
	Allowed    bool     `json:"allowed"`
	MatchedIDs []string `json:"matched_rule_ids"`
}
