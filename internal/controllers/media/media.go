package media

import (
	"context"
	"io"
	"net/http"

	config "github.com/miriam-samuels/loan-management-backend/internal/config/storage"
	"github.com/miriam-samuels/loan-management-backend/internal/helper"
)

func UploadMedia(w http.ResponseWriter, r *http.Request) {
	file, _ := helper.ParseMultipartRequestBody(w, r)

	defer file.Close()

	// GENERATE unique id for media
	mediaName := helper.GenerateUUID()

	// open new file in bucket
	media := config.LoanBucket.Object(mediaName.String())

	wc := media.NewWriter(context.Background())

	_, err := io.Copy(wc, file)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "Unable to upload file to storage", nil)
		return
	}

	// close writer
	if err := wc.Close(); err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "Unable to close storage writer"+err.Error(), nil)
		return
	}

	// Generate a download URL for the uploaded file
	downloadURL, err := media.Attrs(context.Background())
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "Unable to get download URL"+err.Error(), nil)
		return
	}

	// form ok response
	res := map[string]interface{}{
		"url": downloadURL,
	}
	helper.SendJSONResponse(w, http.StatusOK, true, "Media Uploaded Successfully", res)

}
