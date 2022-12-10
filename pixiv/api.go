package pixiv

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"xyzyxJP/meural-link/pixiv/model"
)

func GetUserDetails(userId string) (model.UserDetails, error) {
	log.Println("GetUserDetails: " + userId)

	url := "https://www.pixiv.net/touch/ajax/user/details?id=" + userId + "&lang=ja"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return model.UserDetails{}, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	res, err := client.Do(req)
	if err != nil {
		return model.UserDetails{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.UserDetails{}, err
	}

	var userDetails model.UserDetails
	json.Unmarshal([]byte(string(body)), &userDetails)

	return userDetails, nil
}

func GetUserProfile(userId string) (model.UserProfile, error) {
	log.Println("GetUserProfile: " + userId)

	url := "https://www.pixiv.net/ajax/user/" + userId + "/profile/all?lang=ja"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return model.UserProfile{}, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return model.UserProfile{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.UserProfile{}, err
	}

	var userProfile model.UserProfile
	json.Unmarshal([]byte(string(body)), &userProfile)

	return userProfile, nil
}

func GetIllust(illustId string) (model.Illust, error) {
	log.Println("GetIllust: " + illustId)

	url := "https://www.pixiv.net/ajax/illust/" + illustId + "?lang=ja"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return model.Illust{}, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return model.Illust{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Illust{}, err
	}

	var illust model.Illust
	json.Unmarshal([]byte(string(body)), &illust)

	return illust, nil
}

func GetIllustPages(illustId string) (model.IllustPages, error) {
	log.Println("GetIllustPages: " + illustId)

	url := "https://www.pixiv.net/ajax/illust/" + illustId + "/pages?lang=ja"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return model.IllustPages{}, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return model.IllustPages{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.IllustPages{}, err
	}

	var illustPages model.IllustPages
	json.Unmarshal([]byte(string(body)), &illustPages)

	return illustPages, nil
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

	return nil
}
