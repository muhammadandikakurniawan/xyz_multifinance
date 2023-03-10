package delivery

import (
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/container"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/delivery/http"
)

func Run(dependencyContainer *container.Container) {
	http.Run(dependencyContainer)
}
