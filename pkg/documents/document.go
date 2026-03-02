package documents

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

type DocumentLoader struct {
	uploadDir string
}

func NewDocumentLoader() *DocumentLoader {
	return &DocumentLoader{
		uploadDir: "./uploads",
	}
}

func (d *DocumentLoader) LoadDocument(ctx context.Context, filePath string) ([]schema.Document, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, _ := os.Stat(filePath)

	ext := strings.ToLower(filepath.Ext(filePath))
	var (
		loader   documentloaders.Loader
		splitter textsplitter.TextSplitter
	)

	switch ext {
	case ".pdf":
		loader = documentloaders.NewPDF(file, fileInfo.Size())
		splitter = textsplitter.NewRecursiveCharacter(
			textsplitter.WithChunkSize(1000),
			textsplitter.WithChunkOverlap(200),
		)
	case ".html":
		loader = documentloaders.NewHTML(file)
		splitter = textsplitter.NewRecursiveCharacter(
			textsplitter.WithChunkOverlap(500),
			textsplitter.WithChunkSize(2000),
		)
	default:
		loader = documentloaders.NewText(file)
		splitter = textsplitter.NewRecursiveCharacter(
			textsplitter.WithChunkSize(1000),
			textsplitter.WithChunkOverlap(200),
		)
	}

	docs, err := loader.LoadAndSplit(ctx, splitter)
	if err != nil {
		return nil, err
	}
	return docs, nil
}
