package meural

import (
	"bytes"
	"encoding/json"
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

func GetToken(username string, password string) (string, error) {
	var authenticate model.Authenticate
	log.Println("GetToken: " + username)

	url := "https://api.meural.com/v0/authenticate"
	method := "POST"

	payload := strings.NewReader("username=" + username + "&password=" + password)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	json.Unmarshal([]byte(string(body)), &authenticate)

	token = authenticate.Token
	return authenticate.Token, nil
}

func GetUserDevices() (model.UserDevices, error) {
	var userDevices model.UserDevices
	log.Println("GetUserDevices")

	url := "https://api.meural.com/v0/user/devices?count=10&page=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return userDevices, err
	}
	req.Header.Add("Authorization", "Token "+token)

	res, err := client.Do(req)
	if err != nil {
		return userDevices, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return userDevices, err
	}

	json.Unmarshal([]byte(string(body)), &userDevices)

	return userDevices, nil
}

func GetUserGalleries() (model.UserGalleries, error) {
	var userGalleries model.UserGalleries
	log.Println("GetUserGalleries")

	url := "https://api.meural.com/v0/user/galleries?count=10&page=1"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return userGalleries, err
	}
	req.Header.Add("Authorization", "Token "+token)

	res, err := client.Do(req)
	if err != nil {
		return userGalleries, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return userGalleries, err
	}

	json.Unmarshal([]byte(string(body)), &userGalleries)

	return userGalleries, nil
}

func GetUserGalleryId(galleryName string) (int, error) {
	log.Println("GetUserGalleryId: " + galleryName)

	userGalleries, err := GetUserGalleries()
	if err != nil {
		return -1, err
	}

	for _, gallery := range userGalleries.Data {
		if strings.Contains(gallery.Name, galleryName) {
			return gallery.ID, nil
		}
	}

	return -1, nil
}

func PostGallery(name string, description string, orientation string) (model.Gallery, error) {
	var gallery model.Gallery
	log.Println("PostGallery: " + name)

	url := "https://api.meural.com/v0/galleries"
	method := "POST"

	payload := strings.NewReader("name=" + name + "&description=" + description + "&orientation=" + orientation)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return gallery, err
	}
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return gallery, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return gallery, err
	}

	json.Unmarshal([]byte(string(body)), &gallery)

	return gallery, nil
}

func PostItem(filename string) (model.Item, error) {
	var item model.Item
	log.Println("PostItem: " + filename)

	url := "https://api.meural.com/v0/items"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open("./images/" + filename)
	if err != nil {
		return item, err
	}
	defer file.Close()
	part, err := writer.CreateFormFile("image", "./images/"+filename)
	if err != nil {
		return item, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return item, err
	}
	err = writer.Close()
	if err != nil {
		return item, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return item, err
	}
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Add("Content-Type", "multipart/form-data")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return item, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return item, err
	}

	json.Unmarshal([]byte(string(body)), &item)

	return item, nil
}

func PutItem(id int, name string, author string, description string, medium string, year int) (model.Item, error) {
	var item model.Item
	log.Println("PutItem: " + strconv.Itoa(id))

	url := "https://api.meural.com/v0/items/" + strconv.Itoa(id)
	method := "PUT"

	payload := strings.NewReader("name=" + name + "&author=" + author + "&description=" + description + "&medium=" + medium + "&year=" + strconv.Itoa(year))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return item, err
	}
	req.Header.Add("Authorization", "Token "+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return item, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return item, err
	}

	json.Unmarshal([]byte(string(body)), &item)

	return item, nil
}

func PostItemToGallery(itemId int, galleryId int) (model.Gallery, error) {
	var gallery model.Gallery
	log.Println("PostItemToGallery: " + strconv.Itoa(itemId) + " -> " + strconv.Itoa(galleryId))

	url := "https://api.meural.com/v0/galleries/" + strconv.Itoa(galleryId) + "/items/" + strconv.Itoa(itemId)
	method := "POST"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return gallery, err
	}
	req.Header.Add("Authorization", "Token "+token)

	res, err := client.Do(req)
	if err != nil {
		return gallery, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return gallery, err
	}

	json.Unmarshal([]byte(string(body)), &gallery)

	return gallery, nil
}
