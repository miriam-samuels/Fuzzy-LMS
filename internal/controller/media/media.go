package media

import (
	"context"
	"io"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/storage"
)

func UploadMedia(w http.ResponseWriter, r *http.Request) {
	file, _ := helper.ParseMultipartRequestBody(w, r)

	defer file.Close()

	// GENERATE unique id for media
	mediaName := "MEDIA" + helper.GenerateUniqueId(4)

	// open new file in bucket
	media := storage.LoanBucket.Object(mediaName)

	wc := media.NewWriter(context.Background())

	_, err := io.Copy(wc, file)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "Unable to upload file to storage", nil)
		return
	}

	// close writer
	if err := wc.Close(); err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "Unable to close storage writer"+err.Error(), nil)
		return
	}

	// Generate a download URL for the uploaded file
	downloadURL, err := media.Attrs(context.Background())
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "Unable to get download URL"+err.Error(), nil)
		return
	}

	// form ok response
	res := map[string]interface{}{
		"url": downloadURL.MediaLink,
	}
	helper.SendResponse(w, http.StatusOK, true, "Media Uploaded Successfully", res)

}
