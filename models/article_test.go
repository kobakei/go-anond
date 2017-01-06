package models

import "testing"

func TestSummary(t *testing.T) {
  article := &Article{Title: "Title", Body: "12345678901234567890"}
  if article.Summary() != "1234567890" {
    t.Error("Wrong summary: ", article.Summary())
  }
}
