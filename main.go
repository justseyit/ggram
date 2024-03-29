package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// New creates a new app.
func New() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR: something went wrong with your .env file")
	}

	return &App{
		username:   os.Getenv("USERNAME"),
		password:   os.Getenv("PASSWORD"),
		followings: map[string]bool{},
		followers:  map[string]bool{},
		leeches:    []string{},
	}
}

func main() {
	// outstanding title!
	fmt.Printf("\n███████╗██████╗ ███╗   ███╗███████╗███████╗\n")
	fmt.Printf("██╔════╝██╔══██╗████╗ ████║██╔════╝██╔════╝\n")
	fmt.Printf("█████╗  ██████╔╝██╔████╔██║█████╗  ███████╗\n")
	fmt.Printf("██╔══╝  ██╔══██╗██║╚██╔╝██║██╔══╝  ╚════██║\n")
	fmt.Printf("███████╗██║  ██║██║ ╚═╝ ██║███████╗███████║\n")
	fmt.Printf("╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝╚══════╝\n\n")

	// CLI flags
	// Available commands (./ermes -h):
	// ./ermes -unfollow
	// ./ermes -followers
	// ./ermes -followers -reset
	// ./ermes -user=username -like
	followersPtr := flag.Bool("followers", false, "Like user's followers.")
	followingsPtr := flag.Bool("followings", false, "Like user's followings.")
	noCheckPtr := flag.Bool("no-check", false, "Check saved users.")
	skipPtr := flag.Bool("skip", false, "Skip users checks and start to like/follow.")
	timelinePtr := flag.Bool("timeline", false, "Like timeline. Latest 16 posts.")
	unfollowPtr := flag.Bool("unfollow", false, "Unfollow the ingrates.")
	userPtr := flag.String("user", "empty", "Follow vip's followers.")

	flag.Parse()

	app := New()
	app.Login()

	for _,follower := range app.followers {
		fmt.Println(follower)
	}

	defer app.Logout()

	app.InitDB()

	
	if *timelinePtr {
		app.LikeMyTimeline()
	}

	if *unfollowPtr {
		app.Unfollow()
	}

	if *followersPtr {
		app.LikeFeedFollowers(*skipPtr)
	}

	if *followingsPtr {
		app.LikeFeedFollowings(*skipPtr)
	}

	if *userPtr != "empty" {
		app.ShadowUser(*userPtr, *skipPtr, *noCheckPtr)
	}
	

}
