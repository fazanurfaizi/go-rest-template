package media

var mediaHandlers = make(map[string]MediaHandler)

type MediaHandler interface {
	CouldHandle(media Media) bool
	Handle(media Media, file FileInterface, option *Option) error
}
