package bootstrap

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/option"
)

type Application struct {
	Env   *Env
	Mongo *mongo.Client
	GenAi *genai.GenerativeModel
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)

	client,err:=genai.NewClient(context.TODO(),option.WithAPIKey(app.Env.GeminiApikey))
	if err!=nil{
		log.Fatal(err)
	}
	model:=client.GenerativeModel("gemini-1.5-flash")
	app.GenAi = model
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}

func (app *Application) CloseModelClient() {
	// app.GenAi.Close()
}