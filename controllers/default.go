package controllers

import (
	"goBee/models"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type ContactModel struct {
	Id      uint64 `orm:"auto"`        // this automatically creates an integer primary key
	Name    string `orm:"size(100)"`   // 100 characters max
	Email   string `orm:"size(255)"`   // 255 characters max
	Message string `form:"type(text)"` // any size string
}

var O orm.Ormer

// Visits counts the number of visits to the Ukk page
var Visits int = 0

func InitDB() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./goBee.db")
	// this function can take a list, e.g. orm.RegisterModel(new(M1), new(M2), ...)
	orm.RegisterModel(new(ContactModel))
	O = orm.NewOrm()

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// matches with {{ .Title }} in index.tpl and layout.tpl
	c.Data["Title"] = "Home"
	c.TplName = "index.tpl"
}

type UkkController struct {
	beego.Controller
}

func (c *UkkController) Get() {
	// matches with {{ .Title }} in index.tpl and layout.tpl
	c.Data["Title"] = "Ukk"
	c.Data["Visits"] = Visits
	c.TplName = "ukk.tpl"
	Visits += 1
}

// ===================================

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	// matches with {{ .Title }} in index.tpl and layout.tpl
	c.Data["Title"] = "About"
	c.TplName = "about.tpl"
}

// boating
type BoatController struct {
	beego.Controller
}

func (c *BoatController) Get() {
	// matches with {{ .Title }} in index.tpl and layout.tpl
	c.Data["Title"] = "Boat"
	c.TplName = "boat.tpl"
}

// Contact
type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	c.Data["Title"] = "Contact"
	c.TplName = "contact.tpl"
}

type Contact struct {
	Name    string `form:"name"`    // matches up with <input name="name" ...>
	Email   string `form:"email"`   // matches up with <input name="email" ...>
	Message string `form:"message"` // matches up with <textarea name="message" ...>
}

// new: post handler
func (c *ContactController) Post() {
	c.Data["Title"] = "Contact"
	c.Data["Result"] = "Thank you for your information! We will remember it : )"

	c.TplName = "contact.tpl"
	// handle form data

	contact := Contact{}
	err := c.Ctx.BindForm(&contact) // Pass a pointer to the struct
	if err != nil {
		c.Data["Result"] = "ERROR: " + err.Error()

	} else if contact.Message == "" || contact.Name == "" || contact.Email == "" {
		c.Data["Result"] = "ERROR: Please enter all values."

	} else {
		// for now we will just log to the console.
		log.Default().Println(contact)
		err = sendEmail(contact)
		if err != nil {
			c.Data["Result"] = "ERROR: Could not send email. " + err.Error()
		} else {
			c.Data["Result"] = "Email sent successfully!"
		}
		contactDb := models.ContactModel{
			Name:    contact.Name,
			Email:   contact.Email,
			Message: contact.Message,
		}
		models.O.Insert(&contactDb)
	}
}

// sendEmail sends an email using SMTP
func sendEmail(contact Contact) error {
	from := "merit.apcs@gmail.com"
	EMAILSENDPASSWORD := os.Getenv("EMAILSENDPASSWORD")                   // Replace with your email
	password := EMAILSENDPASSWORD                                         // Replace with your email password
	toArr := [...]string{contact.Email, "david.wiseman@meritknights.com"} // Replace with the recipient's email(s)
	// slice the array
	to := toArr[:]

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com") // Replace with your SMTP server

	// Compose the email message
	subject := "New Contact Form Submission"
	body := "<b>Name:</b> " + contact.Name + "<br><b>Email</b>: " + contact.Email + "<br><b> Here be the Message you sent us</b>: <u>" + contact.Message + "</u> <br>" +
		"Thank you for contacting us! We have considered deeply your message. <i>We sure hope you made sure its going to be readable for us.</i><br>" +
		"If you have anymore questions move along over to the contact us page and we will get back to you... uhh maybe. We haven't figured out how to do that really."

	if strings.Contains(strings.ToLower(body), "hi there hello") {
		body = "<b>Name:</b> " +
			contact.Name +
			"<br><b>Email</b>: " +
			contact.Email +
			"<br><b> Here be the Message you sent us</b>: <u>" +
			contact.Message +
			"<br>" +
			"</u>PARKER, you are on the wrong form! " +
			"Thank you for contacting us! We have considered deeply your message. <i>We sure hope you made sure its going to be readable for us.</i><br>" +
			"If you have anymore questions move along over to the contact us page and we will get back to you... uhh maybe. We haven't figured out how to do that really."
	}

	message := []byte("To: " + strings.Join(to[:], ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" + // This empty line separates the headers from the body
		body)

	// Send the email
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message) // Replace with your SMTP server and port
	return err
}

type ProfileController struct {
	BaseAdminController
}

func (c *ProfileController) Get() {
	c.RequireAuth()
	c.Data["Title"] = "Profile"
	c.TplName = "profile.tpl"
}

type EditProfile struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func (c *ProfileController) Post() {
	c.RequireAuth()
	c.Data["Title"] = "Profile"
	c.TplName = "profile.tpl"

	ep := EditProfile{}
	err := c.Ctx.BindForm(&ep)
	if err != nil {
		c.Data["Result"] = "ERROR: " + err.Error()
	} else if ep.Email == "" || ep.Name == "" {
		c.Data["Result"] = "ERROR: Please enter all values."
	} else {
		user := c.GetCurrentUser()
		user.Name = ep.Name
		user.Email = ep.Email
		_, err := models.O.Update(user, "Name", "Email")
		if err != nil {
			c.Data["Result"] = "ERROR: " + err.Error()
		} else {
			// update the user in the session
			c.SetSession("user", *user)
			c.Data["Result"] = "Profile updated!"
		}
	}

}
