package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/container"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/delivery/http/handler/consumer"
	"github.com/rs/cors"
	"github.com/spf13/cast"
)

func Run(dependencyContainer *container.Container) (err error) {
	router := mux.NewRouter()
	// appChan := make(chan os.Signal, 1)
	// signal.Notify(appChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	consumerHandler := consumer.NewConsumerHandler(dependencyContainer.ConsumerModuleContainer.ConsumerUsecase)

	consumerHandlerRoute := router.PathPrefix("/consumer").Subrouter()
	consumerHandlerRoute.HandleFunc("/register", consumerHandler.Register).Methods(http.MethodPost)
	consumerHandlerRoute.HandleFunc("/request-loan", consumerHandler.RequestLoan).Methods(http.MethodPost)
	consumerHandlerRoute.HandleFunc("/tenor-limit", consumerHandler.AddTenorLimit).Methods(http.MethodPost)

	cors := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := cors.Handler(router)

	port := cast.ToString(dependencyContainer.AppConfig.HttpPort)
	log.Printf("The app is running on port %s.", port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
	log.Println("SHUTING DOWN HTTP SERVER")
	// <-appChan
	log.Println("SHUTDOWN HTTP SERVER")
	return
}
