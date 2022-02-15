package composites

import (
	api "dev-hack-backend/internal/adapters/api"
	user2 "dev-hack-backend/internal/adapters/api/user"
	user3 "dev-hack-backend/internal/adapters/db/mongodb/user"
	user4 "dev-hack-backend/internal/service/user"
	"dev-hack-backend/pkg/manager/jwt"
	"time"
)

type UserComposite struct {
	Storage user4.Storage
	Service user4.Service
	Handler api.Handler
}

func NewUserComposite(composite *MongoDBComposite, jwtKey, userCollection string, accessTTL, refreshTTL time.Duration) *UserComposite {
	jwtManager := jwt.NewManager(jwtKey)
	storage := user3.NewStorage(composite.db, userCollection)
	service := user4.NewService(storage, jwtManager, accessTTL, refreshTTL)
	handler := user2.NewHandler(service)

	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}
}
