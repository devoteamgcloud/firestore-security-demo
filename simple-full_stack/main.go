package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	firestoreUrl = ""
	projectId    = ""
	apiKey       = ""
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Articles struct {
	Articles []Articles `json:"articles"`
}

type StringValue struct {
	Value string `json:"stringValue"` // https://cloud.google.com/firestore/docs/reference/rest/Shared.Types/ArrayValue#Value
}

type FirestoreArticle struct {
	Title   StringValue `json:"title"`
	Content StringValue `json:"content"`
}

type Document struct {
	Fields FirestoreArticle `json:"fields"`
}

type Documents struct {
	List []Document `json:"documents"`
}

type ArticlesHandler struct{}

func (*ArticlesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	switch r.Method {
	case http.MethodGet:
		(&GetArticlesController{response: w, request: r}).Handle()
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type GetArticlesController struct {
	response http.ResponseWriter
	request  *http.Request
}

func (controller *GetArticlesController) Handle() {
	token := controller.token()
	articles := controller.getArticles(token)
	controller.returnArticles(articles)
}

func (controller *GetArticlesController) returnArticles(articles []Article) {
	body, err := json.Marshal(articles)
	panicIfError(err)
	controller.response.WriteHeader(http.StatusOK)
	fmt.Fprint(controller.response, body)
}

func (controller *GetArticlesController) token() string {
	raw := controller.request.Header.Get("Authorization")
	token := raw[7:] // Cut the 'Bearer '
	return token
}

func (controller *GetArticlesController) getArticles(token string) []Article {
	response := controller.getArticlesFromFirestore(token)

	var documents Documents
	panicIfError(json.Unmarshal(response, &documents))
	articles := make([]Article, len(documents.List))
	for i, doc := range documents.List {
		articles[i] = Article{
			Title:   doc.Fields.Title.Value,
			Content: doc.Fields.Content.Value,
		}
	}
	return articles
}

func (controller *GetArticlesController) getArticlesFromFirestore(token string) []byte {
	request := controller.makeGetArticlesFromFirestoreRequest(token)
	response := controller.sendRequest(request)
	body, err := io.ReadAll(response.Body)
	panicIfError(err)
	return body
}

func (controller *GetArticlesController) sendRequest(request *http.Request) *http.Response {
	response, err := http.DefaultClient.Do(request)
	panicIfError(err)
	if response.StatusCode != 200 {
		panic(fmt.Errorf("oops...firestore returned %d", response.StatusCode))
	}
	return response
}

func (controller *GetArticlesController) makeGetArticlesFromFirestoreRequest(token string) *http.Request {
	parent := fmt.Sprintf("projects/%s/databases/(default)/documents/articles", projectId)
	url := fmt.Sprintf("%s%s?%s", firestoreUrl, parent, apiKey)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	panicIfError(err)
	request.Header.Add("Authorization", "Bearer "+token)
	return request
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/articles", &ArticlesHandler{})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
