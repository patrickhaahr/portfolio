package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"GOTH/views/home"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func HandleContactForm(w http.ResponseWriter, r *http.Request) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	from := mail.NewEmail("Portfolio Contact Form", os.Getenv("EMAIL_FROM"))
	subject := "New Message from Portfolio Contact Form"
	to := mail.NewEmail("", os.Getenv("EMAIL_TO"))

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Server-side email validation
	emailPattern := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !emailPattern.MatchString(email) {
		return Render(w, r, home.Alert(false, "Invalid email address. Please try again."))
	}

	plainTextContent := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", name, email, message)
	htmlContent := fmt.Sprintf("<strong>Name:</strong> %s<br><strong>Email:</strong> %s<br><strong>Message:</strong> %s", name, email, message)

	msg := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(msg)
	if err != nil {
		log.Println(err)
		return Render(w, r, home.Alert(false, "There was an error sending your message. Please try again later."))
	}

	return Render(w, r, home.Alert(true, "Your message has been sent successfully!"))
}
func CloseModal(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := w.Write([]byte(""))
		return err
	}))
}
