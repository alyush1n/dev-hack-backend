package composites

import (
	"dev-hack-backend/internal/adapters/api"
	user2 "dev-hack-backend/internal/adapters/api/user"
	user3 "dev-hack-backend/internal/adapters/db/mongodb/user"
	"dev-hack-backend/internal/domain/user"
	"dev-hack-backend/pkg/manager/jwt"
	"time"
)

type UserComposite struct {
	Storage user.Storage
	Service user2.Service
	Handler api.Handler
}

func NewUserComposite(composite *MongoDBComposite, jwtKey, userCollection string, accessTTL, refreshTTL time.Duration) *UserComposite {
	jwtManager := jwt.NewManager(jwtKey)
	storage := user3.NewStorage(composite.db, userCollection)
	service := user.NewService(storage, jwtManager, accessTTL, refreshTTL)
	handler := user2.NewHandler(service)

	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}
}
