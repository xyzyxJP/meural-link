package main

import (
	"log"
	"math/rand"
	"os"
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

	userProfile, err := pixiv.GetUserProfile(os.Getenv("PIXIV_USER_ID"))
	if err != nil {
		return err
	}

	index := rand.Intn(len(userProfile.Body.Illusts))
	illustId := keys(userProfile.Body.Illusts)[index]
	illust, err := pixiv.GetIllust(illustId)
	if err != nil {
		return err
	}

	galleryID, err := meural.GetUserGalleryID(illust.Body.UserName)
	if err != nil {
		log.Fatal(err)
	}

	if galleryID == 0 {
		// TODO: Create gallery
	}

	illustPages, err := pixiv.GetIllustPages(illustId)
	if err != nil {
		return err
	}

	for i := range illustPages.Body {
		filename := pixiv.GetIllustFilename(illustPages, illust, i)
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

		item, err = meural.PutItem(item.Data.ID, illust.Body.Title, illust.Body.UserName, illust.Body.Description, "https://www.pixiv.net/artworks/"+illust.Body.IllustID, illust.Body.CreateDate.Year())
		if err != nil {
			log.Fatal(err)
			continue
		}

		// TODO: Add item to gallery
	}
	return nil
}

func main() {
	_, err := meural.GetToken(os.Getenv("MEURAL_EMAIL"), os.Getenv("MEURAL_PASSWORD"))
	if err != nil {
		log.Fatal(err)
		return
	}

	// err := PostRandomPixivToMeural()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
