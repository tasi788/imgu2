package controllers

import (
	"encoding/json"
	"imgu2/services"
	"imgu2/templates"
	"imgu2/utils"
	"io"
	"log/slog"
	"net/http"
)

type H map[string]any

func render(w io.Writer, name string, data H) {
	title, err := services.Setting.GetSiteName()
	if err != nil {
		slog.Error("render template: site name", "err", err)
		return
	}
	data["title"] = title

	err = templates.Render(w, name, data)
	if err != nil {
		slog.Error("render template", "err", err, "name", name, "data", data)
		return
	}
}

func renderDialog(w io.Writer, title, msg, link, btn string) {
	render(w, "dialog", H{
		"dialog": title,
		"msg":    msg,
		"link":   link,
		"btn":    btn,
	})
}

func writeJSON(w io.Writer, m H) {
	b, err := json.Marshal(m)
	if err != nil {
		slog.Error("write json", "err", err)
		return
	}

	_, err = io.WriteString(w, string(b))
	if err != nil {
		slog.Error("write json", "err", err)
	}
}

func setCookie(w http.ResponseWriter, name string, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})
}

// generate a new csrf token and add it to cookies
func csrfToken(w http.ResponseWriter) string {
	t := utils.RandomHexString(8)
	setCookie(w, "CSRF_TOKEN", t)
	return t
}
