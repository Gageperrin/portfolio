package watermark

import (
    "context"
    "net/http"
    "os"

    "https://github.com/Gageperrin/portfolio/golang/sample-app"

    "github.com/go-kit/kit/log"
    "github.com/lithammer/shortuuid/v3"
)

type watermarkService struct{}

func NewService() Service { return &watermarkService{} }

func (w *watermarkService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
    // query the database using filters and return a list of documents
    doc := internal.Document{
        Content: "book",
        Title: "Harry Potter and the Half Blood Prince",
        Author: "J.K. Rowling",
        Topic: "Fiction and Magic",
    }
    return []internal.Document{doc}, nil
}

func (w *watermarkService) Status(_ context.Context, ticketID string) (internal.Status, error) {
    // query database using the ticketID and return the document info
    return internal.InProgress, nil
}