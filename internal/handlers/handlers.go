package handlers

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/auth"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/repository"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LoginHandler is a demo that returns a JWT token (in production use real auth)
func LoginHandler(c *fiber.Ctx) error {
	type cred struct {
		User string `json:"user"`
	}
	var rc cred
	if err := c.BodyParser(&rc); err != nil {
		return fiber.ErrBadRequest
	}
	// generate token (simple demo)
	token, err := auth.GenerateJWT(rc.User)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{"token": token})
}

func CreateSMS(c *fiber.Ctx) error {
	var s entity.SMS
	if err := c.BodyParser(&s); err != nil {
		return fiber.ErrBadRequest
	}
	s.Status = "pending"
	if err := repository.UpsertSMS(&s); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(s)
}

func GetSMS(c *fiber.Ctx) error {
	id := c.Params("id")
	s, err := repository.GetSMSByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(s)
}

func ListSMS(c *fiber.Ctx) error {
	sms, err := repository.ListSMSes(100)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(sms)
}

func CreateRule(c *fiber.Ctx) error {
	var r entity.Rule
	if err := c.BodyParser(&r); err != nil {
		return fiber.ErrBadRequest
	}
	if r.Action != "allow" && r.Action != "deny" {
		r.Action = "deny"
	}
	if err := repository.UpsertRule(&r); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(r)
}

func GetRule(c *fiber.Ctx) error {
	id := c.Params("id")
	r, err := repository.GetRuleByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(r)
}

func ListRules(c *fiber.Ctx) error {
	rules, err := repository.ListRules(100)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rules)
}

func EvaluateSMS(c *fiber.Ctx) error {
	var s entity.SMS
	if err := c.BodyParser(&s); err != nil {
		return fiber.ErrBadRequest
	}
	// Load rules
	rules, err := repository.GetAllRules()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	res := entity.EvaluationResult{Allowed: true}
	for _, r := range rules {
		matched := false
		// sender match
		if len(r.Senders) > 0 {
			for _, sd := range r.Senders {
				if sd == s.From {
					matched = true
					break
				}
			}
		}
		// receiver match
		if !matched && len(r.Receivers) > 0 {
			for _, rc := range r.Receivers {
				if rc == s.To {
					matched = true
					break
				}
			}
		}
		// body regex
		if !matched && r.BodyRegex != "" {
			re, err := regexp.Compile(r.BodyRegex)
			if err == nil {
				if re.MatchString(s.Body) {
					matched = true
				}
			}
		}
		if matched {
			res.MatchedIDs = append(res.MatchedIDs, r.ID)
			if r.Action == "deny" {
				res.Allowed = false
				break // deny short-circuits
			}
		}
	}
	// optionally persist evaluation result by updating sms status
	if s.ID == "" {
		s.ID = "eval-" + time.Now().Format("20060102150405")
	}
	if res.Allowed {
		s.Status = "allowed"
	} else {
		s.Status = "blocked"
	}
	_ = repository.UpsertSMS(&s)
	return c.JSON(res)
}
