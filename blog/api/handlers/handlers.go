package handlers

type handler struct {
}

type handlerDto struct {
}

func NewHandler(h *handlerDto) *handler {
    return &handler{}
}

func (h *handler) Ping(c echo.Context) error {
    
}