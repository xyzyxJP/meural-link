package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"xyzyxJP/meural-link/meural"
	"xyzyxJP/meural-link/pixiv"
)

func keys(m map[string]struct{}) []string {
	ks := []string{}
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func PostRandomPixivToMeural() error {
	rand.Seed(time.Now().UnixNano())

	userDetails, err := pixiv.GetUserDetails(os.Getenv("PIXIV_USER_ID"))
	if err != nil {
		return err
	}

	userId := userDetails.Body.UserDetails.UserID
	userName := userDetails.Body.UserDetails.UserName

	userProfile, err := pixiv.GetUserProfile(userId)
	if err != nil {
		return err
	}

	gallery, err := meural.GetUserGalleryId(userId)
	if err != nil {
		gallery, err = meural.PostGallery(userId+" - "+userName, userDetails.Body.UserDetails.UserComment, "horizontal")
		if err != nil {
			return err
		}
	}

	galleryItems, err := meural.GetGalleryItems(gallery.Data.ID)
	if err != nil {
		return err
	}

	illustId := ""

	for i := 0; i < len(userProfile.Body.Illusts); i++ {
		illustId = keys(userProfile.Body.Illusts)[i]
		log.Println("illustId: " + illustId)

		flag := true
		for _, item := range galleryItems.Data {
			if strings.Contains(item.Name, illustId) {
				flag = false
				break
			}
		}

		if flag {
			break
		}
	}

	if illustId == "" {
		return fmt.Errorf("error: %s", "no new illusts")
	}

	illust, err := pixiv.GetIllust(illustId)
	if err != nil {
		return err
	}

	illustPages, err := pixiv.GetIllustPages(illustId)
	if err != nil {
		return err
	}

	illustTitle := illust.Body.Title

	for i := range illustPages.Body {
		filename := userName + " - " + illustTitle + " (" + illustId + ") p" + fmt.Sprint(i+1) + filepath.Ext(illustPages.Body[i].Urls.Original)
		err = pixiv.GetIllustImage(illustPages.Body[i].Urls.Original, filename)
		if err != nil {
			log.Fatal(err)
			continue
		}

		item, err := meural.PostItem(filename)
		if err != nil {
			log.Fatal(err)
			continue
		}

		item, err = meural.PutItem(item.Data.ID, illustTitle+" ("+illustId+")", userName, illust.Body.Description, "https://www.pixiv.net/artworks/"+illustId, illust.Body.CreateDate.Year())
		if err != nil {
			log.Fatal(err)
			continue
		}

		_, err = meural.PostItemToGallery(item.Data.ID, gallery.Data.ID)
		if err != nil {
			log.Fatal(err)
			continue
		}
	}
	return nil
}

func main() {
	_, err := meural.GetToken(os.Getenv("MEURAL_EMAIL"), os.Getenv("MEURAL_PASSWORD"))
	if err != nil {
		log.Fatal(err)
		return
	}

	err = PostRandomPixivToMeural()
	if err != nil {
		log.Fatal(err)
	}
}
