package clientmodel

import (
	"errors"
	"regexp"
)

func ClientRegisterValidation(cl ClientRegisterRequest) error {
	if err := ValidateString(cl.Name); err != nil {
		return err
	}
	if err := ValidateString(cl.City); err != nil {
		return err
	}
	if err := ValidateHome(cl.Street); err != nil {
		return err
	}
	if err := ValidateHome(cl.Home_number); err != nil {
		return err
	}
	if err := ValidateEmail(cl.Email); err != nil {
		return err
	}
	if err := ValidatePhone(cl.Phone); err != nil {
		return err
	}
	if err := ValidateBankCard(cl.Bank_card); err != nil {
		return err
	}
	if err := ValidateBalance(cl.Balance); err != nil {
		return err
	}

	return nil
}

func ValidateString(name string) error {
	if name == "" {
		return errors.New("field cannot be empty")
	}
	for _, char := range name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char == ' ')) {
			return errors.New("name/city can only contain letters")
		}
	}
	return nil
}

func ValidateEmail(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailRegex).MatchString(email) {
		return errors.New("invalid email address")
	}
	return nil
}

func ValidatePhone(phone string) error {
	phoneRegex := `^\+\d{3}\s\d{2}\s\d{3}-\d{2}-\d{2}$`
	if !regexp.MustCompile(phoneRegex).MatchString(phone) {
		return errors.New("invalid phone number format")
	}
	return nil
}

func ValidateBankCard(bankCard string) error {
	if len(bankCard) != 16 {
		return errors.New("bank card number must be 16 digits long")
	}
	return nil
}

func ValidateBalance(balance float32) error {
	if balance < 0 {
		return errors.New("balance cannot be negative")
	}
	return nil
}

func ValidateHome(street string) error {
	if street == "" {
		return errors.New("field cannot be empty")
	}
	hasNumber := false
	for _, char := range street {
		if char >= '0' && char <= '9' {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		return errors.New("street/homenumber  must contain at least one number")
	}
	return nil
}
