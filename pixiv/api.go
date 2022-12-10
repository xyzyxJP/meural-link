package pixiv

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"xyzyxJP/meural-link/pixiv/model"
)

func GetUserProfile(userId string) (model.UserProfile, error) {
	var userProfile model.UserProfile
	log.Println("GetUserProfile: " + userId)

	url := "https://www.pixiv.net/ajax/user/" + userId + "/profile/all?lang=ja"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return userProfile, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return userProfile, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return userProfile, err
	}

	json.Unmarshal([]byte(string(body)), &userProfile)

	return userProfile, nil
}

func GetIllust(illustId string) (model.Illust, error) {
	var illust model.Illust
	log.Println("GetIllust: " + illustId)

	url := "https://www.pixiv.net/ajax/illust/" + illustId + "?lang=ja"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return illust, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return illust, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return illust, err
	}

	json.Unmarshal([]byte(string(body)), &illust)

	return illust, nil
}

func GetIllustPages(illustId string) (model.IllustPages, error) {
	var illustPages model.IllustPages
	log.Println("GetIllustPages: " + illustId)

	url := "https://www.pixiv.net/ajax/illust/" + illustId + "/pages?lang=ja"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return illustPages, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return illustPages, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return illustPages, err
	}

	json.Unmarshal([]byte(string(body)), &illustPages)

	return illustPages, nil
}

func GetIllustFilename(illustPages model.IllustPages, illust model.Illust, index int) string {
	return illust.Body.UserName + " - " + illust.Body.Title + " (" + illust.Body.IllustID + ") p" + fmt.Sprint(index+1) + filepath.Ext(illustPages.Body[index].Urls.Original)
}

func GetIllustImage(url string, filename string) error {
	log.Println("GetIllustImage: " + url + " -> " + filename)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	req.Header.Add("Referer", "https://i.pximg.net/")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = os.Stat("./images")
	if err != nil {
		os.Mkdir("./images", 0777)
	}

	file, err := os.Create("./images/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	// exif := map[string]interface{}{}
	// err = update.PrepareAndUpdateExif(file, os.Stdout, exif)
	// if err != nil {
	// 	return err
	// }

	return nil
}
