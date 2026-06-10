package booking
import "net/http"
import "github.com/smeet07/cinemabookingGO/internal/utils"

type handler struct{
	svc *Service		
}
func NewHandler(svc *Service) *handler{
	return &handler{svc};
}
func (h *handler) ListSeats(w http.ResponseWriter,r *http.Request){
	movieID:=r.PathValue("movieID")
	bookings:=h.svc.ListBookings(movieID)
	seats:=make([]seatInfo,0,len(bookings))
	for _,b:=range bookings{
		seats=append(seats,seatInfo{
			seatID:b.SeatID,
			UserID:b.UserID,
			Booked:true,
		})
	}
	utils.WriteJSON(w,http.StatusOK,seats)
	
}
type seatInfo struct{
	seatID string `json:"seat_id"`
	UserID string `json:"user_id"`
	Booked bool `json:"booked"`
}