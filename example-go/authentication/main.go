package main

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func main() {
	// Inicialize o SDK Firebase Admin com as credenciais do seu projeto
	opt := option.WithCredentialsFile("path/to/your/firebase/credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar o SDK Firebase: %v", err)
	}

	// Inicialize o cliente de autenticação Firebase
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Erro ao inicializar o cliente de autenticação Firebase: %v", err)
	}

	// Defina os atributos do novo usuário
	params := (&auth.UserToCreate{}).
		Email("example@example.com").
		Password("password").
		DisplayName("Nome do Usuário").
		PhotoURL("https://example.com/profile.jpg")

	// Crie o novo usuário
	user, err := client.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("Erro ao criar o usuário: %v", err)
	}

	fmt.Printf("Usuário criado com sucesso: %+v\n", user)
}
