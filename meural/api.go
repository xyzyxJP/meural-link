package meural

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"xyzyxJP/meural-link/meural/model"
)

var token string

func GetToken(username string, password string) (model.Authenticate, error) {
	log.Println("GetToken: " + username)

	url := "https://api.meural.com/v0/authenticate"
	method := "POST"

	payload := strings.NewReader("username=" + username + "&password=" + password)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return model.Authenticate{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return model.Authenticate{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Authenticate{}, err
	}

	var authenticate model.Authenticate
	json.Unmarshal([]byte(string(body)), &authenticate)

	token = authenticate.Token
	return authenticate, nil
}

func GetUserDevices() (model.UserDevices, error) {
	log.Println("GetUserDevices")

	url := "https://api.meural.com/v0/user/devices?count=10&page=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return model.UserDevices{}, err
	}
	req.Header.Add("Authorization", "Token "+token)

	res, err := client.Do(req)
	if err != nil {
		return model.UserDevices{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.UserDevices{}, err
	}

	var userDevices model.UserDevices
	json.Unmarshal([]byte(string(body)), &userDevices)

	return userDevices, nil
}

func GetUserGalleries() (model.UserGalleries, error) {
	log.Println("GetUserGalleries")

	url := "https://api.meural.com/v0/user/galleries?count=10&page=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return model.UserGalleries{}, err
	}
	req.Header.Add("Authorization", "Token "+token)

	res, err := client.Do(req)
	if err != nil {
		return model.UserGalleries{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.UserGalleries{}, err
	}

	var userGalleries model.UserGalleries
	json.Unmarshal([]byte(string(body)), &userGalleries)

	return userGalleries, nil
}

func GetUserGalleryId(name string) (model.Gallery, error) {
	log.Println("GetUserGalleryId: " + name)

	userGalleries, err := GetUserGalleries()
	if err != nil {
		return model.Gallery{}, err
	}

	for _, gallery := range userGalleries.Data {
		if strings.Contains(gallery.Name, name) {
			gallery, err := GetGallery(gallery.ID)
			if err != nil {
				return model.Gallery{}, err
			}
			return gallery, nil
		}
	}

	return model.Gallery{}, fmt.Errorf("error: %s", "gallery not found")
}

func GetGallery(galleryId int) (model.Gallery, error) {
	log.Println("GetGallery: " + strconv.Itoa(galleryId))

	url := "https://api.meural.com/v0/galleries/" + strconv.Itoa(galleryId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return model.Gallery{}, err
	}
	req.Header.Add("Authorization", "Token "+token)

	res, err := client.Do(req)
	if err != nil {
		return model.Gallery{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Gallery{}, err
	}

	var gallery model.Gallery
	json.Unmarshal([]byte(string(body)), &gallery)

	return gallery, nil
}

func GetGalleryItems(galleryId int) (model.GalleryItems, error) {
	log.Println("GetGalleryItems: " + strconv.Itoa(galleryId))

	url := "https://api.meural.com/v0/galleries/" + strconv.Itoa(galleryId) + "/items"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return model.GalleryItems{}, err
	}
	req.Header.Add("Authorization", "Token "+token)

	res, err := client.Do(req)
	if err != nil {
		return model.GalleryItems{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.GalleryItems{}, err
	}

	var galleryItems model.GalleryItems
	json.Unmarshal([]byte(string(body)), &galleryItems)

	return galleryItems, nil
}

func PostGallery(name string, description string, orientation string) (model.Gallery, error) {
	log.Println("PostGallery: " + name)

	url := "https://api.meural.com/v0/galleries"
	method := "POST"

	payload := strings.NewReader("name=" + name + "&description=" + description + "&orientation=" + orientation)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return model.Gallery{}, err
	}
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return model.Gallery{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Gallery{}, err
	}

	var gallery model.Gallery
	json.Unmarshal([]byte(string(body)), &gallery)

	return gallery, nil
}

func PostItem(filename string) (model.Item, error) {
	log.Println("PostItem: " + filename)

	url := "https://api.meural.com/v0/items"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open("./images/" + filename)
	if err != nil {
		return model.Item{}, err
	}
	defer file.Close()
	part, err := writer.CreateFormFile("image", "./images/"+filename)
	if err != nil {
		return model.Item{}, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return model.Item{}, err
	}
	err = writer.Close()
	if err != nil {
		return model.Item{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return model.Item{}, err
	}
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Add("Content-Type", "multipart/form-data")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return model.Item{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Item{}, err
	}

	var item model.Item
	json.Unmarshal([]byte(string(body)), &item)

	return item, nil
}

func PutItem(id int, name string, author string, description string, medium string, year int) (model.Item, error) {
	log.Println("PutItem: " + strconv.Itoa(id))

	url := "https://api.meural.com/v0/items/" + strconv.Itoa(id)
	method := "PUT"

	payload := strings.NewReader("name=" + name + "&author=" + author + "&description=" + description + "&medium=" + medium + "&year=" + strconv.Itoa(year))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return model.Item{}, err
	}
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return model.Item{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Item{}, err
	}

	var item model.Item
	json.Unmarshal([]byte(string(body)), &item)

	return item, nil
}

func PostItemToGallery(itemId int, galleryId int) (model.Gallery, error) {
	log.Println("PostItemToGallery: " + strconv.Itoa(itemId) + " -> " + strconv.Itoa(galleryId))

	url := "https://api.meural.com/v0/galleries/" + strconv.Itoa(galleryId) + "/items/" + strconv.Itoa(itemId)
	method := "POST"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return model.Gallery{}, err
	}
	req.Header.Add("Authorization", "Token "+token)

	res, err := client.Do(req)
	if err != nil {
		return model.Gallery{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Gallery{}, err
	}

	var gallery model.Gallery
	json.Unmarshal([]byte(string(body)), &gallery)

	return gallery, nil
}
