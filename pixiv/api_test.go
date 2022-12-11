package pixiv

import (
	"fmt"
	"log"
	"testing"
)

func TestGetUserDetails(*testing.T) {
	userDetails, err := GetUserDetails("1039353")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userDetails)
}

func TestGetUserProfile(*testing.T) {
	userProfile, err := GetUserProfile("1039353")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userProfile)
}

func TestGetIllust(*testing.T) {
	illust, err := GetIllust("103449092")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(illust)
}

func TestGetIllustPages(t *testing.T) {
	illustPages, err := GetIllustPages("103449092")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(illustPages)
}

func TestGetIllustImage(t *testing.T) {
	err := GetIllustImage("https://i.pximg.net/img-original/img/2022/12/09/00/00/10/103449092_p0.png", "test.png")
	if err != nil {
		log.Fatal(err)
	}
}
