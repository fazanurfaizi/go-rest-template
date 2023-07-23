package firebase

import (
	"context"
	"path/filepath"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"google.golang.org/api/option"
)

// NewFirebaseApp creates new firebase app instance
func NewFirebaseApp(logger logger.Logger) *firebase.App {
	ctx := context.Background()

	serviceAccountFilePath, err := filepath.Abs("/config/serviceAccountKey.json")
	if err != nil {
		logger.DPanic("Unable to load serviceAccountKey.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountFilePath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logger.Fatalf("Firebase New App: %v", err)
	}

	logger.Info("Firebase app initialized.")
	return app
}

// NewFirebaseAuth creates new firebase auth client
func NewFirebaseAuth(logger logger.Logger, app *firebase.App) *auth.Client {
	ctx := context.Background()

	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		logger.Fatalf("Firebase Authentication: %v", err)
	}

	return firebaseAuth
}

// NewFirestoreClient creates new firestore client
func NewFirestoreClient(logger logger.Logger, app *firebase.App) *firestore.Client {
	ctx := context.Background()

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		logger.Fatalf("Firestore client: %v", err)
	}

	return firestoreClient
}

// NewFCMClient cteates new firebase cloud messaging client
func NewFCMClient(logger logger.Logger, app *firebase.App) *messaging.Client {
	ctx := context.Background()
	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		logger.Fatalf("Firebase messaging: %v", err)
	}
	return messagingClient
}
