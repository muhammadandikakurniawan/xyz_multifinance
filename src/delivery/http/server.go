package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/muhammadandikakurniawan/xyz_multifinance/cmd/app/docs"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/container"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/delivery/http/handler/consumer"
	"github.com/rs/cors"
	"github.com/spf13/cast"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Run(dependencyContainer *container.Container) (err error) {
	router := mux.NewRouter()
	appChan := make(chan os.Signal, 1)
	signal.Notify(appChan, os.Interrupt, syscall.SIGTERM)

	port := cast.ToString(dependencyContainer.AppConfig.HttpPort)
	docs.SwaggerInfo.Host = ":" + port
	router.PathPrefix("/api-docs").Handler(httpSwagger.WrapHandler)

	consumerHandler := consumer.NewConsumerHandler(dependencyContainer.ConsumerModuleContainer.ConsumerUsecase)

	consumerHandlerRoute := router.PathPrefix("/consumer").Subrouter()
	consumerHandlerRoute.HandleFunc("/register", consumerHandler.Register).Methods(http.MethodPost)
	consumerHandlerRoute.HandleFunc("/request-loan", consumerHandler.RequestLoan).Methods(http.MethodPost)
	consumerHandlerRoute.HandleFunc("/tenor-limit", consumerHandler.AddTenorLimit).Methods(http.MethodPost)
	consumerHandlerRoute.HandleFunc("/search-request-loan", consumerHandler.SearchRequestLoan).Methods(http.MethodGet)
	consumerHandlerRoute.HandleFunc("/approve-request-loan", consumerHandler.ApproveRequestLoan).Methods(http.MethodPut)
	consumerHandlerRoute.HandleFunc("/{id}", consumerHandler.GetDetailConsumer).Methods(http.MethodGet)

	cors := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := cors.Handler(router)

	go func() {
		log.Printf("HTTP running on port %s.", port)
		log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
	}()

	<-appChan
	log.Println("HTTP SERVER CLOSED")
	return
}
