package email

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

func SendGomail(name, email, id string) error {
	body := fmt.Sprintf(`
        <html>
            <body>
                <p>Hello %v,</p>
                <p>I Hope you are doing Well</p>
                <p>You are successfully registered to our Website</p>
                <p>This is your ID: %v</p>
                <p>PLEASE DO NOT SHARE WITH ANYONE</p>
                <p>Thanks and have a nice day</p>
            </body>
        </html>
    `, name, id)

	m := gomail.NewMessage()
	m.SetHeader("From", "juraboevizzatillo5@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Registration status")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "juraboevizzatillo5@gmail.com", "tnkx qlnz anuk exdn")

	if err := d.DialAndSend(m); err != nil {
		log.Println("failed to send an email:", err)
		return err
	}

	return nil
}
