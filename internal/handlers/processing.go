package handlers

import (
	"fmt"
	"net/http"

	"github.com/mwaurathealex/mbumwa3d/internal/middleware"
	"github.com/mwaurathealex/mbumwa3d/internal/store"
	"github.com/mwaurathealex/mbumwa3d/internal/store/dbstore"
	"github.com/mwaurathealex/mbumwa3d/internal/views/processing"
)

func HandleProcessing(w http.ResponseWriter, r *http.Request) error {
	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if ok == false {
		w.Header().Add("HX-Redirect", "/login")
		return nil
	}

	// get orders for user where order status != complete
	orderStore := dbstore.NewOrderStore()
	processingOrders := orderStore.GetNotCompleted(user.ID)

	for _, order := range processingOrders {
		fmt.Printf("Order ID: %d, PaymentComplete: %t, Price: %.2f, Status: %s\n",
			order.ID, order.PaymentComplete, order.Price, order.Status,
		)
		fmt.Printf("FileName: %s, FileID: %d, LocalPath: %s, Color: %s\n",
			order.File.FileName, order.File.ID, order.File.LocalPath, order.File.Color,
		)
		fmt.Println()
	}

	// processing can be = reviewing, printing, shipping

	// return appropriately, dynamic for htmx, full page otherwise
	if IsHtmx(r) {
		return Render(w, r, processing.Content())
	}
	return Render(w, r, processing.Index())
}
