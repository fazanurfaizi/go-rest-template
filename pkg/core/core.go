package core

import (
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/firebase"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(router.NewRouter),
	fx.Provide(postgres.NewConnection),
	fx.Provide(firebase.NewFirebaseApp),
	fx.Provide(firebase.NewFirebaseAuth),
	fx.Provide(firebase.NewFirestoreClient),
	fx.Provide(firebase.NewFCMClient),
)
