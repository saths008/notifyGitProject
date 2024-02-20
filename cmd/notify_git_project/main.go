package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/google/go-github/v59/github"
	"github.com/joho/godotenv"
)

func startDaemon(client *github.Client, owner string, repoName string) {
	lastTimeUpdate := time.Now()
	<-time.Tick(time.Hour * 1)

	for range time.Tick(time.Hour * 1) {
		log.Println(fmt.Sprint("Sending "))
		lastPush, err := callGitHubAPI(client, owner, repoName)
		if err != nil {
			log.Fatal(err)
		}
		err = sendNotification(owner, repoName, lastPush, lastTimeUpdate)
		if err != nil {
			log.Fatal(err)
		}
		lastTimeUpdate = time.Now()
	}
}

// Sends the correct notification according to
// the last push to the repository and the last time
// this notification was sent.
func sendNotification(owner string, repoName string, lastPush time.Time, lastTimeUpdate time.Time) error {
	compareTimeStamps := lastPush.Compare(lastTimeUpdate)
	if (compareTimeStamps == 1) || (compareTimeStamps == 0) {
		log.Println("Repo Updated Notification Sent")
		err := notifyRepoUpdated(owner, repoName, lastPush)
		if err != nil {
			return err
		}
		return nil
	} else {
		log.Println("No Update to Repo Notification Sent")
		err := notifyNoUpdate(owner, repoName)
		if err != nil {
			return err
		}
		return nil

	}

}

// Returns the last push to the repository requested.
func callGitHubAPI(client *github.Client, owner string, repoName string) (time.Time, error) {
	repo, response, err := client.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		return time.Now(), err
	}
	if !((200 <= response.StatusCode) && (response.StatusCode <= 226)) {
		fmt.Println("Response failed: ", response.StatusCode)
		return time.Now(), err
	}
	return repo.GetPushedAt().Time, nil
}

// Notification for no update to a repository.
func notifyNoUpdate(owner string, repoName string) error {
	err := beeep.Notify(fmt.Sprintf("No update for %s/%s", owner, repoName), fmt.Sprintf("Last Checked: %s", time.Now().Format(time.UnixDate)), "assets/information.png")
	if err != nil {
		return err
	}
	return nil
}

// Notification for an update to a repository
func notifyRepoUpdated(owner string, repoName string, timeOfUpdate time.Time) error {
	err := beeep.Notify(fmt.Sprintf("%s/%s Updated!", owner, repoName), fmt.Sprintf("Time of Update: %s", timeOfUpdate.Format(time.UnixDate)), "assets/information.png")
	if err != nil {
		return err
	}
	return nil
}
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	GH_TOKEN, tokenIsDefined := os.LookupEnv("GH_TOKEN")
	if !tokenIsDefined {
		log.Fatal("Make sure GH_TOKEN is defined in the .env")
	}
	owner := "saths008"
	repoName := "notifyGitProject"
	client := github.NewClient(nil).WithAuthToken(GH_TOKEN)
	log.Println("GitHub Client started.")
	startDaemon(client, owner, repoName)
}
