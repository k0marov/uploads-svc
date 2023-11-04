package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/samkomarov/uploads-svc.git/docs"
	"net/http"
)

type IUploadService interface {
	// GetNewFilename uses uploadedFilename to preserve extension if possible
	GetNewFilename(uploadedFilename string) string
	GetFullFSPath(filename string) string
	GetURL(filename string) string
	FSRoot() string
	MaxFileSizeBytes() int64
	MakeErrTooBigFile() error
}

type Server struct {
	r   *gin.Engine
	svc IUploadService
}

func NewServer(svc IUploadService) http.Handler {
	srv := &Server{gin.New(), svc}
	srv.defineEndpoints()
	return srv
}

//	@title			uploads-svc
//	@version		1.0
//	@description	An API for uploading and serving files.

//	@contact.name	Sam Komarov
//	@contact.url	github.com/k0marov
//	@contact.email	sam@skomarov.com

// @host		localhost:8080
// @schemes     https http
func (s *Server) defineEndpoints() {
	s.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	group := s.r.Group("/api/v1/uploads")
	group.POST("/", s.HandleUpload)
	group.Static("/", s.svc.FSRoot())
}

// DownloadFile godoc
//
//		@Summary		Serves a file by its name
//		@Description    Serves a file by its name.
//		@Description	Generally, you wouldn't construct requests to this endpoint yourself,
//	    @Description	because full URLs are returned from the upload endpoint.
//		@Tags			uploads
//		@Param	        name path string true "filename to download"
//		@Success		200			file 	file	"full file contents"
//		@Failure 		404
//		@Router			/api/v1/uploads/{name} [get]
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	// just a stub for swaggo generator, see defineEndpoints()
}

type UploadedResponse struct {
	URL string `json:"url" example:"localhost:8080/api/v1/uploads/afc5a4eb-8dd1-4df0-a3c4-6c2703a3dcb7.png"`
}

// HandleUpload godoc
//
//	@Summary		Upload a file
//	@Description	Upload a file by including it in a multipart request's "file" field.
//	@Description	Before saving on the server, a random name is generated, but file extension is preserved.
//	@Description	Returns full URL for accessing the uploaded file.
//	@Tags			uploads
//	@Produce		plain
//	@Param	        file formData file true "File that you want to upload. Size limits are defined in config"
//	@Success		201			{object} 	UploadedResponse	"URL from which you can download the uploaded file."
//	@Failure 		413 		{object} 	ClientError "When uploaded image exceeds max size limit."
//	@Router			/api/v1/uploads/ [post]
func (s *Server) HandleUpload(c *gin.Context) {
	// limits request size
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, s.svc.MaxFileSizeBytes())

	if c.ContentType() != "multipart/form-data" {
		WriteErrorResponse(c.Writer, ErrInvalidContentType)
		return
	}
	gotFile, err := c.FormFile("file")
	if err != nil {
		if _, ok := err.(*http.MaxBytesError); ok {
			err = s.svc.MakeErrTooBigFile()
		} else if errors.Is(err, http.ErrMissingFile) {
			err = ErrNoFileProvided
		}
		WriteErrorResponse(c.Writer, err)
		return
	}
	newFilename := s.svc.GetNewFilename(gotFile.Filename)
	err = c.SaveUploadedFile(gotFile, s.svc.GetFullFSPath(newFilename))
	if err != nil {
		WriteErrorResponse(c.Writer, err)
		return
	}
	c.JSON(http.StatusCreated, UploadedResponse{
		URL: s.svc.GetURL(newFilename),
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
