package userService

import (
	"bytes"
	"errors"
	"github.com/disintegration/imaging"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
	"image"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func UpdateImage(i *tracer.Infos, email string, file multipart.File) (string, error) {
	i.TraceIt("updating photo")
	defer i.Span.Finish()

	now := time.Now()      // current local time
	name := strconv.Itoa(int(now.Unix())) + ".png"

	photoFile, err := os.Create("./user_photos/" + name)
	if err != nil {
		i.LogError(err)
		return "", err
	}
	defer photoFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		i.LogError(err)
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		i.LogError(err)
		return "", err
	}

	resized := imaging.Resize(img, 200, 200, imaging.Lanczos)

	enc := png.Encoder{
		CompressionLevel: png.BestCompression,
	}

	err = enc.Encode(photoFile, resized)
	if err != nil{
		i.LogError(err)
		return "", err
	}

	uRepo := repository.NewUser()
	err = uRepo.UpdatePhoto(i, email, name)
	if err != nil{
		i.LogError(err)
		return "", err
	}

	return name, nil
}

func ErasePhoto(i *tracer.Infos, email string) error {
	uRepo := repository.NewUser()
	resp, err := uRepo.GetUserByEmail(i, email)
	if err != nil{
		i.LogError(err)
		return err
	}
	filename, ok := resp["photo_path"].(string)

	if !ok || len(filename) == 0 {
		err := errors.New("could not gat user photo")
		i.LogError(err)
		return err
	}

	err = uRepo.UpdatePhoto(i, email, "")
	if err != nil{
		i.LogError(err)
		return err
	}

	//The file could not be removed but the register was updated, must remove manually
	err = os.Remove("./user_photos/" + filename)
	if err != nil{
		i.LogError(err)
	}

	return nil
}